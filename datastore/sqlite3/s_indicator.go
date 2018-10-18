// Copyright 2015-2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"

	"github.com/freetaxii/libstix2/objects/indicator"
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
	return baseDBProperties() + `
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
	return baseDBProperties() + `
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
func (ds *Store) addIndicator(obj *indicator.Indicator) error {
	ds.Logger.Levelln("Function", "FUNC: addIndicator Start")

	datastoreID, err := ds.addBaseObject(&obj.CommonObjectProperties)
	if err != nil {
		ds.Logger.Levelln("Function", "FUNC: addIndicator End with error")
		return err
	}
	ds.Logger.Debugln("DEBUG: Adding Indicator to datastore with database ID", datastoreID)

	// Create SQL Statement
	/*
		INSERT INTO
			s_indicator (
				"datastore_id",
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
	sqlstmt.WriteString(" (datastore_id, name, description, pattern, valid_from, valid_until) ")
	sqlstmt.WriteString("values (?, ?, ?, ?, ?, ?)")
	stmt := sqlstmt.String()

	// Make SQL Call
	_, err1 := ds.DB.Exec(stmt,
		datastoreID,
		obj.Name,
		obj.Description,
		obj.Pattern,
		obj.ValidFrom,
		obj.ValidUntil)

	// TODO if there is an error, we probably need to back out all of the INSERTS
	if err1 != nil {
		ds.Logger.Levelln("Function", "FUNC: addIndicator End with error")
		return fmt.Errorf("database execution error inserting indicator: ", err)
	}

	// ----------------------------------------------------------------------
	// Add Indicator Types
	// ----------------------------------------------------------------------
	if obj.IndicatorTypes != nil {
		for _, itype := range obj.IndicatorTypes {
			err := ds.addIndicatorType(datastoreID, itype)
			if err != nil {
				ds.Logger.Levelln("Function", "FUNC: addIndicator End with error")
				return err
			}
		}
	}

	// ----------------------------------------------------------------------
	// Add Kill Chains
	// ----------------------------------------------------------------------
	if obj.KillChainPhases != nil {
		for _, v := range obj.KillChainPhases {
			err := ds.addKillChainPhase(datastoreID, &v)
			if err != nil {
				ds.Logger.Levelln("Function", "FUNC: addIndicator End with error")
				return err
			}
		}
	}
	ds.Logger.Levelln("Function", "FUNC: addIndicator End")
	return nil
}

/*
getIndicator - This method will get a specific indicator from the database based
on the STIX ID and version.
*/
func (ds *Store) getIndicator(stixid, version string) (*indicator.Indicator, error) {
	ds.Logger.Levelln("Function", "FUNC: getIndicator Start")
	var i indicator.Indicator

	// Get Base Object - this will give us the datastoreID
	// Then copy base object data in to Indicator object
	baseObject, errBase := ds.getBaseObject(stixid, version)
	if errBase != nil {
		ds.Logger.Levelln("Function", "FUNC: getIndicator End with error")
		return nil, errBase
	}
	i.CommonObjectProperties = *baseObject

	// Create SQL Statement
	/*
		SELECT
			s_indicator.name,
			s_indicator.description,
			s_indicator.pattern,
			s_indicator.valid_from,
			s_indicator.valid_until,
			group_concat(s_indicator_types.indicator_type)
		FROM
			s_indicator
		JOIN
			s_indicator_types ON
			s_indicator.datastore_id = s_indicator_types.datastore_id
		WHERE
			s_indicator.datastore_id = ?
	*/
	tblInd := DB_TABLE_STIX_INDICATOR
	tblIndType := DB_TABLE_STIX_INDICATOR_TYPES
	var sqlstmt bytes.Buffer
	sqlstmt.WriteString("SELECT ")
	sqlstmt.WriteString(tblInd)
	sqlstmt.WriteString(".name, ")
	sqlstmt.WriteString(tblInd)
	sqlstmt.WriteString(".description, ")
	sqlstmt.WriteString(tblInd)
	sqlstmt.WriteString(".pattern, ")
	sqlstmt.WriteString(tblInd)
	sqlstmt.WriteString(".valid_from, ")
	sqlstmt.WriteString(tblInd)
	sqlstmt.WriteString(".valid_until, ")
	sqlstmt.WriteString("group_concat(")
	sqlstmt.WriteString(tblIndType)
	sqlstmt.WriteString(".indicator_type) ")
	sqlstmt.WriteString("FROM ")
	sqlstmt.WriteString(tblInd)
	sqlstmt.WriteString(" JOIN ")
	sqlstmt.WriteString(tblIndType)
	sqlstmt.WriteString(" ON ")
	sqlstmt.WriteString(tblInd)
	sqlstmt.WriteString(".datastore_id = ")
	sqlstmt.WriteString(tblIndType)
	sqlstmt.WriteString(".datastore_id ")
	sqlstmt.WriteString("WHERE ")
	sqlstmt.WriteString(tblInd)
	sqlstmt.WriteString(".datastore_id = ?")
	stmt := sqlstmt.String()

	// Make SQL Call
	var name, description, pattern, validFrom, validUntil, indTypes string
	err := ds.DB.QueryRow(stmt, i.DatastoreID).Scan(&name, &description, &pattern, &validFrom, &validUntil, &indTypes)
	if err != nil {
		if err == sql.ErrNoRows {
			ds.Logger.Levelln("Function", "FUNC: getIndicator End with error")
			return nil, errors.New("no indicator record found")
		}
		ds.Logger.Levelln("Function", "FUNC: getIndicator End with error")
		return nil, fmt.Errorf("database execution error getting indicator: ", err)
	}
	i.SetName(name)
	i.SetDescription(description)
	i.AddType(indTypes)
	i.SetPattern(pattern)
	i.SetValidFrom(validFrom)
	i.SetValidUntil(validUntil)

	killChainPhases, errkc := ds.getKillChainPhases(i.DatastoreID)
	if errkc != nil {
		ds.Logger.Levelln("Function", "FUNC: getIndicator End with error")
		return nil, errkc
	}
	i.KillChainPhasesProperty = *killChainPhases

	ds.Logger.Levelln("Function", "FUNC: getIndicator End")
	return &i, nil
}

/*
addIndicatorTypes - This method will add all of the indicator types to the
database for a specific indicator based on its database ID.
*/
func (ds *Store) addIndicatorType(datastoreID int, itype string) error {
	ds.Logger.Levelln("Function", "FUNC: addIndicatorType Start")

	// Create SQL Statement
	/*
		INSERT INTO
			s_indicator_types (
				"datastore_id",
				"indicator_type"
			)
			values (?, ?)
	*/
	tblIndTypes := DB_TABLE_STIX_INDICATOR_TYPES
	var sqlstmt bytes.Buffer
	sqlstmt.WriteString("INSERT INTO ")
	sqlstmt.WriteString(tblIndTypes)
	sqlstmt.WriteString(" (datastore_id, indicator_type) values (?, ?)")
	stmt := sqlstmt.String()

	// Make SQL Call
	_, err := ds.DB.Exec(stmt, datastoreID, itype)
	if err != nil {
		ds.Logger.Levelln("Function", "FUNC: addIndicatorType End with error")
		return fmt.Errorf("database execution error inserting indicator type: ", err)
	}

	ds.Logger.Levelln("Function", "FUNC: addIndicatorType End")
	return nil
}
