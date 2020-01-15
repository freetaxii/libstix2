// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package opinion

import (
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/* Opinion - This type implements the STIX 2 Opinion SDO and defines all of the
properties and methods needed to create and work with this object. All of the
methods not defined local to this type are inherited from the individual
properties. */
type Opinion struct {
	objects.CommonObjectProperties
	Explanation string `json:"explanation,omitempty"`
	properties.AuthorsProperty
	Opinion string `json:"opinion,omitempty"`
	properties.ObjectRefsProperty
}

/* GetProperties - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync. */
func (o *Opinion) GetPropertyList() []string {
	return []string{"explanation", "authors", "opinion", "object_refs"}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/* New - This function will create a new STIX Opinion object and return
it as a pointer. It will also initialize the object by setting all of the basic
properties. */
func New() *Opinion {
	var obj Opinion
	obj.InitSDO("opinion")
	return &obj
}
