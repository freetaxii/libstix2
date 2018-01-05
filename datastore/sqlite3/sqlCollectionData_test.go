// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"bytes"
	"github.com/freetaxii/libstix2/datastore"
	"testing"
)

// ----------------------------------------------------------------------
//
// func (ds *Sqlite3DatastoreType) sqlObjectList(query datastore.QueryType) (string, error)
//
// ----------------------------------------------------------------------
func Test_sqlObjectList(t *testing.T) {
	var ds Sqlite3DatastoreType
	var query datastore.QueryType
	var testdata string

	t.Log("Test 1: get an error for no collection id")
	if _, err := ds.sqlObjectList(query); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 2: get correct sql statement for object list")
	query.CollectionID = "aa"
	testdata = `SELECT t_collection_data.date_added, t_collection_data.stix_id, s_base_object.modified, s_base_object.spec_version FROM t_collection_data JOIN s_base_object ON t_collection_data.stix_id = s_base_object.id WHERE t_collection_data.collection_id = "aa"`
	if v, _ := ds.sqlObjectList(query); testdata != v {
		t.Error("sql statement is not correct")
	}
}

// ----------------------------------------------------------------------
//
// func (ds *Sqlite3DatastoreType) sqlManifestData(query datastore.QueryType) (string, error)
//
// ----------------------------------------------------------------------
func Test_sqlManifestData(t *testing.T) {
	var ds Sqlite3DatastoreType
	var query datastore.QueryType
	var testdata string

	t.Log("Test 1: get an error for no collection id")
	if _, err := ds.sqlManifestData(query); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 2: get correct sql statement for manifest data")
	query.CollectionID = "aa"
	testdata = `SELECT t_collection_data.date_added, t_collection_data.stix_id, group_concat(s_base_object.modified), group_concat(s_base_object.spec_version) FROM t_collection_data JOIN s_base_object ON t_collection_data.stix_id = s_base_object.id WHERE t_collection_data.collection_id = "aa" GROUP BY t_collection_data.date_added`
	if v, _ := ds.sqlManifestData(query); testdata != v {
		t.Error("sql statement is not correct")
	}
}

// ----------------------------------------------------------------------
//
// func (ds *Sqlite3DatastoreType) sqlWhereCollectionID(id string, b *bytes.Buffer) error
//
// ----------------------------------------------------------------------
func Test_sqlCollectionDataWhereCollectionID(t *testing.T) {
	var ds Sqlite3DatastoreType
	var b bytes.Buffer
	var testdata string

	t.Log("Test 1: get an error for no collection id")
	if err := ds.sqlCollectionDataWhereCollectionID("", &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 2: get correct where statement for collection id")
	b.Reset()
	testdata = `t_collection_data.collection_id = "aa"`
	if ds.sqlCollectionDataWhereCollectionID("aa", &b); testdata != b.String() {
		t.Error("sql where statement is not correct")
	}
}

// ----------------------------------------------------------------------
//
// func (ds *Sqlite3DatastoreType) sqlWhereAddedAfter(date []string, b *bytes.Buffer) error
//
// ----------------------------------------------------------------------
func Test_sqlCollectionDataWhereAddedAfter(t *testing.T) {
	var ds Sqlite3DatastoreType
	var b bytes.Buffer
	var testdata string

	t.Log("Test 1: get an error for invalid timestamp")
	if err := ds.sqlCollectionDataWhereAddedAfter([]string{"20011"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 2: do not get an error if year date is used for added after")
	if err := ds.sqlCollectionDataWhereAddedAfter([]string{"2017"}, &b); err != nil {
		t.Error(err)
	}

	t.Log("Test 3: do not get an error if full date is used for added after")
	if err := ds.sqlCollectionDataWhereAddedAfter([]string{"2017-03-02"}, &b); err != nil {
		t.Error(err)
	}

	t.Log("Test 4: do not get an error if a full timestamp (micro) is used for added after")
	if err := ds.sqlCollectionDataWhereAddedAfter([]string{"2017-03-02T01:01:01.123456Z"}, &b); err != nil {
		t.Error(err)
	}

	t.Log("Test 5: do not get an error if a full timestamp (milli) is used for added after")
	if err := ds.sqlCollectionDataWhereAddedAfter([]string{"2017-03-02T01:01:01.123Z"}, &b); err != nil {
		t.Error(err)
	}

	t.Log("Test 6: get an error if the timezone Z is left off")
	if err := ds.sqlCollectionDataWhereAddedAfter([]string{"2017-03-02T01:01:01"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 7: get an error if the timestamp is incorrectly formatted")
	if err := ds.sqlCollectionDataWhereAddedAfter([]string{"2017-03-02 01:01:01"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 8: get correct where statement for added after")
	b.Reset()
	testdata = ` AND t_collection_data.date_added > "2017"`
	if ds.sqlCollectionDataWhereAddedAfter([]string{"2017"}, &b); testdata != b.String() {
		t.Error("sql where statement is not correct")
	}

}

// ----------------------------------------------------------------------
//
// func (ds *Sqlite3DatastoreType) sqlWhereSTIXID(id []string, b *bytes.Buffer) error
//
// ----------------------------------------------------------------------
func Test_sqlCollectionDataWhereSTIXID(t *testing.T) {
	var ds Sqlite3DatastoreType
	var b bytes.Buffer
	var testdata string

	t.Log("Test 1: get an error for invalid stix id")
	if err := ds.sqlCollectionDataWhereSTIXID([]string{"foo--1234"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 2: get correct where statement for single stix id")
	b.Reset()
	testdata = ` AND t_collection_data.stix_id = "indicator--37abef16-7616-439c-86be-23712030c4b7"`
	if ds.sqlCollectionDataWhereSTIXID([]string{"indicator--37abef16-7616-439c-86be-23712030c4b7"}, &b); testdata != b.String() {
		t.Error("sql where statement is not correct")
	}

	t.Log("Test 3: get an error for one valid and one invalid stix id")
	if err := ds.sqlCollectionDataWhereSTIXID([]string{"indicator--37abef16-7616-439c-86be-23712030c4b7", "foo--1234"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 4: get correct where statement for three stix ids")
	b.Reset()
	testdata = ` AND (t_collection_data.stix_id = "indicator--37abef16-7616-439c-86be-23712030c4b7" OR t_collection_data.stix_id = "attack-pattern--c7c8a099-70a9-487b-a95f-2498d2941104" OR t_collection_data.stix_id = "campaign--6f938db5-6648-4ec1-81cb-5b65138c3c66")`
	if ds.sqlCollectionDataWhereSTIXID([]string{"indicator--37abef16-7616-439c-86be-23712030c4b7", "attack-pattern--c7c8a099-70a9-487b-a95f-2498d2941104", "campaign--6f938db5-6648-4ec1-81cb-5b65138c3c66"}, &b); testdata != b.String() {
		t.Error("sql where statement is not correct")
	}
}

// ----------------------------------------------------------------------
//
// func (ds *Sqlite3DatastoreType) sqlCollectionDataWhereSTIXType(t []string, b *bytes.Buffer) error
//
// ----------------------------------------------------------------------
func Test_sqlCollectionDataWhereSTIXType(t *testing.T) {
	var ds Sqlite3DatastoreType
	var b bytes.Buffer
	var testdata string

	t.Log("Test 1: get an error for invalid stix type")
	if err := ds.sqlCollectionDataWhereSTIXType([]string{"indicatorr"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 2: get an error if two stix type values are invalid")
	if err := ds.sqlCollectionDataWhereSTIXType([]string{"foo", "bar"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 3: get an error for one valid and one invalid stix type")
	if err := ds.sqlCollectionDataWhereSTIXType([]string{"indicator", "indicatorr"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 4: get an error if first is invalid and second is valid")
	if err := ds.sqlCollectionDataWhereSTIXType([]string{"foo", "indicator"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 5: do not get an error if a single type value is correct")
	if err := ds.sqlCollectionDataWhereSTIXType([]string{"indicator"}, &b); err != nil {
		t.Error(err)
	}

	t.Log("Test 6: do not get an error if two type values are correct")
	if err := ds.sqlCollectionDataWhereSTIXType([]string{"indicator", "attack-pattern"}, &b); err != nil {
		t.Error(err)
	}

	t.Log("Test 7: get correct where statement for single stix type")
	b.Reset()
	testdata = ` AND t_collection_data.stix_id LIKE "indicator%"`
	if ds.sqlCollectionDataWhereSTIXType([]string{"indicator"}, &b); testdata != b.String() {
		t.Error("sql where statement is not correct")
	}

	t.Log("Test 8: get correct where statement for three stix types")
	b.Reset()
	testdata = ` AND (t_collection_data.stix_id LIKE "indicator%" OR t_collection_data.stix_id LIKE "attack-pattern%" OR t_collection_data.stix_id LIKE "campaign%")`
	if ds.sqlCollectionDataWhereSTIXType([]string{"indicator", "attack-pattern", "campaign"}, &b); testdata != b.String() {
		t.Error("sql where statement is not correct")
	}
}

// ----------------------------------------------------------------------
//
// func (ds *Sqlite3DatastoreType) sqlCollectionDataWhereSTIXVersion(vers []string, b *bytes.Buffer) error
//
// ----------------------------------------------------------------------
func Test_sqlCollectionDataWhereSTIXVersion(t *testing.T) {
	var ds Sqlite3DatastoreType
	var b bytes.Buffer
	var testdata string

	t.Log("Test 1: get an error for invalid stix version")
	if err := ds.sqlCollectionDataWhereSTIXVersion([]string{"200111"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 2: get correct where statement for single stix version")
	b.Reset()
	testdata = ` AND s_base_object.modified = "2017-12-05T02:43:19.783Z"`
	if ds.sqlCollectionDataWhereSTIXVersion([]string{"2017-12-05T02:43:19.783Z"}, &b); testdata != b.String() {
		t.Error("sql where statement is not correct")
	}

	t.Log("Test 3: get correct where statement for first stix version")
	b.Reset()
	testdata = ` AND s_base_object.modified = (select min(modified) from s_base_object where t_collection_data.stix_id = s_base_object.id)`
	if ds.sqlCollectionDataWhereSTIXVersion([]string{"first"}, &b); testdata != b.String() {
		t.Error("sql where statement is not correct")
	}

	t.Log("Test 4: get correct where statement for last stix version")
	b.Reset()
	testdata = ` AND s_base_object.modified = (select max(modified) from s_base_object where t_collection_data.stix_id = s_base_object.id)`
	if ds.sqlCollectionDataWhereSTIXVersion([]string{"last"}, &b); testdata != b.String() {
		t.Error("sql where statement is not correct")
	}

	t.Log("Test 5: get correct where statement when using the all keyword")
	b.Reset()
	testdata = ``
	if ds.sqlCollectionDataWhereSTIXVersion([]string{"all"}, &b); testdata != b.String() {
		t.Error("sql where statement is not correct")
	}

	t.Log("Test 6: get an error for one valid and one invalid stix version")
	if err := ds.sqlCollectionDataWhereSTIXVersion([]string{"2017-12-05T02:43:19.783Z", "12345"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 7: get correct where statement for first and last stix versions")
	b.Reset()
	testdata = ` AND (s_base_object.modified = (select min(modified) from s_base_object where t_collection_data.stix_id = s_base_object.id) OR s_base_object.modified = (select max(modified) from s_base_object where t_collection_data.stix_id = s_base_object.id))`
	if ds.sqlCollectionDataWhereSTIXVersion([]string{"first", "last"}, &b); testdata != b.String() {
		t.Error("sql where statement is not correct")
	}

	t.Log("Test 8: get correct where statement for three stix versions")
	b.Reset()
	testdata = ` AND (s_base_object.modified = "2017-12-05T02:43:19.783Z" OR s_base_object.modified = "2017-12-05T02:43:23.828Z" OR s_base_object.modified = "2017-12-05T02:43:24.835Z")`
	if ds.sqlCollectionDataWhereSTIXVersion([]string{"2017-12-05T02:43:19.783Z", "2017-12-05T02:43:23.828Z", "2017-12-05T02:43:24.835Z"}, &b); testdata != b.String() {
		t.Error("sql where statement is not correct")
	}

	t.Log("Test 9: get an error when using first, last, and all stix versions")
	if err := ds.sqlCollectionDataWhereSTIXVersion([]string{"first", "last", "all"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 10: get an error when using first, all, and last stix versions")
	if err := ds.sqlCollectionDataWhereSTIXVersion([]string{"first", "all", "last"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 11: get an error when using all, first, and last stix versions")
	if err := ds.sqlCollectionDataWhereSTIXVersion([]string{"all", "first", "last"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 12: get an error when using two 'first' keywords for a stix versions")
	if err := ds.sqlCollectionDataWhereSTIXVersion([]string{"first", "first"}, &b); err == nil {
		t.Error("no error returned")
	}

	t.Log("Test 13: get an error when using two 'last' keywords for a stix versions")
	if err := ds.sqlCollectionDataWhereSTIXVersion([]string{"last", "last"}, &b); err == nil {
		t.Error("no error returned")
	}

}

// ----------------------------------------------------------------------
//
// func (ds *Sqlite3DatastoreType) sqlCollectionDataQueryOptions(query datastore.QueryType) (string, error)
//
// ----------------------------------------------------------------------

func Test_sqlCollectionDataQueryOptions(t *testing.T) {
	var ds Sqlite3DatastoreType
	var q datastore.QueryType

	t.Log("Test 1: get an error for no collection id")
	if _, err := ds.sqlCollectionDataQueryOptions(q); err == nil {
		t.Error("no error returned")
	}

	// Setup for remaining tests
	q.CollectionID = "81f6f8c8-061c-4cb0-97e6-98b317ee5c93"

	t.Log("Test 2: get an error for invalid timestamp")
	q.AddedAfter = nil
	q.AddedAfter = []string{"20111"}
	if _, err := ds.sqlCollectionDataQueryOptions(q); err == nil {
		t.Error("no error returned")
	}
	q.AddedAfter = nil

	t.Log("Test 3: get an error for invalid stix id")
	q.STIXID = []string{"foo--1234"}
	if _, err := ds.sqlCollectionDataQueryOptions(q); err == nil {
		t.Error("no error returned")
	}
	q.STIXID = nil

	t.Log("Test 4: get an error for invalid stix type")
	q.STIXType = []string{"indicatorr"}
	if _, err := ds.sqlCollectionDataQueryOptions(q); err == nil {
		t.Error("no error returned")
	}
	q.STIXType = nil

	t.Log("Test 5: get an error for invalid stix version")
	q.STIXVersion = []string{"200111"}
	if _, err := ds.sqlCollectionDataQueryOptions(q); err == nil {
		t.Error("no error returned")
	}
	q.STIXVersion = nil
}
