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
// Note Table
//
// ----------------------------------------------------------------------

/*
noteProperties  - This method will return the properties for note SDOs
*/
func noteProperties() string {
	return baseProperties() + `
	"summary" TEXT,
	"description" TEXT NOT NULL
	`
	// authors
	// object_refs
}
