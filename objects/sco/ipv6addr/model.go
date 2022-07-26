// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package ipv6addr

import (
	"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
IPv6Addr - This type implements the STIX 2 IPv6 Address SCO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties.
*/
type IPv6Addr struct {
	objects.CommonObjectProperties
	objects.ValueProperty
	objects.ResolvesToRefsProperty
	objects.BelongsToRefsProperty
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *IPv6Addr) GetPropertyList() []string {
	return []string{"value", "resolves_to_refs", "belongs_to_refs"}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX IPv4 Address SCO and return it as
a pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *IPv6Addr {
	var obj IPv6Addr
	obj.InitSCO("ipv6-addr")
	return &obj
}
