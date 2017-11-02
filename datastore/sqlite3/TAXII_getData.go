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
)

// GetEnabledCollections - This method will return all of the collections that
// are currently enabled.
func (ds *Sqlite3DatastoreType) GetEnabledCollections() resources.CollectionsType {

	allCollections := resources.NewCollections()

	var getAllEnabledCollections = `
		SELECT 
			id,
	    	title,
	    	description,
	    	can_read,
	    	can_write
	   	FROM ` + defs.DB_TABLE_TAXII_COLLECTION + ` 
	   	WHERE enabled = 1`

	var getMediaTypesForCollection = `
		SELECT media_type 
		FROM ` + defs.DB_TABLE_TAXII_COLLECTION_MEDIA_TYPE + ` 
		WHERE collection_id = ?`

	// Query database for all the collections
	rows, err := ds.DB.Query(getAllEnabledCollections)
	if err != nil {
		log.Fatal("ERROR: Database execution error quering collection: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var iCanRead, iCanWrite int
		var canRead, canWrite bool
		var id, title, description string
		if err := rows.Scan(&id, &title, &description, &iCanRead, &iCanWrite); err != nil {
			log.Fatal(err)
		}

		canRead = intToBool(iCanRead)
		canWrite = intToBool(iCanWrite)

		// Add collection infromation to Colleciton object
		c := allCollections.NewCollection()
		c.SetID(id)
		c.SetTitle(title)
		c.SetDescription(description)
		if canRead == true {
			c.SetCanRead()
		}
		if canWrite == true {
			c.SetCanWrite()
		}

		rows1, err1 := ds.DB.Query(getMediaTypesForCollection, id)
		if err1 != nil {
			log.Fatal("ERROR: Database execution error quering media types: ", err)
		}
		defer rows1.Close()

		for rows1.Next() {
			var mediaType string
			if err := rows1.Scan(&mediaType); err != nil {
				log.Fatal(err)
			}
			c.AddMediaType(mediaType)
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return allCollections
}

func intToBool(i int) bool {
	if i == 1 {
		return true
	}
	return false
}
