// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import "fmt"

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
DescriptionProperty - A property used by one or more STIX objects that
captures the description for the object as a string.
*/
type DescriptionProperty struct {
	Description string `json:"description,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - DescriptionProperty
// ----------------------------------------------------------------------

/*
SetDescription - This method takes in a string value representing a text
description and updates the description property.
*/
func (o *DescriptionProperty) SetDescription(s string) error {
	o.Description = s
	return nil
}

/*
GetDescription - This method returns the description for an object as a string.
*/
func (o *DescriptionProperty) GetDescription() string {
	return o.Description
}

/*
CompareDescriptionProperties - This function will compare two description
properties (object 1 and object 2) to make sure they are the same. This function
will return an integer that tracks the number of problems and a slice of strings
that contain the detailed results, whether good or bad.
*/
func CompareDescriptionProperties(obj1, obj2 *DescriptionProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Description Value
	if obj1.Description != obj2.Description {
		problemsFound++
		str := fmt.Sprintf("-- Descriptions Do Not Match: %s | %s", obj1.Description, obj2.Description)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ Descriptions Match: %s | %s", obj1.Description, obj2.Description)
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
