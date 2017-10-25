// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"log"
)

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

// CreateAllTables - This method will create all of the tables needed to store
// STIX content in the database.
func (ds *Sqlite3DatastoreType) CreateAllTables() {
	ds.createTable("stix_base_object", ds.baseObjectProperties())
	ds.createTable("sdo_attack_pattern", ds.attackPatternProperties())
	ds.createTable("sdo_campaign", ds.campaignProperties())
	ds.createTable("sdo_course_of_action", ds.courseOfActionProperties())
	ds.createTable("sdo_identity", ds.identityProperties())
	ds.createTable("identity_sectors", ds.identitySectorsProperties())
	ds.createTable("sdo_indicator", ds.indicatorProperties())
	ds.createTable("sdo_intrusion_set", ds.intrusionSetProperties())
	ds.createTable("sdo_location", ds.locationProperties())
	ds.createTable("sdo_malware", ds.malwareProperties())
	ds.createTable("sdo_note", ds.noteProperties())
	ds.createTable("sdo_observed_data", ds.observedDataProperties())
	ds.createTable("sdo_opinion", ds.opinionProperties())
	ds.createTable("sdo_report", ds.reportProperties())
	ds.createTable("sdo_threat_actor", ds.threatActorProperties())
	ds.createTable("threat_actor_roles", ds.threatActorRolesProperties())
	ds.createTable("sdo_tool", ds.toolProperties())
	ds.createTable("sdo_vulnerability", ds.vulnerabilityProperties())
	ds.createTable("aliases", ds.commonAliasesProperties())
	ds.createTable("authors", ds.commonAuthorsProperties())
	ds.createTable("external_references", ds.commonExternalReferencesProperties())
	ds.createTable("goals", ds.commonGoalsProperties())
	ds.createTable("hashes", ds.commonHashesProperties())
	ds.createTable("kill_chain_phases", ds.commonKillChainPhasesProperties())
	ds.createTable("labels", ds.commonLabelsProperties())
	ds.createTable("object_marking_refs", ds.commonObjectMarkingRefsProperties())
	ds.createTable("object_refs", ds.commonObjectRefsProperties())
	ds.createTable("secondary_motivations", ds.commonSecondaryMotivationsProperties())
	ds.createTable("primary_motivations", ds.commonPersonalMotivationsProperties())
}

// ----------------------------------------------------------------------
// Private Methods
// ----------------------------------------------------------------------

// createAttackPatternTable - This method will create the actual table
func (ds *Sqlite3DatastoreType) createTable(name, properties string) {
	var stmt = `CREATE TABLE IF NOT EXISTS "` + name + `" (` + properties + `)`
	_, err := ds.DB.Exec(stmt)

	if err != nil {
		log.Println("ERROR: The", name, "table could not be created")
	}
}

// baseProperties - This method will return the base properties for all objects
// row_id    = This is a database tracking number
// object_id = This is a unique identifier for the STIX object based on its ID + created and modified timestamps
func (ds *Sqlite3DatastoreType) baseProperties() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
 	"object_id" TEXT NOT NULL,`
}

// baseObjectProperties - This method will return the the common properties
// version    = STIX specification version
// date_added = TAXII, the date the object was added to the TAXII server
func (ds *Sqlite3DatastoreType) baseObjectProperties() string {
	return ds.baseProperties() + `
 	"version" TEXT NOT NULL,
 	"date_added" TEXT NOT NULL,
 	"type" TEXT NOT NULL,
 	"id" TEXT NOT NULL,
 	"created_by_ref" TEXT,
 	"created" TEXT NOT NULL,
 	"modified" TEXT NOT NULL,
 	"revoked" INTEGER(1,0) DEFAULT 0,
 	"confidence" INTEGER(3,0),
 	"lang" TEXT`
}

// attackPatternProperties  - This method will return the properties for attack pattern SDOs
func (ds *Sqlite3DatastoreType) attackPatternProperties() string {
	return ds.baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT
	`
	// kill_chain_phases
}

// campaignProperties  - This method will return the properties for campaign SDOs
func (ds *Sqlite3DatastoreType) campaignProperties() string {
	return ds.baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT,
	"first_seen" TEXT,
	"last_seen" TEXT,
	"objective" TEXT
	`
	// aliases
}

// courseOfActionProperties  - This method will return the properties for course of action SDOs
func (ds *Sqlite3DatastoreType) courseOfActionProperties() string {
	return ds.baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT
	`
}

// identityProperties  - This method will return the properties for identity SDOs
func (ds *Sqlite3DatastoreType) identityProperties() string {
	return ds.baseProperties() + `
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
func (ds *Sqlite3DatastoreType) identitySectorsProperties() string {
	return ds.baseProperties() + `
	"sectors" TEXT NOT NULL
	`
}

// indicatorProperties  - This method will return the properties for indicator SDOs
func (ds *Sqlite3DatastoreType) indicatorProperties() string {
	return ds.baseProperties() + `
	"name" TEXT,
	"description" TEXT,
	"pattern" TEXT NOT NULL,
	"valid_from" TEXT NOT NULL,
	"valid_until" TEXT
	`
	// kill_chain_phases
}

// intrusionSetProperties  - This method will return the properties for intrusion set SDOs
func (ds *Sqlite3DatastoreType) intrusionSetProperties() string {
	return ds.baseProperties() + `
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
func (ds *Sqlite3DatastoreType) locationProperties() string {
	return ds.baseProperties() + `
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
func (ds *Sqlite3DatastoreType) malwareProperties() string {
	return ds.baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT
	`
	// kill_chain_phases
}

// noteProperties  - This method will return the properties for note SDOs
func (ds *Sqlite3DatastoreType) noteProperties() string {
	return ds.baseProperties() + `
	"summary" TEXT,
	"description" TEXT NOT NULL
	`
	// authors
	// object_refs
}

// observedDataProperties  - This method will return the properties for observed data SDOs
func (ds *Sqlite3DatastoreType) observedDataProperties() string {
	return ds.baseProperties() + `
	"first_observed" TEXT NOT NULL,
	"last_observed" TEXT NOT NULL,
	"number_observed" INTEGER NOT NULL,
	"objects" TEXT NOT NULL
	`
}

// opinionProperties - This method will return the properties for opinion SDOs
func (ds *Sqlite3DatastoreType) opinionProperties() string {
	return ds.baseProperties() + `
	"description" TEXT,
	"opinion" TEXT
	`
	// authors
	// object_refs
}

// reportProperties  - This method will return the properties for report SDOs
func (ds *Sqlite3DatastoreType) reportProperties() string {
	return ds.baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT,
	"published" TEXT NOT NULL
	`
	// object_refs
}

// threatActorProperties  - This method will return the properties for threat actor SDOs
func (ds *Sqlite3DatastoreType) threatActorProperties() string {
	return ds.baseProperties() + `
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
func (ds *Sqlite3DatastoreType) threatActorRolesProperties() string {
	return ds.baseProperties() + `
	"roles" TEXT NOT NULL
	`
}

// toolProperties  - This method will return the properties for tool SDOs
func (ds *Sqlite3DatastoreType) toolProperties() string {
	return ds.baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT,
	"tool_version" TEXT
	`
	// kill_chain_phases
}

// vulnerabilityProperties  - This method will return the properties for vulnerability SDOs
func (ds *Sqlite3DatastoreType) vulnerabilityProperties() string {
	return ds.baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT
	`
}

// ----------------------------------------------------------------------
// Begin Common Properties
// ----------------------------------------------------------------------

// commonAliasesProperties - This method will return the properties for aliases
// Used by:
//   campaign
//   intrusion set
//   threat actor
func (ds *Sqlite3DatastoreType) commonAliasesProperties() string {
	return ds.baseProperties() + `
	"aliases" TEXT NOT NULL
	`
}

// commonAuthorsProperties - This method will return the properties for common authors
// Used by:
//   note
//   opinion
func (ds *Sqlite3DatastoreType) commonAuthorsProperties() string {
	return ds.baseProperties() + `
	"authors" TEXT NOT NULL
	`
}

// commonExternalReferencesProperties - This method will return the properties for external references
// Used by:
//   all SDOs and SROs
func (ds *Sqlite3DatastoreType) commonExternalReferencesProperties() string {
	return ds.baseProperties() + `
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
func (ds *Sqlite3DatastoreType) commonGoalsProperties() string {
	return ds.baseProperties() + `
	"goals" TEXT NOT NULL
	`
}

// commonHashesProperties - This method will return the properties for hashes
// Used by:
//
func (ds *Sqlite3DatastoreType) commonHashesProperties() string {
	return ds.baseProperties() + `
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
func (ds *Sqlite3DatastoreType) commonKillChainPhasesProperties() string {
	return ds.baseProperties() + `
	"kill_chain_name" TEXT NOT NULL,
	"phase_name" TEXT NOT NULL
	`
}

// commonLabelsProperties - This method will return the properties for labels
// Used by:
//   All SDOs and SROs
func (ds *Sqlite3DatastoreType) commonLabelsProperties() string {
	return ds.baseProperties() + `
	"labels" TEXT NOT NULL
	`
}

// commonObjectMarkingRefsProperties - This method will return the properties for object markings
// Used by:
//   All SDOs and SROs
func (ds *Sqlite3DatastoreType) commonObjectMarkingRefsProperties() string {
	return ds.baseProperties() + `
	"object_marking_refs" TEXT NOT NULL
	`
}

// commonObjectRefsProperties - This method will return the properties for object refs
// Used by:
//   note
//   opinion
//   report
func (ds *Sqlite3DatastoreType) commonObjectRefsProperties() string {
	return ds.baseProperties() + `
	"object_refs" TEXT NOT NULL
	`
}

// commonPersonalMotivationsProperties - This method will return the properties for personal motivations
// Used by:
//   threat actor
func (ds *Sqlite3DatastoreType) commonPersonalMotivationsProperties() string {
	return ds.baseProperties() + `
	"personal_motivations" TEXT NOT NULL
	`
}

// commonSecondaryMotivationsProperties - This method will return the properties for secondary motivations
// Used by:
//   intrusion set
//   threat actor
func (ds *Sqlite3DatastoreType) commonSecondaryMotivationsProperties() string {
	return ds.baseProperties() + `
	"secondary_motivations" TEXT NOT NULL
	`
}
