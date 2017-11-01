// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"github.com/freetaxii/libstix2/vocabs"
	"log"
)

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

// CreateAllVocabTables - This method will create all of the tables needed to store
// STIX content in the database.
func (ds *Sqlite3DatastoreType) CreateAllVocabTables() {
	ds.createTable("v_attack_motivation", ds.vocabProperties())
	ds.createTable("v_attack_resource_level", ds.vocabProperties())
	ds.createTable("v_identity_class", ds.vocabProperties())
	ds.createTable("v_indicator_label", ds.vocabProperties())
	ds.createTable("v_industry_sector", ds.vocabProperties())
	ds.createTable("v_malware_label", ds.vocabProperties())
	ds.createTable("v_report_label", ds.vocabProperties())
	ds.createTable("v_threat_actor_label", ds.vocabProperties())
	ds.createTable("v_threat_actor_role", ds.vocabProperties())
	ds.createTable("v_threat_actor_sophistication", ds.vocabProperties())
	ds.createTable("v_tool_label", ds.vocabProperties())
}

// PopulateAllVocabTables - This method will insert all of the vocabulary data
// into the right database tables.
func (ds *Sqlite3DatastoreType) PopulateAllVocabTables() {
	ds.insertVocabData("v_attack_motivation", vocabs.AttackMotivation)
	ds.insertVocabData("v_attack_resource_level", vocabs.AttackResourceLevel)
	ds.insertVocabData("v_identity_class", vocabs.IdentityClass)
	ds.insertVocabData("v_indicator_label", vocabs.IndicatorLabel)
	ds.insertVocabData("v_industry_sector", vocabs.IndustrySector)
	ds.insertVocabData("v_malware_label", vocabs.MalwareLabel)
	ds.insertVocabData("v_report_label", vocabs.ReportLabel)
	ds.insertVocabData("v_threat_actor_label", vocabs.ThreatActorLabel)
	ds.insertVocabData("v_threat_actor_role", vocabs.ThreatActorRole)
	ds.insertVocabData("v_threat_actor_sophistication", vocabs.ThreatActorSophistication)
	ds.insertVocabData("v_tool_label", vocabs.ToolLabel)
}

// ----------------------------------------------------------------------
// Private Methods
// ----------------------------------------------------------------------

// vocabProperties  - This method will return the properties for attack patterns
func (ds *Sqlite3DatastoreType) vocabProperties() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
	"value" text NOT NULL
	`
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
