// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package courseofaction

import (
	"encoding/json"

	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
//
// Define Object Type
//
// ----------------------------------------------------------------------

/*
CourseOfAction - This type implements the STIX 2 Course of Action SDO and defines
all of the properties methods needed to create and work with the STIX Course of Action
SDO. All of the methods not defined local to this type are inherited from
the individual properties.
*/
type CourseOfAction struct {
	objects.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
}

// ----------------------------------------------------------------------
//
// Initialization Functions
//
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Course of Action object and return it
as a pointer.
*/
func New() *CourseOfAction {
	var obj CourseOfAction
	obj.InitObject("course-of-action")
	return &obj
}

// ----------------------------------------------------------------------
// Public Methods - CourseOfAction - Core Functionality
// ----------------------------------------------------------------------

/*
Decode - This function will decode some JSON data encoded as a slice of bytes
into an actual struct. It will return the object as a pointer and any errors found.
*/
func Decode(data []byte) (*CourseOfAction, error) {
	var o CourseOfAction
	err := json.Unmarshal(data, &o)
	if err != nil {
		return nil, err
	}

	if valid, err := o.Valid(); valid != true {
		return nil, err
	}

	o.SetRawData(data)
	return &o, nil
}

/*
Encode - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *CourseOfAction) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
EncodeToString - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *CourseOfAction) EncodeToString() (string, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

/*
Valid - This method will verify all of the properties on the object.
*/
func (o *CourseOfAction) Valid() (bool, error) {

	// Check common base properties first
	if valid, err := o.CommonObjectProperties.Valid(); valid != true {
		return false, err
	}

	return true, nil
}

// ----------------------------------------------------------------------
// Public Methods - CourseOfAction
// ----------------------------------------------------------------------
