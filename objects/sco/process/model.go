// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package process

import (
	"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
Process - This type implements the STIX 2.1 Process SCO and defines all of the
properties and methods needed to create and work with this object. All of the
methods not defined local to this type are inherited from the individual
properties.

Reference: STIX 2.1 specification section 6.11

TODO: Complete implementation of all properties per specification
*/
type Process struct {
	objects.CommonObjectProperties
	// TODO: Add specific properties for Process based on STIX 2.1 spec section 6.11
	// TODO: Add pid, created_time, cwd, command_line, environment_variables, opened_connection_refs, creator_user_ref, image_ref, parent_ref, child_refs
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *Process) GetPropertyList() []string {
	// TODO: Update with actual property names
	return []string{}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Process SCO and return it as a
pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *Process {
	var obj Process
	obj.InitSCO("process")
	return &obj
}
