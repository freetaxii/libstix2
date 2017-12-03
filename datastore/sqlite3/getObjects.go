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
	"github.com/freetaxii/libstix2/resources"
)

/*
GetObject - This method will take in a STIX ID and version timestamp (the
modified timestamp from a STIX object) and return the STIX object.
*/
func (ds *Sqlite3DatastoreType) GetObject(stixid, version string) (interface{}, error) {

	// TODO
	// We first need to look at the STIX ID that was passed in to see what type
	// of object it is. Basically split the ID to get the type and then do a
	// switch statement
	// Need to also be able to handle multiple versions
	i, err := ds.getIndicator(stixid, version)
	return i, err
}

/*
GetObjectsFromCollection - This method will take in query struct with range
parameters for a collection and will return a STIX Bundle that contains all
of the STIX objects that are in that collection that meet those query or range
parameters.
*/
func (ds *Sqlite3DatastoreType) GetObjectsFromCollection(query datastore.QueryType) (*objects.BundleType, *datastore.QueryReturnDataType, error) {

	stixBundle := objects.NewBundle()

	rangeCollectionRawData, metaData, err := ds.GetListOfObjectsFromCollection(query)
	if err != nil {
		return nil, nil, err
	}

	for _, v := range *rangeCollectionRawData {
		// Only get the objects that are part of the response
		obj, err := ds.GetObject(v.STIXID, v.STIXVersion)

		if err != nil {
			return nil, nil, err
		}
		stixBundle.AddObject(obj)
	}

	return &stixBundle, metaData, nil
}

/*
GetListOfObjectsFromCollection - This method will take in query struct with range
parameters for a collection and will return a datastore collection raw data type
that contains all of the STIX IDs and their associated meta data that are in
that collection that meet those query or range parameters.
*/
func (ds *Sqlite3DatastoreType) GetListOfObjectsFromCollection(query datastore.QueryType) (*[]datastore.CollectionRawDataType, *datastore.QueryReturnDataType, error) {
	var metaData datastore.QueryReturnDataType
	var collectionRawData []datastore.CollectionRawDataType
	var rangeCollectionRawData []datastore.CollectionRawDataType

	sqlStmt, err := ds.sqlListOfObjectsFromCollection(query)

	// If an error is found, that means a query parameter was passed incorrectly
	// and we should return an error versus just skipping the option.
	if err != nil {
		return nil, nil, err
	}

	// Query database for all the collection entries
	rows, err := ds.DB.Query(sqlStmt)
	if err != nil {
		return nil, nil, fmt.Errorf("database execution error querying collection content: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var dateAdded, stixid, modified, specVersion string
		if err := rows.Scan(&dateAdded, &stixid, &modified, &specVersion); err != nil {
			return nil, nil, fmt.Errorf("database scan error: ", err)
		}
		var rawData datastore.CollectionRawDataType
		rawData.STIXID = stixid
		rawData.DateAdded = dateAdded
		rawData.STIXVersion = modified
		rawData.SpecVersion = specVersion

		collectionRawData = append(collectionRawData, rawData)
	}

	metaData.Size = len(collectionRawData)

	// If no records are returned, then return an error before processing anything else.
	if metaData.Size == 0 {
		return nil, nil, errors.New("no records returned")
	}

	first, last, errRange := ds.processRangeValues(query.RangeBegin, query.RangeEnd, query.RangeMax, metaData.Size)

	if errRange != nil {
		return nil, nil, errRange
	}

	// Get a new slice based on the range of records
	rangeCollectionRawData = collectionRawData[first:last]
	metaData.DateAddedFirst = rangeCollectionRawData[0].DateAdded
	metaData.DateAddedLast = rangeCollectionRawData[len(rangeCollectionRawData)-1].DateAdded
	metaData.RangeBegin = first
	metaData.RangeEnd = last - 1

	// metaData is already a pointer
	return &rangeCollectionRawData, &metaData, nil
}

/*
GetManifestFromCollection - This method will take in query struct with range
parameters for a collection and will return a TAXII manifest that contains all
of the records that match the query and range parameters.
*/
func (ds *Sqlite3DatastoreType) GetManifestFromCollection(query datastore.QueryType) (*resources.ManifestType, *datastore.QueryReturnDataType, error) {
	manifest := resources.NewManifest()
	rangeManifest := resources.NewManifest()
	var metaData datastore.QueryReturnDataType
	var first, last int
	var errRange error

	sqlStmt, err := ds.sqlManifestDataFromCollection(query)

	// If an error is found, that means a query parameter was passed incorrectly
	// and we should return an error versus just skipping the option.
	if err != nil {
		return nil, nil, err
	}

	// Query database for all the collection entries
	rows, err := ds.DB.Query(sqlStmt)
	if err != nil {
		return nil, nil, fmt.Errorf("database execution error querying collection content: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var dateAdded, stixid, modified, specVersion string
		if err := rows.Scan(&dateAdded, &stixid, &modified, &specVersion); err != nil {
			return nil, nil, fmt.Errorf("database scan error: ", err)
		}
		manifest.CreateManifestEntry(stixid, dateAdded, modified, specVersion)
	}

	metaData.Size = len(manifest.Objects)

	first, last, errRange = ds.processRangeValues(query.RangeBegin, query.RangeEnd, query.RangeMax, metaData.Size)

	if errRange != nil {
		return nil, nil, errRange
	}

	// Get a new slice based on the range of records
	rangeManifest.Objects = manifest.Objects[first:last]
	metaData.DateAddedFirst = rangeManifest.Objects[0].DateAdded
	metaData.DateAddedLast = rangeManifest.Objects[len(rangeManifest.Objects)-1].DateAdded

	return &rangeManifest, &metaData, nil
}

// ----------------------------------------------------------------------
//
// Private Methods
//
// ----------------------------------------------------------------------

/*
processRangeValues - This method will take in the various range parameters and size
of the dataset and will return the correct first and last index values to be used.
*/
func (ds *Sqlite3DatastoreType) processRangeValues(first, last, max, size int) (int, int, error) {

	if first < 0 {
		return 0, 0, errors.New("the starting value can not be negative")
	}

	if first > last {
		return 0, 0, errors.New("the starting range value is larger than the ending range value")
	}

	if first >= size {
		return 0, 0, errors.New("the starting range value is out of scope")
	}

	// If no range is requested and the server is not forcing it, do nothing.
	if last == 0 && first == 0 && max != 0 {
		last = first + max
	} else {
		// We need to be inclusive of the last value that was provided
		last++
	}

	// If the last record requested is bigger than the total size of the data
	// set the last size to be the size of the data
	if last > size {
		last = size
	}

	// If the request is for more records than the max size will allow, then
	// compute where the new last record should be, but only if the server is
	// forcing a max size.
	if max != 0 && (last-first) > max {
		last = first + max
	}

	return first, last, nil
}
