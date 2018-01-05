// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"github.com/freetaxii/libstix2/resources"
	"log"
	"strings"
)

// GetEnabledCollections - This method will return all of the collections that
// are currently enabled.
func (ds *Sqlite3DatastoreType) GetEnabledCollections() resources.CollectionsType {

	allCollections := resources.NewCollections()

	getAllEnabledCollections, nil := ds.sqlEnabledCollections()

	// Query database for all the collections
	rows, err := ds.DB.Query(getAllEnabledCollections)
	if err != nil {
		log.Fatal("ERROR: Database execution error quering collection: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var iCanRead, iCanWrite int
		var id, title, description, mediaType string
		if err := rows.Scan(&id, &title, &description, &iCanRead, &iCanWrite, &mediaType); err != nil {
			log.Fatal(err)
		}

		// Add collection information to Collection object
		c := allCollections.NewCollection()
		c.SetID(id)
		c.SetTitle(title)
		c.SetDescription(description)
		if iCanRead == 1 {
			c.SetCanRead()
		}
		if iCanWrite == 1 {
			c.SetCanWrite()
		}

		mediatypes := strings.Split(mediaType, ",")
		for i, mt := range mediatypes {

			// If the media types are all the same, due to the way the SQL query
			// returns results, then only record one entry.
			if i > 0 && mt == mediatypes[i-1] {
				continue
			}
			c.AddMediaType(mt)
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return allCollections
}
