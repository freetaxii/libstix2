// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"errors"
	"github.com/freetaxii/libstix2/common/timestamp"
	"github.com/freetaxii/libstix2/datastore"
	"github.com/freetaxii/libstix2/objects"
)

/*
sqlListOfObjectsFromCollection - This method will take in a query struct and return
an SQL select statement that matches the requirements and parameters given in the
query struct.
*/
func (ds *Sqlite3DatastoreType) sqlListOfObjectsFromCollection(query datastore.QueryType) (string, error) {
	whereQuery, err := ds.processQueryOptions(query)

	// If an error is found, that means a query parameter was passed incorrectly
	// and we should return an error versus just skipping the option.
	if err != nil {
		return "", err
	}

	var sqlstmt = `
		SELECT
			t_collection_content.date_added,
			t_collection_content.stix_id,
			s_base_object.modified,
			s_base_object.spec_version
		FROM ` + datastore.DB_TABLE_TAXII_COLLECTION_CONTENT + `
		JOIN s_base_object
		ON t_collection_content.stix_id = s_base_object.id
		WHERE ` + whereQuery + `
		`

	// Debug
	//log.Println(sqlstmt)
	return sqlstmt, nil
}

/*
sqlManifestDataFromCollection - This method will take in a query struct and return
an SQL select statement that matches the requirements and parameters given in the
query struct.
*/
func (ds *Sqlite3DatastoreType) sqlManifestDataFromCollection(query datastore.QueryType) (string, error) {
	whereQuery, err := ds.processQueryOptions(query)

	// If an error is found, that means a query parameter was passed incorrectly
	// and we should return an error versus just skipping the option.
	if err != nil {
		return "", err
	}

	var sqlstmt = `
		SELECT
			t_collection_content.date_added,
			t_collection_content.stix_id,
			group_concat(s_base_object.modified),
			group_concat(s_base_object.spec_version)
		FROM ` + datastore.DB_TABLE_TAXII_COLLECTION_CONTENT + `
		JOIN s_base_object
		ON t_collection_content.stix_id = s_base_object.id
		WHERE ` + whereQuery + `
		GROUP BY t_collection_content.stix_id
		`

	// Debug
	//log.Println(sqlstmt)
	return sqlstmt, nil
}

/*
processQueryOptions - This method will take in a query struct and build an SQL
where statement based on all of the provided query parameters.
*/
func (ds *Sqlite3DatastoreType) processQueryOptions(query datastore.QueryType) (string, error) {
	var whereQuery string

	// ----------------------------------------------------------------------
	// Lets first add the collection ID to the where clause.
	// ----------------------------------------------------------------------
	if query.CollectionID != "" {
		whereQuery += datastore.DB_TABLE_TAXII_COLLECTION_CONTENT + `.collection_id = "` + query.CollectionID + `"`
	} else {
		return "", errors.New("no collection ID was provided")
	}

	// ----------------------------------------------------------------------
	// Check to see if an added after query was supplied. There can only be one
	// added after option, it does not make sense to have multiple.
	// ----------------------------------------------------------------------
	if query.AddedAfter != nil {
		if timestamp.Valid(query.AddedAfter[0]) {
			whereQuery += ` AND ` + datastore.DB_TABLE_TAXII_COLLECTION_CONTENT + `.date_added > "` + query.AddedAfter[0] + `"`
		} else {
			return "", errors.New("the provided timestamp for added after is invalid")
		}
	}

	// ----------------------------------------------------------------------
	// Check to see if one or more STIX ID, to query on, was supplied.
	// If there is more than one option given we need to enclose the options in
	// parentheses as the comma represents an OR operator.
	// ----------------------------------------------------------------------
	if query.STIXID != nil {

		if len(query.STIXID) == 1 {
			if objects.IsValidSTIXID(query.STIXID[0]) {
				whereQuery += ` AND ` + datastore.DB_TABLE_TAXII_COLLECTION_CONTENT + `.stix_id = "` + query.STIXID[0] + `"`
			} else {
				return "", errors.New("the provided object id is invalid")
			}
		} else if len(query.STIXID) > 1 {
			whereQuery += ` AND (`
			addOR := false
			for _, v := range query.STIXID {

				// Lets only add the OR after the first object id and not after the last object id
				if addOR == true {
					whereQuery += ` OR `
					addOR = false
				}
				// Lets make sure the value that was passed in is actually a valid id
				if objects.IsValidSTIXID(v) {
					whereQuery += datastore.DB_TABLE_TAXII_COLLECTION_CONTENT + `.stix_id = "` + v + `"`
					addOR = true
				} else {
					return "", errors.New("the provided object id is invalid")
				}
			}
			whereQuery += `)`
		}
	}

	// ----------------------------------------------------------------------
	// Check to see if one or more STIX types, to query on, was supplied.
	// If there is more than one option given we need to enclose the options in
	// parentheses as the comma represents an OR operator.
	// ----------------------------------------------------------------------
	if query.STIXType != nil {

		if len(query.STIXType) == 1 {
			if objects.IsValidSTIXObject(query.STIXType[0]) {
				whereQuery += ` AND ` + datastore.DB_TABLE_TAXII_COLLECTION_CONTENT + `.stix_id LIKE "` + query.STIXType[0] + `%"`
			} else {
				return "", errors.New("the provided object type is invalid")
			}
		} else if len(query.STIXType) > 1 {
			whereQuery += ` AND (`
			addOR := false
			for _, v := range query.STIXType {

				// Lets only add the OR after the first object and not after the last object
				if addOR == true {
					whereQuery += ` OR `
					addOR = false
				}
				// Lets make sure the value that was passed in is actually a valid object
				if objects.IsValidSTIXObject(v) {
					whereQuery += datastore.DB_TABLE_TAXII_COLLECTION_CONTENT + `.stix_id LIKE "` + v + `%"`
					addOR = true
				} else {
					return "", errors.New("the provided object type is invalid")
				}
			}
			whereQuery += `)`
		}
	}

	// ----------------------------------------------------------------------
	// Check to see if one or more STIX query.STIXVersion to query on was supplied.
	// If there is more than one option given, we need to enclose the options in
	// parentheses as the comma represents an OR operator.
	// ----------------------------------------------------------------------
	if query.STIXVersion != nil {

		if len(query.STIXVersion) == 1 {
			if query.STIXVersion[0] == "last" {
				whereQuery += ` AND ` + datastore.DB_TABLE_STIX_BASE_OBJECT + `.modified = (select max(modified) from ` + datastore.DB_TABLE_STIX_BASE_OBJECT + ` where ` + datastore.DB_TABLE_TAXII_COLLECTION_CONTENT + `.stix_id = ` + datastore.DB_TABLE_STIX_BASE_OBJECT + `.id) `
			} else if query.STIXVersion[0] == "first" {
				whereQuery += ` AND ` + datastore.DB_TABLE_STIX_BASE_OBJECT + `.modified = (select min(modified) from ` + datastore.DB_TABLE_STIX_BASE_OBJECT + ` where ` + datastore.DB_TABLE_TAXII_COLLECTION_CONTENT + `.stix_id = ` + datastore.DB_TABLE_STIX_BASE_OBJECT + `.id) `
			} else if query.STIXVersion[0] == "all" {
				// Do nothing, since the default is to return all versions.
			} else {
				//whereQuery = whereQuery + ` AND s_base_object.modified = (select modified from s_base_object where t_collection_content.stix_id = s_base_object.id AND s_base_object.modified = $4) `
				if timestamp.Valid(query.STIXVersion[0]) {
					whereQuery += ` AND ` + datastore.DB_TABLE_STIX_BASE_OBJECT + `.modified = "` + query.STIXVersion[0] + `"`
				} else {
					return "", errors.New("the provided timestamp for the version is invalid")
				}
			}
		} else if len(query.STIXVersion) > 1 {
			whereQuery += ` AND (`
			for i, v := range query.STIXVersion {
				// Lets only add he OR after the first object and not after the last object
				if i > 0 {
					whereQuery += ` OR `
				}
				if v == "last" {
					whereQuery += datastore.DB_TABLE_STIX_BASE_OBJECT + `.modified = (select max(modified) from ` + datastore.DB_TABLE_STIX_BASE_OBJECT + ` where ` + datastore.DB_TABLE_TAXII_COLLECTION_CONTENT + `.stix_id = ` + datastore.DB_TABLE_STIX_BASE_OBJECT + `.id) `
				} else if v == "first" {
					whereQuery += datastore.DB_TABLE_STIX_BASE_OBJECT + `.modified = (select min(modified) from ` + datastore.DB_TABLE_STIX_BASE_OBJECT + ` where ` + datastore.DB_TABLE_TAXII_COLLECTION_CONTENT + `.stix_id = ` + datastore.DB_TABLE_STIX_BASE_OBJECT + `.id) `
				} else if v == "all" {
					// Do nothing as it will do nothing here, or it should not be valid
				} else {
					if timestamp.Valid(v) {
						whereQuery += datastore.DB_TABLE_STIX_BASE_OBJECT + `.modified = "` + v + `"`
					} else {
						return "", errors.New("the provided timestamp for the version is invalid")
					}
				}
			}
			whereQuery += `)`
		}
	}
	return whereQuery, nil
}
