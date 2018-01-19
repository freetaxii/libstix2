// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"errors"
	"fmt"
	"github.com/freetaxii/libstix2/defs"
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
	"strings"
	"time"
)

// ----------------------------------------------------------------------
//
// Public Methods
//
// ----------------------------------------------------------------------

/*
GetObject - This method will take in a STIX ID and version timestamp (the
modified timestamp from a STIX object) and return the STIX object.
*/
func (ds *Sqlite3DatastoreType) GetObject(stixid, version string) (interface{}, error) {
	idparts := strings.Split(stixid, "--")

	if ds.StrictSTIXIDs == true {
		if !objects.IsValidID(stixid) {
			return nil, errors.New("get object error, invalid STIX ID")
		}
	}

	if ds.StrictSTIXTypes == true {
		if !objects.IsValidSTIXObject(stixid) {
			return nil, errors.New("get object error, invalid STIX type")
		}
	}

	switch idparts[0] {
	case "indicator":
		return ds.getIndicator(stixid, version)
	}

	return nil, fmt.Errorf("get object error, the following STIX type is not currently supported: ", idparts[0])
}

// ----------------------------------------------------------------------
//
// Private Methods
//
// ----------------------------------------------------------------------

/*
addBaseObject - This method will add the base properties of an object to the
database and return an integer that tracks the record number for parent child
relationships.
*/
func (ds *Sqlite3DatastoreType) addBaseObject(obj *properties.CommonObjectPropertiesType) (int64, error) {
	dateAdded := time.Now().UTC().Format(defs.TIME_RFC_3339_MICRO)

	objectID := ds.Index
	ds.Index++

	stmt1, _ := sqlAddBaseObject()

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
			stmt2, _ := sqlAddObjectLabel()
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
			stmt3, _ := sqlAddExternalReference()

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
			stmt4, _ := sqlAddObjectMarkingRef()
			_, err4 := ds.DB.Exec(stmt4, objectID, marking)

			if err4 != nil {
				return 0, fmt.Errorf("database execution error inserting object marking ref: ", err4)
			}
		}
	}

	return objectID, nil
}

/*
addIndicator - This method will add an indicator to the database.
*/
func (ds *Sqlite3DatastoreType) addIndicator(obj *objects.IndicatorType) error {

	objectID, err := ds.addBaseObject(&obj.CommonObjectPropertiesType)
	if err != nil {
		return err
	}

	stmt, _ := sqlAddIndicatorObject()
	_, err1 := ds.DB.Exec(stmt,
		objectID,
		obj.Name,
		obj.Description,
		obj.Pattern,
		obj.ValidFrom,
		obj.ValidUntil)

	// TODO if there is an error, we probably need to back out all of the INSERTS
	if err1 != nil {
		return fmt.Errorf("database error inserting indicator: ", err)
	}

	if obj.KillChainPhases != nil {
		err2 := ds.addKillChainPhases(objectID, &obj.KillChainPhasesPropertyType)

		if err2 != nil {
			return err2
		}
	}
	return nil
}

/*
addKillChainPhases - This method will add a kill chain phase for a given object
to the database.
*/
func (ds *Sqlite3DatastoreType) addKillChainPhases(objectID int64, obj *properties.KillChainPhasesPropertyType) error {
	for _, v := range obj.KillChainPhases {
		stmt, _ := sqlAddKillChainPhase()
		_, err := ds.DB.Exec(stmt, objectID, v.KillChainName, v.PhaseName)

		if err != nil {
			return fmt.Errorf("database execution error inserting kill chain phase: ", err)
		}
	}
	return nil
}
