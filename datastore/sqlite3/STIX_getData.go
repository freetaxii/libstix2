// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"github.com/freetaxii/libstix2/datastore"
	"github.com/freetaxii/libstix2/objects/properties"
	"log"
)

// /*
// getbaseObject - This method will get a specific base object based on the STIX ID
// and the version (modified timestamp).  This will return a single base object as
// since we are providing both a STIX ID and a version.
// */
// func (ds *DatastoreType) getBaseObject(stixid, version string) (*properties.CommonObjectPropertiesType, error) {

// 	var baseObject properties.CommonObjectPropertiesType
// 	var objectID, specVersion, dateAdded, objectType, id, createdByRef, created, modified, lang string
// 	var revoked, confidence int

// 	stmt, _ := sqlGetBaseObject()

// 	// Query the database
// 	rows, err := ds.DB.Query(stmt, stixid, version)
// 	if err != nil {
// 		return nil, fmt.Errorf("database error querying the base object: ", err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var base properties.CommonObjectPropertiesType
// 		err := rows.Scan(&objectID, &specVersion, &dateAdded, &objectType, &id, &createdByRef, &created, &modified, &revoked, &confidence, &lang)

// 		if err != nil {
// 			log.Println("ERROR: Database scan error for base object: ", err)
// 		}
// 		base.SetObjectID(objectID)
// 		base.SetSpecVersion(specVersion)
// 		base.SetObjectType(objectType)
// 		base.SetID(id)
// 		base.SetCreatedByRef(createdByRef)
// 		base.SetCreated(created)
// 		base.SetModified(modified)
// 		if revoked == 1 {
// 			base.SetRevoked()
// 		}
// 		base.SetConfidence(confidence)
// 		base.SetLang(lang)

// 		base.LabelsPropertyType = ds.getBaseObjectLabels(objectID)
// 		base.ExternalReferencesPropertyType = ds.getBaseObjectExternalReferences(objectID)

// 		// Capture the base object
// 		baseObjects = append(baseObjects, base)
// 	}
// 	if len(baseObjects) == 0 {
// 		return nil, errors.New("No Records Found")
// 	}
// 	return baseObjects, nil
// }

func (ds *DatastoreType) getBaseObjectExternalReferences(objectID int64) properties.ExternalReferencesPropertyType {
	var extrefs properties.ExternalReferencesPropertyType

	var getExternalReferencesForObject = `
		SELECT 
			source_name,
			description,
			url,
			external_id
		FROM ` + datastore.DB_TABLE_STIX_EXTERNAL_REFERENCES + `
		WHERE object_id = $1`

	rows, err := ds.DB.Query(getExternalReferencesForObject, objectID)
	if err != nil {
		log.Println("ERROR: Database execution error quering for external references: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var sourceName, description, url, externalID string
		e, _ := extrefs.GetNewExternalReference()

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
