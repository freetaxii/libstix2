// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package courseofaction

import (
	"github.com/freetaxii/libstix2/objects/baseobject"
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

The following information comes directly from the STIX 2 specification documents.

A Course of Action is an action taken either to prevent an attack or to respond
to an attack that is in progress. It may describe technical, automatable
responses (applying patches, reconfiguring firewalls) but can also describe
higher level actions like employee training or policy changes. For example, a
course of action to mitigate a vulnerability could describe applying the patch
that fixes it.

The Course of Action SDO contains a textual description of the action; a
reserved action property also serves as placeholder for future inclusion of
machine automatable courses of action. Relationships from the Course of Action
can be used to link it to the Vulnerabilities or behaviors (Tool, Malware,
Attack Pattern) that it mitigates.
*/
type CourseOfAction struct {
	baseobject.CommonObjectProperties
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
//
// Public Methods - CourseOfAction
//
// ----------------------------------------------------------------------
