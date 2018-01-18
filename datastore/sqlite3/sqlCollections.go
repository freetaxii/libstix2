// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"bytes"
	"github.com/freetaxii/libstix2/datastore"
	"log"
)

// ----------------------------------------------------------------------
//
// Private Methods
//
// ----------------------------------------------------------------------

/*
sqlAddCollection - This method will return an SQL statement that will insert
a new collection in to the t_collections table in the database.
*/
func (ds *Sqlite3DatastoreType) sqlAddCollection() (string, error) {
	tblCol := datastore.DB_TABLE_TAXII_COLLECTIONS

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
	s.WriteString(" (\"date_added\", \"id\", \"title\", \"description\", \"can_read\", \"can_write\") values (?, ?, ?, ?, ?, ?) ")

	if ds.LogLevel >= 5 {
		log.Println("DEBUG: Returning SQL statement:", s.String())
	}

	return s.String(), nil
}

/*
sqlAddCollectionMediaType - This method will return an SQL statement that will
insert a media type for a given collection.
*/
func (ds *Sqlite3DatastoreType) sqlAddCollectionMediaType() (string, error) {
	tblColMedia := datastore.DB_TABLE_TAXII_COLLECTION_MEDIA_TYPE

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
	s.WriteString(" (\"collection_id\", \"media_type_id\") values (?, ?) ")

	if ds.LogLevel >= 5 {
		log.Println("DEBUG: Returning SQL statement:", s.String())
	}

	return s.String(), nil
}

/*
sqlAllCollections - This method will return an SQL statement that will return a
list of collections. A byte array is used instead of sting
concatenation as it is the most efficient way to do string concatenation in Go.
*/
func (ds *Sqlite3DatastoreType) sqlAllCollections(whichCollections string) (string, error) {
	tblCol := datastore.DB_TABLE_TAXII_COLLECTIONS
	tblColMedia := datastore.DB_TABLE_TAXII_COLLECTION_MEDIA_TYPE
	tblMediaTypes := datastore.DB_TABLE_TAXII_MEDIA_TYPES

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

	if ds.LogLevel >= 5 {
		log.Println("DEBUG: Returning SQL statement:", s.String())
	}

	return s.String(), nil
}
