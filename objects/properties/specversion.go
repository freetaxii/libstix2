// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import (
	"errors"
	"fmt"
)

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
SpecVersionProperty - A property used by all STIX objects that captures the
STIX specification version.
*/
type SpecVersionProperty struct {
	SpecVersion string `json:"spec_version,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - SpecVersionProperty
// ----------------------------------------------------------------------

/*
Valid - This method will ensure that the spec version property is populated and valid.
It will return a true / false and any error information.
*/
func (o *SpecVersionProperty) Valid() (bool, error) {
	if o.SpecVersion == "" {
		return false, errors.New("the spec version property is required, but missing")
	}
	return true, nil
}

/*
SetSpecVersion20 - This method will set the specification version to 2.0.
*/
func (o *SpecVersionProperty) SetSpecVersion20() error {
	o.SpecVersion = "2.0"
	return nil
}

/*
SetSpecVersion21 - This method will set the specification version to 2.1.
*/
func (o *SpecVersionProperty) SetSpecVersion21() error {
	o.SpecVersion = "2.1"
	return nil
}

/*
SetSpecVersion - This method takes in a string representing a STIX specification
version and updates the Version property.
*/
func (o *SpecVersionProperty) SetSpecVersion(s string) error {
	o.SpecVersion = s
	return nil
}

/*
GetSpecVersion - This method returns the version value as a string.
*/
func (o *SpecVersionProperty) GetSpecVersion() string {
	return o.SpecVersion
}

/*
CompareSpecVersionProperties - This function will compare two spec version
properties (object 1 and object 2) to make sure they are the same. This function
will return an integer that tracks the number of problems and a slice of strings
that contain the detailed results, whether good or bad.
*/
func CompareSpecVersionProperties(obj1, obj2 *SpecVersionProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Spec Version Value
	if obj1.SpecVersion != obj2.SpecVersion {
		problemsFound++
		str := fmt.Sprintf("-- Spec Versions Do Not Match: %s | %s", obj1.SpecVersion, obj2.SpecVersion)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ Spec Versions Match: %s | %s", obj1.SpecVersion, obj2.SpecVersion)
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
