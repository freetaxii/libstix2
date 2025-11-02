// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/freetaxii/libstix2/objects/taxii/collections"
)

// ----------------------------------------------------------------------
//
// func (ds *Sqlite3DatastoreType) sqlGetObjectList(query collections.CollectionQuery) (string, error)
//
// ----------------------------------------------------------------------
func Test_sqlGetObjectList(t *testing.T) {
	var query collections.CollectionQuery
	var testdata string

	t.Log("Test 1: get an error for no collection id")
	if _, err := sqlGetObjectList(query); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 2: get correct sql statement for object list")
	query.CollectionUUID = "aa"
	testdata = `SELECT t_collection_data.date_added, t_collection_data.stix_id, s_base_object.modified, s_base_object.spec_version FROM t_collection_data JOIN s_base_object ON t_collection_data.stix_id = s_base_object.id WHERE t_collection_data.collection_id = "aa"`
	if v, _ := sqlGetObjectList(query); testdata != v {
		t.Error("sql statement is not correct")
	}
}

// ----------------------------------------------------------------------
//
// func (ds *Sqlite3DatastoreType) sqlGetManifestData(query collections.CollectionQuery) (string, error)
//
// ----------------------------------------------------------------------
func Test_sqlGetManifestData(t *testing.T) {
	var query collections.CollectionQuery
	var testdata string

	t.Log("Test 1: get an error for no collection id")
	if _, err := sqlGetManifestData(query); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 2: get correct sql statement for manifest data")
	query.CollectionUUID = "aa"
	testdata = `SELECT t_collection_data.date_added, t_collection_data.stix_id, group_concat(s_base_object.modified), group_concat(s_base_object.spec_version) FROM t_collection_data JOIN s_base_object ON t_collection_data.stix_id = s_base_object.id WHERE t_collection_data.collection_id = "aa" GROUP BY t_collection_data.date_added`
	if v, _ := sqlGetManifestData(query); testdata != v {
		t.Error("sql statement is not correct")
	}
}

// ----------------------------------------------------------------------
//
// func (ds *Sqlite3DatastoreType) sqlCollectionDataQueryOptions(query collections.CollectionQuery) (string, error)
//
// ----------------------------------------------------------------------

func Test_sqlCollectionDataQueryOptions(t *testing.T) {
	var q collections.CollectionQuery

	t.Log("Test 1: get an error for no collection id")
	if _, err := sqlCollectionDataQueryOptions(q); err == nil {
		t.Error("no error returned")
	}

	// Setup for remaining tests
	q.CollectionUUID = "81f6f8c8-061c-4cb0-97e6-98b317ee5c93"

	t.Log("Test 2: get an error for invalid timestamp")
	q.AddedAfter = nil
	q.AddedAfter = []string{"20111"}
	if _, err := sqlCollectionDataQueryOptions(q); err == nil {
		t.Error("no error returned")
	}
	q.AddedAfter = nil

	t.Log("Test 3: get an error for invalid stix id")
	q.STIXID = []string{"foo--1234"}
	if _, err := sqlCollectionDataQueryOptions(q); err == nil {
		t.Error("no error returned")
	}
	q.STIXID = nil

	t.Log("Test 4: get an error for invalid stix type")
	q.STIXType = []string{"indicatorr"}
	if _, err := sqlCollectionDataQueryOptions(q); err == nil {
		t.Error("no error returned")
	}
	q.STIXType = nil

	t.Log("Test 5: get an error for invalid stix version")
	q.STIXVersion = []string{"200111"}
	if _, err := sqlCollectionDataQueryOptions(q); err == nil {
		t.Error("no error returned")
	}
	q.STIXVersion = nil
}

// ----------------------------------------------------------------------
//
// func (ds *Sqlite3DatastoreType) sqlWhereCollectionUUID(id string, b *bytes.Buffer) error
//
// ----------------------------------------------------------------------
func Test_sqlCollectionDataWhereCollectionUUID(t *testing.T) {
	var b bytes.Buffer
	var testdata string

	t.Log("Test 1: get an error for no collection id")
	if err := sqlCollectionDataWhereCollectionUUID("", &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 2: get correct where statement for collection id")
	b.Reset()
	testdata = `t_collection_data.collection_id = "aa"`
	if sqlCollectionDataWhereCollectionUUID("aa", &b); testdata != b.String() {
		t.Error("sql where statement is not correct")
	}
}

// ----------------------------------------------------------------------
//
// func (ds *Sqlite3DatastoreType) sqlWhereAddedAfter(date []string, b *bytes.Buffer) error
//
// ----------------------------------------------------------------------
func Test_sqlCollectionDataWhereAddedAfter(t *testing.T) {
	var b bytes.Buffer
	var testdata string

	t.Log("Test 1: get an error for invalid timestamp")
	if err := sqlCollectionDataWhereAddedAfter([]string{"20011"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 2: do not get an error if year date is used for added after")
	if err := sqlCollectionDataWhereAddedAfter([]string{"2017"}, &b); err != nil {
		t.Error(err)
	}

	t.Log("Test 3: do not get an error if full date is used for added after")
	if err := sqlCollectionDataWhereAddedAfter([]string{"2017-03-02"}, &b); err != nil {
		t.Error(err)
	}

	t.Log("Test 4: do not get an error if a full timestamp (micro) is used for added after")
	if err := sqlCollectionDataWhereAddedAfter([]string{"2017-03-02T01:01:01.123456Z"}, &b); err != nil {
		t.Error(err)
	}

	t.Log("Test 5: do not get an error if a full timestamp (milli) is used for added after")
	if err := sqlCollectionDataWhereAddedAfter([]string{"2017-03-02T01:01:01.123Z"}, &b); err != nil {
		t.Error(err)
	}

	t.Log("Test 6: get an error if the timezone Z is left off")
	if err := sqlCollectionDataWhereAddedAfter([]string{"2017-03-02T01:01:01"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 7: get an error if the timestamp is incorrectly formatted")
	if err := sqlCollectionDataWhereAddedAfter([]string{"2017-03-02 01:01:01"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 8: get correct where statement for added after")
	b.Reset()
	testdata = ` AND t_collection_data.date_added > "2017"`
	if sqlCollectionDataWhereAddedAfter([]string{"2017"}, &b); testdata != b.String() {
		t.Error("sql where statement is not correct")
	}

}

// ----------------------------------------------------------------------
//
// func (ds *Sqlite3DatastoreType) sqlWhereSTIXID(id []string, b *bytes.Buffer) error
//
// ----------------------------------------------------------------------
func Test_sqlCollectionDataWhereSTIXID(t *testing.T) {
	var b bytes.Buffer
	var testdata string

	t.Log("Test 1: get an error for invalid stix id")
	if err := sqlCollectionDataWhereSTIXID([]string{"foo--1234"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 2: get correct where statement for single stix id")
	b.Reset()
	testdata = ` AND t_collection_data.stix_id = "indicator--37abef16-7616-439c-86be-23712030c4b7"`
	if sqlCollectionDataWhereSTIXID([]string{"indicator--37abef16-7616-439c-86be-23712030c4b7"}, &b); testdata != b.String() {
		t.Error("sql where statement is not correct")
	}

	t.Log("Test 3: get an error for one valid and one invalid stix id")
	if err := sqlCollectionDataWhereSTIXID([]string{"indicator--37abef16-7616-439c-86be-23712030c4b7", "foo--1234"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 4: get correct where statement for three stix ids")
	b.Reset()
	testdata = ` AND (t_collection_data.stix_id = "indicator--37abef16-7616-439c-86be-23712030c4b7" OR t_collection_data.stix_id = "attack-pattern--c7c8a099-70a9-487b-a95f-2498d2941104" OR t_collection_data.stix_id = "campaign--6f938db5-6648-4ec1-81cb-5b65138c3c66")`
	if sqlCollectionDataWhereSTIXID([]string{"indicator--37abef16-7616-439c-86be-23712030c4b7", "attack-pattern--c7c8a099-70a9-487b-a95f-2498d2941104", "campaign--6f938db5-6648-4ec1-81cb-5b65138c3c66"}, &b); testdata != b.String() {
		t.Error("sql where statement is not correct")
	}
}

// ----------------------------------------------------------------------
//
// func (ds *Sqlite3DatastoreType) sqlCollectionDataWhereSTIXType(t []string, b *bytes.Buffer) error
//
// ----------------------------------------------------------------------
func Test_sqlCollectionDataWhereSTIXType(t *testing.T) {
	var b bytes.Buffer
	var testdata string

	t.Log("Test 1: get an error for invalid stix type")
	if err := sqlCollectionDataWhereSTIXType([]string{"indicatorr"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 2: get an error if two stix type values are invalid")
	if err := sqlCollectionDataWhereSTIXType([]string{"foo", "bar"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 3: get an error for one valid and one invalid stix type")
	if err := sqlCollectionDataWhereSTIXType([]string{"indicator", "indicatorr"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 4: get an error if first is invalid and second is valid")
	if err := sqlCollectionDataWhereSTIXType([]string{"foo", "indicator"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 5: do not get an error if a single type value is correct")
	if err := sqlCollectionDataWhereSTIXType([]string{"indicator"}, &b); err != nil {
		t.Error(err)
	}

	t.Log("Test 6: do not get an error if two type values are correct")
	if err := sqlCollectionDataWhereSTIXType([]string{"indicator", "attack-pattern"}, &b); err != nil {
		t.Error(err)
	}

	t.Log("Test 7: get correct where statement for single stix type")
	b.Reset()
	testdata = ` AND t_collection_data.stix_id LIKE "indicator%"`
	if sqlCollectionDataWhereSTIXType([]string{"indicator"}, &b); testdata != b.String() {
		t.Error("sql where statement is not correct")
	}

	t.Log("Test 8: get correct where statement for three stix types")
	b.Reset()
	testdata = ` AND (t_collection_data.stix_id LIKE "indicator%" OR t_collection_data.stix_id LIKE "attack-pattern%" OR t_collection_data.stix_id LIKE "campaign%")`
	if sqlCollectionDataWhereSTIXType([]string{"indicator", "attack-pattern", "campaign"}, &b); testdata != b.String() {
		t.Error("sql where statement is not correct")
	}
}

// ----------------------------------------------------------------------
//
// func (ds *Sqlite3DatastoreType) sqlCollectionDataWhereSTIXVersion(vers []string, b *bytes.Buffer) error
//
// ----------------------------------------------------------------------
func Test_sqlCollectionDataWhereSTIXVersion(t *testing.T) {
	var b bytes.Buffer
	var testdata string

	t.Log("Test 1: get an error for invalid stix version")
	if err := sqlCollectionDataWhereSTIXVersion([]string{"200111"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 2: get correct where statement for single stix version")
	b.Reset()
	testdata = ` AND s_base_object.modified = "2017-12-05T02:43:19.783Z"`
	if sqlCollectionDataWhereSTIXVersion([]string{"2017-12-05T02:43:19.783Z"}, &b); testdata != b.String() {
		t.Error("sql where statement is not correct")
	}

	t.Log("Test 3: get correct where statement for first stix version")
	b.Reset()
	testdata = ` AND s_base_object.modified = (select min(modified) from s_base_object where t_collection_data.stix_id = s_base_object.id)`
	if sqlCollectionDataWhereSTIXVersion([]string{"first"}, &b); testdata != b.String() {
		t.Error("sql where statement is not correct")
	}

	t.Log("Test 4: get correct where statement for last stix version")
	b.Reset()
	testdata = ` AND s_base_object.modified = (select max(modified) from s_base_object where t_collection_data.stix_id = s_base_object.id)`
	if sqlCollectionDataWhereSTIXVersion([]string{"last"}, &b); testdata != b.String() {
		t.Error("sql where statement is not correct")
	}

	t.Log("Test 5: get correct where statement when using the all keyword")
	b.Reset()
	testdata = ``
	if sqlCollectionDataWhereSTIXVersion([]string{"all"}, &b); testdata != b.String() {
		t.Error("sql where statement is not correct")
	}

	t.Log("Test 6: get an error for one valid and one invalid stix version")
	if err := sqlCollectionDataWhereSTIXVersion([]string{"2017-12-05T02:43:19.783Z", "12345"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 7: get correct where statement for first and last stix versions")
	b.Reset()
	testdata = ` AND (s_base_object.modified = (select min(modified) from s_base_object where t_collection_data.stix_id = s_base_object.id) OR s_base_object.modified = (select max(modified) from s_base_object where t_collection_data.stix_id = s_base_object.id))`
	if sqlCollectionDataWhereSTIXVersion([]string{"first", "last"}, &b); testdata != b.String() {
		t.Error("sql where statement is not correct")
	}

	t.Log("Test 8: get correct where statement for three stix versions")
	b.Reset()
	testdata = ` AND (s_base_object.modified = "2017-12-05T02:43:19.783Z" OR s_base_object.modified = "2017-12-05T02:43:23.828Z" OR s_base_object.modified = "2017-12-05T02:43:24.835Z")`
	if sqlCollectionDataWhereSTIXVersion([]string{"2017-12-05T02:43:19.783Z", "2017-12-05T02:43:23.828Z", "2017-12-05T02:43:24.835Z"}, &b); testdata != b.String() {
		t.Error("sql where statement is not correct")
	}

	t.Log("Test 9: get an error when using first, last, and all stix versions")
	if err := sqlCollectionDataWhereSTIXVersion([]string{"first", "last", "all"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 10: get an error when using first, all, and last stix versions")
	if err := sqlCollectionDataWhereSTIXVersion([]string{"first", "all", "last"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 11: get an error when using all, first, and last stix versions")
	if err := sqlCollectionDataWhereSTIXVersion([]string{"all", "first", "last"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 12: get an error when using two 'first' keywords for a stix versions")
	if err := sqlCollectionDataWhereSTIXVersion([]string{"first", "first"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 13: get an error when using two 'last' keywords for a stix versions")
	if err := sqlCollectionDataWhereSTIXVersion([]string{"last", "last"}, &b); err == nil {
		t.Error("no error returned")
	}

}

/*
Test_processRangeValues - This function will test the following method:
processRangeValues(first, last, max, size) (first, last, error)
*/
func Test_processRangeValues(t *testing.T) {
	var ds Store
	ds.Logger = slog.Default()

	// Test 1: This test should throw an error
	t.Log("Test 1: get an error if the first value is negative")
	if _, _, err := ds.processRangeValues(-1, 1, 5, 10); err == nil {
		t.Error("no error returned")
	}

	// Test 2: This test should throw an error
	t.Log("Test 2: get an error if the first value is greater than last")
	if _, _, err := ds.processRangeValues(5, 3, 5, 10); err == nil {
		t.Error("no error returned")
	}

	// Test 3: This test should throw an error
	t.Log("Test 3: get an error if the first value is greater than the size of data")
	if _, _, err := ds.processRangeValues(100, 200, 50, 50); err == nil {
		t.Error("no error returned")
	}

	// Test 4:
	t.Log("Test 4: check last value if first and last start as zero and pagination is forced")
	if _, last, _ := ds.processRangeValues(0, 0, 10, 20); last != 10 {
		t.Errorf("incorrect range value returned: got %v want %v", last, 9)
	}

	// Test 5:
	t.Log("Test 5: check last value if first and last start as zero and no pagination is forced")
	if _, last, _ := ds.processRangeValues(0, 0, 0, 20); last != 1 {
		t.Errorf("incorrect range value returned: got %v want %v", last, 1)
	}

	// Test 6:
	t.Log("Test 6: last value is larger than size")
	if _, last, _ := ds.processRangeValues(0, 100, 0, 20); last != 20 {
		t.Errorf("incorrect range value returned: got %v want %v", last, 20)
	}

	// Test 7:
	t.Log("Test 7: last value is larger than max")
	if _, last, _ := ds.processRangeValues(0, 15, 10, 20); last != 10 {
		t.Errorf("incorrect range value returned: got %v want %v", last, 10)
	}

	// Test 8:
	t.Log("Test 8: last value is smaller than max")
	if _, last, _ := ds.processRangeValues(0, 3, 10, 20); last != 4 {
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
