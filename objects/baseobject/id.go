// Copyright 2015-2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package baseobject

import (
	"github.com/pborman/uuid"
)

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
IDProperty - A property used by one or more STIX objects that captures the STIX
That identifier in string format.
*/
type IDProperty struct {
	ID string `json:"id,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - IdProperty
// ----------------------------------------------------------------------

// CreateSTIXUUID - This method takes in a string value representing a STIX object
// type and creates and returns a new ID based on the approved STIX UUIDv4 format.

func (o *IDProperty) CreateSTIXUUID(s string) (string, error) {
	// TODO add check to validate that s is a valid type
	id := s + "--" + uuid.New()
	return id, nil
}

/*
SetNewID - This method takes in a string value representing a STIX object
type and creates a new ID based on the approved STIX UUIDv4 format and update
the id property for the object.
*/
func (o *IDProperty) SetNewID(s string) error {
	// TODO Add check to validate input value
	o.ID, _ = o.CreateSTIXUUID(s)
	return nil
}

/*
SetID - This method takes in a string value representing an existing STIX id
and updates the id property for the object.
*/
func (o *IDProperty) SetID(s string) error {
	o.ID = s
	return nil
}

/*
GetID - This method will return the id for a given STIX object.
*/
func (o *IDProperty) GetID() string {
	return o.ID
}
