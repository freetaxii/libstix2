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
	"time"
)

func (ds *Sqlite3DatastoreType) addBaseObject(obj properties.CommonObjectPropertiesType) (string, error) {
	// TODO change, add to object creation
	ver := "2.0"

	var stmt = `INSERT INTO "stix_base_object" (
	 	"object_id", 
	 	"version", 
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

	dateAdded := time.Now().UTC().Format(defs.TIME_RFC_3339_MICRO)
	objectID := "id" + obj.ID + "created" + obj.Created + "modified" + obj.Modified

	h := sha1.New()
	h.Write([]byte(objectID))
	hashID := base64.URLEncoding.EncodeToString(h.Sum(nil))

	_, err := ds.DB.Exec(stmt,
		hashID,
		ver,
		dateAdded,
		obj.MessageType,
		obj.ID,
		obj.CreatedByRef,
		obj.Created,
		obj.Modified,
		obj.Revoked,
		obj.Confidence,
		obj.Lang)

	return hashID, err
}

func (ds *Sqlite3DatastoreType) addIndicatorToDatabase(obj indicator.IndicatorType) error {

	hashID, err := ds.addBaseObject(obj.CommonObjectPropertiesType)
	if err != nil {
		return err
	}

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
		return err
	}

	if obj.KillChainPhases != nil {
		for _, v := range obj.KillChainPhases {
			var stmt2 = `INSERT INTO "kill_chain_phases" (
			"object_id",
			"kill_chain_name",
			"phase_name"
			)
			values (?, ?, ?)`

			_, err2 := ds.DB.Exec(stmt2, hashID, v.KillChainName, v.PhaseName)

			if err2 != nil {
				return err
			}
		}
	}

	if obj.Labels != nil {
		for _, v1 := range obj.Labels {
			var stmt3 = `INSERT INTO "labels" (
			"object_id",
			"labels"
			)
			values (?, ?)`

			_, err3 := ds.DB.Exec(stmt3, hashID, v1)

			if err3 != nil {
				return err
			}
		}
	}
	return nil
}
