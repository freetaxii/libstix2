// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

import (
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
//
// Define Object Type
//
// ----------------------------------------------------------------------

/*
Identity - This type implements the STIX 2 Identity SDO and defines
all of the properties methods needed to create and work with the STIX Identity
SDO. All of the methods not defined local to this type are inherited from
the individual properties.

The following information comes directly from the STIX 2 specification documents.

Identities can represent actual individuals, organizations, or groups (e.g.,
ACME, Inc.) as well as classes of individuals, organizations, or groups (e.g.,
the finance sector).

The Identity SDO can capture basic identifying information, contact information,
and the sectors that the Identity belongs to. Identity is used in STIX to
represent, among other things, targets of attacks, information sources, object
creators, and threat actor identities.
*/
type Identity struct {
	properties.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	IdentityClass      string   `json:"identity_class,omitempty"`
	Sectors            []string `json:"sectors,omitempty"`
	ContactInformation string   `json:"contact_information,omitempty"`
}

// ----------------------------------------------------------------------
//
// Initialization Functions
//
// ----------------------------------------------------------------------

/*
NewIdentity - This function will create a new STIX Identity object and return
it as a pointer.
*/
func NewIdentity() *Identity {
	var obj Identity
	obj.InitObject("identity")
	return &obj
}

// ----------------------------------------------------------------------
//
// Public Methods - Identity
//
// ----------------------------------------------------------------------

/*
SetIdentityClass - This method takes in a string value representing a STIX
identity class from the vocab identity-class-ov and updates the identity class
property.
*/
func (o *Identity) SetIdentityClass(s string) error {
	o.IdentityClass = s
	return nil
}

/*
AddSector - This method takes in a string value that represents a STIX sector
from the vocab industry-sector-ov and adds it to the identity object.
*/
func (o *Identity) AddSector(s string) error {
	o.Sectors = append(o.Sectors, s)
	return nil
}

/*
SetContactInformation - This method takes in a string value representing
contact information as a text string and updates the contact information
property.
*/
func (o *Identity) SetContactInformation(s string) error {
	o.ContactInformation = s
	return nil
}
