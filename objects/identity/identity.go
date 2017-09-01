// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package identity

import (
	"github.com/freetaxii/libstix2/objects/common"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type IdentityType struct {
	common.CommonObjectPropertiesType
	common.DescriptivePropertiesType
	Identity_class      string   `json:"identity_class,omitempty"`
	Sectors             []string `json:"sectors,omitempty"`
	Contact_information string   `json:"contact_information,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() IdentityType {
	var obj IdentityType
	obj.InitNewObject("identity")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - IdentityType
// ----------------------------------------------------------------------

// SetIdentityClass takes in one parameter
// param: s - a string value representing a STIX identity class from the vocab identity-class-ov
func (this *IdentityType) SetIdentityClass(s string) {
	this.Identity_class = s
}

// AddSector takes in one parameter
// param: s - a string value that represents a STIX sector from the vocab industry-sector-ov
func (this *IdentityType) AddSector(s string) {
	this.Sectors = append(this.Sectors, s)
}

// SetContactInformation takes in one parameter
// param: s - a string value representing contact information as a text string
func (this *IdentityType) SetContactInformation(s string) {
	this.Contact_information = s
}
