// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import (
	"fmt"

	"github.com/freetaxii/libstix2/timestamp"
)

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/* CreatedProperty - A property used by all STIX objects that captures the date
and time that the object was created. */
type CreatedProperty struct {
	Created string `json:"created,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - CreatedProperty - Setters
// ----------------------------------------------------------------------

/* SetCreatedToCurrentTime - This methods sets the object created time to the
current time */
func (o *CreatedProperty) SetCreatedToCurrentTime() error {
	o.Created = timestamp.CurrentTime("milli")
	return nil
}

/* SetCreated - This method takes in a timestamp in either time.Time or string
format and updates the created property with it. The value is stored as a
string, so if the value is in time.Time format, it will be converted to the
correct STIX timestamp format. */
func (o *CreatedProperty) SetCreated(t interface{}) error {
	ts, _ := timestamp.ToString(t, "milli")
	o.Created = ts
	return nil
}

/* GetCreated - This method will return the created timestamp as a string. */
func (o *CreatedProperty) GetCreated() string {
	return o.Created
}

// ----------------------------------------------------------------------
// Public Methods - CreatedProperty - Checks
// ----------------------------------------------------------------------

/* VerifyExists - This method will verify that the created property on an object
is present if required. It will return a boolean, an integer that tracks the
number of problems found, and a slice of strings that contain the detailed
results, whether good or bad. */
func (o *CreatedProperty) VerifyExists() (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 1)

	if o.Created == "" {
		problemsFound++
		resultDetails[0] = fmt.Sprintf("-- The created property is required but missing")
		return false, problemsFound, resultDetails
	}

	resultDetails[0] = fmt.Sprintf("++ The created property is required and is present")
	return true, problemsFound, resultDetails
}

/* Compare - This method will compare two properties to make sure they are the
same and will return a boolean, an integer that tracks the number of problems
found, and a slice of strings that contain the detailed results, whether good or
bad. */
func (o *CreatedProperty) Compare(obj2 *CreatedProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Created Value
	if o.Created != obj2.Created {
		problemsFound++
		str := fmt.Sprintf("-- The created dates do not match: %s | %s", o.Created, obj2.Created)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The created dates match: %s | %s", o.Created, obj2.Created)
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
