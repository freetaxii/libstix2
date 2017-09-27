// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package indicator

import (
	"github.com/freetaxii/libstix2/objects/common/properties"
)

// ----------------------------------------------------------------------
// Define Indicator Type
// ----------------------------------------------------------------------

// IndicatorType -
// This type defines all of the properties associated with the STIX Indicator SDO.
// All of the methods not defined local to this type are inherited from the individual properties.
type IndicatorType struct {
	properties.CommonObjectPropertiesType
	properties.NamePropertyType
	properties.DescriptionPropertyType
	Pattern    string `json:"pattern,omitempty"`
	ValidFrom  string `json:"valid_from,omitempty"`
	ValidUntil string `json:"valid_until,omitempty"`
	properties.KillChainPhasesPropertyType
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// New - This function will create a new indicator object.
func New() IndicatorType {
	var obj IndicatorType
	obj.InitNewObject("indicator")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - IndicatorType
// ----------------------------------------------------------------------

// SetPattern - This method will take in a string value representing a complete
// and valid STIX pattern and set the pattern property to that value.
func (this *IndicatorType) SetPattern(s string) {
	this.Pattern = s
}

// SetValidFrom - This method will take in a timestamp in either time.Time or
// string format and will set the valid from property to that value.
func (this *IndicatorType) SetValidFrom(t interface{}) {
	ts := this.VerifyTimestamp(t)
	this.ValidFrom = ts
}

// SetValidUntil - This method will take in a timestamp in either time.Time or
// string format and will set the valid until property to that value.
func (this *IndicatorType) SetValidUntil(t interface{}) {
	ts := this.VerifyTimestamp(t)

	// TODO check to make sure this is later than the vaild_from
	this.ValidUntil = ts
}
