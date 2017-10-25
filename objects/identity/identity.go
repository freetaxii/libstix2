// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package identity

import (
	"github.com/freetaxii/libstix2/objects/common/properties"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
IdentityType defines all of the properties associated with the STIX Identity
SDO. All of the methods not defined local to this type are inherited from the
individual properties.
*/
type IdentityType struct {
	properties.CommonObjectPropertiesType
	properties.NamePropertyType
	properties.DescriptionPropertyType
	IdentityClass      string   `json:"identity_class,omitempty"`
	Sectors            []string `json:"sectors,omitempty"`
	ContactInformation string   `json:"contact_information,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// New - This function will create a new identity object.
func New(ver string) IdentityType {
	var obj IdentityType
	obj.InitNewObject("identity", ver)
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - IdentityType
// ----------------------------------------------------------------------

// SetIdentityClass - This method takes in a string value representing a STIX
// identity class from the vocab identity-class-ov and updates the identity class
// property.
func (p *IdentityType) SetIdentityClass(s string) {
	p.IdentityClass = s
}

// AddSector - This method takes in a string value that represents a STIX sector
// from the vocab industry-sector-ov and adds it to the identity object.
func (p *IdentityType) AddSector(s string) {
	p.Sectors = append(p.Sectors, s)
}

// SetContactInformation - This method takes in a string value representing
// contact information as a text string and updates the contact information
// property.
func (p *IdentityType) SetContactInformation(s string) {
	p.ContactInformation = s
}
