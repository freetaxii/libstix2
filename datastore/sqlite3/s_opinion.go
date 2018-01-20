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
// Opinion Table
//
// ----------------------------------------------------------------------

/*
opinionProperties - This method will return the properties for opinion SDOs
*/
func opinionProperties() string {
	return baseProperties() + `
	"description" TEXT,
	"opinion" TEXT
	`
	// authors
	// object_refs
}
