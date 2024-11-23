// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/freetaxii/libstix2/defs"
	"github.com/freetaxii/libstix2/objects/taxii/collections"
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

row_id      = This is the datastore ID used to reference this collection
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
// doesCollectionExistInTheCache
// addCollection
// getCollectionDatastoreID
// getCollections
//
// ----------------------------------------------------------------------

/*
doesCollectionExistInTheCache - This method will check to see if the collection
already exists in the cache. This is used by several methods that get data from
the datastore, to make sure the collection is a valid collection.
Called from: addToCollection(), getBundle(), getManifestData(),
*/
func (ds *Store) doesCollectionExistInTheCache(collectionUUID string) bool {
	if _, found := ds.Cache.Collections[collectionUUID]; found {
		return true
	}
	return false
}

/*
addCollection - This method will add a collection to the t_collections table in
the database and return the row_id (datastore ID) for the collection and an
error if there is one.
*/
func (ds *Store) addCollection(obj *collections.Collection) (int, error) {
	ds.Logger.Levelln("Function", "FUNC: addCollection start")

	// Create SQL Statement
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
	tblCol := DB_TABLE_TAXII_COLLECTIONS
	var sqlstmt bytes.Buffer
	sqlstmt.WriteString("INSERT INTO ")
	sqlstmt.WriteString(tblCol)
	sqlstmt.WriteString(" (date_added, id, title, description, can_read, can_write) ")
	sqlstmt.WriteString("values (?, ?, ?, ?, ?, ?)")
	stmt1 := sqlstmt.String()

	dateAdded := time.Now().UTC().Format(defs.TimeRFC3339Micro)

	// Make SQL Call
	val, err1 := ds.DB.Exec(stmt1,
		dateAdded,
		obj.ID,
		obj.Title,
		obj.Description,
		obj.CanRead,
		obj.CanWrite)

	if err1 != nil {
		ds.Logger.Levelln("Function", "FUNC: addCollection exited with an error,", err1)
		return 0, fmt.Errorf("database execution error inserting collection", err1)
	}

	// Get the row_id from the last insert and store in the cache. This comes
	// back from the database as an int64 so we need to convert back to an int.
	rowID, _ := val.LastInsertId()
	datastoreID := int(rowID)

	// Add the media types to the collection
	if obj.MediaTypes != nil {
		for _, media := range obj.MediaTypes {

			// TODO These should really be in the cache where we can look them up
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

			// Create SQL Statement
			/*
				INSERT INTO
					t_collection_media_type (
						collection_id,
						media_type_id
					)
					values (?, ?)
			*/
			tblColMedia := DB_TABLE_TAXII_COLLECTION_MEDIA_TYPE
			var sqlstmt bytes.Buffer
			sqlstmt.WriteString("INSERT INTO ")
			sqlstmt.WriteString(tblColMedia)
			sqlstmt.WriteString(" (collection_id, media_type_id) values (?, ?)")
			stmt2 := sqlstmt.String()

			// Make SQL Call
			_, err2 := ds.DB.Exec(stmt2, obj.ID, mediavalue)

			if err2 != nil {
				ds.Logger.Levelln("Function", "FUNC: addCollection exited with an error,", err2)
				return 0, fmt.Errorf("database execution error inserting collection media type", err2)
			}
		}
	}
	ds.Logger.Levelln("Function", "FUNC: addCollection end")
	return datastoreID, nil
}

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
getCollectionDatastoreID - This method takes in a collection UUID and returns
the collection datastore ID from the database if it is found. If it is not
found it will return an error.
*/
func (ds *Store) getCollectionDatastoreID(uuid string) (int, error) {
	ds.Logger.Levelln("Function", "FUNC: getCollectionDatastoreID start")
	ds.Logger.Debugln("DEBUG: Getting the datastore ID for collection", uuid)
	var datastoreID int

	// Create SQL Statement
	/*
		SELECT
			row_id
		FROM
			t_collections
		WHERE
			id = ?
	*/
	tblCol := DB_TABLE_TAXII_COLLECTIONS
	var sqlstmt bytes.Buffer
	sqlstmt.WriteString("SELECT row_id FROM ")
	sqlstmt.WriteString(tblCol)
	sqlstmt.WriteString(" WHERE id = ?")
	stmt := sqlstmt.String()

	// Make SQL Call
	err := ds.DB.QueryRow(stmt, uuid).Scan(&datastoreID)
	if err != nil {
		if err == sql.ErrNoRows {
			ds.Logger.Levelln("Function", "FUNC: getBaseObject exited with an error,", err)
			return 0, errors.New("collection not found")
		}
		ds.Logger.Levelln("Function", "FUNC: getBaseObject exited with an error,", err)
		return 0, fmt.Errorf("database execution error getting collection: ", err)
	}
	ds.Logger.Debugln("DEBUG: Datastore ID for collection", uuid, "is", datastoreID)
	ds.Logger.Levelln("Function", "FUNC: getCollectionDatastoreID end")
	return datastoreID, nil
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
	ds.Logger.Levelln("Function", "FUNC: getCollections start")
	ds.Logger.Debugln("DEBUG: Which Collections", whichCollections)

	allCollections := collections.New()

	stmt, _ := sqlGetCollections(whichCollections)

	ds.Logger.Traceln("TRACE getCollections(): SQL Statement", stmt)

	// Query database for all the collections
	rows, err := ds.DB.Query(stmt)
	if err != nil {
		ds.Logger.Levelln("Function", "FUNC: getCollections exited with an error,", err)
		return nil, fmt.Errorf("database execution error getting collection: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var datastoreID int
		var enabled, hidden, iCanRead, iCanWrite int
		var dateAdded, id, title, description, mediaType string
		if err := rows.Scan(&datastoreID, &dateAdded, &enabled, &hidden, &id, &title, &description, &iCanRead, &iCanWrite, &mediaType); err != nil {
			rows.Close()
			ds.Logger.Levelln("Function", "FUNC: getCollections exited with an error,", err)
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
		ds.Logger.Levelln("Function", "FUNC: getCollections exited with an error,", err)
		return nil, fmt.Errorf("database row error getting collection: ", err)
	}

	ds.Logger.Levelln("Function", "FUNC: getCollections end")
	return allCollections, nil
}
