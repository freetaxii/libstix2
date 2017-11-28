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
TestGetRangeOfObjects - This function will test the following things:
  1) First value is negative
  2) First value is greater than last value
  3) First value is bigger than size of data
  4) Last value is bigger than size of data
  5) Make sure the size we get back is correct if we get all records
  6) Make sure the size we get back is correct if we only ask for some
  7) Last value minus first value is bigger than the server will allow
  8) Last value minus first value is smaller than the server will allow
  9) The values of first and last are the same
*/
func TestGetRangeOfObjects(t *testing.T) {
	var rangeData []string
	var testData []string
	var size int
	var err error
	var data = []string{"indicator--1", "indicator--2", "indicator--3", "indicator--4", "indicator--5"}
	var ds Sqlite3DatastoreType

	// Test 1: This test should throw an error
	t.Log("Test 1: make sure we get an error if the first value is negative")
	_, _, err = ds.GetRangeOfObjects(data, 5, -1, 1)

	if err == nil {
		t.Error(err)
	}

	// Test 2: This test should throw an error
	t.Log("Test 2: make sure we get an error if the first value is greater than last")
	_, _, err = ds.GetRangeOfObjects(data, 5, 3, 1)

	if err == nil {
		t.Error(err)
	}

	// Test 3: This test should throw an error
	t.Log("Test 3: make sure we get an error if the first value is greater than the size of data")
	_, _, err = ds.GetRangeOfObjects(data, 5, 10, 12)

	if err == nil {
		t.Error(err)
	}

	// Test 4:
	t.Log("Test 4: make sure we get the right data if the last value is greater than the size of data")
	rangeData, _, err = ds.GetRangeOfObjects(data, 20, 0, 12)

	if err != nil {
		t.Error(err)
	}

	testData = []string{"indicator--1", "indicator--2", "indicator--3", "indicator--4", "indicator--5"}
	verifyRangeData(t, rangeData, testData)

	// Test 5:
	t.Log("Test 5: client asks for all data, make sure the size we get back is correct")
	_, size, err = ds.GetRangeOfObjects(data, 20, 0, 4)

	if err != nil {
		t.Error(err)
	}

	if size != 5 {
		t.Errorf("the total size of data is wrong: we got %v and want %v", size, 5)
	}

	// Test 6:
	t.Log("Test 6: client asks for a few records, make sure the size we get back is correct")
	_, size, err = ds.GetRangeOfObjects(data, 20, 1, 3)

	if err != nil {
		t.Error(err)
	}

	if size != 5 {
		t.Errorf("the total size of data is wrong: we got %v and want %v", size, 5)
	}

	// Test 7:
	t.Log("Test 7: client asks for more than server will allow")
	rangeData, _, err = ds.GetRangeOfObjects(data, 2, 1, 3)

	if err != nil {
		t.Error(err)
	}

	testData = []string{"indicator--2", "indicator--3"}
	verifyRangeData(t, rangeData, testData)

	// Test 8:
	t.Log("Test 8: client asks for less than server will allow")
	rangeData, _, err = ds.GetRangeOfObjects(data, 10, 2, 3)

	if err != nil {
		t.Error(err)
	}

	testData = []string{"indicator--3", "indicator--4"}
	verifyRangeData(t, rangeData, testData)

	// Test 9:
	t.Log("Test 9: client asks for a single value")
	rangeData, _, err = ds.GetRangeOfObjects(data, 2, 2, 2)

	if err != nil {
		t.Error(err)
	}

	testData = []string{"indicator--3"}
	verifyRangeData(t, rangeData, testData)

}

func verifyRangeData(t *testing.T, rangeData, testData []string) {
	if rangeData == nil && testData != nil {
		t.Errorf("no results returned: got %v want %v", rangeData, testData)
	} else if len(rangeData) != len(testData) {
		t.Errorf("slice length is not the same: got %v want %v", len(rangeData), len(testData))
	} else {
		for i := range testData {
			if rangeData[i] != testData[i] {
				t.Log("testing index value: ", i)
				t.Errorf("wrong results: got %v want %v", rangeData[i], testData[i])
			}
		}
	}
}

func TestProcessQueryOptions(t *testing.T) {
	var ds Sqlite3DatastoreType
	var q datastore.QueryType
	var err error

	t.Log("Test 1: make sure we get an error if a single type value is wrong")
	q.STIXType = "indicatorr"
	_, err = ds.processQueryOptions(q)

	if err == nil {
		t.Error(err)
	}

	t.Log("Test 2: make sure we get an error if a two type values are wrong")
	q.STIXType = "foo,bar"
	_, err = ds.processQueryOptions(q)

	if err == nil {
		t.Error(err)
	}

	t.Log("Test 3: make sure we get an error if the first type value is correct but the second is wrong")
	q.STIXType = "indicator,bar"
	_, err = ds.processQueryOptions(q)

	if err == nil {
		t.Error(err)
	}

	t.Log("Test 4: make sure we get an error if the first type value is wrong but the second is correct")
	q.STIXType = "foo,indicator"
	_, err = ds.processQueryOptions(q)

	if err == nil {
		t.Error(err)
	}

	t.Log("Test 5: make sure we do not get an error if a single type value is correct")
	q.STIXType = "indicator"
	_, err = ds.processQueryOptions(q)

	if err != nil {
		t.Error(err)
	}

	t.Log("Test 6: make sure we do not get an error if two type values are correct")
	q.STIXType = "indicator,attack-pattern"
	_, err = ds.processQueryOptions(q)

	if err != nil {
		t.Error(err)
	}

}
