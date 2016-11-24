// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package course_of_action

import (
	"github.com/freetaxii/libstix2/objects/common"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type CourseOfActionType struct {
	common.CommonPropertiesType
	common.DescriptivePropertiesType
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() CourseOfActionType {
	var obj CourseOfActionType
	obj.MessageType = "course-of-action"
	obj.Id = obj.NewId("course-of-action")
	obj.Created = obj.GetCurrentTime()
	obj.Modified = obj.Created
	obj.Version = 1
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - CourseOfActionType
// ----------------------------------------------------------------------
