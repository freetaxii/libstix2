// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"bytes"
	"fmt"
	"github.com/freetaxii/libstix2/datastore"
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
	return baseProperties() + `
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
	return baseProperties() + `
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
	return baseProperties() + `
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
	return baseProperties() + `
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
	return baseProperties() + `
	"kill_chain_name" TEXT NOT NULL,
	"phase_name" TEXT NOT NULL
	`
}

/*
sqlAddKillChainPhase - This function will return an SQL statement that will add a
kill chain phase to the database for a given object.
*/
func sqlAddKillChainPhase() (string, error) {
	tblKillChain := datastore.DB_TABLE_STIX_KILL_CHAIN_PHASES

	/*
		INSERT INTO
			s_kill_chain_phases (
				"object_id",
				"kill_chain_name",
				"phase_name"
			)
			values (?, ?)
	*/

	var s bytes.Buffer
	s.WriteString("INSERT INTO ")
	s.WriteString(tblKillChain)
	s.WriteString(" (")
	s.WriteString("\"object_id\", ")
	s.WriteString("\"kill_chain_name\", ")
	s.WriteString("\"phase_name\") ")
	s.WriteString("values (?, ?, ?)")

	return s.String(), nil
}

/*
addKillChainPhases - This method will add a kill chain phase for a given object
to the database.
*/
func (ds *DatastoreType) addKillChainPhases(objectID int64, obj *properties.KillChainPhasesPropertyType) error {
	for _, v := range obj.KillChainPhases {
		stmt, _ := sqlAddKillChainPhase()
		_, err := ds.DB.Exec(stmt, objectID, v.KillChainName, v.PhaseName)

		if err != nil {
			return fmt.Errorf("database execution error inserting kill chain phase: ", err)
		}
	}
	return nil
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
	return baseProperties() + `
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
	return baseProperties() + `
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
	return baseProperties() + `
	"secondary_motivations" TEXT NOT NULL
	`
}
