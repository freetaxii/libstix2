// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/freetaxii/libstix2/defs"
	"github.com/freetaxii/libstix2/resources/collections"
)

// ----------------------------------------------------------------------
//
// Collection Table Private Functions
// Table property names and SQL statements
//
// ----------------------------------------------------------------------

/*
collectionProperties - This function will return the properties that make up the
collection table.

date_added  = The date that this collection was added to the system
enabled     = Is this collection currently enabled
hidden      = Is this collection currently hidden for this directory listing
id 		    = The collection ID, a UUIDv4 value
title 	    = The title of this collection
description = A long description about this collection
can_read    = A boolean flag that indicates if one can read from this collection
can_write   = A boolean flag that indicates if one can write to this collection
*/
func collectionProperties() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
	"date_added" TEXT NOT NULL,
	"enabled" INTEGER(1,0) NOT NULL DEFAULT 1,
	"hidden" INTEGER(1,0) NOT NULL DEFAULT 0,
	"id" TEXT NOT NULL,
	"title" TEXT NOT NULL,
	"description" TEXT,
	"can_read" INTEGER(1,0) NOT NULL DEFAULT 0,
	"can_write" INTEGER(1,0) NOT NULL DEFAULT 0
	`
}

/*
collectionMediaTypeProperties  - This function will return the properties that
make up the collection media type table

collection_id = The collection ID, a UUIDv4 value
media_type_id = The media types supported on this collection
*/
func collectionMediaTypeProperties() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
	"collection_id" TEXT NOT NULL,
	"media_type_id" INTEGER NOT NULL
	`
}

// ----------------------------------------------------------------------
//
// Collection Table Private Functions and Methods
// addCollection
//
// ----------------------------------------------------------------------

/*
sqlAddCollection - This function will return an SQL statement that will insert
a new collection in to the t_collections table in the database.
*/
func sqlAddCollection() (string, error) {
	tblCol := DB_TABLE_TAXII_COLLECTIONS

	/*
		INSERT INTO
			t_collections (
				"date_added",
				"id",
				"title",
				"description",
				"can_read",
				"can_write"
			)
			values (?, ?, ?, ?, ?, ?)
	*/

	var s bytes.Buffer
	s.WriteString("INSERT INTO ")
	s.WriteString(tblCol)
	s.WriteString(" (")
	s.WriteString("date_added, ")
	s.WriteString("id, ")
	s.WriteString("title, ")
	s.WriteString("description, ")
	s.WriteString("can_read, ")
	s.WriteString("can_write) ")
	s.WriteString("values (?, ?, ?, ?, ?, ?)")

	return s.String(), nil
}

/*
sqlAddCollectionMediaType - This function will return an SQL statement that will
insert a media type for a given collection.
*/
func sqlAddCollectionMediaType() (string, error) {
	tblColMedia := DB_TABLE_TAXII_COLLECTION_MEDIA_TYPE

	/*
		INSERT INTO
			t_collection_media_type (
				"collection_id",
				"media_type_id"
			)
			values (?, ?)
	*/

	var s bytes.Buffer
	s.WriteString("INSERT INTO ")
	s.WriteString(tblColMedia)
	s.WriteString(" (")
	s.WriteString("\"collection_id\", ")
	s.WriteString("\"media_type_id\") ")
	s.WriteString("values (?, ?)")

	return s.String(), nil
}

/*
addCollection - This method will add a collection to the t_collections table in
the database.
*/
func (ds *Store) addCollection(obj *collections.Collection) error {
	ds.Logger.Traceln("TRACE addCollection(): Start")

	// Lets first make sure the collection does not already exist in the cache
	if _, found := ds.Cache.Collections[obj.ID]; found {
		return fmt.Errorf("the following collection id was already found in the cache", obj.ID)
	}
	// If the object ID is not found in the cache, then lets initialize it with
	// a TAXII collection object. This NewColleciton() function will return a
	// pointer, which is what we need here.
	ds.Cache.Collections[obj.ID] = collections.NewCollection()

	stmt1, _ := sqlAddCollection()
	dateAdded := time.Now().UTC().Format(defs.TIME_RFC_3339_MICRO)

	val, err1 := ds.DB.Exec(stmt1,
		dateAdded,
		obj.ID,
		obj.Title,
		obj.Description,
		obj.CanRead,
		obj.CanWrite)

	if err1 != nil {
		return fmt.Errorf("database execution error inserting collection", err1)
	}

	indexID, _ := val.LastInsertId()
	ds.Cache.Collections[obj.ID].DatastoreID = int(indexID)

	if obj.MediaTypes != nil {
		for _, media := range obj.MediaTypes {
			stmt2, _ := sqlAddCollectionMediaType()

			// TODO look up in cache
			mediavalue := 0
			switch media {
			case "application/stix+json;version=2.0":
				mediavalue = 1
			case "application/stix+json;version=2.1":
				mediavalue = 2
			case "application/stix+json;version=2.2":
				mediavalue = 3
			case "application/stix+json;version=2.3":
				mediavalue = 4
			}

			_, err2 := ds.DB.Exec(stmt2, obj.ID, mediavalue)

			if err2 != nil {
				return fmt.Errorf("database execution error inserting collection media type", err2)
			}
		}
	}
	return nil
}

// ----------------------------------------------------------------------
//
// Collection Table Private Functions and Methods
// getCollection
//
// ----------------------------------------------------------------------

/*
sqlGetCollections - This function will return an SQL statement that will return a
list of collections. A byte array is used instead of string
concatenation as it is the most efficient way to do string concatenation in Go.
*/
func sqlGetCollections(whichCollections string) (string, error) {
	tblCol := DB_TABLE_TAXII_COLLECTIONS
	tblColMedia := DB_TABLE_TAXII_COLLECTION_MEDIA_TYPE
	tblMediaTypes := DB_TABLE_TAXII_MEDIA_TYPES

	/*
		SELECT
			t_collections.date_added,
			t_collections.enabled,
			t_collections.hidden,
			t_collections.id,
			t_collections.title,
			t_collections.description,
			t_collections.can_read,
			t_collections.can_write,
			group_concat(t_media_types.media_type)
		FROM
			t_collections
		JOIN
			t_collection_media_type ON
			t_collections.id = t_collection_media_type.collection_id
		JOIN
			t_media_types ON
			t_collection_media_type.media_type_id = t_media_types.row_id
		GROUP BY
			t_collections.id
	*/
	var s bytes.Buffer
	s.WriteString("SELECT ")
	s.WriteString(tblCol)
	s.WriteString(".row_id, ")
	s.WriteString(tblCol)
	s.WriteString(".date_added, ")
	s.WriteString(tblCol)
	s.WriteString(".enabled, ")
	s.WriteString(tblCol)
	s.WriteString(".hidden, ")
	s.WriteString(tblCol)
	s.WriteString(".id, ")
	s.WriteString(tblCol)
	s.WriteString(".title, ")
	s.WriteString(tblCol)
	s.WriteString(".description, ")
	s.WriteString(tblCol)
	s.WriteString(".can_read, ")
	s.WriteString(tblCol)
	s.WriteString(".can_write, ")
	s.WriteString("group_concat(")
	s.WriteString(tblMediaTypes)
	s.WriteString(".media_type) ")

	s.WriteString("FROM ")
	s.WriteString(tblCol)

	s.WriteString(" JOIN ")
	s.WriteString(tblColMedia)
	s.WriteString(" ON ")
	s.WriteString(tblCol)
	s.WriteString(".id = ")
	s.WriteString(tblColMedia)
	s.WriteString(".collection_id ")

	s.WriteString("JOIN ")
	s.WriteString(tblMediaTypes)
	s.WriteString(" ON ")
	s.WriteString(tblColMedia)
	s.WriteString(".media_type_id = ")
	s.WriteString(tblMediaTypes)
	s.WriteString(".row_id ")

	if whichCollections == "all" {
		// do nothing
	} else if whichCollections == "allEnabled" {
		s.WriteString("WHERE ")
		s.WriteString(tblCol)
		s.WriteString(".enabled = 1 ")
	} else if whichCollections == "enabledVisible" {
		s.WriteString("WHERE ")
		s.WriteString(tblCol)
		s.WriteString(".enabled = 1 AND ")
		s.WriteString(tblCol)
		s.WriteString(".hidden = 0 ")
	}

	s.WriteString("GROUP BY ")
	s.WriteString(tblCol)
	s.WriteString(".id")

	return s.String(), nil
}

/*
getCollections - This method is called from either GetAllCollections(),
GetAllEnabledCollections(), or GetCollections() and will return all of the
collections that are asked for based on the method that called it.  The options
that can be passed in are:
    "all"
    "allEnabled"
    "enabledVisible"

The "all" option returns every collection, even those that are hidden or disabled.
"allEnabled" will return all enabled collections, even those that are hidden.
"enabledVisible" will return all collections that are both enabled and not
hidden (aka those that are visible).

Administration tools using the database will probably want to see all collections.
The HTTP Router MUX needs to know about all enabled collections, even those that
are hidden, so that it can start an HTTP router for it. The enabled and visible
list is what would be displayed to a client that is pulling a collections resource.
*/
func (ds *Store) getCollections(whichCollections string) (*collections.Collections, error) {
	ds.Logger.Traceln("TRACE getCollections(): Start")
	ds.Logger.Traceln("TRACE getCollections(): Which Collections", whichCollections)

	allCollections := collections.New()

	stmt, _ := sqlGetCollections(whichCollections)

	ds.Logger.Traceln("TRACE getCollections(): SQL Statement", stmt)

	// Query database for all the collections
	rows, err := ds.DB.Query(stmt)
	if err != nil {
		return nil, fmt.Errorf("database execution error getting collection: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var datastoreID, enabled, hidden, iCanRead, iCanWrite int
		var dateAdded, id, title, description, mediaType string
		if err := rows.Scan(&datastoreID, &dateAdded, &enabled, &hidden, &id, &title, &description, &iCanRead, &iCanWrite, &mediaType); err != nil {
			rows.Close()
			return nil, fmt.Errorf("database scan error getting collection: ", err)
		}

		// Add collection information to Collection object
		c, _ := allCollections.NewCollection()
		c.DatastoreID = datastoreID
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
		rows.Close()
		return nil, fmt.Errorf("database row error getting collection: ", err)
	}

	return allCollections, nil
}
