// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package report

import (
	"github.com/freetaxii/libstix2/objects/common"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type ReportType struct {
	common.CommonObjectPropertiesType
	common.DescriptivePropertiesType
	Published   string   `json:"published,omitempty"`
	Object_refs []string `json:"object_refs,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() ReportType {
	var obj ReportType
	obj.InitNewObject("report")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - ReportType
// ----------------------------------------------------------------------

// SetPublished takes in one parameter
// param: t - a timestamp in either time.Time or string format
func (this *ReportType) SetPublished(t interface{}) {
	ts := this.VerifyTimestamp(t)
	this.Published = ts
}

// AddObject takes in one parameter
// param: s - a string value that represents a STIX identifier
func (this *ReportType) AddObject(s string) {
	this.Object_refs = append(this.Object_refs, s)
}
