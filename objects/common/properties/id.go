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

// IDPropertyType - A property used by one or more STIX objects that
// captures the STIX identifier in string format.
type IDPropertyType struct {
	ID string `json:"id,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - IdPropertyType
// ----------------------------------------------------------------------

// NewSTIXID - This method takes in a string value representing a STIX object
// type and create and return a new ID based on the approved STIX UUIDv4 format.
func (this *IDPropertyType) NewSTIXID(s string) string {
	id := s + "--" + uuid.New()
	return id
}

// CreateID - This method takes in a string value representing a STIX object
// type and create a new ID based on the approved STIX UUIDv4 format and update
// the id property for the object.
func (this *IDPropertyType) CreateID(s string) {
	// TODO Add check to validate input value
	this.ID = this.NewSTIXID(s)
}

// SetID - This method takes in a string value representing an existing STIX id
// and updates the id property for the object.
func (this *IDPropertyType) SetID(s string) {
	this.ID = s
}

// GetID - This method will return the id for a given STIX object.
func (this *IDPropertyType) GetID() string {
	return this.ID
}
