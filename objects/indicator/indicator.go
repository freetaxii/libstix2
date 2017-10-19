// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package indicator

import (
	"crypto/sha1"
	"database/sql"
	"encoding/base64"
	"github.com/freetaxii/libstix2/objects/common/properties"
	"github.com/freetaxii/libstix2/objects/defs"
	"time"
)

// ----------------------------------------------------------------------
// Define Indicator Type
// ----------------------------------------------------------------------

/*
IndicatorType defines all of the properties associated with the STIX Indicator
SDO. All of the methods not defined local to this type are inherited from the
individual properties.
*/
type IndicatorType struct {
	properties.CommonObjectPropertiesType
	properties.NamePropertyType
	properties.DescriptionPropertyType
	Pattern    string `json:"pattern,omitempty"`
	ValidFrom  string `json:"valid_from,omitempty"`
	ValidUntil string `json:"valid_until,omitempty"`
	properties.KillChainPhasesPropertyType
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// New - This function will create a new indicator object.
func New() IndicatorType {
	var obj IndicatorType
	obj.InitNewObject("indicator")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - IndicatorType
// ----------------------------------------------------------------------

// SetPattern - This method will take in a string value representing a complete
// and valid STIX pattern and set the pattern property to that value.
func (ezt *IndicatorType) SetPattern(s string) {
	ezt.Pattern = s
}

// SetValidFrom - This method will take in a timestamp in either time.Time or
// string format and will set the valid from property to that value.
func (ezt *IndicatorType) SetValidFrom(t interface{}) {
	ts := ezt.VerifyTimestamp(t)
	ezt.ValidFrom = ts
}

// SetValidUntil - This method will take in a timestamp in either time.Time or
// string format and will set the valid until property to that value.
func (ezt *IndicatorType) SetValidUntil(t interface{}) {
	ts := ezt.VerifyTimestamp(t)

	// TODO check to make sure this is later than the vaild_from
	ezt.ValidUntil = ts
}

func (ezt *IndicatorType) AddToDatabase(db *sql.DB, ver string) error {

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
	objectID := "id" + ezt.ID + "created" + ezt.Created + "modified" + ezt.Modified

	h := sha1.New()
	h.Write([]byte(objectID))
	hashID := base64.URLEncoding.EncodeToString(h.Sum(nil))

	_, err := db.Exec(stmt,
		hashID,
		ver,
		dateAdded,
		ezt.MessageType,
		ezt.ID,
		ezt.CreatedByRef,
		ezt.Created,
		ezt.Modified,
		ezt.Revoked,
		ezt.Confidence,
		ezt.Lang)

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

	_, err1 := db.Exec(stmt1,
		hashID,
		ezt.Name,
		ezt.Description,
		ezt.Pattern,
		ezt.ValidFrom,
		ezt.ValidUntil)

	// TODO if there is an error, we probably need to back out all of the INSERTS
	if err1 != nil {
		return err
	}

	if ezt.KillChainPhases != nil {
		for _, v := range ezt.KillChainPhases {
			var stmt2 = `INSERT INTO "kill_chain_phases" (
			"object_id",
			"kill_chain_name",
			"phase_name"
			)
			values (?, ?, ?)`

			_, err2 := db.Exec(stmt2, hashID, v.KillChainName, v.PhaseName)

			if err2 != nil {
				return err
			}
		}
	}

	if ezt.Labels != nil {
		for _, v1 := range ezt.Labels {
			var stmt3 = `INSERT INTO "labels" (
			"object_id",
			"labels"
			)
			values (?, ?)`

			_, err3 := db.Exec(stmt3, hashID, v1)

			if err3 != nil {
				return err
			}
		}
	}

	return nil
}
