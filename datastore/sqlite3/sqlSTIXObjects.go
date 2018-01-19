// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"bytes"
	"github.com/freetaxii/libstix2/datastore"
)

// ----------------------------------------------------------------------
//
// Private Function
// Each of these functions either returns a list of fields that are used for
// creating a database tables or the SQL statements for interacting with that
// table.
//
// ----------------------------------------------------------------------

// attackPatternProperties  - This method will return the properties for attack pattern SDOs
func attackPatternProperties() string {
	return baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT
	`
	// kill_chain_phases
}

// campaignProperties  - This method will return the properties for campaign SDOs
func campaignProperties() string {
	return baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT,
	"first_seen" TEXT,
	"last_seen" TEXT,
	"objective" TEXT
	`
	// aliases
}

// courseOfActionProperties  - This method will return the properties for course of action SDOs
func courseOfActionProperties() string {
	return baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT
	`
}

// identityProperties  - This method will return the properties for identity SDOs
func identityProperties() string {
	return baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT,
	"identity_class" TEXT NOT NULL,
	"contact_information" TEXT
	`
	// sectors
}

// identitySectorsProperties  - This method will return the properties for identity sectors
// Used by:
//   identity
func identitySectorsProperties() string {
	return baseProperties() + `
	"sectors" TEXT NOT NULL
	`
}

// ----------------------------------------------------------------------
//
// Indicator Table
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
func sqlAddIndicatorObject() (string, error) {
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

// intrusionSetProperties  - This method will return the properties for intrusion set SDOs
func intrusionSetProperties() string {
	return baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT,
	"first_seen" TEXT,
	"last_seen" TEXT,
	"resource_level" TEXT,
	"primary_motivation" TEXT
	`
	// aliases
	// goals
	// secondary_motivations
}

// locationProperties - This method will return the properties for location SDOs
func locationProperties() string {
	return baseProperties() + `
	"description" TEXT,
	"latitude" TEXT,
	"longitude" TEXT,
	"precision" TEXT,
	"region" TEXT,
	"country" TEXT,
	"administrative_area" TEXT,
	"city" TEXT,
	"street_address" TEXT,
	"postal_code" TEXT
	`
}

// malwareProperties  - This method will return the properties for malware SDOs
func malwareProperties() string {
	return baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT
	`
	// kill_chain_phases
}

// noteProperties  - This method will return the properties for note SDOs
func noteProperties() string {
	return baseProperties() + `
	"summary" TEXT,
	"description" TEXT NOT NULL
	`
	// authors
	// object_refs
}

// observedDataProperties  - This method will return the properties for observed data SDOs
func observedDataProperties() string {
	return baseProperties() + `
	"first_observed" TEXT NOT NULL,
	"last_observed" TEXT NOT NULL,
	"number_observed" INTEGER NOT NULL,
	"objects" TEXT NOT NULL
	`
}

// opinionProperties - This method will return the properties for opinion SDOs
func opinionProperties() string {
	return baseProperties() + `
	"description" TEXT,
	"opinion" TEXT
	`
	// authors
	// object_refs
}

// reportProperties  - This method will return the properties for report SDOs
func reportProperties() string {
	return baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT,
	"published" TEXT NOT NULL
	`
	// object_refs
}

// threatActorProperties  - This method will return the properties for threat actor SDOs
func threatActorProperties() string {
	return baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT,
	"sophistication" TEXT,
	"resource_level" TEXT,
	"primary_motivation" TEXT
	`
	// aliases
	// roles
	// goals
	// secondary_motivations
	// personal_motivations
}

// threatActorRolesProperties  - This method will return the properties for threat actor roles
// Used by:
//   threat actor
func threatActorRolesProperties() string {
	return baseProperties() + `
	"roles" TEXT NOT NULL
	`
}

// toolProperties  - This method will return the properties for tool SDOs
func toolProperties() string {
	return baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT,
	"tool_version" TEXT
	`
	// kill_chain_phases
}

// vocabProperties  - This method will return the properties for attack patterns
func vocabProperties() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
	"value" text NOT NULL
	`
}

// vulnerabilityProperties  - This method will return the properties for vulnerability SDOs
func vulnerabilityProperties() string {
	return baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT
	`
}

// ----------------------------------------------------------------------
//
// Begin Secondary STIX Tables
//
// ----------------------------------------------------------------------

// commonAliasesProperties - This method will return the properties for aliases
// Used by:
//   campaign
//   intrusion set
//   threat actor
func commonAliasesProperties() string {
	return baseProperties() + `
	"aliases" TEXT NOT NULL
	`
}

// commonAuthorsProperties - This method will return the properties for common authors
// Used by:
//   note
//   opinion
func commonAuthorsProperties() string {
	return baseProperties() + `
	"authors" TEXT NOT NULL
	`
}

// commonGoalsProperties  - This method will return the properties for goals
// Used by:
//   intrusion set
//   threat actor
func commonGoalsProperties() string {
	return baseProperties() + `
	"goals" TEXT NOT NULL
	`
}

// commonHashesProperties - This method will return the properties for hashes
// Used by:
//
//   external references
// TODO need find a way to link this back to an actual external reference instance
// maybe this should be called external references hashes.  Otherwise  how will you
// know which object in the database it is tied to.
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

// commonObjectRefsProperties - This method will return the properties for object refs
// Used by:
//   note
//   opinion
//   report
func commonObjectRefsProperties() string {
	return baseProperties() + `
	"object_refs" TEXT NOT NULL
	`
}

// commonPersonalMotivationsProperties - This method will return the properties for personal motivations
// Used by:
//   threat actor
func commonPersonalMotivationsProperties() string {
	return baseProperties() + `
	"personal_motivations" TEXT NOT NULL
	`
}

// commonSecondaryMotivationsProperties - This method will return the properties for secondary motivations
// Used by:
//   intrusion set
//   threat actor
func commonSecondaryMotivationsProperties() string {
	return baseProperties() + `
	"secondary_motivations" TEXT NOT NULL
	`
}
