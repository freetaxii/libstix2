// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package decode_example

import "github.com/freetaxii/libstix2/objects"

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
Example - This is an example STIX 2 object.
*/
type Example struct {
	objects.CommonObjectProperties
	Name    string `json:"name,omitempty" bson:"name,omitempty"`
	Example bool   `json:"example,omitempty" bson:"example,omitempty"`
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *Example) GetPropertyList() []string {
	return []string{
		"name",
		"example",
	}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Malware Example object and return
it as a pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *Example {
	var obj Example
	obj.InitSDO("example")
	return &obj
}
