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
CreatedByRefProperty - A property used by all STIX objects that captures the STIX
identifier of the identity that created this object.
*/
type CreatedByRefProperty struct {
	CreatedByRef string `json:"created_by_ref,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - CreatedByRefProperty
// ----------------------------------------------------------------------

/*
SetCreatedByRef - This method takes in a string value representing a STIX
identifier and updates the Created By Ref property.
*/
func (o *CreatedByRefProperty) SetCreatedByRef(s string) error {
	o.CreatedByRef = s
	return nil
}

/*
GetCreatedByRef - This method returns the STIX identifier for the identity
that created this object.
*/
func (o *CreatedByRefProperty) GetCreatedByRef() string {
	return o.CreatedByRef
}

/*
CompareCreatedByRefProperties - This function will compare two created by ref
properties (object 1 and object 2) to make sure they are the same. This function
will return an integer that tracks the number of problems and a slice of strings
that contain the detailed results, whether good or bad.
*/
func CompareCreatedByRefProperties(obj1, obj2 *CreatedByRefProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Created By Ref Value
	if obj1.CreatedByRef != obj2.CreatedByRef {
		problemsFound++
		str := fmt.Sprintf("-- Created By Refs Do Not Match: %s | %s", obj1.CreatedByRef, obj2.CreatedByRef)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ Created By Refs Match: %s | %s", obj1.CreatedByRef, obj2.CreatedByRef)
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
