// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"crypto/sha1"
	"encoding/base64"
	"github.com/freetaxii/libstix2/defs"
	"github.com/freetaxii/libstix2/objects/common/properties"
	"github.com/freetaxii/libstix2/objects/indicator"
	"log"
	"time"
)

func (ds *Sqlite3DatastoreType) addBaseObject(obj properties.CommonObjectPropertiesType) string {
	dateAdded := time.Now().UTC().Format(defs.TIME_RFC_3339_MICRO)
	objectID := "id" + obj.ID + "created" + obj.Created + "modified" + obj.Modified

	h := sha1.New()
	h.Write([]byte(objectID))
	hashID := base64.URLEncoding.EncodeToString(h.Sum(nil))

	var stmt1 = `INSERT INTO "stix_base_object" (
	 	"object_id", 
	 	"spec_version", 
	 	"date_added", 
	 	"type", 
	 	"id", 
	 	"created_by_ref",
	 	"created", 
	 	"modified", 
	 	"revoked", 
	 	"confidence", 
	 	"lang"
	 	)
		values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err1 := ds.DB.Exec(stmt1,
		hashID,
		obj.SpecVersion,
		dateAdded,
		obj.MessageType,
		obj.ID,
		obj.CreatedByRef,
		obj.Created,
		obj.Modified,
		obj.Revoked,
		obj.Confidence,
		obj.Lang)

	if err1 != nil {
		log.Println("ERROR: Database execution error inserting base record", err1)
	}

	if obj.Labels != nil {
		for _, label := range obj.Labels {
			var stmt2 = `INSERT INTO "labels" (
			"object_id",
			"labels"
			)
			values (?, ?)`

			_, err2 := ds.DB.Exec(stmt2, hashID, label)

			if err2 != nil {
				log.Println("ERROR: Database execution error inserting labels", err2)
			}
		}
	}

	if obj.ExternalReferences != nil {
		for _, reference := range obj.ExternalReferences {
			var stmt3 = `INSERT INTO "external_references" (
			"object_id",
			"source_name",
			"description"
			"url",
			"external_id"
			)
			values (?, ?, ?, ?, ?)`

			_, err3 := ds.DB.Exec(stmt3,
				hashID,
				reference.SourceName,
				reference.Description,
				reference.URL,
				reference.ExternalID)

			if err3 != nil {
				log.Println("ERROR: Database execution error inserting external references", err3)
			}
		}
	}

	if obj.ObjectMarkingRefs != nil {
		for _, marking := range obj.ObjectMarkingRefs {
			var stmt4 = `INSERT INTO "object_marking_refs" (
			"object_id",
			"object_marking_refs"
			)
			values (?, ?)`

			_, err4 := ds.DB.Exec(stmt4, hashID, marking)

			if err4 != nil {
				log.Println("ERROR: Database execution error inserting object marking refs", err4)
			}
		}
	}

	return hashID
}

// addKillChainPhases
func (ds *Sqlite3DatastoreType) addKillChainPhases(hashID string, obj properties.KillChainPhasesPropertyType) {
	for _, v := range obj.KillChainPhases {
		var stmt = `INSERT INTO "kill_chain_phases" (
			"object_id",
			"kill_chain_name",
			"phase_name"
			)
			values (?, ?, ?)`

		_, err := ds.DB.Exec(stmt, hashID, v.KillChainName, v.PhaseName)

		if err != nil {
			log.Println("ERROR: Database execution error inserting kill chain phases", err)
		}
	}
}

// addIndicator
func (ds *Sqlite3DatastoreType) addIndicator(obj indicator.IndicatorType) error {

	hashID := ds.addBaseObject(obj.CommonObjectPropertiesType)

	var stmt1 = `INSERT INTO "sdo_indicator" (
		"object_id",
		"name",
		"description",
		"pattern",
		"valid_from",
		"valid_until"
		)
		values (?, ?, ?, ?, ?, ?)`

	_, err1 := ds.DB.Exec(stmt1,
		hashID,
		obj.Name,
		obj.Description,
		obj.Pattern,
		obj.ValidFrom,
		obj.ValidUntil)

	// TODO if there is an error, we probably need to back out all of the INSERTS
	if err1 != nil {
		return err1
	}

	if obj.KillChainPhases != nil {
		ds.addKillChainPhases(hashID, obj.KillChainPhasesPropertyType)
	}

	return nil
}
