// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

import (
	"github.com/pborman/uuid"
)

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type IdPropertyType struct {
	Id string `json:"id,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - IdPropertyType
// ----------------------------------------------------------------------

// NewId takes in one parameter
// This function will create a new ID based on the approved STIX UUIDv4 format
// This can be called from functions that just need to make new IDs
// param: s - a string value representing a STIX Object type
func (this *IdPropertyType) NewSTIXId(s string) string {
	id := s + "--" + uuid.New()
	return id
}

// CreateNewId takes in one parameter
// This function will create a new ID based on the approved STIX UUIDv4 format
// param: s - a string value representing a STIX Object type
func (this *IdPropertyType) CreateId(s string) {
	// TODO Add check to validate input value
	this.Id = this.NewSTIXId(s)
}

// SetId takes in one parameter
// param: s - a string value representing a STIX ID
func (this *IdPropertyType) SetId(s string) {
	this.Id = s
}

func (this *IdPropertyType) GetId() string {
	return this.Id
}
