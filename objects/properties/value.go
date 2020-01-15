// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import "fmt"

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/*
ValueProperty -
*/
type ValueProperty struct {
	Value string `json:"value,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - ValueProperty - Setters
// ----------------------------------------------------------------------

/*
SetValue -
*/
func (o *ValueProperty) SetValue(val string) error {
	o.Value = val
	return nil
}

// ----------------------------------------------------------------------
// Public Functions - NameProperty - Checks
// ----------------------------------------------------------------------

/*
VerifyExists - This method will verify that the value property on an object
is present. It will return a boolean, an integer that tracks the number of
problems found, and a slice of strings that contain the detailed results,
whether good or bad.
*/
func (o *ValueProperty) VerifyExists() (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 1)

	if o.Value == "" {
		problemsFound++
		resultDetails[0] = fmt.Sprintf("-- The value property is required but missing")
		return false, problemsFound, resultDetails
	}

	resultDetails[0] = fmt.Sprintf("++ The value property is required and is present")
	return true, problemsFound, resultDetails
}
