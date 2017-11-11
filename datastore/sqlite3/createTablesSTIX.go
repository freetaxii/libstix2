// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"github.com/freetaxii/libstix2/datastore"
)

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

// CreateAllSTIXTables - This method will create all of the tables needed to store
// STIX content in the database.
func (ds *sqlite3DatastoreType) CreateAllSTIXTables() {
	ds.createTable(datastore.DB_TABLE_STIX_BASE_OBJECT, ds.baseObjectProperties())
	ds.createTable(datastore.DB_TABLE_STIX_ATTACK_PATTERN, ds.attackPatternProperties())
	ds.createTable(datastore.DB_TABLE_STIX_CAMPAIGN, ds.campaignProperties())
	ds.createTable(datastore.DB_TABLE_STIX_COURSE_OF_ACTION, ds.courseOfActionProperties())
	ds.createTable(datastore.DB_TABLE_STIX_IDENTITY, ds.identityProperties())
	ds.createTable(datastore.DB_TABLE_STIX_IDENTITY_SECTORS, ds.identitySectorsProperties())
	ds.createTable(datastore.DB_TABLE_STIX_INDICATOR, ds.indicatorProperties())
	ds.createTable(datastore.DB_TABLE_STIX_INTRUSION_SET, ds.intrusionSetProperties())
	ds.createTable(datastore.DB_TABLE_STIX_LOCATION, ds.locationProperties())
	ds.createTable(datastore.DB_TABLE_STIX_MALWARE, ds.malwareProperties())
	ds.createTable(datastore.DB_TABLE_STIX_NOTE, ds.noteProperties())
	ds.createTable(datastore.DB_TABLE_STIX_OBSERVED_DATA, ds.observedDataProperties())
	ds.createTable(datastore.DB_TABLE_STIX_OPINION, ds.opinionProperties())
	ds.createTable(datastore.DB_TABLE_STIX_REPORT, ds.reportProperties())
	ds.createTable(datastore.DB_TABLE_STIX_THREAT_ACTOR, ds.threatActorProperties())
	ds.createTable(datastore.DB_TABLE_STIX_THREAT_ACTOR_ROLES, ds.threatActorRolesProperties())
	ds.createTable(datastore.DB_TABLE_STIX_TOOL, ds.toolProperties())
	ds.createTable(datastore.DB_TABLE_STIX_VULNERABILITY, ds.vulnerabilityProperties())
	ds.createTable(datastore.DB_TABLE_STIX_ALIASES, ds.commonAliasesProperties())
	ds.createTable(datastore.DB_TABLE_STIX_AUTHORS, ds.commonAuthorsProperties())
	ds.createTable(datastore.DB_TABLE_STIX_EXTERNAL_REFERENCES, ds.commonExternalReferencesProperties())
	ds.createTable(datastore.DB_TABLE_STIX_GOALS, ds.commonGoalsProperties())
	ds.createTable(datastore.DB_TABLE_STIX_HASHES, ds.commonHashesProperties())
	ds.createTable(datastore.DB_TABLE_STIX_KILL_CHAIN_PHASES, ds.commonKillChainPhasesProperties())
	ds.createTable(datastore.DB_TABLE_STIX_LABELS, ds.commonLabelsProperties())
	ds.createTable(datastore.DB_TABLE_STIX_OBJECT_MARKING_REFS, ds.commonObjectMarkingRefsProperties())
	ds.createTable(datastore.DB_TABLE_STIX_OBJECT_REFS, ds.commonObjectRefsProperties())
	ds.createTable(datastore.DB_TABLE_STIX_SECONDARY_MOTIVATIONS, ds.commonSecondaryMotivationsProperties())
	ds.createTable(datastore.DB_TABLE_STIX_PERSONAL_MOTIVATIONS, ds.commonPersonalMotivationsProperties())
}

// ----------------------------------------------------------------------
// Private Methods
// ----------------------------------------------------------------------

// baseProperties - This method will return the base properties for all objects
// row_id    = This is a database tracking number
// object_id = This is a unique identifier for the STIX object based on its ID + created and modified timestamps
func (ds *sqlite3DatastoreType) baseProperties() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
 	"object_id" TEXT NOT NULL,`
}

// baseObjectProperties - This method will return the the common properties
// spec_version = STIX specification version
// date_added = TAXII, the date the object was added to the TAXII server
func (ds *sqlite3DatastoreType) baseObjectProperties() string {
	return ds.baseProperties() + `
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
func (ds *sqlite3DatastoreType) attackPatternProperties() string {
	return ds.baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT
	`
	// kill_chain_phases
}

// campaignProperties  - This method will return the properties for campaign SDOs
func (ds *sqlite3DatastoreType) campaignProperties() string {
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
func (ds *sqlite3DatastoreType) courseOfActionProperties() string {
	return ds.baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT
	`
}

// identityProperties  - This method will return the properties for identity SDOs
func (ds *sqlite3DatastoreType) identityProperties() string {
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
func (ds *sqlite3DatastoreType) identitySectorsProperties() string {
	return ds.baseProperties() + `
	"sectors" TEXT NOT NULL
	`
}

// indicatorProperties  - This method will return the properties for indicator SDOs
func (ds *sqlite3DatastoreType) indicatorProperties() string {
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
func (ds *sqlite3DatastoreType) intrusionSetProperties() string {
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
func (ds *sqlite3DatastoreType) locationProperties() string {
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
func (ds *sqlite3DatastoreType) malwareProperties() string {
	return ds.baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT
	`
	// kill_chain_phases
}

// noteProperties  - This method will return the properties for note SDOs
func (ds *sqlite3DatastoreType) noteProperties() string {
	return ds.baseProperties() + `
	"summary" TEXT,
	"description" TEXT NOT NULL
	`
	// authors
	// object_refs
}

// observedDataProperties  - This method will return the properties for observed data SDOs
func (ds *sqlite3DatastoreType) observedDataProperties() string {
	return ds.baseProperties() + `
	"first_observed" TEXT NOT NULL,
	"last_observed" TEXT NOT NULL,
	"number_observed" INTEGER NOT NULL,
	"objects" TEXT NOT NULL
	`
}

// opinionProperties - This method will return the properties for opinion SDOs
func (ds *sqlite3DatastoreType) opinionProperties() string {
	return ds.baseProperties() + `
	"description" TEXT,
	"opinion" TEXT
	`
	// authors
	// object_refs
}

// reportProperties  - This method will return the properties for report SDOs
func (ds *sqlite3DatastoreType) reportProperties() string {
	return ds.baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT,
	"published" TEXT NOT NULL
	`
	// object_refs
}

// threatActorProperties  - This method will return the properties for threat actor SDOs
func (ds *sqlite3DatastoreType) threatActorProperties() string {
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
func (ds *sqlite3DatastoreType) threatActorRolesProperties() string {
	return ds.baseProperties() + `
	"roles" TEXT NOT NULL
	`
}

// toolProperties  - This method will return the properties for tool SDOs
func (ds *sqlite3DatastoreType) toolProperties() string {
	return ds.baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT,
	"tool_version" TEXT
	`
	// kill_chain_phases
}

// vulnerabilityProperties  - This method will return the properties for vulnerability SDOs
func (ds *sqlite3DatastoreType) vulnerabilityProperties() string {
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
func (ds *sqlite3DatastoreType) commonAliasesProperties() string {
	return ds.baseProperties() + `
	"aliases" TEXT NOT NULL
	`
}

// commonAuthorsProperties - This method will return the properties for common authors
// Used by:
//   note
//   opinion
func (ds *sqlite3DatastoreType) commonAuthorsProperties() string {
	return ds.baseProperties() + `
	"authors" TEXT NOT NULL
	`
}

// commonExternalReferencesProperties - This method will return the properties for external references
// Used by:
//   all SDOs and SROs
func (ds *sqlite3DatastoreType) commonExternalReferencesProperties() string {
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
func (ds *sqlite3DatastoreType) commonGoalsProperties() string {
	return ds.baseProperties() + `
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
func (ds *sqlite3DatastoreType) commonHashesProperties() string {
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
func (ds *sqlite3DatastoreType) commonKillChainPhasesProperties() string {
	return ds.baseProperties() + `
	"kill_chain_name" TEXT NOT NULL,
	"phase_name" TEXT NOT NULL
	`
}

// commonLabelsProperties - This method will return the properties for labels
// Used by:
//   All SDOs and SROs
func (ds *sqlite3DatastoreType) commonLabelsProperties() string {
	return ds.baseProperties() + `
	"labels" TEXT NOT NULL
	`
}

// commonObjectMarkingRefsProperties - This method will return the properties for object markings
// Used by:
//   All SDOs and SROs
func (ds *sqlite3DatastoreType) commonObjectMarkingRefsProperties() string {
	return ds.baseProperties() + `
	"object_marking_refs" TEXT NOT NULL
	`
}

// commonObjectRefsProperties - This method will return the properties for object refs
// Used by:
//   note
//   opinion
//   report
func (ds *sqlite3DatastoreType) commonObjectRefsProperties() string {
	return ds.baseProperties() + `
	"object_refs" TEXT NOT NULL
	`
}

// commonPersonalMotivationsProperties - This method will return the properties for personal motivations
// Used by:
//   threat actor
func (ds *sqlite3DatastoreType) commonPersonalMotivationsProperties() string {
	return ds.baseProperties() + `
	"personal_motivations" TEXT NOT NULL
	`
}

// commonSecondaryMotivationsProperties - This method will return the properties for secondary motivations
// Used by:
//   intrusion set
//   threat actor
func (ds *sqlite3DatastoreType) commonSecondaryMotivationsProperties() string {
	return ds.baseProperties() + `
	"secondary_motivations" TEXT NOT NULL
	`
}
