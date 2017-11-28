// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"errors"
	"fmt"
	"github.com/freetaxii/libstix2/common/timestamp"
	"github.com/freetaxii/libstix2/datastore"
	"github.com/freetaxii/libstix2/objects"
	"log"
	"strings"
)

// GetObject - This method will take in a STIX ID and return the STIX object.
func (ds *Sqlite3DatastoreType) GetObject(stixid string) (interface{}, error) {

	// TODO
	// We first need to look at the STIX ID that was passed in to see what type
	// of object it is. Basically split the ID to get the type and then do a
	// switch statement
	// Need to also be able to handle multiple versions
	i, err := ds.getIndicator(stixid)
	return i, err
}

/*
GetListOfObjectsInCollection - This method will take in query struct and range
parameters for a collection and will return a slice of strings that contains all
of the STIX IDs that are in that collection that meet those query or range
parameters.

Return:
	allObjects ([]string) - A list of all STIX objects that match the query parameters
    error
*/
func (ds *Sqlite3DatastoreType) GetListOfObjectsInCollection(query datastore.QueryType) ([]string, error) {
	var allObjects []string

	whereQuery, err := ds.processQueryOptions(query)

	// If an error is found, that means a query parameter was passed incorrectly
	// and we should return an error versus just skipping the option.
	if err != nil {
		return nil, err
	}

	var getAllObjectsInCollection = `
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
	log.Println(getAllObjectsInCollection)

	// Query database for all the collection entries
	rows, err := ds.DB.Query(getAllObjectsInCollection)
	if err != nil {
		return nil, fmt.Errorf("database execution error querying collection content: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var dateAdded, stixid, modified, specVersion string
		if err := rows.Scan(&dateAdded, &stixid, &modified, &specVersion); err != nil {
			log.Fatal(err)
		}
		log.Println(stixid, " ", modified)
		allObjects = append(allObjects, stixid)
	}

	return allObjects, nil
}

/*
GetRangeOfObjects - This method will take in a slice of strings and two index
values that represent the number of records to return. The method will return a
new slice of strings that meet the range requirements or an error.
Retval:
	[]string = list of objects
    int = number of total objects, not number in range
	error
*/
func (ds *Sqlite3DatastoreType) GetRangeOfObjects(allObjects []string, maxsize, first, last int) ([]string, int, error) {

	if first < 0 {
		return nil, 0, errors.New("The starting value can not be negative")
	}

	if first > last {
		return nil, 0, errors.New("The starting range value is larger than the ending range value")
	}

	// We need to be inclusive of the last value
	last++

	if (last - first) > maxsize {
		last = first + maxsize
	}

	size := len(allObjects)

	if first >= size {
		return nil, size, errors.New("The starting range value is out of scope")
	}
	if last > size {
		last = size
	}

	rangeObjects := allObjects[first:last]
	return rangeObjects, size, nil
}

/*
GetObjectsInCollection - This method will take in an ID for a collection and
will return a STIX Bundle that contains all of the STIX objects that are in that
collection that meet the range requirements.
Retval:
  STIX Bundle Type
  error
*/
// func (ds *Sqlite3DatastoreType) GetObjectsInCollection(collectionid string, paginate bool, maxsize, first, last int) (objects.BundleType, error) {
// 	// TODO need the ability to take in a query struct of list of parameters

// 	var rangeOfObjects []string
// 	var err error
// 	stixBundle := objects.NewBundle()
// 	allObjects, _, err := ds.GetListOfObjectsInCollection(collectionid)

// 	if err != nil {
// 		return stixBundle, err
// 	}

// 	if paginate == true {
// 		rangeOfObjects, _, err = ds.GetRangeOfObjects(allObjects, maxsize, first, last)
// 	} else {
// 		rangeOfObjects = allObjects
// 	}

// 	for _, stixid := range rangeOfObjects {
// 		obj, _ := ds.GetObject(stixid)
// 		stixBundle.AddObject(obj)
// 	}
// 	return stixBundle, nil
// }

// ----------------------------------------------------------------------
// Private Methods
// ----------------------------------------------------------------------

/*
processQueryOptions - This method will take in a query struct and build an SQL
where statement based on all of the provided query parameters.

Return:
    webQuery (string) - The where statement for the SQL query
    error
*/
func (ds *Sqlite3DatastoreType) processQueryOptions(query datastore.QueryType) (string, error) {
	var whereQuery string

	// ----------------------------------------------------------------------
	// Lets first add the collection ID to the where clause.
	// ----------------------------------------------------------------------
	if query.CollectionID != "" {
		whereQuery += datastore.DB_TABLE_TAXII_COLLECTION_CONTENT + `.collection_id = "` + query.CollectionID + `"`
	}

	// ----------------------------------------------------------------------
	// Check to see if an added after query was supplied. There can only be one
	// added after option, it does not make sense to have multiple.
	// ----------------------------------------------------------------------
	if query.AddedAfter != "" {
		if timestamp.Valid(query.AddedAfter) {
			whereQuery += ` AND ` + datastore.DB_TABLE_TAXII_COLLECTION_CONTENT + `.date_added > "` + query.AddedAfter + `"`
		} else {
			return "", errors.New("the provided timestamp for added after is invalid")
		}
	}

	// ----------------------------------------------------------------------
	// Check to see if one or more STIX types, to query on, was supplied.
	// If there is more than one option given, split with a comma, we need to
	// enclose the options in parentheses as the comma represents an OR operator.
	// ----------------------------------------------------------------------
	if query.STIXType != "" {
		// If there is more than one type, split it out. If there is only one it
		// will be element [0] in the slice.
		types := strings.Split(query.STIXType, ",")

		if len(types) == 1 {
			if objects.ValidSTIXObject(types[0]) {
				whereQuery += ` AND ` + datastore.DB_TABLE_TAXII_COLLECTION_CONTENT + `.stix_id LIKE "` + types[0] + `%"`
			} else {
				return "", errors.New("the provided object type is invalid")
			}
		} else if len(types) > 1 {
			whereQuery += ` AND (`
			addOR := false
			for _, v := range types {

				// Lets only add the OR after the first object and not after the last object
				if addOR == true {
					whereQuery += ` OR `
					addOR = false
				}
				// Lets make sure the value that was passed in is actually a valid object
				if objects.ValidSTIXObject(v) {
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
	// Check to see if one or more STIX versions to query on was supplied.
	// If there is more than one option given, split with a comma, we need to
	// enclose the options in parentheses as the comma represents an OR operator.
	// ----------------------------------------------------------------------
	if query.STIXVersion != "" {
		// If there is more than one version, split it out. If there is only one
		// it will be element [0] in the slice.
		versions := strings.Split(query.STIXVersion, ",")

		if len(versions) == 1 {
			if versions[0] == "last" {
				whereQuery += ` AND ` + datastore.DB_TABLE_STIX_BASE_OBJECT + `.modified = (select max(modified) from ` + datastore.DB_TABLE_STIX_BASE_OBJECT + ` where ` + datastore.DB_TABLE_TAXII_COLLECTION_CONTENT + `.stix_id = ` + datastore.DB_TABLE_STIX_BASE_OBJECT + `.id) `
			} else if versions[0] == "first" {
				whereQuery += ` AND ` + datastore.DB_TABLE_STIX_BASE_OBJECT + `.modified = (select min(modified) from ` + datastore.DB_TABLE_STIX_BASE_OBJECT + ` where ` + datastore.DB_TABLE_TAXII_COLLECTION_CONTENT + `.stix_id = ` + datastore.DB_TABLE_STIX_BASE_OBJECT + `.id) `
			} else if versions[0] == "all" {
				// Do nothing, since the default is to return all versions.
			} else {
				//whereQuery = whereQuery + ` AND s_base_object.modified = (select modified from s_base_object where t_collection_content.stix_id = s_base_object.id AND s_base_object.modified = $4) `
				if timestamp.Valid(versions[0]) {
					whereQuery += ` AND ` + datastore.DB_TABLE_STIX_BASE_OBJECT + `.modified = "` + versions[0] + `"`
				} else {
					return "", errors.New("the provided timestamp for the version is invalid")
				}
			}
		} else if len(versions) > 1 {
			whereQuery += ` AND (`
			for i, v := range versions {
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
