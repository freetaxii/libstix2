// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package courseofaction

import (
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Type
// ----------------------------------------------------------------------

/* CourseOfAction - This type implements the STIX 2 Course Of Action SDO and
defines all of the properties and methods needed to create and work with this
object. All of the methods not defined local to this type are inherited from the
individual properties. */
type CourseOfAction struct {
	objects.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
}

// TODO Finish fleshing out this model to 2.1

/* GetProperties - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync. */
func (o *CourseOfAction) GetPropertyList() []string {
	return []string{"name", "description"}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/* New - This function will create a new STIX Course of Action object and return
it as a pointer. */
func New() *CourseOfAction {
	var obj CourseOfAction
	obj.InitSDO("course-of-action")
	return &obj
}
