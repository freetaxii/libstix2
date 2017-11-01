// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"github.com/freetaxii/libstix2/defs"
	"log"
)

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

// CreateAllTables - This method will create all of the tables needed to store
// STIX content in the database.
func (ds *Sqlite3DatastoreType) CreateAllTables() {
	ds.createTable(defs.DB_TABLE_TAXII_COLLECTION_CONTENT, ds.collectionContent())
	ds.createTable(defs.DB_TABLE_TAXII_COLLECTION, ds.collection())
	ds.createTable(defs.DB_TABLE_TAXII_COLLECTION_MEDIA_TYPE, ds.collectionMediaType())
}

// ----------------------------------------------------------------------
// Private Methods
// ----------------------------------------------------------------------

// createAttackPatternTable - This method will create the actual table
func (ds *Sqlite3DatastoreType) createTable(name, properties string) {
	var stmt = `CREATE TABLE IF NOT EXISTS "` + name + `" (` + properties + `)`
	_, err := ds.DB.Exec(stmt)

	if err != nil {
		log.Println("ERROR: The", name, "table could not be created due to error:", err)
	}
}

// baseProperties - This method will return the base properties for all objects
// row_id    = This is a database tracking number
func (ds *Sqlite3DatastoreType) baseProperties() string {
	return `
	"row_id" INTEGER PRIMARY KEY,`
}

// collectionContent - This method will return the properties for the collectionContent table
// collection_id = The collection ID that this object is tied to
// stix_id       = The STIX ID for the object that is being mapped to a collection.
//   We do not use the object_id here or the row_id as that would point to a
//   specific version and we need to be able to find all versions of an object
// date_added    = The date that this object was added to the collection
func (ds *Sqlite3DatastoreType) collectionContent() string {
	return ds.baseProperties() + `
	"date_added" TEXT NOT NULL,
 	"collection_id" TEXT NOT NULL,
 	"stix_id" TEXT NOT NULL`
}

// collection  - This method will return the properties for the collections table
// date_added = The date that this collection was added to the system
func (ds *Sqlite3DatastoreType) collection() string {
	return ds.baseProperties() + `
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
	return ds.baseProperties() + `
	"collection_id" TEXT NOT NULL,
	"media_type" TEXT NOT NULL
	`
}
