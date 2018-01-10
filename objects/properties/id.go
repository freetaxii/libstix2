// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import (
	"github.com/pborman/uuid"
)

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
IDPropertyType - A property used by one or more STIX objects that
captures the STIX identifier in string format.
*/
type IDPropertyType struct {
	ID string `json:"id,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - IdPropertyType
// ----------------------------------------------------------------------

/*
CreateSTIXUUID - This method takes in a string value representing a STIX object
type and create and return a new ID based on the approved STIX UUIDv4 format.
*/
func (ezt *IDPropertyType) CreateSTIXUUID(s string) (string, error) {
	id := s + "--" + uuid.New()
	return id, nil
}

/*
GetNewID - This method takes in a string value representing a STIX object
type and create a new ID based on the approved STIX UUIDv4 format and update
the id property for the object.
*/
func (ezt *IDPropertyType) GetNewID(s string) error {
	// TODO Add check to validate input value
	ezt.ID, _ = ezt.CreateSTIXUUID(s)
	return nil
}

/*
SetID - This method takes in a string value representing an existing STIX id
and updates the id property for the object.
*/
func (ezt *IDPropertyType) SetID(s string) error {
	ezt.ID = s
	return nil
}

/*
GetID - This method will return the id for a given STIX object.
*/
func (ezt *IDPropertyType) GetID() (string, error) {
	return ezt.ID, nil
}
