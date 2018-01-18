// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"fmt"
	"github.com/freetaxii/libstix2/defs"
	"github.com/freetaxii/libstix2/resources"
	"strings"
	"time"
)

// ----------------------------------------------------------------------
//
// Public Methods
//
// ----------------------------------------------------------------------

/*
GetAllCollections - This method will return all collections, even those that
are disabled and hidden. This is primarily used for administration tools that
need to see all collections.
*/
func (ds *Sqlite3DatastoreType) GetAllCollections() (*resources.CollectionsType, error) {
	return ds.getCollections("all")
}

/*
GetAllEnabledCollections - This method will return only enabled collections,
even those that are hidden. This is used for setup up the HTTP MUX routers.
*/
func (ds *Sqlite3DatastoreType) GetAllEnabledCollections() (*resources.CollectionsType, error) {
	return ds.getCollections("allEnabled")
}

/*
GetCollections - This method will return just those collections that are both
enabled and visible. This is primarily used for client that pull a collections
resource.
*/
func (ds *Sqlite3DatastoreType) GetCollections() (*resources.CollectionsType, error) {
	return ds.getCollections("enabledVisible")
}

// ----------------------------------------------------------------------
//
// Private Methods
//
// ----------------------------------------------------------------------

/*
addCollection - This method will add a collection to the t_collections table in
the database.
*/
func (ds *Sqlite3DatastoreType) addCollection(obj *resources.CollectionType) error {
	dateAdded := time.Now().UTC().Format(defs.TIME_RFC_3339_MICRO)

	stmt1, _ := ds.sqlAddCollection()

	_, err1 := ds.DB.Exec(stmt1,
		dateAdded,
		obj.ID,
		obj.Title,
		obj.Description,
		obj.CanRead,
		obj.CanWrite)

	if err1 != nil {
		return fmt.Errorf("database execution error inserting collection", err1)
	}

	if obj.MediaTypes != nil {
		for _, media := range obj.MediaTypes {
			stmt2, _ := ds.sqlAddCollectionMediaType()

			// TODO look up in cache
			mediavalue := 0
			if media == "application/vnd.oasis.stix+json" {
				mediavalue = 1
			}
			_, err2 := ds.DB.Exec(stmt2, obj.ID, mediavalue)

			if err2 != nil {
				return fmt.Errorf("database execution error inserting collection media type", err2)
			}
		}
	}
	return nil
}

/*
addObjectToColleciton - This method will add an object to a collection by adding
an entry in the taxii_collection_data table. In this table we use the STIX ID
not the Object ID because we need to make sure we include all versions of an
object. So we need to store just the STIX ID.
*/
func (ds *Sqlite3DatastoreType) addObjectToCollection(obj *resources.CollectionRecordType) error {
	dateAdded := time.Now().UTC().Format(defs.TIME_RFC_3339_MICRO)

	stmt, _ := ds.sqlAddObjectToCollection()
	_, err := ds.DB.Exec(stmt, dateAdded, obj.CollectionID, obj.STIXID)

	if err != nil {
		return fmt.Errorf("database execution error inserting collection data", err)
	}
	return nil
}

/*
getCollections - This method is called from either GetAllCollections(),
GetAllEnabledCollections(), or GetCollections() and will return all of the
collections that are asked for based on the method that called it.  The options
that can be passed in are: "all", "allEnabled", and "enabledVisible". The "all"
option returns every collection, even those that are hidden or disabled.
"allEnabled" will return all enabled collections, even those that are hidden.
"enabledVisible" will return all collections that are both enabled and not
hidden (aka those that are visible). Administration tools using the database
will probably want to see all collections. The HTTP Router MUX needs to know
about all enabled collections, even those that are hidden, so that it can start
an HTTP router for it. The enabled and visible list is what would be displayed
to a client that is pulling a collections resource.
*/
func (ds *Sqlite3DatastoreType) getCollections(whichCollections string) (*resources.CollectionsType, error) {

	allCollections := resources.InitCollections()

	getAllCollectionsStmt, _ := ds.sqlGetAllCollections(whichCollections)

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
		c, _ := allCollections.GetNewCollection()
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
