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

/*
ReportType defines all of the properties associated with the STIX Report
SDO. All of the methods not defined local to this type are inherited from the
individual properties.
*/
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
func New(ver string) ReportType {
	var obj ReportType
	obj.InitNewObject("report", ver)
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - ReportType
// ----------------------------------------------------------------------

// SetPublished - This method takes in a timestamp in either time.Time or string
// format and updates the published timestamp property.
func (p *ReportType) SetPublished(t interface{}) {
	ts := p.VerifyTimestamp(t)
	p.Published = ts
}

// AddObject - This methods takes in a string value that represents a STIX
// identifier and adds it to the objects ref property.
func (p *ReportType) AddObject(s string) {
	p.ObjectRefs = append(p.ObjectRefs, s)
}
