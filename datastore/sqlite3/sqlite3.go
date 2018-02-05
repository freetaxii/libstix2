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
	Cache           datastore.DatabaseCacheType
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

	err = ds.connect()
	if err != nil {
		log.Fatalln(err)
	}

	// Get current index value so new records being added can use it.
	ds.getBaseObjectIndex()
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

// /*
// GetBundle - This method will take in a query struct with range
// parameters for a collection and will return a STIX Bundle that contains all
// of the STIX objects that are in that collection that meet those query or range
// parameters.
// */
// func (ds *DatastoreType) GetBundle(query datastore.CollectionQueryType) (*objects.BundleType, *datastore.CollectionQueryResultType, error) {
// 	stixBundle := objects.InitBundle()

// 	// First get a list of all of the objects that are in the collection that
// 	// meet the query requirements
// 	rangeCollectionRawData, metaData, err := ds.GetObjectList(query)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	// Loop through all of the STIX IDs in the list and get the actual object
// 	for _, v := range *rangeCollectionRawData {
// 		log.Println("STIX ID: ", v.STIXID, " Version: ", v.STIXVersion)
// 		obj, err := ds.GetSTIXObject(v.STIXID, v.STIXVersion)

// 		if err != nil {
// 			return nil, nil, err
// 		}
// 		stixBundle.AddObject(obj)
// 	}

// 	return stixBundle, metaData, nil
// }

// /*
// GetObjectList - This method will take in a query struct with range
// parameters for a collection and will return a datastore collection raw data type
// that contains all of the STIX IDs and their associated meta data that are in
// that collection that meet those query or range parameters.
// */
// func (ds *DatastoreType) GetObjectList(query datastore.CollectionQueryType) (*[]datastore.CollectionRawDataType, *datastore.CollectionQueryResultType, error) {
// 	var metaData datastore.CollectionQueryResultType
// 	var collectionRawData []datastore.CollectionRawDataType
// 	var rangeCollectionRawData []datastore.CollectionRawDataType

// 	sqlStmt, err := sqlGetObjectList(query)

// 	// If an error is found, that means a query parameter was passed incorrectly
// 	// and we should return an error versus just skipping the option.
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	// Query database for all the collection entries
// 	rows, err := ds.DB.Query(sqlStmt)
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("database execution error getting collection data: ", err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var dateAdded, stixid, modified, specVersion string
// 		if err := rows.Scan(&stixid, &dateAdded, &modified, &specVersion); err != nil {
// 			rows.Close()
// 			return nil, nil, fmt.Errorf("database scan error: ", err)
// 		}
// 		var rawData datastore.CollectionRawDataType
// 		rawData.STIXID = stixid
// 		rawData.DateAdded = dateAdded
// 		rawData.STIXVersion = modified
// 		rawData.SpecVersion = specVersion

// 		collectionRawData = append(collectionRawData, rawData)
// 	}

// 	// Errors can cause the rows.Next() to exit prematurely, if this happens lets
// 	// check for the error and handle it.
// 	if err := rows.Err(); err != nil {
// 		rows.Close()
// 		return nil, nil, fmt.Errorf("database rows error getting collection data: ", err)
// 	}

// 	metaData.Size = len(collectionRawData)

// 	// If no records are returned, then return an error before processing anything else.
// 	if metaData.Size == 0 {
// 		return nil, nil, errors.New("no records returned")
// 	}

// 	first, last, errRange := ds.processRangeValues(query.RangeBegin, query.RangeEnd, query.RangeMax, metaData.Size)

// 	if errRange != nil {
// 		return nil, nil, errRange
// 	}

// 	// Get a new slice based on the range of records
// 	rangeCollectionRawData = collectionRawData[first:last]
// 	metaData.DateAddedFirst = rangeCollectionRawData[0].DateAdded
// 	metaData.DateAddedLast = rangeCollectionRawData[len(rangeCollectionRawData)-1].DateAdded
// 	metaData.RangeBegin = first
// 	metaData.RangeEnd = last - 1

// 	// metaData is already a pointer
// 	return &rangeCollectionRawData, &metaData, nil
// }

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

func (ds *DatastoreType) initCache() error {
	ds.Cache.Collections = make(map[string]*resources.CollectionType)

	// Lets initialize the collections cache from the datastore
	allCollections, err := ds.GetAllCollections()

	if err != nil {
		return err
	}

	for _, c := range allCollections.Collections {
		ds.Cache.Collections[c.ID] = &c
	}
	return nil
}
