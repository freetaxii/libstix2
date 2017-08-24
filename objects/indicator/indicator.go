// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package indicator

import (
	"github.com/freetaxii/libstix2/objects/common"
)

// ----------------------------------------------------------------------
// Define Indicator Type
// ----------------------------------------------------------------------

type IndicatorType struct {
	common.CommonObjectPropertiesType
	common.DescriptivePropertiesType
	Pattern     string `json:"pattern,omitempty"`
	Valid_from  string `json:"valid_from,omitempty"`
	Valid_until string `json:"valid_until,omitempty"`
	common.KillChainPhasesPropertyType
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() IndicatorType {
	var obj IndicatorType
	obj.InitNewObject("indicator")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - IndicatorType
// ----------------------------------------------------------------------

func (this *IndicatorType) SetPattern(s string) {
	this.Pattern = s
}

// SetValidForm takes in one parameter
// param: t a timestamp in either time.Time or string format
func (this *IndicatorType) SetValidFrom(t interface{}) {
	ts := this.VerifyTimestamp(t)
	this.Valid_from = ts
}

// SetValidUntil takes in one parameter
// param: t a timestamp in either time.Time or string format
func (this *IndicatorType) SetValidUntil(t interface{}) {
	ts := this.VerifyTimestamp(t)

	// TODO check to make sure this is later than the vaild_from
	this.Valid_until = ts
}
