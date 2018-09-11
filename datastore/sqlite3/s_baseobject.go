// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/freetaxii/libstix2/defs"
	"github.com/freetaxii/libstix2/objects/baseobject"
)

// ----------------------------------------------------------------------
//
// Base Object Tables Private Functions
// Table property names and SQL statements
//
// ----------------------------------------------------------------------

/*
baseDBProperties - This method will return the base properties for all objects
row_id    = This is a database tracking number
object_id = This is a unique integer for the STIX object
*/
func baseDBProperties() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
 	"object_id" INTEGER NOT NULL,`
}

/*
baseObjectProperties - This method will return the the common properties
date_added = TAXII, the date the object was added to the TAXII server
type = STIX Object Type
spec_version = STIX specification version
id = STIX ID in type--uuidv4 format
created_by_ref = A STIX ID that points to an Identity Object
created = RFC 3339 timestamp with microsecond precision
modified = RFC 3339 timestamp with microsecond precision
revoked = a boolean of true / false
confidence = an integer 0-100
lang = An ISO language code
*/
func baseObjectProperties() string {
	return baseDBProperties() + `
 	"date_added" TEXT NOT NULL,
 	"type" TEXT NOT NULL,
 	"spec_version" TEXT NOT NULL,
 	"id" TEXT NOT NULL,
 	"created_by_ref" TEXT,
 	"created" TEXT NOT NULL,
 	"modified" TEXT NOT NULL,
 	"revoked" INTEGER(1,0) DEFAULT 0,
 	"confidence" INTEGER(3,0),
 	"lang" TEXT`
	// labels
	// external_references
	// object_marking_refs
}

/*
commonLabelsProperties - This method will return the properties for labels
Used by: All SDOs and SROs
*/
func commonLabelsProperties() string {
	return baseDBProperties() + `
	"label" TEXT NOT NULL
	`
}

/*
commonExternalReferencesProperties - This method will return the properties for external references
Used by: All SDOs and SROs
*/
func commonExternalReferencesProperties() string {
	return baseDBProperties() + `
	"source_name" TEXT NOT NULL,
	"description" TEXT,
	"url" TEXT,
	"external_id" TEXT
	`
}

/*
commonObjectMarkingRefsProperties - This method will return the properties for object markings
Used by: All SDOs and SROs
*/
func commonObjectMarkingRefsProperties() string {
	return baseDBProperties() + `
	"object_marking_refs" TEXT NOT NULL
	`
}

// ----------------------------------------------------------------------
//
// Base Object Table Private Functions
// getBaseObjectIndex
//
// ----------------------------------------------------------------------

/*
getBaseObjectIndex - This method will return the last object index value from the
database base object table.
*/
func (ds *Store) getBaseObjectIndex() (int, error) {
	var index int

	// Create SQL Statement
	/*
		SELECT
			object_id
		FROM
			t_base_object
		ORDER BY
			object_id DESC LIMIT 1
	*/
	tblBaseObj := DB_TABLE_STIX_BASE_OBJECT
	var sqlstmt bytes.Buffer
	sqlstmt.WriteString("SELECT ")
	sqlstmt.WriteString("object_id ")
	sqlstmt.WriteString("FROM ")
	sqlstmt.WriteString(tblBaseObj)
	sqlstmt.WriteString(" ORDER BY object_id DESC LIMIT 1")
	stmt := sqlstmt.String()

	// Make SQL Call
	err := ds.DB.QueryRow(stmt).Scan(&index)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("no base object record found")
		}
		return 0, fmt.Errorf("database execution error getting base index: ", err)
	}

	return index, nil
}

// ----------------------------------------------------------------------
//
// Base Object Table Private Functions
// addBaseObject()
// getBaseObject()
//
// ----------------------------------------------------------------------

/*
addBaseObject - This method will add the base properties of an object to the
database and return an integer that tracks the record number for parent child
relationships.
*/
func (ds *Store) addBaseObject(obj *baseobject.CommonObjectProperties) (int, error) {
	dateAdded := time.Now().UTC().Format(defs.TIME_RFC_3339_MICRO)

	objectID := ds.Cache.BaseObjectIDIndex
	ds.Cache.BaseObjectIDIndex++

	ds.Logger.Debugln("DEBUG: Adding Base Object to datastore with object ID", objectID, "and STIX ID", obj.ID)

	// Create SQL Statement
	/*
		INSERT INTO
			s_base_object (
				"object_id",
				"date_added",
				"type",
				"spec_version",
				"id",
				"created_by_ref",
				"created",
				"modified",
				"revoked",
				"confidence",
				"lang"
			)
			values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	*/
	tblBaseObj := DB_TABLE_STIX_BASE_OBJECT
	var sqlstmt bytes.Buffer
	sqlstmt.WriteString("INSERT INTO ")
	sqlstmt.WriteString(tblBaseObj)
	sqlstmt.WriteString(" (object_id, date_added, type, spec_version, id, created_by_ref, created, modified, revoked, confidence, lang) ")
	sqlstmt.WriteString("values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	stmt1 := sqlstmt.String()

	// Make SQL Call
	_, err1 := ds.DB.Exec(stmt1,
		objectID,
		dateAdded,
		obj.ObjectType,
		obj.SpecVersion,
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
			err2 := ds.addLabel(objectID, label)
			if err2 != nil {
				return 0, err2
			}
		}
	}

	// ----------------------------------------------------------------------
	// Add External References
	// ----------------------------------------------------------------------
	if obj.ExternalReferences != nil {
		for _, reference := range obj.ExternalReferences {
			err3 := ds.addExternalReference(objectID, reference)
			if err3 != nil {
				return 0, err3
			}
		}
	}

	// ----------------------------------------------------------------------
	// Add External References
	// ----------------------------------------------------------------------
	if obj.ObjectMarkingRefs != nil {
		for _, marking := range obj.ObjectMarkingRefs {
			err4 := ds.addObjectMarkingRef(objectID, marking)

			if err4 != nil {
				return 0, err4
			}
		}
	}

	return objectID, nil
}

/*
getbaseObject - This method will get a specific base object based on the STIX ID
and the version (modified timestamp).  This method is most often called from
a get method on a STIX object (for example: getIndicator).
*/
func (ds *Store) getBaseObject(stixid, version string) (*baseobject.CommonObjectProperties, error) {

	var baseObj baseobject.CommonObjectProperties
	var objectID int
	var dateAdded, objectType, specVersion, id, createdByRef, created, modified, lang string

	// Since not every object will have a label, and since we are using group_concat
	// we need to define the label as a pointer so it can be a null value.
	var label *string
	var revoked, confidence int

	// Create SQL Statement
	/*
		SELECT
			s_base_object.object_id,
			s_base_object.date_added,
			s_base_object.type,
			s_base_object.spec_version,
			s_base_object.id,
			s_base_object.created_by_ref,
			s_base_object.created,
			s_base_object.modified,
			s_base_object.revoked,
			s_base_object.confidence,
			s_base_object.lang,
			group_concat(s_labels.label)
		FROM
			s_base_object
		JOIN
			s_labels ON
			s_base_object.object_id = s_labels.object_id
		WHERE
			s_base_object.id = $1 AND
			s_base_object.modified = $2
	*/
	tblBaseObj := DB_TABLE_STIX_BASE_OBJECT
	tblLabels := DB_TABLE_STIX_LABELS
	var sqlstmt bytes.Buffer
	sqlstmt.WriteString("SELECT ")
	sqlstmt.WriteString(tblBaseObj)
	sqlstmt.WriteString(".object_id, ")
	sqlstmt.WriteString(tblBaseObj)
	sqlstmt.WriteString(".date_added, ")
	sqlstmt.WriteString(tblBaseObj)
	sqlstmt.WriteString(".type, ")
	sqlstmt.WriteString(tblBaseObj)
	sqlstmt.WriteString(".spec_version, ")
	sqlstmt.WriteString(tblBaseObj)
	sqlstmt.WriteString(".id, ")
	sqlstmt.WriteString(tblBaseObj)
	sqlstmt.WriteString(".created_by_ref, ")
	sqlstmt.WriteString(tblBaseObj)
	sqlstmt.WriteString(".created, ")
	sqlstmt.WriteString(tblBaseObj)
	sqlstmt.WriteString(".modified, ")
	sqlstmt.WriteString(tblBaseObj)
	sqlstmt.WriteString(".revoked, ")
	sqlstmt.WriteString(tblBaseObj)
	sqlstmt.WriteString(".confidence, ")
	sqlstmt.WriteString(tblBaseObj)
	sqlstmt.WriteString(".lang, ")
	sqlstmt.WriteString("group_concat(")
	sqlstmt.WriteString(tblLabels)
	sqlstmt.WriteString(".label) ")
	sqlstmt.WriteString("FROM ")
	sqlstmt.WriteString(tblBaseObj)
	sqlstmt.WriteString(" LEFT JOIN ")
	sqlstmt.WriteString(tblLabels)
	sqlstmt.WriteString(" ON ")
	sqlstmt.WriteString(tblBaseObj)
	sqlstmt.WriteString(".object_id = ")
	sqlstmt.WriteString(tblLabels)
	sqlstmt.WriteString(".object_id ")
	sqlstmt.WriteString("WHERE ")
	sqlstmt.WriteString(tblBaseObj)
	sqlstmt.WriteString(".id = $1 AND ")
	sqlstmt.WriteString(tblBaseObj)
	sqlstmt.WriteString(".modified = $2")
	stmt := sqlstmt.String()

	// Make SQL Call
	err := ds.DB.QueryRow(stmt, stixid, version).Scan(&objectID, &dateAdded, &objectType, &specVersion, &id, &createdByRef, &created, &modified, &revoked, &confidence, &lang, &label)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no base object record found")
		}
		return nil, fmt.Errorf("database execution error getting base object: ", err)
	}
	baseObj.SetObjectID(objectID)
	baseObj.SetObjectType(objectType)
	baseObj.SetSpecVersion(specVersion)
	baseObj.SetID(id)
	baseObj.SetCreatedByRef(createdByRef)
	baseObj.SetCreated(created)
	baseObj.SetModified(modified)
	if revoked == 1 {
		baseObj.SetRevoked()
	}
	baseObj.SetConfidence(confidence)
	baseObj.SetLang(lang)
	if label != nil {
		baseObj.AddLabel(*label)
	}

	externalRefData, err1 := ds.getExternalReferences(objectID)
	if err1 != nil {
		return nil, err1
	}
	baseObj.ExternalReferencesProperty = *externalRefData

	return &baseObj, nil
}

// ----------------------------------------------------------------------
//
// Labels Table Private Functions
// addLabel()
//
// ----------------------------------------------------------------------

/*
addLabel - This method will add a label to the database for a specific object ID.
*/
func (ds *Store) addLabel(objectID int, label string) error {

	// Create SQL Statement
	/*
		INSERT INTO
			s_labels (
				"object_id",
				"label"
			)
			values (?, ?)
	*/
	tblLabels := DB_TABLE_STIX_LABELS
	var sqlstmt bytes.Buffer
	sqlstmt.WriteString("INSERT INTO ")
	sqlstmt.WriteString(tblLabels)
	sqlstmt.WriteString(" (object_id, label) values (?, ?)")
	stmt := sqlstmt.String()

	// Make SQL Call
	_, err := ds.DB.Exec(stmt, objectID, label)

	if err != nil {
		return fmt.Errorf("database execution error inserting object label: ", err)
	}
	return nil
}

// ----------------------------------------------------------------------
//
// External References Table Private Functions
// addExternalReference()
// getExternalReference()
//
// ----------------------------------------------------------------------

/*
addExternalReference - This method will add an external reference to the
database for a specific object ID.
*/
func (ds *Store) addExternalReference(objectID int, extref baseobject.ExternalReference) error {

	// Create SQL Statement
	/*
		INSERT INTO
			s_external_references (
				"object_id",
				"source_name",
				"description",
				"url",
				"external_id"
			)
			values (?, ?, ?, ?, ?)
	*/
	tblExtRef := DB_TABLE_STIX_EXTERNAL_REFERENCES
	var sqlstmt bytes.Buffer
	sqlstmt.WriteString("INSERT INTO ")
	sqlstmt.WriteString(tblExtRef)
	sqlstmt.WriteString(" (object_id, source_name, description, url, external_id) ")
	sqlstmt.WriteString("values (?, ?, ?, ?, ?)")
	stmt := sqlstmt.String()

	// Make SQL Call
	_, err := ds.DB.Exec(stmt,
		objectID,
		extref.SourceName,
		extref.Description,
		extref.URL,
		extref.ExternalID)

	if err != nil {
		return fmt.Errorf("database execution error inserting external reference: ", err)
	}
	return nil
}

/*
getExternalReferences - This method will return all external references that are
part of a specific object ID.
*/
func (ds *Store) getExternalReferences(objectID int) (*baseobject.ExternalReferencesProperty, error) {
	var extrefs baseobject.ExternalReferencesProperty

	// Create SQL Statement
	/*
		SELECT
			source_name,
			description,
			url,
			external_id
		FROM
			s_external_references
		WHERE
			object_id = $1
	*/
	tblExtRef := DB_TABLE_STIX_EXTERNAL_REFERENCES
	var sqlstmt bytes.Buffer
	sqlstmt.WriteString("SELECT ")
	sqlstmt.WriteString("source_name, description, url, external_id ")
	sqlstmt.WriteString("FROM ")
	sqlstmt.WriteString(tblExtRef)
	sqlstmt.WriteString(" WHERE object_id = $1")
	stmt := sqlstmt.String()

	// Make SQL Call
	rows, err := ds.DB.Query(stmt, objectID)
	if err != nil {
		return nil, fmt.Errorf("database execution error getting external reference: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var sourceName, description, url, externalID string
		e, _ := extrefs.NewExternalReference()

		if err := rows.Scan(&sourceName, &description, &url, &externalID); err != nil {
			rows.Close()
			return nil, fmt.Errorf("database scan error getting external references: ", err)
		}
		e.SetSourceName(sourceName)
		e.SetDescription(description)
		e.SetURL(url)
		e.SetExternalID(externalID)
	}

	// Errors can cause the rows.Next() to exit prematurely, if this happens lets
	// check for the error and handle it.
	if err := rows.Err(); err != nil {
		rows.Close()
		return nil, fmt.Errorf("database rows error getting external references: ", err)
	}

	return &extrefs, nil
}

// ----------------------------------------------------------------------
//
// Object Marking Refs Table Private Functions
// addObjectMarkingRef()
//
// ----------------------------------------------------------------------

/*
addObjectMarkingRef - This method will add an object marking ref to the
database for a specific object ID.
*/
func (ds *Store) addObjectMarkingRef(objectID int, marking string) error {

	// Create SQL Statement
	/*
		INSERT INTO
			s_object_marking_refs (
				"object_id",
				"object_marking_refs"
			)
			values (?, ?)
	*/
	tblObjMarking := DB_TABLE_STIX_OBJECT_MARKING_REFS
	var sqlstmt bytes.Buffer
	sqlstmt.WriteString("INSERT INTO ")
	sqlstmt.WriteString(tblObjMarking)
	sqlstmt.WriteString(" (object_id, object_marking_refs) values (?, ?)")
	stmt := sqlstmt.String()

	// Make SQL Call
	_, err := ds.DB.Exec(stmt, objectID, marking)

	if err != nil {
		return fmt.Errorf("database execution error inserting object marking ref: ", err)
	}
	return nil
}
