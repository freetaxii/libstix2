// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"github.com/freetaxii/libstix2/objects"
	"log"
)

// ListObjectsInCollection - This method will take in an ID for a collection
// and return a slice of strings that contains all of the STIX IDs that are in
// the colleciton.
func (ds *Sqlite3DatastoreType) ListObjectsInCollection(cid string) []string {
	var allObjects []string

	var getAllObjectsInCollection = `
		SELECT stix_id
	   	FROM ` + DB_TABLE_TAXII_COLLECTION_CONTENT + ` 
	   	WHERE collection_id = $1`

	// Query database for all the collection entries
	rows, err := ds.DB.Query(getAllObjectsInCollection, cid)
	if err != nil {
		log.Fatal("ERROR: Database execution error quering collection content: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var stixid string
		if err := rows.Scan(&stixid); err != nil {
			log.Fatal(err)
		}
		allObjects = append(allObjects, stixid)
	}
	return allObjects
}

// GetObjectsInCollection - This method will take in an ID for a collection
// and return a STIX Bundle that contains all of the STIX objects that are in
// the collection.
func (ds *Sqlite3DatastoreType) GetObjectsInCollection(cid string) objects.BundleType {
	stixBundle := objects.NewBundle()

	allObjects := ds.ListObjectsInCollection(cid)
	for _, stixid := range allObjects {
		obj, _ := ds.GetObject(stixid)
		stixBundle.AddObject(obj)
	}
	return stixBundle
}
