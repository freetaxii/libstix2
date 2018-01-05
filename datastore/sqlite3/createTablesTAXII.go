// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"github.com/freetaxii/libstix2/datastore"
	"log"
)

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

// CreateAllTAXIITables - This method will create all of the tables needed to store
// STIX content in the database.
func (ds *Sqlite3DatastoreType) CreateAllTAXIITables() {
	ds.createTAXIITable(datastore.DB_TABLE_TAXII_COLLECTION_DATA, collectionData())
	ds.createTAXIITable(datastore.DB_TABLE_TAXII_COLLECTIONS, collections())
	ds.createTAXIITable(datastore.DB_TABLE_TAXII_COLLECTION_MEDIA_TYPE, collectionMediaType())
	ds.createTAXIITable(datastore.DB_TABLE_TAXII_MEDIA_TYPES, mediaTypes())
	ds.createTAXIIIndexes(datastore.DB_TABLE_TAXII_COLLECTION_DATA)
	ds.insertMediaTypes(datastore.DB_TABLE_TAXII_MEDIA_TYPES)
}

// ----------------------------------------------------------------------
// Private Methods
// ----------------------------------------------------------------------

func (ds *Sqlite3DatastoreType) createTAXIITable(name, properties string) {
	var stmt = `CREATE TABLE IF NOT EXISTS "` + name + `" (` + properties + `)`
	_, err := ds.DB.Exec(stmt)

	if err != nil {
		log.Println("ERROR: The", name, "table could not be created due to error:", err)
	}
}

func (ds *Sqlite3DatastoreType) createTAXIIIndexes(name string) {
	var stmt string

	if name == datastore.DB_TABLE_TAXII_COLLECTION_DATA {
		stmt = `CREATE INDEX "idx_` + name + `" ON ` + name + ` ("collection_id" COLLATE BINARY ASC, "stix_id" COLLATE BINARY ASC)`
	}

	if stmt != "" {
		_, err := ds.DB.Exec(stmt)

		if err != nil {
			log.Println("ERROR: The indexes for table", name, "could not be created due to error:", err)
		}
	}
}

func (ds *Sqlite3DatastoreType) insertMediaTypes(name string) {
	var stmt = `INSERT INTO "` + name + `" (media_type) values (?)`

	var err error
	_, err = ds.DB.Exec(stmt, "application/vnd.oasis.stix+json")

	if err != nil {
		log.Println("ERROR: The media type item could not be inserted in to the", name, "table")
	}
}

// ----------------------------------------------------------------------
//
// Each of these functions returns a list of fields that are used for creating
// a database tables.
//
// ----------------------------------------------------------------------

/*
collectionData - This method will return the properties that make up the
collection content table

date_added    = The date that this object was added to the collection
collection_id = The collection ID that this object is tied to
stix_id       = The STIX ID for the object that is being mapped to a collection.
  We do not use the object_id here or the row_id as that would point to a
  specific version and we need to be able to find all versions of an object.
  and if we used row_id for example, it would require two queries, the first
  to get the SITX ID and then the second to get all objects with that STIX ID.
*/
func collectionData() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
	"date_added" TEXT NOT NULL,
 	"collection_id" TEXT NOT NULL,
 	"stix_id" TEXT NOT NULL
 	`
}

/*
collections - This method will return the properties that make up the collection
table

date_added  = The date that this collection was added to the system
enabled     = Is this collection currently enabled
hidden      = Is this collection currently hidden for this directory listing
id 		    = The collection ID, a UUIDv4 value
title 	    = The title of this collection
description = A long description about this collection
can_read    = A boolean flag that indicates if one can read from this collection
can_write   = A boolean flag that indicates if one can write to this collection
*/
func collections() string {
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
collectionMediaType  - This method will return the properties that make up the
collection media type table

collection_id = The collection ID, a UUIDv4 value
media_type_id = The media types supported on this collection
*/
func collectionMediaType() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
	"collection_id" TEXT NOT NULL,
	"media_type_id" INTEGER NOT NULL
	`
}

/*
mediaTypes  - This method will return the properties that make up the media
types table

media_type    = A media type
*/
func mediaTypes() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
	"media_type" TEXT NOT NULL
	`
}
