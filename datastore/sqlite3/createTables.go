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

func (ds *Sqlite3DatastoreType) createTable(name, properties string) {
	var stmt = `CREATE TABLE IF NOT EXISTS "` + name + `" (` + properties + `)`
	_, err := ds.DB.Exec(stmt)

	if err != nil {
		log.Println("ERROR: The", name, "table could not be created due to error:", err)
	}
}

// createTAXIITable - This method will create the actual table
func (ds *Sqlite3DatastoreType) createTAXIITable(name, properties string) {
	ds.createTable(name, properties)
	ds.createTAXIIIndexes(name)
}

func (ds *Sqlite3DatastoreType) createSTIXTable(name, properties string) {
	ds.createTable(name, properties)
	ds.createSTIXIndexes(name)
}

func (ds *Sqlite3DatastoreType) createSTIXIndexes(name string) {
	var stmt string

	if name == datastore.DB_TABLE_STIX_BASE_OBJECT {
		stmt = `CREATE INDEX "idx_` + name + `" ON ` + name + ` ("object_id" COLLATE BINARY ASC, "id" COLLATE BINARY ASC)`
	} else if name == datastore.DB_TABLE_TAXII_COLLECTION_CONTENT {
		stmt = `CREATE INDEX "idx_` + name + `" ON ` + name + ` ("collection_id" COLLATE BINARY ASC, "stix_id" COLLATE BINARY ASC)`
	} else {
		stmt = `CREATE INDEX "idx_` + name + `" ON ` + name + ` ("object_id" COLLATE BINARY ASC)`
	}

	_, err := ds.DB.Exec(stmt)

	if err != nil {
		log.Println("ERROR: The indexes for table", name, "could not be created due to error:", err)
	}
}

func (ds *Sqlite3DatastoreType) createTAXIIIndexes(name string) {
	var stmt string

	if name == datastore.DB_TABLE_TAXII_COLLECTION_CONTENT {
		stmt = `CREATE INDEX "idx_` + name + `" ON ` + name + ` ("collection_id" COLLATE BINARY ASC, "stix_id" COLLATE BINARY ASC)`
	}

	if stmt != "" {
		_, err := ds.DB.Exec(stmt)

		if err != nil {
			log.Println("ERROR: The indexes for table", name, "could not be created due to error:", err)
		}
	}
}
