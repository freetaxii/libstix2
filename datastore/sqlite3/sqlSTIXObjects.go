// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"bytes"
	"github.com/freetaxii/libstix2/datastore"
	"log"
)

// ----------------------------------------------------------------------
//
// Private Methods
//
// ----------------------------------------------------------------------

/*
sqlAddBaseObject - This method will return an SQL statement that will
add the base object properties to the database.
*/
func (ds *Sqlite3DatastoreType) sqlAddBaseObject() (string, error) {
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
	s.WriteString("\"object_id\", \"spec_version\", \"date_added\", ")
	s.WriteString("\"type\", \"id\", \"created_by_ref\", \"created\", ")
	s.WriteString("\"modified\", \"revoked\", \"confidence\", \"lang\") ")
	s.WriteString("values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

	if ds.LogLevel >= 5 {
		log.Println("DEBUG: Returning SQL statement:", s.String())
	}

	return s.String(), nil
}

/*
sqlAddObjectLabel - This method will return an SQL statement that will add a
label to the database for a given object.
*/
func (ds *Sqlite3DatastoreType) sqlAddObjectLabel() (string, error) {
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

	if ds.LogLevel >= 5 {
		log.Println("DEBUG: Returning SQL statement:", s.String())
	}

	return s.String(), nil
}

/*
sqlAddExternalReference - This method will return an SQL statement that will add
an external reference to the database for a given object.
*/
func (ds *Sqlite3DatastoreType) sqlAddExternalReference() (string, error) {
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
	s.WriteString(" (\"object_id\", \"source_name\", \"description\", \"url\", \"external_id\") values (?, ?, ?, ?, ?)")

	if ds.LogLevel >= 5 {
		log.Println("DEBUG: Returning SQL statement:", s.String())
	}

	return s.String(), nil
}

/*
sqlAddObjectMarkingRef - This method will return an SQL statement that will add
an object marking ref to the database for a given object.
*/
func (ds *Sqlite3DatastoreType) sqlAddObjectMarkingRef() (string, error) {
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

	if ds.LogLevel >= 5 {
		log.Println("DEBUG: Returning SQL statement:", s.String())
	}

	return s.String(), nil
}
