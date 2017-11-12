// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"database/sql"
	"fmt"
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/resources"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

// Sqlite3DatastoreType defines all of the properties and information associated
// with connecting and talking to the database.
type Sqlite3DatastoreType struct {
	Filename string
	DB       *sql.DB
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// New - This function will return a datastore.Datastorer which is any datastore
// that implements all of the methods that are defined in the Datastorer type.
func New(filename string) Sqlite3DatastoreType {
	var ds Sqlite3DatastoreType
	ds.Filename = filename

	err := ds.connect()
	if err != nil {
		log.Fatalln(err)
	}

	return ds
}

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

func (ds *Sqlite3DatastoreType) GetObject(stixid string) (interface{}, error) {

	// TODO
	// We first need to look at the STIX ID that was passed in to see what type
	// of object it is. Basically split the ID to get the type and then do a
	// switch statement
	i, err := ds.getIndicator(stixid)
	return i, err
}

func (ds *Sqlite3DatastoreType) Put(obj interface{}) {
	switch o := obj.(type) {
	case resources.CollectionType:
		ds.addCollection(o)
	case resources.CollectionRecordType:
		ds.addObjectToCollection(o)
	case objects.IndicatorType:
		ds.addIndicator(o)
	}
}

// Close - This method will close the database connection
func (ds *Sqlite3DatastoreType) Close() error {
	err := ds.DB.Close()
	if err != nil {
		return err
	}
	return nil
}

// ----------------------------------------------------------------------
// Private Methods
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
