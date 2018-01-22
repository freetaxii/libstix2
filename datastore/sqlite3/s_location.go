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
// Private Functions - Location Table
// Table property names and SQL statements
//
// ----------------------------------------------------------------------

/*
locationProperties - This method will return the properties for location SDOs
*/
func locationProperties() string {
	return baseProperties() + `
	"description" TEXT,
	"latitude" TEXT,
	"longitude" TEXT,
	"precision" TEXT,
	"region" TEXT,
	"country" TEXT,
	"administrative_area" TEXT,
	"city" TEXT,
	"street_address" TEXT,
	"postal_code" TEXT
	`
}
