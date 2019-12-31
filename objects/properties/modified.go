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

/* ModifiedProperty - A property used by all STIX objects that captures the date
and time that the object was modified or changed. This property effectively
tracks the version of the object. */
type ModifiedProperty struct {
	Created  string `json:"created,omitempty"`
	Modified string `json:"modified,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - ModifiedProperty - Setters
// ----------------------------------------------------------------------

/* SetModifiedToCurrentTime - This methods sets the object created time to the
current time */
func (o *ModifiedProperty) SetModifiedToCurrentTime() error {
	o.Created = timestamp.CurrentTime("milli")
	return nil
}

/* SetModified - This method takes in a timestamp in either time.Time or string
format and updates the modifed property with it. The value is stored as a
string, so if the value is in time.Time format, it will be converted to the
correct STIX timestamp format. */
func (o *ModifiedProperty) SetModified(t interface{}) error {
	ts, _ := timestamp.ToString(t, "milli")
	o.Modified = ts
	return nil
}

/* GetModified - This method will return the modified timestamp as a string. If
the value is the same as the created timestamp, then this object is the first
version of the object. */
func (o *ModifiedProperty) GetModified() string {
	return o.Modified
}

// ----------------------------------------------------------------------
// Public Methods - ModifiedProperty - Checks
// ----------------------------------------------------------------------

/* VerifyExists - This method will verify that the modified property on an object
is present if required. It will return a boolean, an integer that tracks the
number of problems found, and a slice of strings that contain the detailed
results, whether good or bad. */
func (o *ModifiedProperty) VerifyExists() (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 1)

	if o.Modified == "" {
		problemsFound++
		resultDetails[0] = fmt.Sprintf("-- The modified property is required but missing")
		return false, problemsFound, resultDetails
	}

	resultDetails[0] = fmt.Sprintf("++ The modified property is required and is present")
	return true, problemsFound, resultDetails
}

/* Compare - This method will compare two properties to make sure they are the
same and will return a boolean, an integer that tracks the number of problems
found, and a slice of strings that contain the detailed results, whether good or
bad. */
func (o *ModifiedProperty) Compare(obj2 *ModifiedProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Modified Value
	if o.Modified != obj2.Modified {
		problemsFound++
		str := fmt.Sprintf("-- The modified dates do not match: %s | %s", o.Modified, obj2.Modified)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The modified dates match: %s | %s", o.Modified, obj2.Modified)
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
