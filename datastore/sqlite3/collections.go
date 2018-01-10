// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"fmt"
	"github.com/freetaxii/libstix2/resources"
	"strings"
)

// GetAllCollections - This method will return all collections, even those that
// are not enabled and hidden
func (ds *Sqlite3DatastoreType) GetAllCollections() (*resources.CollectionsType, error) {
	return ds.getCollections("all")
}

// GetAllEnabledCollections - This method will return all enabled collections,
// even those that are hidden. This is used for setup up the http routers
func (ds *Sqlite3DatastoreType) GetAllEnabledCollections() (*resources.CollectionsType, error) {
	return ds.getCollections("allEnabled")
}

// GetCollections - This method will return all enabled and visible collections
func (ds *Sqlite3DatastoreType) GetCollections() (*resources.CollectionsType, error) {
	return ds.getCollections("enabledVisible")
}

// getCollections - This method will return all of the collections that
// are currently enabled.
func (ds *Sqlite3DatastoreType) getCollections(whichCollections string) (*resources.CollectionsType, error) {

	allCollections := resources.NewCollections()

	getAllCollectionsStmt, _ := ds.sqlAllCollections(whichCollections)

	// Query database for all the collections
	rows, err := ds.DB.Query(getAllCollectionsStmt)
	if err != nil {
		return nil, fmt.Errorf("database execution error querying collection: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var enabled, hidden, iCanRead, iCanWrite int
		var dateAdded, id, title, description, mediaType string
		if err := rows.Scan(&dateAdded, &enabled, &hidden, &id, &title, &description, &iCanRead, &iCanWrite, &mediaType); err != nil {
			return nil, fmt.Errorf("database scan error querying collection: ", err)
		}

		// Add collection information to Collection object
		c := allCollections.NewCollection()
		c.DateAdded = dateAdded
		if enabled == 1 {
			c.SetEnabled()
		} else {
			c.SetDisabled()
		}

		if hidden == 1 {
			c.SetHidden()
		} else {
			c.SetVisible()
		}

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
		return nil, fmt.Errorf("database row error querying collection: ", err)
	}

	return allCollections, nil
}
