// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"bytes"
	"fmt"

	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
//
// Common Tables Private Functions and Methods
//
// ----------------------------------------------------------------------

// ----------------------------------------------------------------------
//
// Aliases Table
//
// ----------------------------------------------------------------------

/*
commonAliasesProperties - This method will return the properties for aliases
Used by:
  campaign
  intrusion set
  threat actor
*/
func commonAliasesProperties() string {
	return baseDBProperties() + `
	"aliases" TEXT NOT NULL
	`
}

// ----------------------------------------------------------------------
//
// Authors Table
//
// ----------------------------------------------------------------------

/*
commonAuthorsProperties - This method will return the properties for common authors
Used by:
  note
  opinion
*/
func commonAuthorsProperties() string {
	return baseDBProperties() + `
	"authors" TEXT NOT NULL
	`
}

// ----------------------------------------------------------------------
//
// Goals Table
//
// ----------------------------------------------------------------------

/*
commonGoalsProperties  - This method will return the properties for goals
Used by:
  intrusion set
  threat actor
*/
func commonGoalsProperties() string {
	return baseDBProperties() + `
	"goals" TEXT NOT NULL
	`
}

// ----------------------------------------------------------------------
//
// Hashes Table
//
// ----------------------------------------------------------------------

/*
commonHashesProperties - This method will return the properties for hashes
Used by:
  external references
TODO need find a way to link this back to an actual external reference instance
maybe this should be called external references hashes.  Otherwise  how will you
know which object in the database it is tied to.
*/
func commonHashesProperties() string {
	return baseDBProperties() + `
	"hash" TEXT NOT NULL,
	"value" TEXT NOT NULL
	`
}

// ----------------------------------------------------------------------
//
// Kill Chain Phases Table
//
// ----------------------------------------------------------------------

/*
commonKillChainPhasesProperties - This method will return the properties for kill chain phases
Used by:
  attack pattern
  indicator
  malware
  tool
*/
func commonKillChainPhasesProperties() string {
	return baseDBProperties() + `
	"kill_chain_name" TEXT NOT NULL,
	"phase_name" TEXT NOT NULL
	`
}

/*
addKillChainPhase - This method will add a kill chain phase for a given object
to the database.
*/
func (ds *Store) addKillChainPhase(objectID int, obj *properties.KillChainPhase) error {

	// Create SQL Statement
	/*
		INSERT INTO
			s_kill_chain_phases (
				"datastore_id",
				"kill_chain_name",
				"phase_name"
			)
			values (?, ?)
	*/
	tblKillChain := DB_TABLE_STIX_KILL_CHAIN_PHASES
	var sqlstmt bytes.Buffer
	sqlstmt.WriteString("INSERT INTO ")
	sqlstmt.WriteString(tblKillChain)
	sqlstmt.WriteString(" (datastore_id, kill_chain_name, phase_name) values (?, ?, ?)")
	stmt := sqlstmt.String()

	// Make SQL Call
	_, err := ds.DB.Exec(stmt, objectID, obj.KillChainName, obj.PhaseName)
	if err != nil {
		return fmt.Errorf("database execution error inserting kill chain phase: ", err)
	}

	return nil
}

/*
getKillChainPhases - This method will get the kill chain phases for a given
object ID.
*/
func (ds *Store) getKillChainPhases(objectID int) (*properties.KillChainPhasesProperty, error) {
	var kcPhases properties.KillChainPhasesProperty

	// Create SQL Statement
	/*
		SELECT
			kill_chain_name,
			phase_name
		FROM
			s_kill_chain_phases
		WHERE
			datastore_id = ?
	*/
	tblKCP := DB_TABLE_STIX_KILL_CHAIN_PHASES
	var sqlstmt bytes.Buffer
	sqlstmt.WriteString("SELECT ")
	sqlstmt.WriteString("kill_chain_name, phase_name ")
	sqlstmt.WriteString("FROM ")
	sqlstmt.WriteString(tblKCP)
	sqlstmt.WriteString(" WHERE datastore_id = ?")
	stmt := sqlstmt.String()

	// Make SQL Call
	rows, err := ds.DB.Query(stmt, objectID)
	if err != nil {
		return nil, fmt.Errorf("database execution error getting kill chain phases: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var name, phase string
		p, _ := kcPhases.NewKillChainPhase()

		if err := rows.Scan(&name, &phase); err != nil {
			rows.Close()
			return nil, fmt.Errorf("database scan error getting kill chain phases: ", err)
		}
		p.SetName(name)
		p.SetPhase(phase)
	}

	// Errors can cause the rows.Next() to exit prematurely, if this happens lets
	// check for the error and handle it.
	if err := rows.Err(); err != nil {
		rows.Close()
		return nil, fmt.Errorf("database rows error getting kill chain phases: ", err)
	}

	return &kcPhases, nil
}

// ----------------------------------------------------------------------
//
// Object Refs Table
//
// ----------------------------------------------------------------------

/*
commonObjectRefsProperties - This method will return the properties for object refs
Used by:
  note
  opinion
  report
*/
func commonObjectRefsProperties() string {
	return baseDBProperties() + `
	"object_refs" TEXT NOT NULL
	`
}

// ----------------------------------------------------------------------
//
// Personal Motivations Table
//
// ----------------------------------------------------------------------

/*
commonPersonalMotivationsProperties - This method will return the properties for personal motivations
Used by:
  threat actor
*/
func commonPersonalMotivationsProperties() string {
	return baseDBProperties() + `
	"personal_motivations" TEXT NOT NULL
	`
}

// ----------------------------------------------------------------------
//
// Secondary Motivations Table
//
// ----------------------------------------------------------------------

/*
commonSecondaryMotivationsProperties - This method will return the properties for secondary motivations
Used by:
  intrusion set
  threat actor
*/
func commonSecondaryMotivationsProperties() string {
	return baseDBProperties() + `
	"secondary_motivations" TEXT NOT NULL
	`
}
