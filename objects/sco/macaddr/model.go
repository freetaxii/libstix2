// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package macaddr

import (
	"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
MACAddr - This type implements the STIX 2.1 MAC Address SCO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties.

Reference: STIX 2.1 specification section 6.8
*/
type MACAddr struct {
	objects.CommonObjectProperties
	objects.ValueProperty
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *MACAddr) GetPropertyList() []string {
	return []string{"value"}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX MAC Address SCO and return it as a
pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *MACAddr {
	var obj MACAddr
	obj.InitSCO("mac-addr")
	return &obj
}
