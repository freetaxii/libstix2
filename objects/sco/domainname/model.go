// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package domainname

import (
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
DomainName - This type implements the STIX 2 Domain Name SCO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties.
*/
type DomainName struct {
	objects.CommonObjectProperties
	properties.ValueProperty
	properties.ResolvesToRefsProperty
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *DomainName) GetPropertyList() []string {
	return []string{"value", "resolves_to_refs"}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Domain Name SCO and return it as a
pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *DomainName {
	var obj DomainName
	obj.InitSCO("domain-name")
	return &obj
}
