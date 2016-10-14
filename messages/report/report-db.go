// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package report

import (
	"database/sql"
	"github.com/freetaxii/libstix2/messages/defs"
	"github.com/freetaxii/libstix2/messages/stix"
	"log"
	"time"
)

// ----------------------------------------------------------------------
// Public Database Methods - ReportType
// ----------------------------------------------------------------------

func (this *ReportType) AddToDatabase(db *sql.DB) {
	dateAdded := time.Now().UTC().Format(defs.TIME_RFC_3339)

	var stmt1 = `
		INSERT INTO 'reports'
		(date, id, version, created_by_ref, created, modified, version, revoked, name, description, published)
		values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	res, err := db.Exec(stmt1,
		dateAdded,
		this.Id,
		this.Version,
		this.Created_by_ref,
		this.Created,
		this.Modified,
		this.Version,
		this.Revoked,
		this.Name,
		this.Description,
		this.Published)

	parentId, _ := res.LastInsertId()

	if err != nil {
		log.Printf("M: Unable to insert record due to error %v", err)
	}

	if this.Labels != nil {
		stix.AddLabelsToDatabase(db, &parentId, &dateAdded, &this.Id, &this.Version, &this.Labels)
	}

	if this.Object_refs != nil {
		this.addReportObjectRefsToDatabase(db, &parentId, &dateAdded)
	}
}

// This function will take in the following values that are passed in as
// references and write them to the correct database table.
//
// db      = Pointer to the open database connection
// pid     = The auto_id record from the parent to be put in the parent_id field
// date    = The date added value so you know when this record was added
func (this *ReportType) addReportObjectRefsToDatabase(db *sql.DB, pid *int64, date *string) {
	for _, object := range this.Object_refs {
		var stmt = `
			INSERT INTO 'reports_object_refs'
			(parent_id, date, id, version, object_refs)
			values (?, ?, ?, ?, ?)`
		_, err := db.Exec(stmt, pid, date, this.Id, this.Version, object)

		if err != nil {
			log.Printf("M: Unable to insert record due to error %v", err)
		}
	}
}
