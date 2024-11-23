// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"github.com/freetaxii/libstix2/vocabs"
)

// ----------------------------------------------------------------------
//
// Public Methods
//
// ----------------------------------------------------------------------

/*
CreateTAXIITables - This method will create all of the tables needed to store
STIX content in the database.
*/
func (ds *Store) CreateTAXIITables() {
	ds.createTAXIITable(DB_TABLE_TAXII_COLLECTION_DATA, collectionDataProperties())
	ds.createTAXIITable(DB_TABLE_TAXII_COLLECTIONS, collectionProperties())
	ds.createTAXIITable(DB_TABLE_TAXII_COLLECTION_MEDIA_TYPE, collectionMediaTypeProperties())
	ds.createTAXIITable(DB_TABLE_TAXII_MEDIA_TYPES, mediaTypeProperties())
	ds.createTAXIIIndexes(DB_TABLE_TAXII_COLLECTION_DATA)
	ds.insertMediaTypes(DB_TABLE_TAXII_MEDIA_TYPES)
}

/*
CreateSTIXTables - This method will create all of the tables needed to store
STIX content in the database.
*/
func (ds *Store) CreateSTIXTables() {
	ds.createSTIXTable(DB_TABLE_STIX_BASE_OBJECT, baseObjectProperties())
	ds.createSTIXTable(DB_TABLE_STIX_ATTACK_PATTERN, attackPatternProperties())
	ds.createSTIXTable(DB_TABLE_STIX_CAMPAIGN, campaignProperties())
	ds.createSTIXTable(DB_TABLE_STIX_COURSE_OF_ACTION, courseOfActionProperties())
	ds.createSTIXTable(DB_TABLE_STIX_IDENTITY, identityProperties())
	ds.createSTIXTable(DB_TABLE_STIX_IDENTITY_SECTORS, identitySectorsProperties())
	ds.createSTIXTable(DB_TABLE_STIX_INDICATOR, indicatorProperties())
	ds.createSTIXTable(DB_TABLE_STIX_INDICATOR_TYPES, indicatorTypesProperties())
	ds.createSTIXTable(DB_TABLE_STIX_INTRUSION_SET, intrusionSetProperties())
	ds.createSTIXTable(DB_TABLE_STIX_LOCATION, locationProperties())
	ds.createSTIXTable(DB_TABLE_STIX_MALWARE, malwareProperties())
	ds.createSTIXTable(DB_TABLE_STIX_NOTE, noteProperties())
	ds.createSTIXTable(DB_TABLE_STIX_OBSERVED_DATA, observedDataProperties())
	ds.createSTIXTable(DB_TABLE_STIX_OPINION, opinionProperties())
	ds.createSTIXTable(DB_TABLE_STIX_REPORT, reportProperties())
	ds.createSTIXTable(DB_TABLE_STIX_THREAT_ACTOR, threatActorProperties())
	ds.createSTIXTable(DB_TABLE_STIX_THREAT_ACTOR_ROLES, threatActorRolesProperties())
	ds.createSTIXTable(DB_TABLE_STIX_TOOL, toolProperties())
	ds.createSTIXTable(DB_TABLE_STIX_VULNERABILITY, vulnerabilityProperties())
	ds.createSTIXTable(DB_TABLE_STIX_ALIASES, commonAliasesProperties())
	ds.createSTIXTable(DB_TABLE_STIX_AUTHORS, commonAuthorsProperties())
	ds.createSTIXTable(DB_TABLE_STIX_EXTERNAL_REFERENCES, commonExternalReferencesProperties())
	ds.createSTIXTable(DB_TABLE_STIX_GOALS, commonGoalsProperties())
	ds.createSTIXTable(DB_TABLE_STIX_HASHES, commonHashesProperties())
	ds.createSTIXTable(DB_TABLE_STIX_KILL_CHAIN_PHASES, commonKillChainPhasesProperties())
	ds.createSTIXTable(DB_TABLE_STIX_LABELS, commonLabelsProperties())
	ds.createSTIXTable(DB_TABLE_STIX_OBJECT_MARKING_REFS, commonObjectMarkingRefsProperties())
	ds.createSTIXTable(DB_TABLE_STIX_OBJECT_REFS, commonObjectRefsProperties())
	ds.createSTIXTable(DB_TABLE_STIX_SECONDARY_MOTIVATIONS, commonSecondaryMotivationsProperties())
	ds.createSTIXTable(DB_TABLE_STIX_PERSONAL_MOTIVATIONS, commonPersonalMotivationsProperties())
}

/*
CreateVocabTables - This method will create all of the tables needed to store
STIX content in the database.
*/
func (ds *Store) CreateVocabTables() {
	ds.createVocabTable(DB_TABLE_VOCAB_ATTACK_MOTIVATIONS, vocabProperties())
	ds.createVocabTable(DB_TABLE_VOCAB_ATTACK_RESOURCE_LEVEL, vocabProperties())
	ds.createVocabTable(DB_TABLE_VOCAB_IDENTITY_CLASS, vocabProperties())
	ds.createVocabTable(DB_TABLE_VOCAB_INDICATOR_LABEL, vocabProperties())
	ds.createVocabTable(DB_TABLE_VOCAB_INDUSTRY_SECTOR, vocabProperties())
	ds.createVocabTable(DB_TABLE_VOCAB_MALWARE_LABEL, vocabProperties())
	ds.createVocabTable(DB_TABLE_VOCAB_REPORT_LABEL, vocabProperties())
	ds.createVocabTable(DB_TABLE_VOCAB_THREAT_ACTOR_LABEL, vocabProperties())
	ds.createVocabTable(DB_TABLE_VOCAB_THREAT_ACTOR_ROLE, vocabProperties())
	ds.createVocabTable(DB_TABLE_VOCAB_THREAT_ACTOR_SOPHISTICATION, vocabProperties())
	ds.createVocabTable(DB_TABLE_VOCAB_TOOL_LABEL, vocabProperties())
}

/*
PopulateVocabTables - This method will insert all of the vocabulary data
into the right database tables.
*/
func (ds *Store) PopulateVocabTables() {
	ds.insertVocabData(DB_TABLE_VOCAB_ATTACK_MOTIVATIONS, vocabs.GetKeys(vocabs.GetAttackMotivationVocab()))
	ds.insertVocabData(DB_TABLE_VOCAB_ATTACK_RESOURCE_LEVEL, vocabs.GetKeys(vocabs.GetAttackResourceLevelVocab()))
	ds.insertVocabData(DB_TABLE_VOCAB_IDENTITY_CLASS, vocabs.GetKeys(vocabs.GetIdentityClassVocab()))
	ds.insertVocabData(DB_TABLE_VOCAB_INDICATOR_LABEL, vocabs.GetKeys(vocabs.GetIndicatorTypeVocab()))
	ds.insertVocabData(DB_TABLE_VOCAB_INDUSTRY_SECTOR, vocabs.GetKeys(vocabs.GetIndustrySectorVocab()))
	ds.insertVocabData(DB_TABLE_VOCAB_MALWARE_LABEL, vocabs.GetKeys(vocabs.GetMalwareTypeVocab()))
	ds.insertVocabData(DB_TABLE_VOCAB_REPORT_LABEL, vocabs.GetKeys(vocabs.GetReportTypeVocab()))
	ds.insertVocabData(DB_TABLE_VOCAB_THREAT_ACTOR_LABEL, vocabs.GetKeys(vocabs.GetThreatActorTypeVocab()))
	ds.insertVocabData(DB_TABLE_VOCAB_THREAT_ACTOR_ROLE, vocabs.GetKeys(vocabs.GetThreatActorRoleVocab()))
	ds.insertVocabData(DB_TABLE_VOCAB_THREAT_ACTOR_SOPHISTICATION, vocabs.GetKeys(vocabs.GetThreatActorSophisticationVocab()))
	ds.insertVocabData(DB_TABLE_VOCAB_TOOL_LABEL, vocabs.GetKeys(vocabs.GetToolTypeVocab()))
}

// ----------------------------------------------------------------------
//
// Private Methods
//
// ----------------------------------------------------------------------

func (ds *Store) createTAXIITable(name, properties string) {
	var stmt = `CREATE TABLE IF NOT EXISTS "` + name + `" (` + properties + `)`
	_, err := ds.DB.Exec(stmt)

	if err != nil {
		ds.Logger.Println("ERROR: The", name, "table could not be created due to error:", err)
	}
}

func (ds *Store) createTAXIIIndexes(name string) {
	var stmt string

	if name == DB_TABLE_TAXII_COLLECTION_DATA {
		stmt = `CREATE INDEX "idx_` + name + `" ON ` + name + ` ("collection_id" COLLATE BINARY ASC, "stix_id" COLLATE BINARY ASC)`
	}

	if stmt != "" {
		_, err := ds.DB.Exec(stmt)

		if err != nil {
			ds.Logger.Println("ERROR: The indexes for table", name, "could not be created due to error:", err)
		}
	}
}

func (ds *Store) insertMediaTypes(name string) {
	var stmt = `INSERT INTO "` + name + `" (media_type) values (?)`

	mediaTypes := []string{"application/stix+json;version=2.0", "application/stix+json;version=2.1", "application/stix+json;version=2.2", "application/stix+json;version=2.3"}
	for _, value := range mediaTypes {
		var err error
		_, err = ds.DB.Exec(stmt, value)

		if err != nil {
			ds.Logger.Println("ERROR: The media type item could not be inserted in to the", name, "table")
		}
	}
}

func (ds *Store) createSTIXTable(name, properties string) {
	var stmt = `CREATE TABLE IF NOT EXISTS "` + name + `" (` + properties + `)`
	_, err := ds.DB.Exec(stmt)

	if err != nil {
		ds.Logger.Println("ERROR: The", name, "table could not be created due to error:", err)
	}
	ds.createSTIXIndexes(name)
}

func (ds *Store) createSTIXIndexes(name string) {
	var stmt string

	if name == DB_TABLE_STIX_BASE_OBJECT {
		stmt = `CREATE INDEX "idx_` + name + `" ON ` + name + ` ("object_id" COLLATE BINARY ASC, "id" COLLATE BINARY ASC)`
	} else {
		stmt = `CREATE INDEX "idx_` + name + `" ON ` + name + ` ("object_id" COLLATE BINARY ASC)`
	}

	_, err := ds.DB.Exec(stmt)

	if err != nil {
		ds.Logger.Println("ERROR: The indexes for table", name, "could not be created due to error:", err)
	}
}

func (ds *Store) createVocabTable(name, properties string) {
	var stmt = `CREATE TABLE IF NOT EXISTS "` + name + `" (` + properties + `)`
	_, err := ds.DB.Exec(stmt)

	if err != nil {
		ds.Logger.Println("ERROR: The", name, "table could not be created due to error:", err)
	}
}

// InsertVocabData - This method will add a vocabulary item to its table
func (ds *Store) insertVocabData(name string, data []string) {
	var stmt = `INSERT INTO "` + name + `" (value) values (?)`

	var err error
	for _, value := range data {
		_, err = ds.DB.Exec(stmt, value)
	}

	if err != nil {
		ds.Logger.Println("ERROR: The vocabulary item could not be inserted in to the", name, "table")
	}
}
