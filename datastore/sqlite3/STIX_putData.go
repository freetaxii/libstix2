// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"fmt"
	"github.com/freetaxii/libstix2/datastore"
	"github.com/freetaxii/libstix2/defs"
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
	"log"
	"time"
)

/*
addBaseObject - This method will add the base properties of an object to the
database and return an integer that tracks the record number for parent child
relationships.
*/
func (ds *Sqlite3DatastoreType) addBaseObject(obj properties.CommonObjectPropertiesType) (int64, error) {
	dateAdded := time.Now().UTC().Format(defs.TIME_RFC_3339_MICRO)

	objectID := ds.Index
	ds.Index++

	stmt1, _ := ds.sqlAddBaseObject()

	_, err1 := ds.DB.Exec(stmt1,
		objectID,
		obj.SpecVersion,
		dateAdded,
		obj.ObjectType,
		obj.ID,
		obj.CreatedByRef,
		obj.Created,
		obj.Modified,
		obj.Revoked,
		obj.Confidence,
		obj.Lang)

	if err1 != nil {
		return 0, fmt.Errorf("database execution error inserting base object: ", err1)
	}

	// ----------------------------------------------------------------------
	// Add Labels
	// ----------------------------------------------------------------------
	if obj.Labels != nil {
		for _, label := range obj.Labels {
			stmt2, _ := ds.sqlAddObjectLabel()
			_, err2 := ds.DB.Exec(stmt2, objectID, label)

			if err2 != nil {
				return 0, fmt.Errorf("database execution error inserting object label: ", err2)
			}
		}
	}

	// ----------------------------------------------------------------------
	// Add External References
	// ----------------------------------------------------------------------
	if obj.ExternalReferences != nil {
		for _, reference := range obj.ExternalReferences {
			stmt3, _ := ds.sqlAddExternalReference()

			_, err3 := ds.DB.Exec(stmt3,
				objectID,
				reference.SourceName,
				reference.Description,
				reference.URL,
				reference.ExternalID)

			if err3 != nil {
				return 0, fmt.Errorf("database execution error inserting external reference: ", err3)
			}
		}
	}

	// ----------------------------------------------------------------------
	// Add External References
	// ----------------------------------------------------------------------
	if obj.ObjectMarkingRefs != nil {
		for _, marking := range obj.ObjectMarkingRefs {
			stmt4, _ := ds.sqlAddObjectMarkingRef()
			_, err4 := ds.DB.Exec(stmt4, objectID, marking)

			if err4 != nil {
				return 0, fmt.Errorf("database execution error inserting object marking ref: ", err4)
			}
		}
	}

	return objectID, nil
}

// addKillChainPhases
func (ds *Sqlite3DatastoreType) addKillChainPhases(objectID int64, obj properties.KillChainPhasesPropertyType) {
	for _, v := range obj.KillChainPhases {
		var stmt = `INSERT INTO "` + datastore.DB_TABLE_STIX_KILL_CHAIN_PHASES + `" (
			"object_id",
			"kill_chain_name",
			"phase_name"
			)
			values (?, ?, ?)`

		_, err := ds.DB.Exec(stmt, objectID, v.KillChainName, v.PhaseName)

		if err != nil {
			log.Println("ERROR: Database execution error inserting kill chain phases", err)
		}
	}
}

// addIndicator
func (ds *Sqlite3DatastoreType) addIndicator(obj *objects.IndicatorType) error {

	objectID, err := ds.addBaseObject(obj.CommonObjectPropertiesType)
	if err != nil {
		return fmt.Errorf("database error inserting base object: ", err)
	}

	var stmt1 = `INSERT INTO "` + datastore.DB_TABLE_STIX_INDICATOR + `" (
		"object_id",
		"name",
		"description",
		"pattern",
		"valid_from",
		"valid_until"
		)
		values (?, ?, ?, ?, ?, ?)`

	_, err1 := ds.DB.Exec(stmt1,
		objectID,
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
		ds.addKillChainPhases(objectID, obj.KillChainPhasesPropertyType)
	}

	return nil
}
