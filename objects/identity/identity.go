// Copyright 2016 Bret Jordan, All rights reserved.
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
	common.CommonPropertiesType
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
	obj.MessageType = "identity"
	obj.Id = obj.NewId("identity")
	obj.Created = obj.GetCurrentTime()
	obj.Modified = obj.Created
	obj.Version = 1
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - IdentityType
// ----------------------------------------------------------------------

func (this *IdentityType) SetIdentityClass(s string) {
	this.Identity_class = s
}

func (this *IdentityType) AddSector(value string) {
	if this.Sectors == nil {
		a := make([]string, 0)
		this.Sectors = a
	}
	this.Sectors = append(this.Sectors, value)
}

func (this *IdentityType) SetContactInformation(s string) {
	this.Contact_information = s
}
