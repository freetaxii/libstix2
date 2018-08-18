// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"

	"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
//
// Private Functions - Indicator Table
// Table property names and SQL statements
//
// ----------------------------------------------------------------------

/*
indicatorProperties  - This function will return the properties for the
indicator SDO table
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
indicatorTypeProperties - This function will return the properties for the
indicator types table
*/
func indicatorTypesProperties() string {
	return baseProperties() + `
	"indicator_type" TEXT NOT NULL
	`
}

/*
sqlAddIndicator - This function will return an SQL statement that will add an
indicator to the database.
*/
func sqlAddIndicator() (string, error) {
	tblInd := DB_TABLE_STIX_INDICATOR

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

// ----------------------------------------------------------------------
//
// Private Methods - Indicator Table
//
// ----------------------------------------------------------------------

/*
addIndicator - This method will add an indicator to the database.
*/
func (ds *Datastore) addIndicator(obj *objects.Indicator) error {

	objectID, err := ds.addBaseObject(&obj.CommonObjectProperties)
	if err != nil {
		return err
	}

	ds.Logger.Debugln("DEBUG: Adding Indicator to datastore with object ID", objectID)

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
		return fmt.Errorf("database execution error inserting indicator: ", err)
	}

	if obj.KillChainPhases != nil {
		err2 := ds.addKillChainPhases(objectID, &obj.KillChainPhasesProperty)

		if err2 != nil {
			return err2
		}
	}
	return nil
}

/*
getIndicator - This method will get a specific indicator from the database based
on the STIX ID and version.
*/
func (ds *Datastore) getIndicator(stixid, version string) (*objects.Indicator, error) {
	var i objects.Indicator

	// Get Base Object - this will give us the objectID
	// Then copy base object data in to Indicator object
	baseObject, errBase := ds.getBaseObject(stixid, version)
	if errBase != nil {
		return nil, errBase
	}
	i.CommonObjectProperties = *baseObject

	// Create SQL Statement
	/*
		SELECT
			name,
			description,
			pattern,
			valid_from,
			valid_until
		FROM
			s_indicator
		WHERE
			object_id = ?
	*/
	tblInd := DB_TABLE_STIX_INDICATOR
	var sqlstmt bytes.Buffer
	sqlstmt.WriteString("SELECT ")
	sqlstmt.WriteString("name, description, pattern, valid_from, valid_until ")
	sqlstmt.WriteString("FROM ")
	sqlstmt.WriteString(tblInd)
	sqlstmt.WriteString(" WHERE object_id = ?")
	stmt := sqlstmt.String()

	// Make SQL Call
	var description, pattern, validFrom, validUntil string
	err := ds.DB.QueryRow(stmt, i.ObjectID).Scan(&i.Name, &description, &pattern, &validFrom, &validUntil)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no indicator record found")
		}
		return nil, fmt.Errorf("database execution error getting indicator: ", err)
	}
	//i.SetName(name)
	i.SetDescription(description)
	i.SetPattern(pattern)
	i.SetValidFrom(validFrom)
	i.SetValidUntil(validUntil)

	return &i, nil
}
