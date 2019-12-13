// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import (
	"errors"
	"fmt"

	"github.com/freetaxii/libstix2/timestamp"
)

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
CreatedModifiedProperty - Timestamps to track the created and modified times.

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
// Public Methods - CreatedModifiedProperty
// ----------------------------------------------------------------------

/*
Valid - This method will ensure that the created and modified properties are populated and valid.
It will return a true / false and any error information.
*/
func (o *CreatedModifiedProperty) Valid() (bool, error) {

	if o.Created == "" {
		return false, errors.New("the created property is required, but missing")
	}

	if o.Modified == "" {
		return false, errors.New("the created property is required, but missing")
	}

	// TODO check to make sure timestamp is a valid STIX timestamp but only if it is defined
	return true, nil
}

/*
SetCreatedToCurrentTime - This methods sets the object created time to the
current time
*/
func (o *CreatedModifiedProperty) SetCreatedToCurrentTime() error {
	o.Created = timestamp.CurrentTime("milli")
	return nil
}

/*
SetCreated - This method takes in a timestamp in either time.Time or string
format and updates the created property with it. The value is stored as a
string, so if the value is in time.Time format, it will be converted to the
correct STIX timestamp format.
*/
func (o *CreatedModifiedProperty) SetCreated(t interface{}) error {
	ts, _ := timestamp.ToString(t, "milli")
	o.Created = ts
	return nil
}

/*
GetCreated - This method will return the created timestamp as a string.
*/
func (o *CreatedModifiedProperty) GetCreated() string {
	return o.Created
}

/*
SetModifiedToCreated sets the object modified time to be the same as the
created time.
*/
func (o *CreatedModifiedProperty) SetModifiedToCreated() error {
	o.Modified = o.Created
	return nil
}

/*
SetModifiedToCurrentTime - This methods sets the object modified time to the
current time
*/
func (o *CreatedModifiedProperty) SetModifiedToCurrentTime() error {
	o.Modified = timestamp.CurrentTime("milli")
	return nil
}

/*
SetModified - This method takes in a timestamp in either time.Time or string
format and updates the modifed property with it. The value is stored as a
string, so if the value is in time.Time format, it will be converted to the
correct STIX timestamp format.
*/
func (o *CreatedModifiedProperty) SetModified(t interface{}) error {
	ts, _ := timestamp.ToString(t, "milli")
	o.Modified = ts
	return nil
}

/*
GetModified - This method will return the modified timestamp as a string. If
the value is the same as the created timestamp, then this object is the
first version of the object.
*/
func (o *CreatedModifiedProperty) GetModified() string {
	return o.Modified
}

/*
CompareCreatedModifiedProperties - This function will compare two created and
modified properties (object 1 and object 2) to make sure they are the same. This
function will return an integer that tracks the number of problems and a slice
of strings that contain the detailed results, whether good or bad.
*/
func CompareCreatedModifiedProperties(obj1, obj2 *CreatedModifiedProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Created Value
	if obj1.Created != obj2.Created {
		problemsFound++
		str := fmt.Sprintf("-- Created Dates Do Not Match: %s | %s", obj1.Created, obj2.Created)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ Created Dates Match: %s | %s", obj1.Created, obj2.Created)
		resultDetails = append(resultDetails, str)
	}

	// Check Modified Value
	if obj1.Modified != obj2.Modified {
		problemsFound++
		str := fmt.Sprintf("-- Modified Dates Do Not Match: %s | %s", obj1.Modified, obj2.Modified)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ Modified Dates Match: %s | %s", obj1.Modified, obj2.Modified)
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
