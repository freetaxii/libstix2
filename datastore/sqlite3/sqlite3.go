// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/freetaxii/libstix2/common/stixid"
	"github.com/freetaxii/libstix2/datastore"
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/resources"
	"github.com/gologme/log"
	_ "github.com/mattn/go-sqlite3"
)

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
DatastoreType defines all of the properties and information associated
with connecting and talking to the database.

When Strict.IDs = false, then the system will allow vanity STIX IDs like:
indicator--1, indicator--2

When Strict.Types = false, then the system will allow un-known STIX types
*/
type DatastoreType struct {
	Filename string
	DB       *sql.DB
	Logger   *log.Logger
	Cache    datastore.MemCache
	Strict   struct {
		IDs   bool
		Types bool
	}
}

// ----------------------------------------------------------------------
//
// Public Create Functions
//
// ----------------------------------------------------------------------

/*
New - This function will return a DatastoreType.
*/
func New(logger *log.Logger, filename string) *DatastoreType {
	var err error
	var ds DatastoreType
	ds.Filename = filename
	ds.Strict.IDs = false
	ds.Strict.Types = true

	if logger == nil {
		ds.Logger = log.New(os.Stderr, "", log.LstdFlags)
	} else {
		ds.Logger = logger
	}

	err = ds.connect()
	if err != nil {
		log.Fatalln(err)
	}

	// Initialize the MemCache which will have the current object ID
	// and the collections data.
	err = ds.initCache()
	if err != nil {
		log.Fatalln(err)
	}

	return &ds
}

/*
Close - This method will close the database connection
*/
func (ds *DatastoreType) Close() error {
	err := ds.DB.Close()
	if err != nil {
		return err
	}
	return nil
}

// ----------------------------------------------------------------------
//
// Public Methods
//
// ----------------------------------------------------------------------

/*
GetSTIXObject - This method will take in a STIX ID and version timestamp (the
modified timestamp from a STIX object) and return the matching STIX object.
*/
func (ds *DatastoreType) GetSTIXObject(id, version string) (interface{}, error) {

	idparts := strings.Split(id, "--")

	// If the datastore requires the id portion of the STIX id to be a valid
	// UUIDv4, then lets test that.
	if ds.Strict.IDs == true {
		if !stixid.ValidUUID(idparts[1]) {
			return nil, errors.New("get STIX object error, invalid STIX ID")
		}
	}

	if ds.Strict.Types == true {
		if !stixid.ValidSTIXObjectType(idparts[0]) {
			return nil, errors.New("get STIX object error, invalid STIX type")
		}
	}

	switch idparts[0] {
	case "indicator":
		return ds.getIndicator(id, version)
	}

	return nil, fmt.Errorf("get object error, the following STIX type is not currently supported: ", idparts[0])
}

/*
AddSTIXObject - This method will take in a STIX object and add it to the
database.
*/
func (ds *DatastoreType) AddSTIXObject(obj interface{}) error {
	switch o := obj.(type) {
	case *objects.Indicator:
		ds.Logger.Debugln("DEBUG: Found Indicator to add to datastore")
		err := ds.addIndicator(o)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("add object error, the following STIX type is not currently supported: ", o)
	}
	return nil
}

/*
AddTAXIIObject - This method will take in a TAXII object and add it to the
database.
*/
func (ds *DatastoreType) AddTAXIIObject(obj interface{}) error {
	var err error

	switch o := obj.(type) {
	case *resources.Collection:
		ds.Logger.Debugln("DEBUG: Adding TAXII Collection to datastore")
		err = ds.addCollection(o)
	case *resources.CollectionRecord:
		ds.Logger.Debugln("DEBUG: Adding TAXII Collection Record to datastore")
		err = ds.addObjectToCollection(o)
	default:
		err = fmt.Errorf("does not match any known types ", o)
	}
	if err != nil {
		return err
	}
	return nil
}

// ----------------------------------------------------------------------
//
// Collection Table Public Methods
//
// ----------------------------------------------------------------------

// TODO - These probably need to be done by API Root

/*
GetAllCollections - This method will return all collections, even those that
are disabled and hidden. This is primarily used for administration tools that
need to see all collections.
*/
func (ds *DatastoreType) GetAllCollections() (*resources.Collections, error) {
	return ds.getCollections("all")
}

/*
GetAllEnabledCollections - This method will return only enabled collections,
even those that are hidden. This is used for setup up the HTTP MUX routers.
*/
func (ds *DatastoreType) GetAllEnabledCollections() (*resources.Collections, error) {
	return ds.getCollections("allEnabled")
}

/*
GetCollections - This method will return just those collections that are both
enabled and visible. This is primarily used to populate the results for clients
that pull a collections resource. Clients may be able to talk to a hidden
collection, but they should not see it in the list.
*/
func (ds *DatastoreType) GetCollections() (*resources.Collections, error) {
	return ds.getCollections("enabledVisible")
}

// ----------------------------------------------------------------------
//
// Collection Data Table Public Methods
//
// ----------------------------------------------------------------------

/*
GetBundle - This method will take in a query struct with range
parameters for a collection and will return a STIX Bundle that contains all
of the STIX objects that are in that collection that meet those query or range
parameters.
*/
func (ds *DatastoreType) GetBundle(query resources.CollectionQuery) (*resources.CollectionQueryResult, error) {
	return ds.getBundle(query)
}

/*
GetManifestData - This method will take in query struct with range
parameters for a collection and will return a TAXII manifest that contains all
of the records that match the query and range parameters.
*/
func (ds *DatastoreType) GetManifestData(query resources.CollectionQuery) (*resources.CollectionQueryResult, error) {
	return ds.getManifestData(query)
}

// ----------------------------------------------------------------------
//
// Private Methods
//
// ----------------------------------------------------------------------

/*
connect - This method is used to connect to an sqlite3 database
*/
func (ds *DatastoreType) connect() error {
	var err error

	if ds.Filename == "" {
		return fmt.Errorf("A valid filename is required for connecting to the sqlite3 datastore")
	}

	err = ds.verifyFileExists()
	if err != nil {
		return err
	}

	//ds.Logger.Println("Connecting to sqlite3 datastore at filename", ds.Filename)

	db, sqlerr := sql.Open("sqlite3", ds.Filename)
	if sqlerr != nil {
		return fmt.Errorf("Unable to open file %s due to error: %v", ds.Filename, sqlerr)
	}
	ds.DB = db

	return nil
}

/*
verifyFileExists - This method will check to make sure the sqlite3 database file
is found on the file system
*/
func (ds *DatastoreType) verifyFileExists() error {
	if _, err := os.Stat(ds.Filename); os.IsNotExist(err) {
		w, err2 := os.Create(ds.Filename)
		w.Close()
		if err2 != nil {
			return fmt.Errorf("ERROR: The sqlite3 database cannot be opened due to error: %v", err2)
		}
	}
	return nil
}

/*
initCache - This method will populate the datastore cache.
*/
func (ds *DatastoreType) initCache() error {
	ds.Logger.Traceln("TRACE initCache(): Start ")
	ds.Cache.Collections = make(map[string]*resources.Collection)

	// Get current index value of the s_base_object table so new records being
	// added can use it as their object_id. By using an integer here instead
	// of the full STIX ID, we can save significant amounts of space.
	// TODO - fix this once I setup my own error type
	ds.Logger.Traceln("TRACE: Call getBaseObjectIndex()")
	baseObjectIndex, err := ds.getBaseObjectIndex()
	if err != nil && err.Error() != "no base object record found" {
		return err
	}
	ds.Cache.BaseObjectIDIndex = baseObjectIndex + 1

	ds.Logger.Debugln("DEBUG: Base object index ID", ds.Cache.BaseObjectIDIndex)

	// Populate the collections cache
	ds.Cache.Collections = make(map[string]*resources.Collection)

	// Lets initialize the collections cache from the datastore
	allCollections, err := ds.GetAllCollections()
	if err != nil {
		return err
	}

	for k, c := range allCollections.Collections {
		ds.Cache.Collections[c.ID] = &allCollections.Collections[k]
		// get the size of the collection
		ds.Logger.Traceln("TRACE: Call getCollectionSize() for collection", c.ID)
		size, err3 := ds.getCollectionSize(c.ID)
		if err3 != nil {
			return err3
		}
		// If there was no error, set the size of the collection in the cache
		ds.Cache.Collections[c.ID].Size = size
	}

	for k, v := range ds.Cache.Collections {
		ds.Logger.Debugln("DEBUG: Collection Cache Key", k, "Collection ID", v.ID, "Size", v.Size)
	}

	return nil
}
