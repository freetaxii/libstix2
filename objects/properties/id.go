// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import (
	"errors"
	"fmt"

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

/*
Valid - This method will ensure that the ID property is populated and valid.
It will return a true / false and any error information.
*/
func (o *IDProperty) Valid() (bool, error) {
	if o.ID == "" {
		return false, errors.New("the ID property is required, but missing")
	}

	// TODO check to make sure ID is a valid STIX ID
	return true, nil
}

/*
CreateSTIXUUID - This method takes in a string value representing a STIX object
type and creates and returns a new ID based on the approved STIX UUIDv4 format.
*/
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

/*
CompareIDProperties - This function will compare two id properties (object 1 and
object 2) to make sure they are the same. This function will return an integer
that tracks the number of problems and a slice of strings that contain the
detailed results, whether good or bad.
*/
func CompareIDProperties(obj1, obj2 *IDProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check ID Value
	if obj1.ID != obj2.ID {
		problemsFound++
		str := fmt.Sprintf("-- IDs Do Not Match: %s | %s", obj1.ID, obj2.ID)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ IDs Match: %s | %s", obj1.ID, obj2.ID)
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
