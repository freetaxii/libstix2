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
TypeProperty - A property used by one or more STIX objects that
captures the STIX object type in string format.
*/
type TypeProperty struct {
	ObjectType string `json:"type,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - TypeProperty
// ----------------------------------------------------------------------

/*
Valid - This method will ensure that the type property is populated and valid.
It will return a true / false and any error information.
*/
func (o *TypeProperty) Valid() (bool, error) {
	if o.ObjectType == "" {
		return false, errors.New("the type property is required, but missing")
	}
	return true, nil
}

/*
SetObjectType - This method takes in a string value representing a STIX object
type and updates the type property.
*/
func (o *TypeProperty) SetObjectType(s string) error {
	o.ObjectType = s
	return nil
}

/*
GetObjectType - This method returns the object type.
*/
func (o *TypeProperty) GetObjectType() string {
	return o.ObjectType
}

/*
CompareTypeProperties - This function will compare two type properties (object 1
and object 2) to make sure they are the same. This function will return an
integer that tracks the number of problems and a slice of strings that contain
the detailed results, whether good or bad.
*/
func CompareTypeProperties(obj1, obj2 *TypeProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Type Value
	if obj1.ObjectType != obj2.ObjectType {
		problemsFound++
		str := fmt.Sprintf("-- Types Do Not Match: %s | %s", obj1.ObjectType, obj2.ObjectType)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ Types Match: %s | %s", obj1.ObjectType, obj2.ObjectType)
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
