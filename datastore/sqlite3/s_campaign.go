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
// Campaign Table
//
// ----------------------------------------------------------------------

/*
campaignProperties  - This method will return the properties for campaign SDOs
*/
func campaignProperties() string {
	return baseProperties() + `
	"name" TEXT NOT NULL,
	"description" TEXT,
	"first_seen" TEXT,
	"last_seen" TEXT,
	"objective" TEXT
	`
	// aliases
}
