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
// Private Functions - Threat Actor Table
// Table property names and SQL statements
//
// ----------------------------------------------------------------------

/*
threatActorProperties  - This method will return the properties for threat actor SDOs
*/
func threatActorProperties() string {
	return baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT,
	"sophistication" TEXT,
	"resource_level" TEXT,
	"primary_motivation" TEXT
	`
	// aliases
	// roles
	// goals
	// secondary_motivations
	// personal_motivations
}

/*
threatActorRolesProperties  - This method will return the properties for threat actor roles
Used by: threat actor
*/
func threatActorRolesProperties() string {
	return baseProperties() + `
	"roles" TEXT NOT NULL
	`
}
