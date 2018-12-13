// Copyright 2015-2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"

	"github.com/freetaxii/libstix2/defs"
	"github.com/freetaxii/libstix2/resources/collections"
	"github.com/freetaxii/libstix2/resources/manifest"
	"github.com/freetaxii/libstix2/stixid"
	"github.com/freetaxii/libstix2/timestamp"
)

// ----------------------------------------------------------------------
//
// Collection Data Table Private Functions and Methods
// getManifestData
//
// ----------------------------------------------------------------------

/*
getManifestData - This method will return manifest data based on the query provided.

The SQL statement that is built in this method will return a list of objects
from a given collection and all of the information needed to create the manifest
resource. It will use the query struct to determine the requirements and
parameters for the where clause of the SQL statement. A byte array is used
instead of sting concatenation as it is the most efficient way to do string
concatenation in Go.
*/
func (ds *Store) getManifestData(query collections.CollectionQuery) (*collections.CollectionQueryResult, error) {
	ds.Logger.Levelln("Function", "FUNC: getManifestData start")

	// Lets first make sure the collection exists in the cache
	if found := ds.doesCollectionExistInTheCache(query.CollectionUUID); !found {
		ds.Logger.Levelln("Function", "FUNC: getManifestData exited with an error")
		return nil, fmt.Errorf("the following collection id was not found in the cache", query.CollectionUUID)
	}
	query.CollectionDatastoreID = ds.Cache.Collections[query.CollectionUUID].DatastoreID

	var resultData collections.CollectionQueryResult
	manifestData := manifest.New()

	// Create SQL Statement
	/*
		SELECT
			t_collection_data.stix_id,
			s_base_object.date_added,
			s_base_object.modified,
			s_base_object.spec_version
		FROM
			t_collection_data
		JOIN
			s_base_object ON
			t_collection_data.stix_id = s_base_object.id
		WHERE
			t_collection_data.collection_id = ?
		ORDER BY
			s_base_object.date_added
		LIMIT 5
	*/

	// If an error is found, that means a query parameter was passed incorrectly
	// and we should return an error versus just skipping the option.
	whereQuery, err := ds.sqlCollectionDataQueryOptions(query)

	if err != nil {
		ds.Logger.Levelln("Function", "FUNC: getManifestData exited with an error,", err)
		return nil, err
	}

	// If the client passes in an invalid value, then the server limit is used
	limitQuery := ds.sqlQueryLimit(query)

	tblColData := DB_TABLE_TAXII_COLLECTION_DATA
	tblBaseObj := DB_TABLE_STIX_BASE_OBJECT
	var sqlstmt bytes.Buffer
	sqlstmt.WriteString("SELECT ")
	sqlstmt.WriteString(tblColData)
	sqlstmt.WriteString(".stix_id, ")
	sqlstmt.WriteString(tblBaseObj)
	sqlstmt.WriteString(".date_added, ")
	sqlstmt.WriteString(tblBaseObj)
	sqlstmt.WriteString(".modified, ")
	sqlstmt.WriteString(tblBaseObj)
	sqlstmt.WriteString(".spec_version ")

	sqlstmt.WriteString("FROM ")
	sqlstmt.WriteString(tblColData)

	sqlstmt.WriteString(" JOIN ")
	sqlstmt.WriteString(tblBaseObj)
	sqlstmt.WriteString(" ON ")
	sqlstmt.WriteString(tblColData)
	sqlstmt.WriteString(".stix_id = ")
	sqlstmt.WriteString(tblBaseObj)
	sqlstmt.WriteString(".id ")

	sqlstmt.WriteString("WHERE ")
	sqlstmt.WriteString(whereQuery)

	sqlstmt.WriteString(" ORDER BY ")
	sqlstmt.WriteString(tblBaseObj)
	sqlstmt.WriteString(".date_added ASC ")

	if limitQuery != 0 {
		sqlstmt.WriteString(" LIMIT ")
		// We need to ask for limit + 1 to see if there are more records available
		// but only if a limit is being enforced. Later on we will check the number
		// of records that comes back to see if we got limit or limit+1
		limitQueryPagination := limitQuery + 1
		i := strconv.Itoa(limitQueryPagination)
		sqlstmt.WriteString(i)
	}
	stmt := sqlstmt.String()

	// Make SQL Call
	// Query database for all the collection entries
	rows, err := ds.DB.Query(stmt)

	if err != nil {
		ds.Logger.Levelln("Function", "FUNC: getManifestData exited with an error,", err)
		return nil, fmt.Errorf("database execution error getting collection data: ", err)
	}
	defer rows.Close()

	// Loop through all records returned and build a manifest resource
	counter := 0
	for rows.Next() {
		var stixid, dateAdded, modified, specVersion string
		if err := rows.Scan(&stixid, &dateAdded, &modified, &specVersion); err != nil {
			rows.Close()
			ds.Logger.Levelln("Function", "FUNC: getManifestData exited with an error,", err)
			return nil, fmt.Errorf("database scan error getting collection data: ", err)
		}

		if counter == limitQuery {
			// If we have added the number of records that the limit query is
			// set to, then lets not add the last record, but rather set the
			// pagination value of more to true and eject from the for loop
			manifestData.SetMore()
			continue
		}

		switch specVersion {
		case "2.0":
			specVersion = defs.MEDIA_TYPE_STIX20
		case "2.1":
			specVersion = defs.MEDIA_TYPE_STIX21
		case "2.2":
			specVersion = defs.MEDIA_TYPE_STIX22
		case "2.3":
			specVersion = defs.MEDIA_TYPE_STIX23
		case "2.4":
			specVersion = defs.MEDIA_TYPE_STIX24
		default:
			specVersion = defs.MEDIA_TYPE_STIX
		}
		manifestData.CreateRecord(stixid, dateAdded, modified, specVersion)
		counter++
	}

	// Errors can cause the rows.Next() to exit prematurely, if this happens lets
	// check for the error and handle it.
	if err := rows.Err(); err != nil {
		rows.Close()
		ds.Logger.Levelln("Function", "FUNC: getManifestData exited with an error,", err)
		return nil, fmt.Errorf("database rows error getting manifest data: ", err)
	}

	if len(manifestData.Objects) == 0 {
		ds.Logger.Levelln("Function", "FUNC: getManifestData exited with an error")
		return nil, fmt.Errorf("no records returned getting manifest data")
	}

	resultData.Size = ds.Cache.Collections[query.CollectionUUID].Size

	ds.Logger.Debugln("DEBUG: Query Collection ID", query.CollectionUUID)
	ds.Logger.Debugln("DEBUG: Cache ID", ds.Cache.Collections[query.CollectionUUID].ID, "Cache Datastore ID", ds.Cache.Collections[query.CollectionUUID].DatastoreID, "Size in Cache", ds.Cache.Collections[query.CollectionUUID].Size)

	resultData.ManifestData.More = manifestData.More
	resultData.ManifestData.Objects = manifestData.Objects

	// ----------------------------------------------------------------------
	// This is the old pagination code
	// first, last, errRange = ds.processRangeValues(query.RangeBegin, query.RangeEnd, query.ServerRecordLimit, resultData.Size)

	// if errRange != nil {
	// 	return nil, errRange
	// }

	// Get a new slice based on the range of records
	// resultData.ManifestData.Objects = manifest.Objects[first:last]
	// ----------------------------------------------------------------------

	resultData.DateAddedFirst = resultData.ManifestData.Objects[0].DateAdded
	resultData.DateAddedLast = resultData.ManifestData.Objects[len(resultData.ManifestData.Objects)-1].DateAdded
	ds.Logger.Levelln("Function", "FUNC: getManifestData end")
	return &resultData, nil
}

// ----------------------------------------------------------------------
//
// LIMIT statements for Collection Data Queries
//
// ----------------------------------------------------------------------

/*
sqlQueryLimit - This method will take in a query struct and build an SQL LIMIT
statement based on the values provided in the query object.
*/
func (ds *Store) sqlQueryLimit(query collections.CollectionQuery) int {
	srv := 0
	client := 0
	var err error

	// In the configuration, if the server record limit is 0 or a negative number
	// than the value is zero and thus no limit will be applied.
	if query.ServerRecordLimit > 0 {
		srv = query.ServerRecordLimit
	}

	// Lets check to see if the client passed in a valid value over the URL
	if query.Limit != nil {
		client, err = strconv.Atoi(query.Limit[0])
	}
	if err != nil {
		ds.Logger.Debugln("DEBUG: Client limit value is not valid: ", err)
		return srv
	}

	if client > srv {
		ds.Logger.Debugln("DEBUG: Client limit value is greater than the server limit, using server limit of", srv)
		return srv
	} else if client < 0 {
		ds.Logger.Debugln("DEBUG: Client limit value is less than zero, using server limit of", srv)
		return srv
	} else if client == 0 {
		ds.Logger.Debugln("DEBUG: Client limit value is equal to zero, using server limit of", srv)
		return srv
	} else if client == srv {
		ds.Logger.Debugln("DEBUG: Client limit value is equal to server limit, using server limit of", srv)
		return srv
	} else if client < srv {
		ds.Logger.Debugln("DEBUG: Client limit value is less than server limit, using client limit of", client)
		return client
	}

	return 0
}

// ----------------------------------------------------------------------
//
// WHERE statements for Collection Data Queries
//
// ----------------------------------------------------------------------

/*
sqlCollectionDataQueryOptions - This method will take in a query struct and
build an SQL where statement based on all of the provided query parameters.
*/
func (ds *Store) sqlCollectionDataQueryOptions(query collections.CollectionQuery) (string, error) {
	var wherestmt bytes.Buffer
	var err error

	// ----------------------------------------------------------------------
	// Lets first add the collection ID to the where clause.
	// ----------------------------------------------------------------------
	if err = sqlCollectionDataWhereCollectionID(query.CollectionDatastoreID, &wherestmt); err != nil {
		return "", err
	}

	// ----------------------------------------------------------------------
	// Check to see if an added after query was supplied. There can only be one
	// added after option, it does not make sense to have multiple.
	// ----------------------------------------------------------------------
	if err = sqlCollectionDataWhereAddedAfter(query.AddedAfter, &wherestmt); err != nil {
		return "", err
	}

	// ----------------------------------------------------------------------
	// Check to see if one or more STIX IDs was supplied.
	// If there is more than one option given we need to enclose the options in
	// parentheses as the comma represents an OR operator.
	// ----------------------------------------------------------------------
	if err = sqlCollectionDataWhereSTIXID(query.STIXID, &wherestmt); err != nil {
		return "", err
	}

	// ----------------------------------------------------------------------
	// Check to see if one or more STIX types, to query on, was supplied.
	// If there is more than one option given we need to enclose the options in
	// parentheses as the comma represents an OR operator.
	// ----------------------------------------------------------------------
	if err = sqlCollectionDataWhereSTIXType(query.STIXType, &wherestmt); err != nil {
		return "", err
	}

	// ----------------------------------------------------------------------
	// Check to see if one or more STIX versions to query on was supplied.
	// If there is more than one option given, we need to enclose the options in
	// parentheses as the comma represents an OR operator.
	// ----------------------------------------------------------------------
	if err = sqlCollectionDataWhereSTIXVersion(query.STIXVersion, &wherestmt); err != nil {
		return "", err
	}

	// ----------------------------------------------------------------------
	// Check to see if one or more sepc versions to query on was supplied.
	// If there is more than one option given, we need to enclose the options in
	// parentheses as the comma represents an OR operator.
	// ----------------------------------------------------------------------
	if err = sqlCollectionDataWhereSpecVersion(query.SpecVersion, &wherestmt); err != nil {
		return "", err
	}

	return wherestmt.String(), nil
}

/*
sqlCollectionDataWhereCollectionID - This function will build the correct WHERE
statement for a provided collection ID (datastore id) value and is called from
func sqlCollectionDataQueryOptions(query collections.CollectionQueryType) (string, error)
*/
func sqlCollectionDataWhereCollectionID(datastoreID int, b *bytes.Buffer) error {
	tblColData := DB_TABLE_TAXII_COLLECTION_DATA

	/*
		This sql where statement should look like:
		t_collection_data.collection_id = "some collection id"
	*/
	if datastoreID != 0 {
		b.WriteString(tblColData)
		b.WriteString(`.collection_id = `)
		// Even though we are matching an integer in the database we need to
		// write a string out since the sql statement itself is a string.
		b.WriteString(strconv.Itoa(datastoreID))
	} else {
		return errors.New("no collection ID was provided")
	}
	return nil
}

/*
sqlCollectionDataWhereAddedAfter - This function will build the correct WHERE
statement for a provided added after value and is called from
func sqlCollectionDataQueryOptions(query collections.CollectionQueryType) (string, error)

This method only supports a single added after value, since more than one does
not make sense.
*/
func sqlCollectionDataWhereAddedAfter(date []string, b *bytes.Buffer) error {
	tblBaseObj := DB_TABLE_STIX_BASE_OBJECT

	/*
		This sql where statement should look like:
		t_collection_data.collection_id = "aa" AND
		t_collection_data.date_added > "2017"
	*/
	if date != nil {
		// We are only allowing a single added after value, since having more does
		// not make sense.
		if timestamp.Valid(date[0]) {
			b.WriteString(" AND ")
			b.WriteString(tblBaseObj)
			b.WriteString(`.date_added > "`)
			b.WriteString(date[0])
			b.WriteString(`"`)
		} else {
			return errors.New("the provided timestamp for added after is invalid")
		}
	}
	return nil
}

/*
sqlCollectionDataWhereSTIXID - This function will build the correct WHERE
statement when one or more STIX IDs is provided and is called from
func sqlCollectionDataQueryOptions(query collections.CollectionQueryType) (string, error)
*/
func sqlCollectionDataWhereSTIXID(id []string, b *bytes.Buffer) error {
	tblColData := DB_TABLE_TAXII_COLLECTION_DATA

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
			if stixid.ValidSTIXID(id[0]) {
				b.WriteString(" AND ")
				b.WriteString(tblColData)
				b.WriteString(`.stix_id = "`)
				b.WriteString(id[0])
				b.WriteString(`"`)
			} else {
				return errors.New("invalid SQL where statement, the provided object id is invalid")
			}
		} else if len(id) > 1 {
			b.WriteString(" AND (")
			addOR := false
			for _, v := range id {

				// Lets only add the OR after the first object id and not after the last object id
				if addOR == true {
					b.WriteString(" OR ")
					addOR = false
				}
				// Lets make sure the value that was passed in is actually a valid id

				if stixid.ValidSTIXID(v) {
					b.WriteString(tblColData)
					b.WriteString(`.stix_id = "`)
					b.WriteString(v)
					b.WriteString(`"`)
					addOR = true
				} else {
					return errors.New("invalid SQL where statement, the provided object id is invalid")
				}
			}
			b.WriteString(")")
		}
	}
	return nil
}

/*
sqlCollectionDataWhereSTIXType - This function will build the correct WHERE
statement when one or more STIX types is provided and is called from
func sqlCollectionDataQueryOptions(query collections.CollectionQueryType) (string, error)
*/
func sqlCollectionDataWhereSTIXType(t []string, b *bytes.Buffer) error {
	tblColData := DB_TABLE_TAXII_COLLECTION_DATA

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
			if stixid.ValidSTIXObjectType(t[0]) {
				b.WriteString(" AND ")
				b.WriteString(tblColData)
				b.WriteString(`.stix_id LIKE "`)
				b.WriteString(t[0])
				b.WriteString(`%"`)
			} else {
				return errors.New("the provided object type is invalid")
			}
		} else if len(t) > 1 {
			b.WriteString(" AND (")
			addOR := false
			for _, v := range t {

				// Lets only add the OR after the first object and not after the last object
				if addOR == true {
					b.WriteString(" OR ")
					addOR = false
				}
				// Lets make sure the value that was passed in is actually a valid object
				if stixid.ValidSTIXObjectType(v) {
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

/*
sqlCollectionDataWhereSTIXVersion - This function will build the correct WHERE
statement when one or more STIX versions is provided and is called from
func sqlCollectionDataQueryOptions(query collections.CollectionQueryType) (string, error).

It will return an error if multiple "all", "first", or "last" values is provided.
*/
func sqlCollectionDataWhereSTIXVersion(vers []string, b *bytes.Buffer) error {
	tblColData := DB_TABLE_TAXII_COLLECTION_DATA
	tblBaseObj := DB_TABLE_STIX_BASE_OBJECT

	// If no version parameter was supplied, then set "last" as the default
	if vers == nil {
		vers = append(vers, "last")
	}

	// Lets check the multiple version use case and see if the options are valid
	if len(vers) > 1 {
		first := 0
		last := 0
		all := 0

		for _, v := range vers {
			if v == "last" {
				last++
			} else if v == "first" {
				first++
			} else if v == "all" {
				all++
			}
		}

		if last > 1 {
			return errors.New("can not use the 'last' key word multiple time in the version selector")
		}
		if first > 1 {
			return errors.New("can not use the 'first' key word multiple time in the version selector")
		}
		if all > 0 {
			return errors.New("can not use the 'all' key word with a multiple version selector")
		}
	}

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

	if len(vers) == 1 {
		if vers[0] == "last" {
			// s_base_object.modified = (select max(modified) from s_base_object where t_collection_data.stix_id = s_base_object.id)
			b.WriteString(" AND ")
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
			b.WriteString(" AND ")
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
				b.WriteString(" AND ")
				b.WriteString(tblBaseObj)
				b.WriteString(`.modified = "`)
				b.WriteString(vers[0])
				b.WriteString(`"`)

			} else {
				return errors.New("the provided timestamp for the version is invalid")
			}
		}
	} else if len(vers) > 1 {
		b.WriteString(" AND (")
		for i, v := range vers {
			// Lets only add he OR after the first object and not after the
			// last object. Since skipOr starts as true, this takes care of
			// the first run case where i == 0

			if i > 0 {
				b.WriteString(" OR ")
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

	return nil
}

/*
sqlCollectionDataWhereSpecVersion - This function will build the correct WHERE
statement when one or more STIX versions is provided and is called from
func sqlCollectionDataQueryOptions(query collections.CollectionQueryType) (string, error).
*/
func sqlCollectionDataWhereSpecVersion(vers []string, b *bytes.Buffer) error {
	tblBaseObj := DB_TABLE_STIX_BASE_OBJECT

	/*
		This sql where statement should look like one of the following:
		t_collection_data.collection_id = "aa" AND
		s_base_object.spec_version = "2.0"

		t_collection_data.collection_id = "aa" AND
		(s_base_object.spec_version = "2.0" OR
		s_base_object.spec_version = "2.1")
	*/

	if len(vers) == 1 {

		if vers[0] == "2.0" {
			b.WriteString(" AND ")
			b.WriteString(tblBaseObj)
			b.WriteString(`.spec_version = "2.0" `)
		} else {
			b.WriteString(" AND ")
			b.WriteString(tblBaseObj)
			b.WriteString(`.spec_version = "2.1" `)
		}

	} else if len(vers) > 1 {
		b.WriteString(" AND (")
		for i, v := range vers {
			// Lets only add he OR after the first object and not after the
			// last object. Since skipOr starts as true, this takes care of
			// the first run case where i == 0

			if i > 0 {
				b.WriteString(" OR ")
			}

			if v == "2.0" {
				b.WriteString(tblBaseObj)
				b.WriteString(`.spec_version = "2.0"`)
			} else {
				b.WriteString(tblBaseObj)
				b.WriteString(`.spec_version = "2.1"`)
			}

		}
		b.WriteString(`)`)
	}

	return nil
}
