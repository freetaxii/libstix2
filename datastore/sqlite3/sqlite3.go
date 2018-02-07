// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/freetaxii/libstix2/datastore"
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/resources"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strings"
)

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
DatastoreType defines all of the properties and information associated
with connecting and talking to the database.

When StrictSTIXIDs = false, then the system will allow vanity STIX IDs like:
indicator--1, indicator--2
*/
type DatastoreType struct {
	Filename        string
	DB              *sql.DB
	StrictSTIXIDs   bool
	StrictSTIXTypes bool
	LogLevel        int
	Cache           datastore.DatastoreCacheType
}

// ----------------------------------------------------------------------
//
// Public Create Functions
//
// ----------------------------------------------------------------------

/*
New - This function will return a DatastoreType.
*/
func New(filename string) DatastoreType {
	var err error
	var ds DatastoreType
	ds.Filename = filename
	ds.StrictSTIXIDs = false
	ds.StrictSTIXTypes = true
	ds.LogLevel = 0

	err = ds.connect()
	if err != nil {
		log.Fatalln(err)
	}

	// Initialize the datastore cache which will have the current object ID
	// and the collections data.
	err = ds.initCache()
	if err != nil {
		log.Fatalln(err)
	}

	return ds
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
func (ds *DatastoreType) GetSTIXObject(stixid, version string) (interface{}, error) {

	idparts := strings.Split(stixid, "--")

	if ds.StrictSTIXIDs == true {
		if !objects.IsValidID(stixid) {
			return nil, errors.New("get STIX object error, invalid STIX ID")
		}
	}

	if ds.StrictSTIXTypes == true {
		if !objects.IsValidSTIXObject(stixid) {
			return nil, errors.New("get STIX object error, invalid STIX type")
		}
	}

	switch idparts[0] {
	case "indicator":
		return ds.getIndicator(stixid, version)
	}

	return nil, fmt.Errorf("get object error, the following STIX type is not currently supported: ", idparts[0])
}

/*
AddSTIXObject - This method will take in a STIX object and add it to the
database.
*/
func (ds *DatastoreType) AddSTIXObject(obj interface{}) error {
	switch o := obj.(type) {
	case *objects.IndicatorType:
		ds.addIndicator(o)
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
	case *resources.CollectionType:
		err = ds.addCollection(o)
	case *resources.CollectionRecordType:
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

/*
GetAllCollections - This method will return all collections, even those that
are disabled and hidden. This is primarily used for administration tools that
need to see all collections.
*/
func (ds *DatastoreType) GetAllCollections() (*resources.CollectionsType, error) {
	return ds.getCollections("all")
}

/*
GetAllEnabledCollections - This method will return only enabled collections,
even those that are hidden. This is used for setup up the HTTP MUX routers.
*/
func (ds *DatastoreType) GetAllEnabledCollections() (*resources.CollectionsType, error) {
	return ds.getCollections("allEnabled")
}

/*
GetCollections - This method will return just those collections that are both
enabled and visible. This is primarily used for clients that pull a collections
resource.
*/
func (ds *DatastoreType) GetCollections() (*resources.CollectionsType, error) {
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
func (ds *DatastoreType) GetBundle(query datastore.CollectionQueryType) (*datastore.CollectionQueryResultType, error) {
	return ds.getBundle(query)
}

/*
GetManifestData - This method will take in query struct with range
parameters for a collection and will return a TAXII manifest that contains all
of the records that match the query and range parameters.
*/
func (ds *DatastoreType) GetManifestData(query datastore.CollectionQueryType) (*datastore.CollectionQueryResultType, error) {
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

	//log.Println("Connecting to sqlite3 datastore at filename", ds.Filename)

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
	ds.Cache.Collections = make(map[string]*resources.CollectionType)

	if ds.LogLevel >= 5 {
		log.Println("DEBUG: Entering initCache()")
	}

	// Get current index value so new records being added can use it.
	objectIndex, err := ds.getBaseObjectIndex()
	if err != nil && err.Error() != "no base object record found" {
		return err
	}
	ds.Cache.BaseObjectIDIndex = objectIndex + 1

	if ds.LogLevel >= 5 {
		log.Println("DEBUG: Base object index ID", ds.Cache.BaseObjectIDIndex)
	}

	// Populate the collections cache
	ds.Cache.Collections = make(map[string]*resources.CollectionType)

	// Lets initialize the collections cache from the datastore
	allCollections, err := ds.GetAllCollections()

	if err != nil {
		return err
	}

	for k, c := range allCollections.Collections {
		ds.Cache.Collections[c.ID] = &allCollections.Collections[k]
		// get the size of the collection
		size, err3 := ds.getCollectionSize(c.ID)
		if err3 != nil {
			return err3
		}
		// If there was no error, set the size of the collection in the cache
		ds.Cache.Collections[c.ID].Size = size
	}

	if ds.LogLevel >= 5 {
		for k, v := range ds.Cache.Collections {
			log.Println("DEBUG: Collection Cache Key", k, "Collection ID", v.ID, "Size", v.Size)
		}
	}

	return nil
}
