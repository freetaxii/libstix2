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
// Private Functions - Observed Data Table
// Table property names and SQL statements
//
// ----------------------------------------------------------------------

/*
observedDataProperties  - This method will return the properties for observed data SDOs
*/
func observedDataProperties() string {
	return baseProperties() + `
	"first_observed" TEXT NOT NULL,
	"last_observed" TEXT NOT NULL,
	"number_observed" INTEGER NOT NULL,
	"objects" TEXT NOT NULL
	`
}
