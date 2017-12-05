// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"github.com/freetaxii/libstix2/datastore"
	"log"
)

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

// CreateAllSTIXTables - This method will create all of the tables needed to store
// STIX content in the database.
func (ds *Sqlite3DatastoreType) CreateAllSTIXTables() {
	ds.createSTIXTable(datastore.DB_TABLE_STIX_BASE_OBJECT, baseObjectProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_ATTACK_PATTERN, attackPatternProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_CAMPAIGN, campaignProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_COURSE_OF_ACTION, courseOfActionProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_IDENTITY, identityProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_IDENTITY_SECTORS, identitySectorsProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_INDICATOR, indicatorProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_INTRUSION_SET, intrusionSetProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_LOCATION, locationProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_MALWARE, malwareProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_NOTE, noteProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_OBSERVED_DATA, observedDataProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_OPINION, opinionProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_REPORT, reportProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_THREAT_ACTOR, threatActorProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_THREAT_ACTOR_ROLES, threatActorRolesProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_TOOL, toolProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_VULNERABILITY, vulnerabilityProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_ALIASES, commonAliasesProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_AUTHORS, commonAuthorsProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_EXTERNAL_REFERENCES, commonExternalReferencesProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_GOALS, commonGoalsProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_HASHES, commonHashesProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_KILL_CHAIN_PHASES, commonKillChainPhasesProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_LABELS, commonLabelsProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_OBJECT_MARKING_REFS, commonObjectMarkingRefsProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_OBJECT_REFS, commonObjectRefsProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_SECONDARY_MOTIVATIONS, commonSecondaryMotivationsProperties())
	ds.createSTIXTable(datastore.DB_TABLE_STIX_PERSONAL_MOTIVATIONS, commonPersonalMotivationsProperties())
}

// ----------------------------------------------------------------------
// Private Methods
// ----------------------------------------------------------------------

func (ds *Sqlite3DatastoreType) createSTIXTable(name, properties string) {
	var stmt = `CREATE TABLE IF NOT EXISTS "` + name + `" (` + properties + `)`
	_, err := ds.DB.Exec(stmt)

	if err != nil {
		log.Println("ERROR: The", name, "table could not be created due to error:", err)
	}
	ds.createSTIXIndexes(name)
}

func (ds *Sqlite3DatastoreType) createSTIXIndexes(name string) {
	var stmt string

	if name == datastore.DB_TABLE_STIX_BASE_OBJECT {
		stmt = `CREATE INDEX "idx_` + name + `" ON ` + name + ` ("object_id" COLLATE BINARY ASC, "id" COLLATE BINARY ASC)`
	} else {
		stmt = `CREATE INDEX "idx_` + name + `" ON ` + name + ` ("object_id" COLLATE BINARY ASC)`
	}

	_, err := ds.DB.Exec(stmt)

	if err != nil {
		log.Println("ERROR: The indexes for table", name, "could not be created due to error:", err)
	}
}

// ----------------------------------------------------------------------
//
// Each of these functions returns a list of fields that are used for creating
// a database tables.
//
// ----------------------------------------------------------------------

// baseProperties - This method will return the base properties for all objects
// row_id    = This is a database tracking number
// object_id = This is a unique identifier for the STIX object based on its ID + created and modified timestamps
func baseProperties() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
 	"object_id" TEXT NOT NULL,`
}

// baseObjectProperties - This method will return the the common properties
// spec_version = STIX specification version
// date_added = TAXII, the date the object was added to the TAXII server
func baseObjectProperties() string {
	return baseProperties() + `
 	"spec_version" TEXT NOT NULL,
 	"date_added" TEXT NOT NULL,
 	"type" TEXT NOT NULL,
 	"id" TEXT NOT NULL,
 	"created_by_ref" TEXT,
 	"created" TEXT NOT NULL,
 	"modified" TEXT NOT NULL,
 	"revoked" INTEGER(1,0) DEFAULT 0,
 	"confidence" INTEGER(3,0),
 	"lang" TEXT`
	// labels
	// external_references
	// object_marking_refs
}

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

// indicatorProperties  - This method will return the properties for indicator SDOs
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

// vulnerabilityProperties  - This method will return the properties for vulnerability SDOs
func vulnerabilityProperties() string {
	return baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT
	`
}

// ----------------------------------------------------------------------
// Begin Secondary STIX Tables
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

// commonExternalReferencesProperties - This method will return the properties for external references
// Used by:
//   all SDOs and SROs
func commonExternalReferencesProperties() string {
	return baseProperties() + `
	"source_name" TEXT NOT NULL,
	"description" TEXT,
	"url" TEXT,
	"external_id" TEXT
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

// commonKillChainPhasesProperties - This method will return the properties for kill chain phases
// Used by:
//   attack pattern
//   indicator
//   malware
//   tool
func commonKillChainPhasesProperties() string {
	return baseProperties() + `
	"kill_chain_name" TEXT NOT NULL,
	"phase_name" TEXT NOT NULL
	`
}

// commonLabelsProperties - This method will return the properties for labels
// Used by:
//   All SDOs and SROs
func commonLabelsProperties() string {
	return baseProperties() + `
	"labels" TEXT NOT NULL
	`
}

// commonObjectMarkingRefsProperties - This method will return the properties for object markings
// Used by:
//   All SDOs and SROs
func commonObjectMarkingRefsProperties() string {
	return baseProperties() + `
	"object_marking_refs" TEXT NOT NULL
	`
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
