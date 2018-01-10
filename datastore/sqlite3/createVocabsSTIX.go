// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"github.com/freetaxii/libstix2/datastore"
	"github.com/freetaxii/libstix2/vocabs"
	"log"
)

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

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
// Private Methods
// ----------------------------------------------------------------------

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

// vocabProperties  - This method will return the properties for attack patterns
func vocabProperties() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
	"value" text NOT NULL
	`
}
