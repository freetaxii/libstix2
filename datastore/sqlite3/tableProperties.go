// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

/*
attackPatternProperties  - This method will return the properties for the
attack pattern SDO
*/
func attackPatternProperties() string {
	return baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT
	`
	// kill_chain_phases
}

/*
campaignProperties  - This method will return the properties for campaign SDOs
*/
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

/*
courseOfActionProperties  - This method will return the properties for course of action SDOs
*/
func courseOfActionProperties() string {
	return baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT
	`
}

/*
identityProperties  - This method will return the properties for identity SDOs
*/
func identityProperties() string {
	return baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT,
	"identity_class" TEXT NOT NULL,
	"contact_information" TEXT
	`
	// sectors
}

/*
identitySectorsProperties  - This method will return the properties for identity sectors
Used by: identity
*/
func identitySectorsProperties() string {
	return baseProperties() + `
	"sectors" TEXT NOT NULL
	`
}

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

/*
intrusionSetProperties  - This method will return the properties for intrusion set SDOs
*/
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

/*
locationProperties - This method will return the properties for location SDOs
*/
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

/*
malwareProperties  - This method will return the properties for malware SDOs
*/
func malwareProperties() string {
	return baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT
	`
	// kill_chain_phases
}

/*
noteProperties  - This method will return the properties for note SDOs
*/
func noteProperties() string {
	return baseProperties() + `
	"summary" TEXT,
	"description" TEXT NOT NULL
	`
	// authors
	// object_refs
}

/*
observedDataProperties  - This method will return the properties for observed data SDOs
*/
func observedDataProperties() string {
	return baseProperties() + `
	"first_observed" TEXT NOT NULL,
	"last_observed" TEXT NOT NULL,
	"number_observed" INTEGER NOT NULL,
	"objects" TEXT NOT NULL
	`
}

/*
opinionProperties - This method will return the properties for opinion SDOs
*/
func opinionProperties() string {
	return baseProperties() + `
	"description" TEXT,
	"opinion" TEXT
	`
	// authors
	// object_refs
}

/*
reportProperties  - This method will return the properties for report SDOs
*/
func reportProperties() string {
	return baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT,
	"published" TEXT NOT NULL
	`
	// object_refs
}

/*
threatActorProperties  - This method will return the properties for threat actor SDOs
*/
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

/*
threatActorRolesProperties  - This method will return the properties for threat actor roles
Used by: threat actor
*/
func threatActorRolesProperties() string {
	return baseProperties() + `
	"roles" TEXT NOT NULL
	`
}

/*
toolProperties  - This method will return the properties for tool SDOs
*/
func toolProperties() string {
	return baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT,
	"tool_version" TEXT
	`
	// kill_chain_phases
}

/*
vocabProperties  - This method will return the properties for attack patterns
*/
func vocabProperties() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
	"value" text NOT NULL
	`
}

/*
vulnerabilityProperties  - This method will return the properties for vulnerability SDOs
*/
func vulnerabilityProperties() string {
	return baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT
	`
}
