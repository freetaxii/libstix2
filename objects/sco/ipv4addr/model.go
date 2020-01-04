// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package ipv4addr

import (
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/* IPv4Addr - This type implements the STIX 2 IPv4 Address SCO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties. */
type IPv4Addr struct {
	objects.CommonObjectProperties
	properties.ValueProperty
	properties.ResolvesToRefsProperty
	properties.BelongsToRefsProperty
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/* New - This function will create a new STIX IPv4 Address SCO and return it as
a pointer. It will also initialize the object by setting all of the basic
properties. */
func New() *IPv4Addr {
	var obj IPv4Addr
	obj.InitSCO("ipv4-addr")
	return &obj
}
