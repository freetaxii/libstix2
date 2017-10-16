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

// CreateTables - This method will create all of the tables needed to store
// STIX content in the database.
func (ds *Sqlite3DatastoreType) CreateAllTables() {
	ds.createTable("stix_base_object", ds.commonProperties())
	ds.createTable("sdo_attack_pattern", ds.attackPatternProperties())
	ds.createTable("sdo_campaign", ds.campaignProperties())
	ds.createTable("sdo_course_of_action", ds.courseOfActionProperties())
	ds.createTable("sdo_identity", ds.identityProperties())
	ds.createTable("sdo_indicator", ds.indicatorProperties())
	ds.createTable("sdo_intrusion_set", ds.intrusionSetProperties())
	ds.createTable("sdo_malware", ds.malwareProperties())
	ds.createTable("sdo_observed_data", ds.observedDataProperties())
	ds.createTable("sdo_report", ds.reportProperties())
	ds.createTable("sdo_report_objects", ds.reportObjectsProperties())
	ds.createTable("sdo_threat_actor", ds.threatActorProperties())
	ds.createTable("sdo_threat_actor_goals", ds.threatActorGoalsProperties())
	ds.createTable("sdo_tool", ds.toolProperties())
	ds.createTable("sdo_vulnerability", ds.vulnerabilityProperties())
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

func (ds *Sqlite3DatastoreType) baseProperties() string {
	return ds.baseProperties() + `
	"row_id" INTEGER PRIMARY KEY,
 	"object_id" TEXT NOT NULL,`
}

// commonProperties - This method will return the the common properties
func (ds *Sqlite3DatastoreType) commonProperties() string {
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
 	"lang" TEXT,`
}

// attackPatternProperties  - This method will return the properties for attack pattern SDOs
func (ds *Sqlite3DatastoreType) attackPatternProperties() string {
	return ds.baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT
	`
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
	"identity_class" INTEGER NOT NULL,
	"contact_information" TEXT
	`
}

// indicatorProperties  - This method will return the properties for indicator SDOs
func (ds *Sqlite3DatastoreType) indicatorProperties() string {
	return ds.baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT,
	"pattern" TEXT NOT NULL,
	"valid_from" TEXT NOT NULL,
	"valid_until" TEXT
	`
}

// intrusionSetProperties  - This method will return the properties for intrusion set SDOs
func (ds *Sqlite3DatastoreType) intrusionSetProperties() string {
	return ds.baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT,
	"first_seen" TEXT,
	"last_seen" TEXT,
	"resource_level" INTEGER,
	"primary_motivation" TEXT
	`
}

// malwareProperties  - This method will return the properties for malware SDOs
func (ds *Sqlite3DatastoreType) malwareProperties() string {
	return ds.baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT
	`
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

// reportProperties  - This method will return the properties for report SDOs
func (ds *Sqlite3DatastoreType) reportProperties() string {
	return ds.baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT,
	"published" TEXT NOT NULL
	`
}

// reportObjectsProperties  - This method will return the properties for report objects
func (ds *Sqlite3DatastoreType) reportObjectsProperties() string {
	return ds.baseProperties() + `
	"object_refs" TEXT NOT NULL
	`
}

// threatActorProperties  - This method will return the properties for threat actor SDOs
func (ds *Sqlite3DatastoreType) threatActorProperties() string {
	return ds.baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT,
	"sophistication" INTEGER,
	"resource_level" INTEGER,
	"primary_motivation" TEXT
	`
}

// threatActorGoalsProperties  - This method will return the properties for threat actor goals
func (ds *Sqlite3DatastoreType) threatActorGoalsProperties() string {
	return ds.baseProperties() + `
	"goals" TEXT NOT NULL
	`
}

// toolProperties  - This method will return the properties for tool SDOs
func (ds *Sqlite3DatastoreType) toolProperties() string {
	return ds.baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT,
	"tool_version" TEXT
	`
}

// vulnerabilityProperties  - This method will return the properties for vulnerability SDOs
func (ds *Sqlite3DatastoreType) vulnerabilityProperties() string {
	return ds.baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT
	`
}

// commonAliasesProperties - This method will return the properties for aliases
func (ds *Sqlite3DatastoreType) commonAliasesProperties() string {
	return ds.baseProperties() + `
	"aliases" TEXT NOT NULL
	`
}

// commonExternalReferencesProperties - This method will return the properties for external references
func (ds *Sqlite3DatastoreType) commonExternalReferencesProperties() string {
	return ds.baseProperties() + `
	"source_name" TEXT NOT NULL,
	"description" TEXT,
	"url" TEXT,
	"external_id" TEXT
	`
}

// commonHashesProperties - This method will return the properties for hashes
func (ds *Sqlite3DatastoreType) commonHashesProperties() string {
	return ds.baseProperties() + `
	"hash" TEXT NOT NULL,
	"value" TEXT NOT NULL
	`
}

// commonKillChainPhasesProperties - This method will return the properties for kill chain phases
func (ds *Sqlite3DatastoreType) commonKillChainPhasesProperties() string {
	return ds.baseProperties() + `
	"kill_chain_name" TEXT NOT NULL,
	"phase_name" TEXT NOT NULL
	`
}

// commonLabelsProperties - This method will return the properties for labels
func (ds *Sqlite3DatastoreType) commonLabelsProperties() string {
	return ds.baseProperties() + `
	"labels" TEXT
	`
}

// commonSecondaryMotivationsProperties - This method will return the properties for secondary motivations
func (ds *Sqlite3DatastoreType) commonSecondaryMotivationsProperties() string {
	return ds.baseProperties() + `
	"secondary_motivations" TEXT
	`
}

// commonObjectMarkingRefsProperties - This method will return the properties for object markings
func (ds *Sqlite3DatastoreType) commonObjectMarkingRefsProperties() string {
	return ds.baseProperties() + `
	"object_marking_refs" TEXT NOT NULL
	`
}
