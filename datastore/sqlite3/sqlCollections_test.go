// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"testing"
)

// ----------------------------------------------------------------------
//
// func (ds *Sqlite3DatastoreType) sqlEnabledCollections() (string, error)
//
// ----------------------------------------------------------------------
func Test_sqlEnabledCollections(t *testing.T) {
	var ds Sqlite3DatastoreType
	var testdata string

	t.Log("Test 1: get correct sql statement for enabled collections")
	testdata = `SELECT t_collections.id, t_collections.title, t_collections.description, t_collections.can_read, t_collections.can_write, group_concat(t_media_types.media_type) FROM t_collections JOIN t_collection_media_type ON t_collections.id = t_collection_media_type.collection_id JOIN t_media_types ON t_collection_media_type.media_type_id = t_media_types.row_id WHERE t_collections.enabled == '1' && t_collections.hidden != '1' GROUP BY t_collections.id`
	if v, _ := ds.sqlEnabledCollections(); testdata != v {
		t.Error("sql statement is not correct")
	}
}
