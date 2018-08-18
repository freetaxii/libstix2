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
	// indicator_types
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

// ----------------------------------------------------------------------
//
// Private Methods - Indicator Table
// addIndicator()
// getIndicator()
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

	// Create SQL Statement
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
	tblInd := DB_TABLE_STIX_INDICATOR
	var sqlstmt bytes.Buffer
	sqlstmt.WriteString("INSERT INTO ")
	sqlstmt.WriteString(tblInd)
	sqlstmt.WriteString(" (object_id, name, description, pattern, valid_from, valid_until) ")
	sqlstmt.WriteString("values (?, ?, ?, ?, ?, ?)")
	stmt := sqlstmt.String()

	// Make SQL Call
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

	// ----------------------------------------------------------------------
	// Add Indicator Types
	// ----------------------------------------------------------------------
	if obj.IndicatorTypes != nil {
		for _, itype := range obj.IndicatorTypes {
			err := ds.addIndicatorType(objectID, itype)
			if err != nil {
				return err
			}
		}
	}

	// ----------------------------------------------------------------------
	// Add Kill Chains
	// ----------------------------------------------------------------------
	if obj.KillChainPhases != nil {
		err := ds.addKillChainPhases(objectID, &obj.KillChainPhasesProperty)
		if err != nil {
			return err
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

/*
addIndicatorTypes - This method will add all of the indicator types to the
database for a specific indicator based on its object ID.
*/
func (ds *Datastore) addIndicatorType(objectID int, itype string) error {

	// Create SQL Statement
	/*
		INSERT INTO
			s_indicator_types (
				"object_id",
				"indicator_type"
			)
			values (?, ?)
	*/
	tblIndTypes := DB_TABLE_STIX_INDICATOR_TYPES
	var sqlstmt bytes.Buffer
	sqlstmt.WriteString("INSERT INTO ")
	sqlstmt.WriteString(tblIndTypes)
	sqlstmt.WriteString(" (object_id, indicator_type) values (?, ?)")
	stmt := sqlstmt.String()

	// Make SQL Call
	_, err := ds.DB.Exec(stmt, objectID, itype)
	if err != nil {
		return fmt.Errorf("database execution error inserting indicator type: ", err)
	}

	return nil
}
