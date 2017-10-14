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
func (ds *Sqlite3DatastoreType) CreateTables() {
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

// commonProperties - This method will return the the common properties
func (ds *Sqlite3DatastoreType) commonProperties() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
 	"spec_version" TEXT NOT NULL,
 	"taxii_date_added" TEXT NOT NULL,
 	"type" TEXT NOT NULL,
 	"id" TEXT NOT NULL,
 	"created_by_ref" TEXT,
 	"created" TEXT NOT NULL,
 	"modified" TEXT NOT NULL,
 	"revoked" integer(1,0) DEFAULT 0,
 	"confidence" integer(3,0),
 	"lang" text,`
}

func (ds *Sqlite3DatastoreType) childProperties() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
 	"parent_id" INTEGER NOT NULL,`
}

// attackPatternProperties  - This method will return the properties for attack patterns
func (ds *Sqlite3DatastoreType) attackPatternProperties() string {
	return ds.commonProperties() + `
	"name" text NOT NULL,
	"description" text
	`
}

// campaignProperties  - This method will return the properties for attack patterns
func (ds *Sqlite3DatastoreType) campaignProperties() string {
	return ds.commonProperties() + `
	"name" text NOT NULL,
	"description" text,
	"first_seen" text,
	"last_seen" text,
	"objective" text
	`
}

// courseOfActionProperties  - This method will return the properties for attack patterns
func (ds *Sqlite3DatastoreType) courseOfActionProperties() string {
	return ds.commonProperties() + `
	"name" text NOT NULL,
	"description" text
	`
}

// identityProperties  - This method will return the properties for attack patterns
func (ds *Sqlite3DatastoreType) identityProperties() string {
	return ds.commonProperties() + `
	"name" text NOT NULL,
	"description" text,
	"identity_class" integer NOT NULL,
	"identity_class_ov_text" text,
	"contact_information" text
	`
}

// indicatorProperties  - This method will return the properties for attack patterns
func (ds *Sqlite3DatastoreType) indicatorProperties() string {
	return ds.commonProperties() + `
	"name" text NOT NULL,
	"description" text,
	"pattern" text NOT NULL,
	"valid_from" text NOT NULL,
	"valid_until" text
	`
}

// intrusionSetProperties  - This method will return the properties for attack patterns
func (ds *Sqlite3DatastoreType) intrusionSetProperties() string {
	return ds.commonProperties() + `
	"name" text NOT NULL,
	"description" text,
	"first_seen" text,
	"last_seen" text,
	"resource_level_id" integer,
	"primary_motivation_id" integer
	`
}

// malwareProperties  - This method will return the properties for attack patterns
func (ds *Sqlite3DatastoreType) malwareProperties() string {
	return ds.commonProperties() + `
	"name" text NOT NULL,
	"description" text
	`
}

// observedDataProperties  - This method will return the properties for attack patterns
func (ds *Sqlite3DatastoreType) observedDataProperties() string {
	return ds.commonProperties() + `
	"first_observed" text NOT NULL,
	"last_observed" text NOT NULL,
	"number_observed" integer NOT NULL,
	"objects" text NOT NULL
	`
}

// reportProperties  - This method will return the properties for attack patterns
func (ds *Sqlite3DatastoreType) reportProperties() string {
	return ds.commonProperties() + `
	"name" text NOT NULL,
	"description" text,
	"published" text NOT NULL
	`
}

// reportObjectsProperties  - This method will return the properties for attack patterns
func (ds *Sqlite3DatastoreType) reportObjectsProperties() string {
	return ds.childProperties() + `
	"object_refs" text NOT NULL
	`
}

// threatActorProperties  - This method will return the properties for attack patterns
func (ds *Sqlite3DatastoreType) threatActorProperties() string {
	return ds.commonProperties() + `
	"name" text NOT NULL,
	"description" text,
	"sophistication_id" integer,
	"resource_level_id" integer,
	"primary_motivation_id" integer
	`
}

// threatActorGoalsProperties  - This method will return the properties for attack patterns
func (ds *Sqlite3DatastoreType) threatActorGoalsProperties() string {
	return ds.childProperties() + `
	"goals" text NOT NULL
	`
}

// toolProperties  - This method will return the properties for attack patterns
func (ds *Sqlite3DatastoreType) toolProperties() string {
	return ds.commonProperties() + `
	"name" text NOT NULL,
	"description" text,
	"tool_version" text
	`
}

// vulnerabilityProperties  - This method will return the properties for attack patterns
func (ds *Sqlite3DatastoreType) vulnerabilityProperties() string {
	return ds.commonProperties() + `
	"name" text NOT NULL,
	"description" text
	`
}
