// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package courseOfAction

import (
	"github.com/freetaxii/libstix2/objects/common/properties"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
CourseOfActionType defines all of the properties associated with the STIX Course
of Action SDO. All of the methods not defined local to this type are inherited
from the individual properties.
*/
type CourseOfActionType struct {
	properties.CommonObjectPropertiesType
	properties.NamePropertyType
	properties.DescriptionPropertyType
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// New - This function will create a new course of action object.
func New() CourseOfActionType {
	var obj CourseOfActionType
	obj.InitNewObject("course-of-action")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - CourseOfActionType
// ----------------------------------------------------------------------
