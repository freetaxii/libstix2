// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"github.com/freetaxii/libstix2/datastore"
)

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

// CreateAllTAXIITables - This method will create all of the tables needed to store
// STIX content in the database.
func (ds *Sqlite3DatastoreType) CreateAllTAXIITables() {
	ds.createTable(datastore.DB_TABLE_TAXII_COLLECTION_CONTENT, ds.collectionContent())
	ds.createTable(datastore.DB_TABLE_TAXII_COLLECTION, ds.collection())
	ds.createTable(datastore.DB_TABLE_TAXII_COLLECTION_MEDIA_TYPE, ds.collectionMediaType())
}

// ----------------------------------------------------------------------
// Private Methods
// ----------------------------------------------------------------------

// collectionContent - This method will return the properties for the collectionContent table
// collection_id = The collection ID that this object is tied to
// stix_id       = The STIX ID for the object that is being mapped to a collection.
//   We do not use the object_id here or the row_id as that would point to a
//   specific version and we need to be able to find all versions of an object
// date_added    = The date that this object was added to the collection
func (ds *Sqlite3DatastoreType) collectionContent() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
	"date_added" TEXT NOT NULL,
 	"collection_id" TEXT NOT NULL,
 	"stix_id" TEXT NOT NULL`
}

// collection  - This method will return the properties for the collections table
// date_added = The date that this collection was added to the system
func (ds *Sqlite3DatastoreType) collection() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
	"date_added" TEXT NOT NULL,
	"enabled" INTEGER(1,0) NOT NULL DEFAULT 1,
	"id" TEXT NOT NULL,
	"title" TEXT NOT NULL,
	"description" TEXT,
	"can_read" INTEGER(1,0) NOT NULL DEFAULT 0,
	"can_write" INTEGER(1,0) NOT NULL DEFAULT 0
	`
}

// collectionMediaType  - This method will return the properties for the collection media type table
func (ds *Sqlite3DatastoreType) collectionMediaType() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
	"collection_id" TEXT NOT NULL,
	"media_type" TEXT NOT NULL
	`
}
