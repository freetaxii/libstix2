// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package software

import (
	"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
Software - This type implements the STIX 2.1 Software SCO and defines all of the
properties and methods needed to create and work with this object. All of the
methods not defined local to this type are inherited from the individual
properties.

Reference: STIX 2.1 specification section 6.13

TODO: Complete implementation of all properties per specification
*/
type Software struct {
	objects.CommonObjectProperties
	// TODO: Add specific properties for Software based on STIX 2.1 spec section 6.13
	objects.NameProperty
	CPE       string   `json:"cpe,omitempty" bson:"cpe,omitempty"`
	SWID      string   `json:"swid,omitempty" bson:"swid,omitempty"`
	Languages []string `json:"languages,omitempty" bson:"languages,omitempty"`
	Vendor    string   `json:"vendor,omitempty" bson:"vendor,omitempty"`
	Version   string   `json:"version,omitempty" bson:"version,omitempty"`
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *Software) GetPropertyList() []string {
	// TODO: Update with actual property names
	return []string{}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Software SCO and return it as a
pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *Software {
	var obj Software
	obj.InitSCO("software")
	return &obj
}
