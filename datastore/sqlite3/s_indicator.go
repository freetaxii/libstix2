// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"bytes"
	"fmt"
	"github.com/freetaxii/libstix2/datastore"
	"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
//
// Indicator Table Private Functions and Methods
//
// ----------------------------------------------------------------------

/*
indicatorProperties  - This method will return the properties for indicator SDOs
*/
func indicatorProperties() string {
	return baseProperties() + `
	"name" TEXT,
	"description" TEXT,
	"pattern" TEXT NOT NULL,
	"valid_from" TEXT NOT NULL,
	"valid_until" TEXT
	`
	// kill_chain_phases
}

/*
sqlAddIndicatorObject - This function will return an SQL statement that will add
an indicator to the database.
*/
func sqlAddIndicator() (string, error) {
	tblInd := datastore.DB_TABLE_STIX_INDICATOR

	/*
		INSERT INTO
			s_indicator (
				"object_id",
				"name",
				"description",
				"pattern",
				"valid_from",
				"valid_until"
			)
			values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	*/

	var s bytes.Buffer
	s.WriteString("INSERT INTO ")
	s.WriteString(tblInd)
	s.WriteString(" (")
	s.WriteString("\"object_id\", ")
	s.WriteString("\"name\", ")
	s.WriteString("\"description\", ")
	s.WriteString("\"pattern\", ")
	s.WriteString("\"valid_from\", ")
	s.WriteString("\"valid_until\") ")
	s.WriteString("values (?, ?, ?, ?, ?, ?)")

	return s.String(), nil
}

/*
addIndicator - This method will add an indicator to the database.
*/
func (ds *Sqlite3DatastoreType) addIndicator(obj *objects.IndicatorType) error {

	objectID, err := ds.addBaseObject(&obj.CommonObjectPropertiesType)
	if err != nil {
		return err
	}

	stmt, _ := sqlAddIndicator()
	_, err1 := ds.DB.Exec(stmt,
		objectID,
		obj.Name,
		obj.Description,
		obj.Pattern,
		obj.ValidFrom,
		obj.ValidUntil)

	// TODO if there is an error, we probably need to back out all of the INSERTS
	if err1 != nil {
		return fmt.Errorf("database error inserting indicator: ", err)
	}

	if obj.KillChainPhases != nil {
		err2 := ds.addKillChainPhases(objectID, &obj.KillChainPhasesPropertyType)

		if err2 != nil {
			return err2
		}
	}
	return nil
}
