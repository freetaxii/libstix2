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
// Identity Table
//
// ----------------------------------------------------------------------

/*
identityProperties  - This method will return the properties for identity SDOs
*/
func identityProperties() string {
	return baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT,
	"identity_class" TEXT NOT NULL,
	"contact_information" TEXT
	`
	// sectors
}

/*
identitySectorsProperties  - This method will return the properties for identity sectors
Used by: identity
*/
func identitySectorsProperties() string {
	return baseProperties() + `
	"sectors" TEXT NOT NULL
	`
}
