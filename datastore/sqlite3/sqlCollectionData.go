// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"bytes"
	"errors"
	"github.com/freetaxii/libstix2/common/timestamp"
	"github.com/freetaxii/libstix2/datastore"
	"github.com/freetaxii/libstix2/objects"
	//"log"
)

/*
sqlObjectList - This method will return an SQL statement that will
return a list of objects from a given collection. It will use the query struct to
determine the requirements and parameters for the where clause of the SQL
statement. A byte array is used instead of sting concatenation as it is the most
efficient way to do string concatenation in Go.
*/
func (ds *Sqlite3DatastoreType) sqlObjectList(query datastore.QueryType) (string, error) {
	tblColData := datastore.DB_TABLE_TAXII_COLLECTION_DATA
	tblBaseObj := datastore.DB_TABLE_STIX_BASE_OBJECT

	whereQuery, err := ds.sqlCollectionDataQueryOptions(query)

	// If an error is found, that means a query parameter was passed incorrectly
	// and we should return an error versus just skipping the option.
	if err != nil {
		return "", err
	}

	/*
		SELECT
			t_collection_data.date_added,
			t_collection_data.stix_id,
			s_base_object.modified,
			s_base_object.spec_version
		FROM
			t_collection_data
		JOIN
			s_base_object ON
			t_collection_data.stix_id = s_base_object.id
		WHERE
	*/
	var s bytes.Buffer
	s.WriteString("SELECT \n\t")
	s.WriteString(tblColData)
	s.WriteString(".date_added, \n\t")
	s.WriteString(tblColData)
	s.WriteString(".stix_id, \n\t")
	s.WriteString(tblBaseObj)
	s.WriteString(".modified, \n\t")
	s.WriteString(tblBaseObj)
	s.WriteString(".spec_version \n")
	s.WriteString("FROM \n\t")
	s.WriteString(tblColData)
	s.WriteString("\n")
	s.WriteString("JOIN \n\t")
	s.WriteString(tblBaseObj)
	s.WriteString(" ON \n\t")
	s.WriteString(tblColData)
	s.WriteString(".stix_id = ")
	s.WriteString(tblBaseObj)
	s.WriteString(".id \n")
	s.WriteString("WHERE \n\t")
	s.WriteString(whereQuery)

	//log.Println("DEBUG: \n", s.String())
	return s.String(), nil
}

/*
sqlManifestData - This method will return an SQL statement that will
return a list of objects from a given collection and all of the information
needed to create the manifest resource. It will use the query struct to
determine the requirements and parameters for the where clause of the SQL
statement. A byte array is used instead of sting concatenation as it is the most
efficient way to do string concatenation in Go.

Since the manifest resource in TAXII lists out all of the versions of the
object, we are using a group_concat SQL function to give us a string of values
separated by a comma. This will prevent us from having to query the database
multiple times to get all of the needed data.

If you do not use the GROUP BY filter when using the group_concat function then
you get a single row returned with all of the versions listed in the
corresponding modified and spec_version fields.
*/
func (ds *Sqlite3DatastoreType) sqlManifestData(query datastore.QueryType) (string, error) {
	tblColData := datastore.DB_TABLE_TAXII_COLLECTION_DATA
	tblBaseObj := datastore.DB_TABLE_STIX_BASE_OBJECT

	whereQuery, err := ds.sqlCollectionDataQueryOptions(query)

	// If an error is found, that means a query parameter was passed incorrectly
	// and we should return an error versus just skipping the option.
	if err != nil {
		return "", err
	}

	/*
		SELECT
			t_collection_data.date_added,
			t_collection_data.stix_id,
			group_concat(s_base_object.modified),
			group_concat(s_base_object.spec_version)
		FROM
			t_collection_data
		JOIN
			s_base_object ON
			t_collection_data.stix_id = s_base_object.id
		WHERE
			t_collection_data.collection_id = "aa"
		GROUP BY
			t_collection_data.date_added
	*/
	var s bytes.Buffer
	s.WriteString("SELECT \n\t")
	s.WriteString(tblColData)
	s.WriteString(".date_added, \n\t")
	s.WriteString(tblColData)
	s.WriteString(".stix_id, \n\t")
	s.WriteString("group_concat(")
	s.WriteString(tblBaseObj)
	s.WriteString(".modified), \n\t")
	s.WriteString("group_concat(")
	s.WriteString(tblBaseObj)
	s.WriteString(".spec_version) \n")
	s.WriteString("FROM \n\t")
	s.WriteString(tblColData)
	s.WriteString("\n")
	s.WriteString("JOIN \n\t")
	s.WriteString(tblBaseObj)
	s.WriteString(" ON \n\t")
	s.WriteString(tblColData)
	s.WriteString(".stix_id = ")
	s.WriteString(tblBaseObj)
	s.WriteString(".id \n")
	s.WriteString("WHERE \n\t")
	s.WriteString(whereQuery)
	s.WriteString("\n")
	s.WriteString("GROUP BY \n\t")
	s.WriteString(tblColData)
	s.WriteString(".date_added")

	//log.Println("DEBUG: \n", s.String())
	return s.String(), nil
}

func (ds *Sqlite3DatastoreType) sqlCollectionDataWhereCollectionID(id string, b *bytes.Buffer) error {
	tblColData := datastore.DB_TABLE_TAXII_COLLECTION_DATA

	/*
		This sql where statement should look like:
		t_collection_data.collection_id = "some collection id"
	*/
	if id != "" {
		b.WriteString(tblColData)
		b.WriteString(`.collection_id = "`)
		b.WriteString(id)
		b.WriteString(`"`)
	} else {
		return errors.New("no collection ID was provided")
	}
	return nil
}

func (ds *Sqlite3DatastoreType) sqlCollectionDataWhereAddedAfter(date []string, b *bytes.Buffer) error {
	tblColData := datastore.DB_TABLE_TAXII_COLLECTION_DATA

	/*
		This sql where statement should look like:
		t_collection_data.collection_id = "aa" AND
		t_collection_data.date_added > "2017"
	*/
	if date != nil {
		// We are only allowing a single added after value, since having more does
		// not make sense.
		if timestamp.Valid(date[0]) {
			b.WriteString(" AND \n\t")
			b.WriteString(tblColData)
			b.WriteString(`.date_added > "`)
			b.WriteString(date[0])
			b.WriteString(`"`)
		} else {
			return errors.New("the provided timestamp for added after is invalid")
		}
	}
	return nil
}

func (ds *Sqlite3DatastoreType) sqlCollectionDataWhereSTIXID(id []string, b *bytes.Buffer) error {
	tblColData := datastore.DB_TABLE_TAXII_COLLECTION_DATA

	/*
		This sql where statement should look like one of these two:
		t_collection_data.collection_id = "aa" AND
		t_collection_data.stix_id = "indicator--37abef16-7616-439c-86be-23712030c4b7"

		t_collection_data.collection_id = "aa" AND
		(t_collection_data.stix_id = "indicator--37abef16-7616-439c-86be-23712030c4b7" OR
		t_collection_data.stix_id = "attack-pattern--c7c8a099-70a9-487b-a95f-2498d2941104" OR
		t_collection_data.stix_id = "campaign--6f938db5-6648-4ec1-81cb-5b65138c3c66")
	*/
	if id != nil {
		if len(id) == 1 {
			if objects.IsValidSTIXID(id[0]) {
				b.WriteString(" AND \n\t")
				b.WriteString(tblColData)
				b.WriteString(`.stix_id = "`)
				b.WriteString(id[0])
				b.WriteString(`"`)
			} else {
				return errors.New("the provided object id is invalid")
			}
		} else if len(id) > 1 {
			b.WriteString(" AND \n\t(")
			addOR := false
			for _, v := range id {

				// Lets only add the OR after the first object id and not after the last object id
				if addOR == true {
					b.WriteString(" OR \n\t")
					addOR = false
				}
				// Lets make sure the value that was passed in is actually a valid id
				if objects.IsValidSTIXID(v) {
					b.WriteString(tblColData)
					b.WriteString(`.stix_id = "`)
					b.WriteString(v)
					b.WriteString(`"`)
					addOR = true
				} else {
					return errors.New("the provided object id is invalid")
				}
			}
			b.WriteString(")")
		}
	}
	return nil
}

func (ds *Sqlite3DatastoreType) sqlCollectionDataWhereSTIXType(t []string, b *bytes.Buffer) error {
	tblColData := datastore.DB_TABLE_TAXII_COLLECTION_DATA

	/*
		This sql where statement should look like one of these two:
		t_collection_data.collection_id = "aa" AND
		t_collection_data.stix_id LIKE "indicator%"

		t_collection_data.collection_id = "aa" AND
		(t_collection_data.stix_id LIKE "indicator%" OR
		t_collection_data.stix_id LIKE "attack-pattern%" OR
		t_collection_data.stix_id LIKE "campaign%")
	*/
	if t != nil {
		if len(t) == 1 {
			if objects.IsValidSTIXObject(t[0]) {
				b.WriteString(" AND \n\t")
				b.WriteString(tblColData)
				b.WriteString(`.stix_id LIKE "`)
				b.WriteString(t[0])
				b.WriteString(`%"`)
			} else {
				return errors.New("the provided object type is invalid")
			}
		} else if len(t) > 1 {
			b.WriteString(" AND \n\t(")
			addOR := false
			for _, v := range t {

				// Lets only add the OR after the first object and not after the last object
				if addOR == true {
					b.WriteString(" OR \n\t")
					addOR = false
				}
				// Lets make sure the value that was passed in is actually a valid object
				if objects.IsValidSTIXObject(v) {
					b.WriteString(tblColData)
					b.WriteString(`.stix_id LIKE "`)
					b.WriteString(v)
					b.WriteString(`%"`)
					addOR = true
				} else {
					return errors.New("the provided object type is invalid")
				}
			}
			b.WriteString(`)`)
		}
	}
	return nil
}

func (ds *Sqlite3DatastoreType) sqlCollectionDataWhereSTIXVersion(vers []string, b *bytes.Buffer) error {
	tblColData := datastore.DB_TABLE_TAXII_COLLECTION_DATA
	tblBaseObj := datastore.DB_TABLE_STIX_BASE_OBJECT

	/*
		This sql where statement should look like one of the following:
		t_collection_data.collection_id = "aa" AND
		s_base_object.modified = "2017-12-05T02:43:19.783Z"

		t_collection_data.collection_id = "aa" AND
		s_base_object.modified = (select max(modified) from s_base_object where t_collection_data.stix_id = s_base_object.id)

		t_collection_data.collection_id = "aa" AND
		s_base_object.modified = (select min(modified) from s_base_object where t_collection_data.stix_id = s_base_object.id)

		t_collection_data.collection_id = "aa" AND
		(s_base_object.modified = (select min(modified) from s_base_object where t_collection_data.stix_id = s_base_object.id)  OR
		s_base_object.modified = (select max(modified) from s_base_object where t_collection_data.stix_id = s_base_object.id))

		t_collection_data.collection_id = "aa" AND
		(s_base_object.modified = "2017-12-05T02:43:19.783Z" OR
		s_base_object.modified = "2017-12-05T02:43:23.822Z" OR
		s_base_object.modified = "2017-12-05T02:43:24.835Z")
	*/
	if vers != nil {
		if len(vers) == 1 {
			if vers[0] == "last" {

				// s_base_object.modified = (select max(modified) from s_base_object where t_collection_data.stix_id = s_base_object.id)
				b.WriteString(" AND \n\t")
				b.WriteString(tblBaseObj)
				b.WriteString(`.modified = (select max(modified) from `)
				b.WriteString(tblBaseObj)
				b.WriteString(` where `)
				b.WriteString(tblColData)
				b.WriteString(`.stix_id = `)
				b.WriteString(tblBaseObj)
				b.WriteString(`.id)`)
			} else if vers[0] == "first" {

				// s_base_object.modified = (select min(modified) from s_base_object where t_collection_data.stix_id = s_base_object.id)
				b.WriteString(" AND \n\t")
				b.WriteString(tblBaseObj)
				b.WriteString(`.modified = (select min(modified) from `)
				b.WriteString(tblBaseObj)
				b.WriteString(` where `)
				b.WriteString(tblColData)
				b.WriteString(`.stix_id = `)
				b.WriteString(tblBaseObj)
				b.WriteString(`.id)`)
			} else if vers[0] == "all" {
				// Do nothing, since the default is to return all versions.
			} else {
				if timestamp.Valid(vers[0]) {
					b.WriteString(" AND \n\t")
					b.WriteString(tblBaseObj)
					b.WriteString(`.modified = "`)
					b.WriteString(vers[0])
					b.WriteString(`"`)
				} else {
					return errors.New("the provided timestamp for the version is invalid")
				}
			}
		} else if len(vers) > 1 {
			b.WriteString(" AND \n\t(")
			for i, v := range vers {
				// Lets only add he OR after the first object and not after the
				// last object. Since skipOr starts as true, this takes care of
				// the first run case where i == 0

				// TODO need to check to make sure someone does not put last or first in twice and then build a test case for it

				if i > 0 {
					b.WriteString(" OR \n\t")
				}

				if v == "last" {
					b.WriteString(tblBaseObj)
					b.WriteString(`.modified = (select max(modified) from `)
					b.WriteString(tblBaseObj)
					b.WriteString(` where `)
					b.WriteString(tblColData)
					b.WriteString(`.stix_id = `)
					b.WriteString(tblBaseObj)
					b.WriteString(`.id)`)

				} else if v == "first" {
					b.WriteString(tblBaseObj)
					b.WriteString(`.modified = (select min(modified) from `)
					b.WriteString(tblBaseObj)
					b.WriteString(` where `)
					b.WriteString(tblColData)
					b.WriteString(`.stix_id = `)
					b.WriteString(tblBaseObj)
					b.WriteString(`.id)`)

				} else if v == "all" {
					// Do nothing as it will do nothing here, or it should not be valid
					return errors.New("can not use the all key word with a multiple version selector")

				} else {
					if timestamp.Valid(v) {
						b.WriteString(tblBaseObj)
						b.WriteString(`.modified = "`)
						b.WriteString(v)
						b.WriteString(`"`)
					} else {
						return errors.New("the provided timestamp for the version is invalid")
					}
				}
			}
			b.WriteString(`)`)
		}
	}
	//log.Println("DEBUG: \n", b.String())
	return nil
}

/*
sqlCollectionDataQueryOptions - This method will take in a query struct and build an SQL
where statement based on all of the provided query parameters.
*/
func (ds *Sqlite3DatastoreType) sqlCollectionDataQueryOptions(query datastore.QueryType) (string, error) {
	var wherestmt bytes.Buffer
	var err error

	// ----------------------------------------------------------------------
	// Lets first add the collection ID to the where clause.
	// ----------------------------------------------------------------------
	if err = ds.sqlCollectionDataWhereCollectionID(query.CollectionID, &wherestmt); err != nil {
		return "", err
	}

	// ----------------------------------------------------------------------
	// Check to see if an added after query was supplied. There can only be one
	// added after option, it does not make sense to have multiple.
	// ----------------------------------------------------------------------
	if err = ds.sqlCollectionDataWhereAddedAfter(query.AddedAfter, &wherestmt); err != nil {
		return "", err
	}

	// ----------------------------------------------------------------------
	// Check to see if one or more STIX ID, to query on, was supplied.
	// If there is more than one option given we need to enclose the options in
	// parentheses as the comma represents an OR operator.
	// ----------------------------------------------------------------------
	if err = ds.sqlCollectionDataWhereSTIXID(query.STIXID, &wherestmt); err != nil {
		return "", err
	}

	// ----------------------------------------------------------------------
	// Check to see if one or more STIX types, to query on, was supplied.
	// If there is more than one option given we need to enclose the options in
	// parentheses as the comma represents an OR operator.
	// ----------------------------------------------------------------------
	if err = ds.sqlCollectionDataWhereSTIXType(query.STIXType, &wherestmt); err != nil {
		return "", err
	}

	// ----------------------------------------------------------------------
	// Check to see if one or more STIX versions to query on was supplied.
	// If there is more than one option given, we need to enclose the options in
	// parentheses as the comma represents an OR operator.
	// ----------------------------------------------------------------------
	if err = ds.sqlCollectionDataWhereSTIXVersion(query.STIXVersion, &wherestmt); err != nil {
		return "", err
	}

	return wherestmt.String(), nil
}
