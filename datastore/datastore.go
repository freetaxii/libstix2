// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package datastore

import (
	"github.com/freetaxii/libstix2/datastore/sqlite3"
)

// NewSqlite3 - This is a helper function that will return a sqlite3.Sqlite3DatastoreType
func NewSqlite3(filename string) sqlite3.Sqlite3DatastoreType {
	return sqlite3.New(filename)
}
