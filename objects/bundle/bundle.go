// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package bundle

import (
	"github.com/freetaxii/libstix2/objects/common/properties"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

// BundleType -
// This type defines all of the properties associated with the STIX Bundle.
// All of the methods not defined local to this type are inherited from the individual properties.
type BundleType struct {
	properties.CommonBundlePropertiesType
	SpecVersion string        `json:"spec_version,omitempty"`
	Objects     []interface{} `json:"objects,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// New - This function will create a new bundle object.
// This function can not use the InitNewObject() function as a Bundle does not
// have all of the fields that are common to a standard object.
func New() BundleType {
	var obj BundleType
	obj.SetMessageType("bundle")
	obj.CreateID("bundle")
	obj.SetSpecVersion21()
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - BundleType
// ----------------------------------------------------------------------

// SetSpecVersion21 - This method will set the specification version to 2.1.
func (this *BundleType) SetSpecVersion21() {
	this.SpecVersion = "2.1"
}

// AddObject - This method will take in an object as an interface and add it to
// the list of objects in the bundle.
func (this *BundleType) AddObject(i interface{}) {
	this.Objects = append(this.Objects, i)
}
