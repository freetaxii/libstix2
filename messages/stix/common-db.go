// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package stix

import (
	"database/sql"
	"log"
)

// ----------------------------------------------------------------------
// Public Database Methods - CommonProperties
// ----------------------------------------------------------------------

// This function will take in the following values that are passed in as
// references and write them to the correct database table.
//
// db      = Pointer to the open database connection
// pid     = The auto_id record from the parent to be put in the parent_id field
// date    = The date added value so you know when this record was added
// id      = The STIX ID
// ver     = The STIX object version
// labels  = The list of labels
func AddLabelsToDatabase(db *sql.DB, pid *int64, date, id *string, ver *int, labels *[]string) {
	for _, label := range *labels {
		var stmt = `
			INSERT INTO 'common_labels'
			(parent_id, date, id, version, labels)
			values (?, ?, ?, ?, ?)`
		_, err := db.Exec(stmt, pid, date, id, ver, label)

		if err != nil {
			log.Printf("M: Unable to insert record due to error %v", err)
		}
	}
}

// This function will take in the following values that are passed in as
// references and write them to the correct database table.
//
// db        = Pointer to the open database connection
// pid       = The auto_id record from the parent to be put in the parent_id field
// date      = The date added value so you know when this record was added
// id        = The STIX ID
// ver       = The STIX object version
// markings  = The list of markings
func AddObjectMarkingsToDatabase(db *sql.DB, pid *int64, date, id *string, ver *int, markings *[]string) {
	for _, markingref := range *markings {
		var stmt = `
			INSERT INTO 'common_object_marking_refs'
			(parent_id, date, id, version, object_marking_refs)
			values (?, ?, ?, ?, ?)`
		_, err := db.Exec(stmt, pid, date, id, ver, markingref)

		if err != nil {
			log.Printf("M: Unable to insert record due to error %v", err)
		}
	}
}
