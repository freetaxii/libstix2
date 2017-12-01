// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"github.com/freetaxii/libstix2/datastore"
	"testing"
)

/*
Test_getRangeValues - This function will test the following method:
getRangeValues(first, last, max, size) (first, last, error)
*/
func Test_getRangeValues(t *testing.T) {
	var ds Sqlite3DatastoreType

	// Test 1: This test should throw an error
	t.Log("Test 1: get an error if the first value is negative")
	if _, _, err := ds.getRangeValues(-1, 1, 5, 10); err == nil {
		t.Error("no error returned")
	}

	// Test 2: This test should throw an error
	t.Log("Test 2: get an error if the first value is greater than last")
	if _, _, err := ds.getRangeValues(5, 3, 5, 10); err == nil {
		t.Error("no error returned")
	}

	// Test 3: This test should throw an error
	t.Log("Test 3: get an error if the first value is greater than the size of data")
	if _, _, err := ds.getRangeValues(100, 200, 50, 50); err == nil {
		t.Error("no error returned")
	}

	// Test 4:
	t.Log("Test 4: check last value if first and last start as zero and pagination is forced")
	if _, last, _ := ds.getRangeValues(0, 0, 10, 20); last != 10 {
		t.Errorf("incorrect range value returned: got %v want %v", last, 9)
	}

	// Test 5:
	t.Log("Test 5: check last value if first and last start as zero and no pagination is forced")
	if _, last, _ := ds.getRangeValues(0, 0, 0, 20); last != 1 {
		t.Errorf("incorrect range value returned: got %v want %v", last, 1)
	}

	// Test 6:
	t.Log("Test 6: last value is larger than size")
	if _, last, _ := ds.getRangeValues(0, 100, 0, 20); last != 20 {
		t.Errorf("incorrect range value returned: got %v want %v", last, 20)
	}

	// Test 7:
	t.Log("Test 7: last value is larger than max")
	if _, last, _ := ds.getRangeValues(0, 15, 10, 20); last != 10 {
		t.Errorf("incorrect range value returned: got %v want %v", last, 10)
	}

	// Test 8:
	t.Log("Test 8: last value is smaller than max")
	if _, last, _ := ds.getRangeValues(0, 3, 10, 20); last != 4 {
		t.Errorf("incorrect range value returned: got %v want %v", last, 10)
	}
}

// 	var rangeData []string
// 	var testData []string
// 	var data = []string{"indicator--1", "indicator--2", "indicator--3", "indicator--4", "indicator--5"}
//
// 	testData = []string{"indicator--1", "indicator--2", "indicator--3", "indicator--4", "indicator--5"}
// 	verifyRangeData(t, rangeData, testData)
// 	testData = []string{"indicator--3"}
// 	verifyRangeData(t, rangeData, testData)
// }

// func verifyRangeData(t *testing.T, rangeData, testData []string) {
// 	if rangeData == nil && testData != nil {
// 		t.Errorf("no results returned: got %v want %v", rangeData, testData)
// 	} else if len(rangeData) != len(testData) {
// 		t.Errorf("slice length is not the same: got %v want %v", len(rangeData), len(testData))
// 	} else {
// 		for i := range testData {
// 			if rangeData[i] != testData[i] {
// 				t.Log("testing index value: ", i)
// 				t.Errorf("wrong results: got %v want %v", rangeData[i], testData[i])
// 			}
// 		}
// 	}
// }

/*
Test_processQueryOptions - This function will test the following method:
processQueryOptions(query) (whereQuery, error)
*/
func Test_processQueryOptions(t *testing.T) {
	var ds Sqlite3DatastoreType
	var q datastore.QueryType

	t.Log("Test 1.1: get an error if no collection id is provided")
	if _, err := ds.processQueryOptions(q); err == nil {
		t.Error("no error returned")
	}

	// Setup for remaining tests
	q.CollectionID = "81f6f8c8-061c-4cb0-97e6-98b317ee5c93"

	t.Log("Test 2.1: do not get an error if year date is used for added after")
	q.AddedAfter = "2017"
	if _, err := ds.processQueryOptions(q); err != nil {
		t.Error(err)
	}

	t.Log("Test 2.2: do not get an error if full date is used for added after")
	q.AddedAfter = "2017-03-02"
	if _, err := ds.processQueryOptions(q); err != nil {
		t.Error(err)
	}

	t.Log("Test 2.3: do not get an error if a full timestamp (micro) is used for added after")
	q.AddedAfter = "2017-03-02T01:01:01.123456Z"
	if _, err := ds.processQueryOptions(q); err != nil {
		t.Error(err)
	}

	t.Log("Test 2.4: do not get an error if a full timestamp (milli) is used for added after")
	q.AddedAfter = "2017-03-02T01:01:01.123Z"
	if _, err := ds.processQueryOptions(q); err != nil {
		t.Error(err)
	}

	t.Log("Test 2.5: get an error if the timezone Z is left off")
	q.AddedAfter = "2017-03-02T01:01:01"
	if _, err := ds.processQueryOptions(q); err == nil {
		t.Error(err)
	}

	t.Log("Test 2.6: get an error if the timestamp is incorrectly formatted")
	q.AddedAfter = "2017-03-02 01:01:01"
	if _, err := ds.processQueryOptions(q); err == nil {
		t.Error(err)
	}

	// Clear out value
	q.AddedAfter = ""

	// TODO test STIX ID

	t.Log("Test 4.1: get an error if a single type value is wrong")
	q.STIXType = "indicatorr"
	if _, err := ds.processQueryOptions(q); err == nil {
		t.Error(err)
	}

	t.Log("Test 4.2: get an error if a two type values are wrong")
	q.STIXType = "foo,bar"
	if _, err := ds.processQueryOptions(q); err == nil {
		t.Error(err)
	}

	t.Log("Test 4.3: get an error if the first type value is correct but the second is wrong")
	q.STIXType = "indicator,bar"
	if _, err := ds.processQueryOptions(q); err == nil {
		t.Error(err)
	}

	t.Log("Test 4.4: get an error if the first type value is wrong but the second is correct")
	q.STIXType = "foo,indicator"
	if _, err := ds.processQueryOptions(q); err == nil {
		t.Error(err)
	}

	t.Log("Test 4.5: do not get an error if a single type value is correct")
	q.STIXType = "indicator"
	if _, err := ds.processQueryOptions(q); err != nil {
		t.Error(err)
	}

	t.Log("Test 4.6: do not get an error if two type values are correct")
	q.STIXType = "indicator,attack-pattern"
	if _, err := ds.processQueryOptions(q); err != nil {
		t.Error(err)
	}

	// TODO test STIX versions

}
