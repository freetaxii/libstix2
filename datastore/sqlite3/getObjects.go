// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"github.com/freetaxii/libstix2/datastore"
	"github.com/freetaxii/libstix2/objects"
	"log"
)

// GetObject - This method will take in a STIX ID and return the STIX object.
func (ds *Sqlite3DatastoreType) GetObject(stixid string) (interface{}, error) {

	// TODO
	// We first need to look at the STIX ID that was passed in to see what type
	// of object it is. Basically split the ID to get the type and then do a
	// switch statement
	// Need to also be able to handle multiple versions
	i, err := ds.getIndicator(stixid)
	return i, err
}

// GetListOfObjectInCollection - This method will take in an ID for a collection
// and return a slice of strings that contains all of the STIX IDs that are in
// that colleciton.
func (ds *Sqlite3DatastoreType) GetListOfObjectInCollection(collectionid string) []string {
	var allObjects []string

	var getAllObjectsInCollection = `
		SELECT stix_id
	   	FROM ` + datastore.DB_TABLE_TAXII_COLLECTION_CONTENT + ` 
	   	WHERE collection_id = $1`

	// Query database for all the collection entries
	rows, err := ds.DB.Query(getAllObjectsInCollection, collectionid)
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
// that collection.
func (ds *Sqlite3DatastoreType) GetObjectsInCollection(collectionid string) objects.BundleType {
	// TODO need the ability to take in a query struct of list of parameters

	stixBundle := objects.NewBundle()

	allObjects := ds.GetListOfObjectInCollection(collectionid)
	for _, stixid := range allObjects {
		obj, _ := ds.GetObject(stixid)
		stixBundle.AddObject(obj)
	}
	return stixBundle
}
