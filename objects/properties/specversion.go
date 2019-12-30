// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import (
	"fmt"
)

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/* SpecVersionProperty - A property used by all STIX objects that captures the
STIX specification version. */
type SpecVersionProperty struct {
	SpecVersion string `json:"spec_version,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - SpecVersionProperty - Setters
// ----------------------------------------------------------------------

/* SetSpecVersion20 - This method will set the specification version to 2.0. */
func (o *SpecVersionProperty) SetSpecVersion20() error {
	o.SpecVersion = "2.0"
	return nil
}

/* SetSpecVersion21 - This method will set the specification version to 2.1. */
func (o *SpecVersionProperty) SetSpecVersion21() error {
	o.SpecVersion = "2.1"
	return nil
}

/* SetSpecVersion - This method takes in a string representing a STIX
specification version and updates the Version property. */
func (o *SpecVersionProperty) SetSpecVersion(s string) error {
	o.SpecVersion = s
	return nil
}

/* GetSpecVersion - This method returns the version value as a string. */
func (o *SpecVersionProperty) GetSpecVersion() string {
	return o.SpecVersion
}

// ----------------------------------------------------------------------
// Public Methods - SpecVersionProperty - Checks
// ----------------------------------------------------------------------

/* VerifyPresent - This method will verify that the spec version property on an
object is present. It will return a boolean, an integer that tracks the number
of problems found, and a slice of strings that contain the detailed results,
whether good or bad. */
func (o *SpecVersionProperty) VerifyPresent() (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 1)

	if o.SpecVersion == "" {
		problemsFound++
		resultDetails[0] = fmt.Sprintf("-- The spec version property is required but missing")
		return false, problemsFound, resultDetails
	}

	resultDetails[0] = fmt.Sprintf("++ The spec version property is required and is present")
	return true, problemsFound, resultDetails
}

/* Compare - This method will compare two properties to make sure they are the
same and will return a boolean, an integer that tracks the number of problems
found, and a slice of strings that contain the detailed results, whether good or
bad. */
func (o *SpecVersionProperty) Compare(obj2 *SpecVersionProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Spec Version Value
	if o.SpecVersion != obj2.SpecVersion {
		problemsFound++
		str := fmt.Sprintf("-- The spec version values do not match: %s | %s", o.SpecVersion, obj2.SpecVersion)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The spec version values match: %s | %s", o.SpecVersion, obj2.SpecVersion)
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
