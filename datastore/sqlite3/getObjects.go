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
GetListOfObjectsInCollection - This method will take in an ID for a collection
and will return a slice of strings that contains all of the STIX IDs that are in
that collection that meet the range requirements or an error.
*/
func (ds *Sqlite3DatastoreType) GetListOfObjectsInCollection(collectionid string) ([]string, error) {
	var allObjects []string

	var getAllObjectsInCollection = `
		SELECT stix_id
	   	FROM ` + datastore.DB_TABLE_TAXII_COLLECTION_CONTENT + ` 
	   	WHERE collection_id = $1`

	// Query database for all the collection entries
	rows, err := ds.DB.Query(getAllObjectsInCollection, collectionid)
	if err != nil {
		return nil, fmt.Errorf("Database execution error querying collection content: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var stixid string
		if err := rows.Scan(&stixid); err != nil {
			log.Fatal(err)
		}
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
func (ds *Sqlite3DatastoreType) GetObjectsInCollection(collectionid string, paginate bool, maxsize, first, last int) (objects.BundleType, error) {
	// TODO need the ability to take in a query struct of list of parameters

	var rangeOfObjects []string
	var err error
	stixBundle := objects.NewBundle()
	allObjects, err := ds.GetListOfObjectsInCollection(collectionid)

	if err != nil {
		return stixBundle, err
	}

	if paginate == true {
		rangeOfObjects, _, err = ds.GetRangeOfObjects(allObjects, maxsize, first, last)
	} else {
		rangeOfObjects = allObjects
	}

	for _, stixid := range rangeOfObjects {
		obj, _ := ds.GetObject(stixid)
		stixBundle.AddObject(obj)
	}
	return stixBundle, nil
}
