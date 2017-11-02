// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"log"
)

// createAttackPatternTable - This method will create the actual table
func (ds *Sqlite3DatastoreType) createTable(name, properties string) {
	var stmt = `CREATE TABLE IF NOT EXISTS "` + name + `" (` + properties + `)`
	_, err := ds.DB.Exec(stmt)

	if err != nil {
		log.Println("ERROR: The", name, "table could not be created due to error:", err)
	}
}
