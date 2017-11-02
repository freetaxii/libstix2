// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"github.com/freetaxii/libstix2/defs"
	"github.com/freetaxii/libstix2/resources"
	"log"
	"time"
)

// addCollection
func (ds *Sqlite3DatastoreType) addCollection(obj resources.CollectionType) {
	dateAdded := time.Now().UTC().Format(defs.TIME_RFC_3339_MICRO)

	var stmt1 = `INSERT INTO ` + defs.DB_TABLE_TAXII_COLLECTION + ` (
	 	"date_added", 
	 	"id", 
	 	"title",
	 	"description", 
	 	"can_read", 
	 	"can_write"
	 	)
		values (?, ?, ?, ?, ?, ?)`

	_, err1 := ds.DB.Exec(stmt1,
		dateAdded,
		obj.ID,
		obj.Title,
		obj.Description,
		obj.CanRead,
		obj.CanWrite)

	if err1 != nil {
		log.Println("ERROR: Database execution error inserting collection", err1)
	}

	if obj.MediaTypes != nil {
		for _, media := range obj.MediaTypes {
			var stmt2 = `INSERT INTO ` + defs.DB_TABLE_TAXII_COLLECTION_MEDIA_TYPE + ` (
			"collection_id",
			"media_type"
			)
			values (?, ?)`

			_, err2 := ds.DB.Exec(stmt2, obj.ID, media)

			if err2 != nil {
				log.Println("ERROR: Database execution error collection media type", err2)
			}
		}
	}
}

// addObjectToColleciton - This method will add an object to a collection
// by adding an entry in the taxii_collection_content table. In this table
// we use the STIX ID not the Object ID because we need to make sure we
// include all versions of an object. So we need to store just the STIX ID
func (ds *Sqlite3DatastoreType) addObjectToCollection(cID, sID string) {
	dateAdded := time.Now().UTC().Format(defs.TIME_RFC_3339_MICRO)

	var stmt1 = `INSERT INTO ` + defs.DB_TABLE_TAXII_COLLECTION_CONTENT + ` (
	 	"date_added", 
	 	"collection_id", 
	 	"stix_id"
	 	)
		values (?, ?, ?)`

	_, err1 := ds.DB.Exec(stmt1,
		dateAdded,
		cID,
		sID)

	if err1 != nil {
		log.Println("ERROR: Database execution error inserting collection content", err1)
	}
}
