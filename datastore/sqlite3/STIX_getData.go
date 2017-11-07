// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"database/sql"
	"errors"
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
	"log"
)

func (ds *Sqlite3DatastoreType) GetObject(stixid string) (interface{}, error) {

	// We first need to look at the STIX ID that was passed in to see what type of object it is
	i, err := ds.getIndicator(stixid)
	return i, err
}

func (ds *Sqlite3DatastoreType) getBaseObjects(stixid string) ([]properties.CommonObjectPropertiesType, error) {

	var baseObjects []properties.CommonObjectPropertiesType
	var objectID, specVersion, dateAdded, objectType, id, createdByRef, created, modified, lang string
	var revoked, confidence int

	// By ordering the records in descending order, the newest record will by at location 0
	var getBaseObject = `
		SELECT 
 			object_id,
		 	spec_version,
		 	date_added,
		 	type,
		 	id,
		 	created_by_ref,
		 	created,
		 	modified,
		 	revoked,
		 	confidence,
		 	lang
	   	FROM ` + DB_TABLE_STIX_BASE_OBJECT + ` 
	   	WHERE id = $1
	   	ORDER BY modified DESC`

	// Query the database
	rows, err := ds.DB.Query(getBaseObject, stixid)
	if err != nil {
		log.Println("ERROR: Database execution error quering for base object: ", err)
	}
	defer rows.Close()

	// There might be more than one record that matches that STIX ID. This can happen
	// when the object is versioned. Lets get all of them and just return an array of
	// objects
	for rows.Next() {
		var base properties.CommonObjectPropertiesType
		err := rows.Scan(&objectID, &specVersion, &dateAdded, &objectType, &id, &createdByRef, &created, &modified, &revoked, &confidence, &lang)

		if err != nil {
			log.Println("ERROR: Database scan error for base object: ", err)
		}
		base.SetObjectID(objectID)
		base.SetSpecVersion(specVersion)
		base.SetObjectType(objectType)
		base.SetID(id)
		base.SetCreatedByRef(createdByRef)
		base.SetCreated(created)
		base.SetModified(modified)
		if revoked == 1 {
			base.SetRevoked()
		}
		base.SetConfidence(confidence)
		base.SetLang(lang)

		base.LabelsPropertyType = ds.getBaseObjectLabels(objectID)
		base.ExternalReferencesPropertyType = ds.getBaseObjectExternalReferences(objectID)

		// Capture the base object
		baseObjects = append(baseObjects, base)
	}
	if len(baseObjects) == 0 {
		return nil, errors.New("No Records Found")
	}
	return baseObjects, nil
}

func (ds *Sqlite3DatastoreType) getBaseObjectLabels(objectID string) properties.LabelsPropertyType {
	var objectLabels properties.LabelsPropertyType

	var getLabelForObject = `
		SELECT labels
		FROM ` + DB_TABLE_STIX_LABELS + `
		WHERE object_id = $1`

	rows, err := ds.DB.Query(getLabelForObject, objectID)
	if err != nil {
		log.Println("ERROR: Database execution error quering for labels: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var label string

		err := rows.Scan(&label)
		if err != nil {
			log.Println("ERROR: Database scan error for labels: ", err)
		}
		objectLabels.AddLabel(label)
	}
	return objectLabels
}

func (ds *Sqlite3DatastoreType) getBaseObjectExternalReferences(objectID string) properties.ExternalReferencesPropertyType {
	var extrefs properties.ExternalReferencesPropertyType

	var getExternalReferencesForObject = `
		SELECT 
			source_name,
			description,
			url,
			external_id
		FROM ` + DB_TABLE_STIX_EXTERNAL_REFERENCES + `
		WHERE object_id = $1`

	rows, err := ds.DB.Query(getExternalReferencesForObject, objectID)
	if err != nil {
		log.Println("ERROR: Database execution error quering for external references: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var sourceName, description, url, externalID string
		e := extrefs.NewExternalReference()

		err := rows.Scan(&sourceName, &description, &url, &externalID)
		if err != nil {
			log.Println("ERROR: Database scan error for external references: ", err)
		}
		e.SetSourceName(sourceName)
		e.SetDescription(description)
		e.SetURL(url)
		e.SetExternalID(externalID)
	}
	return extrefs
}

func (ds *Sqlite3DatastoreType) getIndicator(stixid string) (objects.IndicatorType, error) {
	var i objects.IndicatorType

	// Lets first get the base object so we know the objectID
	baseObjects, errBase := ds.getBaseObjects(stixid)
	if errBase != nil {
		return i, errBase
	}

	// TODO this needs to be changed so we can make use of more than one version of an object
	i.CommonObjectPropertiesType = baseObjects[0]

	var getIndicatorObject = `
		SELECT
			name,
			description,
			pattern,
			valid_from,
			valid_until
		FROM ` + DB_TABLE_STIX_INDICATOR + `
		WHERE object_id = $1`

	var name, description, pattern, validFrom, validUntil string
	err := ds.DB.QueryRow(getIndicatorObject, i.ObjectID).Scan(&name, &description, &pattern, &validFrom, &validUntil)
	if err != nil {
		if err == sql.ErrNoRows {
			return i, errors.New("No Records Found")
		}
		log.Println("ERROR", err)
	}
	i.SetName(name)
	i.SetDescription(description)
	i.SetPattern(pattern)
	i.SetValidFrom(validFrom)
	i.SetValidUntil(validUntil)
	return i, nil
}
