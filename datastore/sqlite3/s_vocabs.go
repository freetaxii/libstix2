// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

// import (
// 	"bytes"
// 	"github.com/freetaxii/libstix2/datastore"
// )

// ----------------------------------------------------------------------
//
// Private Functions - Vocab Table
// Table property names and SQL statements
//
// ----------------------------------------------------------------------

/*
vocabProperties  - This method will return the properties for attack patterns
*/
func vocabProperties() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
	"value" text NOT NULL
	`
}
