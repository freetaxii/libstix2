// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import "fmt"

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
RevokedProperty - A property used by one or more STIX objects that
captures whether or not this STIX object has been revoked by the object
// creator.
*/
type RevokedProperty struct {
	Revoked bool `json:"revoked,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - RevokedProperty
// ----------------------------------------------------------------------

/*
SetRevoked - This method sets the revoked boolean to true
*/
func (o *RevokedProperty) SetRevoked() error {
	o.Revoked = true
	return nil
}

/*
GetRevoked - This method returns the current value of the revoked property.
*/
func (o *RevokedProperty) GetRevoked() bool {
	return o.Revoked
}

// ----------------------------------------------------------------------
// Public Functions - RevokedProperty
// ----------------------------------------------------------------------

/* CompareRevokedProperties - This function will compare two properties to make
sure they are the same and will return a boolean, an integer that tracks the
number of problems found, and a slice of strings that contain the detailed
results, whether good or bad. */
func CompareRevokedProperties(obj1, obj2 *RevokedProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Revoked Value
	if obj1.Revoked != obj2.Revoked {
		problemsFound++
		str := fmt.Sprintf("-- The Revoked values do not match: %t | %t", obj1.Revoked, obj2.Revoked)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The Revoked values match: %t | %t", obj1.Revoked, obj2.Revoked)
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
