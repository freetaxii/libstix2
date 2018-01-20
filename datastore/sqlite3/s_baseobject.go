// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"bytes"
	"fmt"
	"github.com/freetaxii/libstix2/datastore"
	"github.com/freetaxii/libstix2/defs"
	"github.com/freetaxii/libstix2/objects/properties"
	"time"
)

// ----------------------------------------------------------------------
//
// Base Object Table Private Functions and Methods
//
// ----------------------------------------------------------------------

/*
baseProperties - This method will return the base properties for all objects
row_id    = This is a database tracking number
object_id = This is a unique identifier for the STIX object based on its ID + created and modified timestamps
*/
func baseProperties() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
 	"object_id" INTEGER NOT NULL,`
}

/*
baseObjectProperties - This method will return the the common properties
spec_version = STIX specification version
date_added = TAXII, the date the object was added to the TAXII server
*/
func baseObjectProperties() string {
	return baseProperties() + `
 	"spec_version" TEXT NOT NULL,
 	"date_added" TEXT NOT NULL,
 	"type" TEXT NOT NULL,
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
sqlAddBaseObject - This function will return an SQL statement that will add the
base object properties to the database.
*/
func sqlAddBaseObject() (string, error) {
	tblBaseObj := datastore.DB_TABLE_STIX_BASE_OBJECT

	/*
		INSERT INTO
			s_base_object (
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
			values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	*/

	var s bytes.Buffer
	s.WriteString("INSERT INTO ")
	s.WriteString(tblBaseObj)
	s.WriteString(" (")
	s.WriteString("\"object_id\", ")
	s.WriteString("\"spec_version\", ")
	s.WriteString("\"date_added\", ")
	s.WriteString("\"type\", ")
	s.WriteString("\"id\", ")
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
			stmt2, _ := sqlAddLabel()
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

// ----------------------------------------------------------------------
//
// Labels Table
//
// ----------------------------------------------------------------------

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
sqlAddLabel - This function will return an SQL statement that will add a
label to the database for a given object.
*/
func sqlAddLabel() (string, error) {
	tblLabels := datastore.DB_TABLE_STIX_LABELS

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

// ----------------------------------------------------------------------
//
// External References Table
//
// ----------------------------------------------------------------------

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
sqlAddExternalReference - This function will return an SQL statement that will add
an external reference to the database for a given object.
*/
func sqlAddExternalReference() (string, error) {
	tblExtRef := datastore.DB_TABLE_STIX_EXTERNAL_REFERENCES

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

// ----------------------------------------------------------------------
//
// Object Marking Refs Table
//
// ----------------------------------------------------------------------

/*
commonObjectMarkingRefsProperties - This method will return the properties for object markings
Used by: All SDOs and SROs
*/
func commonObjectMarkingRefsProperties() string {
	return baseProperties() + `
	"object_marking_refs" TEXT NOT NULL
	`
}

/*
sqlAddObjectMarkingRef - This function will return an SQL statement that will add
an object marking ref to the database for a given object.
*/
func sqlAddObjectMarkingRef() (string, error) {
	tblObjMarking := datastore.DB_TABLE_STIX_OBJECT_MARKING_REFS

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
