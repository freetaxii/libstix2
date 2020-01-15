// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import (
	"fmt"

	"github.com/pborman/uuid"
)

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/*
IDProperty - A property used by one or more STIX objects that captures the
STIX That identifier in string format.
*/
type IDProperty struct {
	ID string `json:"id,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - IDProperty - Setters
// ----------------------------------------------------------------------

/*
CreateSTIXUUID - This method takes in a string value representing a STIX
object type and creates and returns a new ID based on the approved STIX UUIDv4
format.
*/
func (o *IDProperty) CreateSTIXUUID(s string) (string, error) {
	// TODO add check to validate that s is a valid type
	id := s + "--" + uuid.New()
	return id, nil
}

/*
CreateTAXIIUUID - This method does not take in any parameters. It is used to
create a new ID based on the approved TAXII UUIDv4 format.
*/
func (o *IDProperty) CreateTAXIIUUID() (string, error) {
	id := uuid.New()
	return id, nil
}

/*
SetNewTAXIIID - This method does not take in any parameters. It is used to
create a new ID based on the approved TAXII UUIDv4 format and assigns it to the
ID property.
*/
func (o *IDProperty) SetNewTAXIIID() error {
	o.ID, _ = o.CreateTAXIIUUID()
	return nil
}

/*
SetNewSTIXID - This method takes in a string value representing a STIX object
type and creates a new ID based on the approved STIX UUIDv4 format and update
the id property for the object.
*/
func (o *IDProperty) SetNewSTIXID(s string) error {
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

// ----------------------------------------------------------------------
// Public Methods - IDProperty - Checks
// ----------------------------------------------------------------------

/*
VerifyExists - This method will verify that the id property on an object is
present if required. It will return a boolean, an integer that tracks the number
of problems found, and a slice of strings that contain the detailed results,
whether good or bad.
*/
func (o *IDProperty) VerifyExists() (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 1)

	if o.ID == "" {
		problemsFound++
		resultDetails[0] = fmt.Sprintf("-- The id property is required but missing")
		return false, problemsFound, resultDetails
	}

	resultDetails[0] = fmt.Sprintf("++ The id property is required and is present")
	return true, problemsFound, resultDetails
}

/*
Compare - This method will compare two properties to make sure they are the
same and will return a boolean, an integer that tracks the number of problems
found, and a slice of strings that contain the detailed results, whether good or
bad.
*/
func (o *IDProperty) Compare(obj2 *IDProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check ID Value
	if o.ID != obj2.ID {
		problemsFound++
		str := fmt.Sprintf("-- The id values do not match: %s | %s", o.ID, obj2.ID)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The id values match: %s | %s", o.ID, obj2.ID)
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
