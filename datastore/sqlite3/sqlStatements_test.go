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
