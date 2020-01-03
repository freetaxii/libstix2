// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package identity

import (
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Type
// ----------------------------------------------------------------------

/*
Identity - This type implements the STIX 2 Identity SDO and
defines all of the properties and methods needed to create and work with this
object. All of the methods not defined local to this type are inherited from the
individual properties. */
type Identity struct {
	objects.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	properties.RolesProperty
	IdentityClass      string   `json:"identity_class,omitempty"`
	Sectors            []string `json:"sectors,omitempty"`
	ContactInformation string   `json:"contact_information,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/* New - This function will create a new STIX Identity object and return it as a
pointer. */
func New() *Identity {
	var obj Identity
	obj.InitSDO("identity")
	return &obj
}
