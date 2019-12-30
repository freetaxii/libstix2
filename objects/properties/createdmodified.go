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
// Define Object Model
// ----------------------------------------------------------------------

/* CreatedModifiedProperty - Timestamps to track the created and modified times.

Created - A property used by all STIX objects that captures the date and time
that the object was created.

Modified - A property used by all STIX objects that captures the date and time
that the object was modified or changed. This property effectively tracks the
version of the object.
*/
type CreatedModifiedProperty struct {
	Created  string `json:"created,omitempty"`
	Modified string `json:"modified,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - CreatedModifiedProperty - Setters
// ----------------------------------------------------------------------

/* SetCreatedToCurrentTime - This methods sets the object created time to the
current time */
func (o *CreatedModifiedProperty) SetCreatedToCurrentTime() error {
	o.Created = timestamp.CurrentTime("milli")
	return nil
}

/* SetCreated - This method takes in a timestamp in either time.Time or string
format and updates the created property with it. The value is stored as a
string, so if the value is in time.Time format, it will be converted to the
correct STIX timestamp format. */
func (o *CreatedModifiedProperty) SetCreated(t interface{}) error {
	ts, _ := timestamp.ToString(t, "milli")
	o.Created = ts
	return nil
}

/* GetCreated - This method will return the created timestamp as a string. */
func (o *CreatedModifiedProperty) GetCreated() string {
	return o.Created
}

/* SetModifiedToCreated sets the object modified time to be the same as the
created time. */
func (o *CreatedModifiedProperty) SetModifiedToCreated() error {
	o.Modified = o.Created
	return nil
}

/* SetModified - This method takes in a timestamp in either time.Time or string
format and updates the modifed property with it. The value is stored as a
string, so if the value is in time.Time format, it will be converted to the
correct STIX timestamp format. */
func (o *CreatedModifiedProperty) SetModified(t interface{}) error {
	ts, _ := timestamp.ToString(t, "milli")
	o.Modified = ts
	return nil
}

/* GetModified - This method will return the modified timestamp as a string. If
the value is the same as the created timestamp, then this object is the first
version of the object. */
func (o *CreatedModifiedProperty) GetModified() string {
	return o.Modified
}

// ----------------------------------------------------------------------
// Public Methods - SpecVersionProperty - Checks
// ----------------------------------------------------------------------

/* VerifyPresent - This method will verify that the created and modified
properties on an object is present. It will return a boolean, an integer that
tracks the number of problems found, and a slice of strings that contain the
detailed results, whether good or bad. */
func (o *CreatedModifiedProperty) VerifyPresent() (bool, int, []string) {
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
func (o *CreatedModifiedProperty) Compare(obj1, obj2 *CreatedModifiedProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Created Value
	if o.Created != obj2.Created {
		problemsFound++
		str := fmt.Sprintf("-- The Created dates do not match: %s | %s", o.Created, obj2.Created)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The Created dates match: %s | %s", o.Created, obj2.Created)
		resultDetails = append(resultDetails, str)
	}

	// Check Modified Value
	if o.Modified != obj2.Modified {
		problemsFound++
		str := fmt.Sprintf("-- The Modified dates do not match: %s | %s", o.Modified, obj2.Modified)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The Modified dates match: %s | %s", o.Modified, obj2.Modified)
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
