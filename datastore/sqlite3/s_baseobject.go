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
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
//
// Base Object Tables Private Functions
// Table property names and SQL statements
//
// ----------------------------------------------------------------------

/*
baseProperties - This method will return the base properties for all objects
row_id    = This is a database tracking number
object_id = This is a unique integer for the STIX object
*/
func baseProperties() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
 	"object_id" INTEGER NOT NULL,`
}

/*
baseObjectProperties - This method will return the the common properties
date_added = TAXII, the date the object was added to the TAXII server
type = STIX Object Type
id = STIX ID in type--uuidv4 format
spec_version = STIX specification version
created_by_ref = A STIX ID that points to an Identity Object
created = RFC 3339 timestamp with microsecond precision
modified = RFC 3339 timestamp with microsecond precision
revoked = a boolean of true / false
confidence = an integer 0-100
lang = An ISO language code
*/
func baseObjectProperties() string {
	return baseProperties() + `
 	"date_added" TEXT NOT NULL,
 	"type" TEXT NOT NULL,
 	"id" TEXT NOT NULL,
 	"spec_version" TEXT NOT NULL,
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
	return baseProperties() + `
	"label" TEXT NOT NULL
	`
}

/*
commonExternalReferencesProperties - This method will return the properties for external references
Used by: All SDOs and SROs
*/
func commonExternalReferencesProperties() string {
	return baseProperties() + `
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
	return baseProperties() + `
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
sqlGetBaseObjectIndex - This function will return an SQL statement that will
get the last object_id in the base objects table to be used as an index for
new records.
*/
func sqlGetBaseObjectIndex() (string, error) {
	tblBaseObj := DB_TABLE_STIX_BASE_OBJECT

	/*
		SELECT
			object_id
		FROM
			t_base_object
		ORDER BY
			object_id DESC LIMIT 1
	*/

	var s bytes.Buffer
	s.WriteString("SELECT ")
	s.WriteString("object_id ")
	s.WriteString("FROM ")
	s.WriteString(tblBaseObj)
	s.WriteString(" ORDER BY object_id DESC LIMIT 1")

	return s.String(), nil
}

/*
getBaseObjectIndex - This method will return the last object index value from the
database base object table.
*/
func (ds *Datastore) getBaseObjectIndex() (int, error) {
	var index int
	sqlStmt, _ := sqlGetBaseObjectIndex()

	err := ds.DB.QueryRow(sqlStmt).Scan(&index)
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
// addBaseObject
//
// ----------------------------------------------------------------------

/*
sqlAddBaseObject - This function will return an SQL statement that will add the
base object properties to the database.
*/
func sqlAddBaseObject() (string, error) {
	tblBaseObj := DB_TABLE_STIX_BASE_OBJECT

	/*
		INSERT INTO
			s_base_object (
				"object_id",
				"date_added",
				"type",
				"id",
				"spec_version",
				"created_by_ref",
				"created",
				"modified",
				"revoked",
				"confidence",
				"lang"
			)
			values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	*/

	var s bytes.Buffer
	s.WriteString("INSERT INTO ")
	s.WriteString(tblBaseObj)
	s.WriteString(" (")
	s.WriteString("\"object_id\", ")
	s.WriteString("\"date_added\", ")
	s.WriteString("\"type\", ")
	s.WriteString("\"id\", ")
	s.WriteString("\"spec_version\", ")
	s.WriteString("\"created_by_ref\", ")
	s.WriteString("\"created\", ")
	s.WriteString("\"modified\", ")
	s.WriteString("\"revoked\", ")
	s.WriteString("\"confidence\", ")
	s.WriteString("\"lang\") ")
	s.WriteString("values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

	return s.String(), nil
}

/*
addBaseObject - This method will add the base properties of an object to the
database and return an integer that tracks the record number for parent child
relationships.
*/
func (ds *Datastore) addBaseObject(obj *properties.CommonObjectProperties) (int, error) {
	dateAdded := time.Now().UTC().Format(defs.TIME_RFC_3339_MICRO)

	objectID := ds.Cache.BaseObjectIDIndex
	ds.Cache.BaseObjectIDIndex++

	ds.Logger.Debugln("DEBUG: Adding Base Object to datastore with object ID", objectID, "and STIX ID", obj.ID)

	stmt1, _ := sqlAddBaseObject()

	_, err1 := ds.DB.Exec(stmt1,
		objectID,
		dateAdded,
		obj.ObjectType,
		obj.ID,
		obj.SpecVersion,
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

// ----------------------------------------------------------------------
//
// Base Object Table Private Functions
// getBaseObject
//
// ----------------------------------------------------------------------

/*
sqlGetbaseObject - This function will return an SQL statement that will get a
specific base object from the database.
*/
func sqlGetBaseObject() (string, error) {
	tblBaseObj := DB_TABLE_STIX_BASE_OBJECT
	tblLabels := DB_TABLE_STIX_LABELS

	/*
		SELECT
			s_base_object.object_id,
			s_base_object.date_added,
			s_base_object.type,
			s_base_object.id,
			s_base_object.spec_version,
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

	var s bytes.Buffer
	s.WriteString("SELECT ")
	s.WriteString(tblBaseObj)
	s.WriteString(".object_id, ")
	s.WriteString(tblBaseObj)
	s.WriteString(".date_added, ")
	s.WriteString(tblBaseObj)
	s.WriteString(".type, ")
	s.WriteString(tblBaseObj)
	s.WriteString(".id, ")
	s.WriteString(tblBaseObj)
	s.WriteString(".spec_version, ")
	s.WriteString(tblBaseObj)
	s.WriteString(".created_by_ref, ")
	s.WriteString(tblBaseObj)
	s.WriteString(".created, ")
	s.WriteString(tblBaseObj)
	s.WriteString(".modified, ")
	s.WriteString(tblBaseObj)
	s.WriteString(".revoked, ")
	s.WriteString(tblBaseObj)
	s.WriteString(".confidence, ")
	s.WriteString(tblBaseObj)
	s.WriteString(".lang, ")
	s.WriteString("group_concat(")
	s.WriteString(tblLabels)
	s.WriteString(".label) ")
	s.WriteString("FROM ")
	s.WriteString(tblBaseObj)
	s.WriteString(" LEFT JOIN ")
	s.WriteString(tblLabels)
	s.WriteString(" ON ")
	s.WriteString(tblBaseObj)
	s.WriteString(".object_id = ")
	s.WriteString(tblLabels)
	s.WriteString(".object_id ")
	s.WriteString("WHERE ")
	s.WriteString(tblBaseObj)
	s.WriteString(".id = $1 AND ")
	s.WriteString(tblBaseObj)
	s.WriteString(".modified = $2")

	return s.String(), nil
}

/*
getbaseObject - This method will get a specific base object based on the STIX ID
and the version (modified timestamp).  This method is most often called from
a get method on a STIX object (for example: getIndicator).
*/
func (ds *Datastore) getBaseObject(stixid, version string) (*properties.CommonObjectProperties, error) {

	var baseObject properties.CommonObjectProperties
	var objectID int
	var specVersion, dateAdded, objectType, id, createdByRef, created, modified, lang string

	// Since not every object will have a label, and since we are using group_concat
	// we need to define the label as a pointer so it can be a null value.
	var label *string
	var revoked, confidence int

	stmt, _ := sqlGetBaseObject()
	err := ds.DB.QueryRow(stmt, stixid, version).Scan(&objectID, &specVersion, &dateAdded, &objectType, &id, &createdByRef, &created, &modified, &revoked, &confidence, &lang, &label)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no base object record found")
		}
		return nil, fmt.Errorf("database execution error getting base object: ", err)
	}
	baseObject.SetObjectID(objectID)
	baseObject.SetObjectType(objectType)
	baseObject.SetID(id)
	baseObject.SetSpecVersion(specVersion)
	baseObject.SetCreatedByRef(createdByRef)
	baseObject.SetCreated(created)
	baseObject.SetModified(modified)
	if revoked == 1 {
		baseObject.SetRevoked()
	}
	baseObject.SetConfidence(confidence)
	baseObject.SetLang(lang)
	if label != nil {
		baseObject.AddLabel(*label)
	}

	externalRefData, err1 := ds.getExternalReferences(objectID)
	if err1 != nil {
		return nil, err1
	}
	baseObject.ExternalReferencesProperty = *externalRefData

	return &baseObject, nil
}

// ----------------------------------------------------------------------
//
// Labels Table Private Functions
// addLabel
//
// ----------------------------------------------------------------------

/*
sqlAddLabel - This function will return an SQL statement that will add a
label to the database for a given object.
*/
func sqlAddLabel() (string, error) {
	tblLabels := DB_TABLE_STIX_LABELS

	/*
		INSERT INTO
			s_labels (
				"object_id",
				"label"
			)
			values (?, ?)
	*/

	var s bytes.Buffer
	s.WriteString("INSERT INTO ")
	s.WriteString(tblLabels)
	s.WriteString(" (\"object_id\", \"label\") values (?, ?)")

	return s.String(), nil
}

/*
addLabel - This method will add a label to the database for a specific object ID.
*/
func (ds *Datastore) addLabel(objectID int, label string) error {
	stmt, _ := sqlAddLabel()
	_, err := ds.DB.Exec(stmt, objectID, label)

	if err != nil {
		return fmt.Errorf("database execution error inserting object label: ", err)
	}
	return nil
}

// ----------------------------------------------------------------------
//
// External References Table Private Functions
// addExternalReference
//
// ----------------------------------------------------------------------

/*
sqlAddExternalReference - This function will return an SQL statement that will add
an external reference to the database for a given object.
*/
func sqlAddExternalReference() (string, error) {
	tblExtRef := DB_TABLE_STIX_EXTERNAL_REFERENCES

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

	var s bytes.Buffer
	s.WriteString("INSERT INTO ")
	s.WriteString(tblExtRef)
	s.WriteString(" (")
	s.WriteString("\"object_id\", ")
	s.WriteString("\"source_name\", ")
	s.WriteString("\"description\", ")
	s.WriteString("\"url\", ")
	s.WriteString("\"external_id\") ")
	s.WriteString("values (?, ?, ?, ?, ?)")

	return s.String(), nil
}

/*
addExternalReference - This method will add an external reference to the
database for a specific object ID.
*/
func (ds *Datastore) addExternalReference(objectID int, reference properties.ExternalReference) error {
	stmt, _ := sqlAddExternalReference()

	_, err := ds.DB.Exec(stmt,
		objectID,
		reference.SourceName,
		reference.Description,
		reference.URL,
		reference.ExternalID)

	if err != nil {
		return fmt.Errorf("database execution error inserting external reference: ", err)
	}
	return nil
}

// ----------------------------------------------------------------------
//
// External References Table Private Functions
// getExternalReference
//
// ----------------------------------------------------------------------

/*
sqlGetExternalReference - This function will return an SQL statement that will
get an external reference from the database for a specific object ID.
*/
func sqlGetExternalReference() (string, error) {
	tblExtRef := DB_TABLE_STIX_EXTERNAL_REFERENCES

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

	var s bytes.Buffer
	s.WriteString("SELECT ")
	s.WriteString("source_name, ")
	s.WriteString("description, ")
	s.WriteString("url, ")
	s.WriteString("external_id ")
	s.WriteString("FROM ")
	s.WriteString(tblExtRef)
	s.WriteString(" WHERE object_id = $1")

	return s.String(), nil
}

/*
getExternalReferences - This method will return all external references that are
part of a specific object ID.
*/
func (ds *Datastore) getExternalReferences(objectID int) (*properties.ExternalReferencesProperty, error) {
	var extrefs properties.ExternalReferencesProperty
	stmt, _ := sqlGetExternalReference()

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
// addObjectMarkingRef
//
// ----------------------------------------------------------------------

/*
sqlAddObjectMarkingRef - This function will return an SQL statement that will add
an object marking ref to the database for a given object.
*/
func sqlAddObjectMarkingRef() (string, error) {
	tblObjMarking := DB_TABLE_STIX_OBJECT_MARKING_REFS

	/*
		INSERT INTO
			s_object_marking_refs (
				"object_id",
				"object_marking_refs"
			)
			values (?, ?)
	*/

	var s bytes.Buffer
	s.WriteString("INSERT INTO ")
	s.WriteString(tblObjMarking)
	s.WriteString(" (\"object_id\", \"object_marking_refs\") values (?, ?)")

	return s.String(), nil
}

/*
addObjectMarkingRef - This method will add an object marking ref to the
database for a specific object ID.
*/
func (ds *Datastore) addObjectMarkingRef(objectID int, marking string) error {
	stmt, _ := sqlAddObjectMarkingRef()
	_, err := ds.DB.Exec(stmt, objectID, marking)

	if err != nil {
		return fmt.Errorf("database execution error inserting object marking ref: ", err)
	}
	return nil
}
