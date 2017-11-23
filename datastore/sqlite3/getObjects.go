// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"errors"
	"fmt"
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
GetListOfObjectsInCollection - This method will take in query and range
parameters for a collection and will return a slice of strings that contains all
of the STIX IDs that are in that collection that meet those query or range
parameters.

Retval:
	[]string = list of objects
    int = number of total objects, not number in range
    error
*/
func (ds *Sqlite3DatastoreType) GetListOfObjectsInCollection(query datastore.QueryType) ([]string, int, error) {
	var allObjects []string
	var whereQuery string

	if query.AddedAfter != "" {
		whereQuery = whereQuery + ` AND t_collection_content.date_added > $2 `
	}

	// ----------------------------------------------------------------------
	// Check for one or more STIX types to query on
	// ----------------------------------------------------------------------
	if query.STIXType != "" {
		// If there is more than one type, split it out
		types := strings.Split(query.STIXType, ",")

		if len(types) == 1 {
			if objects.ValidSTIXObject(query.STIXType) {
				whereQuery += ` AND t_collection_content.stix_id LIKE "` + query.STIXType + `%"`
			}
		} else if len(types) > 1 {
			whereQuery += ` AND (`
			for i, v := range types {
				// Lets only add the OR after the first object and not after the last object
				if i > 0 {
					whereQuery += ` OR `
				}
				// Lets make sure the value that was passed in is actually a valid object
				if objects.ValidSTIXObject(v) {
					whereQuery += `t_collection_content.stix_id LIKE "` + v + `%"`
				}
			}
			whereQuery += `)`
		}
	}

	if query.STIXVersion != "" {
		if query.STIXVersion == "last" {
			whereQuery = whereQuery + ` AND s_base_object.modified = (select max(modified) from s_base_object where t_collection_content.stix_id = s_base_object.id) `
			// We need to zero out the value since we will be passing it in to the query function below for the else use case
			query.STIXVersion = ""
		} else if query.STIXVersion == "first" {
			whereQuery = whereQuery + ` AND s_base_object.modified = (select min(modified) from s_base_object where t_collection_content.stix_id = s_base_object.id) `
			// We need to zero out the value since we will be passing it in to the query function below for the else use case
			query.STIXVersion = ""
		} else if query.STIXVersion == "all" {
			// We need to zero out the value since we will be passing it in to the query function below for the else use case
			query.STIXVersion = ""
		} else {
			//whereQuery = whereQuery + ` AND s_base_object.modified = (select modified from s_base_object where t_collection_content.stix_id = s_base_object.id AND s_base_object.modified = $4) `
			whereQuery = whereQuery + ` AND s_base_object.modified = $3 `
		}
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
		WHERE 
			t_collection_content.collection_id = $1 ` + whereQuery +
		` GROUP BY t_collection_content.stix_id
	`

	log.Println(getAllObjectsInCollection)

	// Query database for all the collection entries
	rows, err := ds.DB.Query(getAllObjectsInCollection, query.CollectionID, query.AddedAfter, query.STIXVersion)
	if err != nil {
		return nil, 0, fmt.Errorf("Database execution error querying collection content: ", err)
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

	size := len(allObjects)
	return allObjects, size, nil
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
