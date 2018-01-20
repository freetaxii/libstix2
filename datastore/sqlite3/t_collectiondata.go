// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"bytes"
	"fmt"
	"github.com/freetaxii/libstix2/datastore"
	"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
//
// Collection Data Table Private Functions and Methods
//
// ----------------------------------------------------------------------

/*
collectionDataProperties - This function will return the properties that make up
the collection content table

date_added    = The date that this object was added to the collection
collection_id = The collection ID that this object is tied to
stix_id       = The STIX ID for the object that is being mapped to a collection.
  We do not use the object_id here or the row_id as that would point to a
  specific version and we need to be able to find all versions of an object.
  and if we used row_id for example, it would require two queries, the first
  to get the SITX ID and then the second to get all objects with that STIX ID.
*/
func collectionDataProperties() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
	"date_added" TEXT NOT NULL,
 	"collection_id" TEXT NOT NULL,
 	"stix_id" TEXT NOT NULL
 	`
}
