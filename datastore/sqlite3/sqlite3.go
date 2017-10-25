// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"crypto/sha1"
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/freetaxii/libstix2/defs"
	"github.com/freetaxii/libstix2/objects/indicator"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"time"
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

// New - This function will return a datastore.STIXDatastorer
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

func (ds *Sqlite3DatastoreType) Put(obj interface{}) {
	switch o := obj.(type) {
	case indicator.IndicatorType:
		ds.addIndicatorToDatabase(o)
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

func (ds *Sqlite3DatastoreType) addIndicatorToDatabase(obj indicator.IndicatorType) error {
	// TODO change, add to object creation
	ver := "2.0"

	var stmt = `INSERT INTO "stix_base_object" (
	 	"object_id",
	 	"version",
	 	"date_added",
	 	"type",
	 	"id",
	 	"created_by_ref",
	 	"created",
	 	"modified",
	 	"revoked",
	 	"confidence",
	 	"lang"
		)
		values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	dateAdded := time.Now().UTC().Format(defs.TIME_RFC_3339_MICRO)
	objectID := "id" + obj.ID + "created" + obj.Created + "modified" + obj.Modified

	h := sha1.New()
	h.Write([]byte(objectID))
	hashID := base64.URLEncoding.EncodeToString(h.Sum(nil))

	_, err := ds.DB.Exec(stmt,
		hashID,
		ver,
		dateAdded,
		obj.MessageType,
		obj.ID,
		obj.CreatedByRef,
		obj.Created,
		obj.Modified,
		obj.Revoked,
		obj.Confidence,
		obj.Lang)

	if err != nil {
		return err
	}

	var stmt1 = `INSERT INTO "sdo_indicator" (
		"object_id",
		"name",
		"description",
		"pattern",
		"valid_from",
		"valid_until"
		)
		values (?, ?, ?, ?, ?, ?)`

	_, err1 := ds.DB.Exec(stmt1,
		hashID,
		obj.Name,
		obj.Description,
		obj.Pattern,
		obj.ValidFrom,
		obj.ValidUntil)

	// TODO if there is an error, we probably need to back out all of the INSERTS
	if err1 != nil {
		return err
	}

	if obj.KillChainPhases != nil {
		for _, v := range obj.KillChainPhases {
			var stmt2 = `INSERT INTO "kill_chain_phases" (
			"object_id",
			"kill_chain_name",
			"phase_name"
			)
			values (?, ?, ?)`

			_, err2 := ds.DB.Exec(stmt2, hashID, v.KillChainName, v.PhaseName)

			if err2 != nil {
				return err
			}
		}
	}

	if obj.Labels != nil {
		for _, v1 := range obj.Labels {
			var stmt3 = `INSERT INTO "labels" (
			"object_id",
			"labels"
			)
			values (?, ?)`

			_, err3 := ds.DB.Exec(stmt3, hashID, v1)

			if err3 != nil {
				return err
			}
		}
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

	log.Println("Connecting to sqlite3 datastore at filename", ds.Filename)

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
		return fmt.Errorf("ERROR: The sqlite3 database cannot be opened due to error: %v", err)
	}
	return nil
}
