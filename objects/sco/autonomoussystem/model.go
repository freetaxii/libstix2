// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package autonomoussystem

import (
	"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
AutonomousSystem - This type implements the STIX 2.1 AutonomousSystem SCO and defines all of the
properties and methods needed to create and work with this object. All of the
methods not defined local to this type are inherited from the individual
properties.

Reference: STIX 2.1 specification section 6.2

TODO: Complete implementation of all properties per specification
*/
type AutonomousSystem struct {
	objects.CommonObjectProperties
	// TODO: Add specific properties for AutonomousSystem based on STIX 2.1 spec section 6.2
	Number int `json:"number" bson:"number"`
	objects.NameProperty
	RIR string `json:"rir,omitempty" bson:"rir,omitempty"`
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *AutonomousSystem) GetPropertyList() []string {
	// TODO: Update with actual property names
	return []string{}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX AutonomousSystem SCO and return it as a
pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *AutonomousSystem {
	var obj AutonomousSystem
	obj.InitSCO("autonomoussystem")
	return &obj
}
