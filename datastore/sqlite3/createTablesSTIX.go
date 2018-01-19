// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"github.com/freetaxii/libstix2/datastore"
	"log"
)

// ----------------------------------------------------------------------
//
// Public Methods
//
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

// CreateAllVocabTables - This method will create all of the tables needed to store
// STIX content in the database.
func (ds *Sqlite3DatastoreType) CreateAllVocabTables() {
	ds.createVocabTable(datastore.DB_TABLE_VOCAB_ATTACK_MOTIVATIONS, vocabProperties())
	ds.createVocabTable(datastore.DB_TABLE_VOCAB_ATTACK_RESOURCE_LEVEL, vocabProperties())
	ds.createVocabTable(datastore.DB_TABLE_VOCAB_IDENTITY_CLASS, vocabProperties())
	ds.createVocabTable(datastore.DB_TABLE_VOCAB_INDICATOR_LABEL, vocabProperties())
	ds.createVocabTable(datastore.DB_TABLE_VOCAB_INDUSTRY_SECTOR, vocabProperties())
	ds.createVocabTable(datastore.DB_TABLE_VOCAB_MALWARE_LABEL, vocabProperties())
	ds.createVocabTable(datastore.DB_TABLE_VOCAB_REPORT_LABEL, vocabProperties())
	ds.createVocabTable(datastore.DB_TABLE_VOCAB_THREAT_ACTOR_LABEL, vocabProperties())
	ds.createVocabTable(datastore.DB_TABLE_VOCAB_THREAT_ACTOR_ROLE, vocabProperties())
	ds.createVocabTable(datastore.DB_TABLE_VOCAB_THREAT_ACTOR_SOPHISTICATION, vocabProperties())
	ds.createVocabTable(datastore.DB_TABLE_VOCAB_TOOL_LABEL, vocabProperties())
}

// PopulateAllVocabTables - This method will insert all of the vocabulary data
// into the right database tables.
func (ds *Sqlite3DatastoreType) PopulateAllVocabTables() {
	ds.insertVocabData(datastore.DB_TABLE_VOCAB_ATTACK_MOTIVATIONS, vocabs.AttackMotivation)
	ds.insertVocabData(datastore.DB_TABLE_VOCAB_ATTACK_RESOURCE_LEVEL, vocabs.AttackResourceLevel)
	ds.insertVocabData(datastore.DB_TABLE_VOCAB_IDENTITY_CLASS, vocabs.IdentityClass)
	ds.insertVocabData(datastore.DB_TABLE_VOCAB_INDICATOR_LABEL, vocabs.IndicatorLabel)
	ds.insertVocabData(datastore.DB_TABLE_VOCAB_INDUSTRY_SECTOR, vocabs.IndustrySector)
	ds.insertVocabData(datastore.DB_TABLE_VOCAB_MALWARE_LABEL, vocabs.MalwareLabel)
	ds.insertVocabData(datastore.DB_TABLE_VOCAB_REPORT_LABEL, vocabs.ReportLabel)
	ds.insertVocabData(datastore.DB_TABLE_VOCAB_THREAT_ACTOR_LABEL, vocabs.ThreatActorLabel)
	ds.insertVocabData(datastore.DB_TABLE_VOCAB_THREAT_ACTOR_ROLE, vocabs.ThreatActorRole)
	ds.insertVocabData(datastore.DB_TABLE_VOCAB_THREAT_ACTOR_SOPHISTICATION, vocabs.ThreatActorSophistication)
	ds.insertVocabData(datastore.DB_TABLE_VOCAB_TOOL_LABEL, vocabs.ToolLabel)
}

// ----------------------------------------------------------------------
//
// Private Methods
//
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

func (ds *Sqlite3DatastoreType) createVocabTable(name, properties string) {
	var stmt = `CREATE TABLE IF NOT EXISTS "` + name + `" (` + properties + `)`
	_, err := ds.DB.Exec(stmt)

	if err != nil {
		log.Println("ERROR: The", name, "table could not be created due to error:", err)
	}
}

// InsertVocabData - This method will add a vocabulary item to its table
func (ds *Sqlite3DatastoreType) insertVocabData(name string, data []string) {
	var stmt = `INSERT INTO "` + name + `" (value) values (?)`

	var err error
	for _, value := range data {
		_, err = ds.DB.Exec(stmt, value)
	}

	if err != nil {
		log.Println("ERROR: The vocabulary item could not be inserted in to the", name, "table")
	}
}
