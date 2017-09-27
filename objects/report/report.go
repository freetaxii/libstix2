// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package report

import (
	"github.com/freetaxii/libstix2/objects/common/properties"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

// ReportType -
// This type defines all of the properties associated with the STIX Report SDO.
// All of the methods not defined local to this type are inherited from the individual properties.
type ReportType struct {
	properties.CommonObjectPropertiesType
	properties.NamePropertyType
	properties.DescriptionPropertyType
	Published  string   `json:"published,omitempty"`
	ObjectRefs []string `json:"object_refs,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// New - This function will create a new report object.
func New() ReportType {
	var obj ReportType
	obj.InitNewObject("report")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - ReportType
// ----------------------------------------------------------------------

// SetPublished - This method takes in a timestamp in either time.Time or string
// format and updates the published timestamp property.
func (this *ReportType) SetPublished(t interface{}) {
	ts := this.VerifyTimestamp(t)
	this.Published = ts
}

// AddObject - This methods takes in a string value that represents a STIX
// identifier and adds it to the objects ref property.
func (this *ReportType) AddObject(s string) {
	this.ObjectRefs = append(this.ObjectRefs, s)
}
