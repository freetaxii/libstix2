// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"database/sql"
	"errors"
	"fmt"
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
Sqlite3DatastoreType defines all of the properties and information associated
with connecting and talking to the database.

When StrictSTIXIDs = false, then the system will allow vanity STIX IDs like:
indicator--1, indicator--2
*/
type Sqlite3DatastoreType struct {
	Filename        string
	DB              *sql.DB
	LogLevel        int
	StrictSTIXIDs   bool
	StrictSTIXTypes bool
	Index           int64
}

// ----------------------------------------------------------------------
//
// Public Create Functions
//
// ----------------------------------------------------------------------

/*
New - This function will return a Sqlite3DatastoreType.
*/
func New(filename string) Sqlite3DatastoreType {
	var ds Sqlite3DatastoreType
	ds.Filename = filename
	ds.LogLevel = 5
	ds.StrictSTIXIDs = false
	ds.StrictSTIXTypes = true

	// TODO get current index from database or other stored area
	// we also need a way of updating the database or getting the value
	// from the database somehow.
	ds.Index = 1

	err := ds.connect()
	if err != nil {
		log.Fatalln(err)
	}

	return ds
}

/*
Close - This method will close the database connection
*/
func (ds *Sqlite3DatastoreType) Close() error {
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
modified timestamp from a STIX object) and return the STIX object.
*/
func (ds *Sqlite3DatastoreType) GetSTIXObject(stixid, version string) (interface{}, error) {
	idparts := strings.Split(stixid, "--")

	if ds.StrictSTIXIDs == true {
		if !objects.IsValidID(stixid) {
			return nil, errors.New("get object error, invalid STIX ID")
		}
	}

	if ds.StrictSTIXTypes == true {
		if !objects.IsValidSTIXObject(stixid) {
			return nil, errors.New("get object error, invalid STIX type")
		}
	}

	switch idparts[0] {
	case "indicator":
		return ds.getIndicator(stixid, version)
	}

	return nil, fmt.Errorf("get object error, the following STIX type is not currently supported: ", idparts[0])
}

func (ds *Sqlite3DatastoreType) Add(obj interface{}) {
	switch o := obj.(type) {
	case *resources.CollectionType:
		ds.addCollection(o)
	case *resources.CollectionRecordType:
		ds.addObjectToCollection(o)
	case *objects.IndicatorType:
		ds.addIndicator(o)
	default:
		log.Println("ERROR: Does not match any known types ", o)
	}

}

// ----------------------------------------------------------------------
//
// Private Methods
//
// ----------------------------------------------------------------------

// connect - This method is used to connect to an sqlite3 database
func (ds *Sqlite3DatastoreType) connect() error {
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

// verifyFileExists - This method will check to make sure the file is found on the filesystem
func (ds *Sqlite3DatastoreType) verifyFileExists() error {
	if _, err := os.Stat(ds.Filename); os.IsNotExist(err) {
		w, err2 := os.Create(ds.Filename)
		w.Close()
		if err2 != nil {
			return fmt.Errorf("ERROR: The sqlite3 database cannot be opened due to error: %v", err2)
		}
	}
	return nil
}
