// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package bundle

import (
	"github.com/freetaxii/libstix2/objects/common"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type BundleType struct {
	common.CommonBasePropertiesType
	Spec_version string   `json:"spec_version,omitempty"`
	Objects      []string `json:"objects,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// This function can not use the InitNewObject() function as a Bundle does not
// have all of the fields that are common to a standard object
func New() BundleType {
	var obj BundleType
	obj.SetMessageType("bundle")
	obj.CreateId("bundle")
	obj.SetSpecVersion21()
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - BundleType
// ----------------------------------------------------------------------

// SetSepcVersion21 set the specification version to 2.1
func (this *BundleType) SetSpecVersion21() {
	this.Spec_version = "2.1"
}

// AddObject takes in one parameter
// param: s - a string value representing a STIX Identifier
func (this *BundleType) AddObject(s string) {
	if this.Objects == nil {
		a := make([]string, 0)
		this.Objects = a
	}
	this.Objects = append(this.Objects, s)
}
