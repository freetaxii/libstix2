// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package indicator

import (
	"github.com/freetaxii/libstix2/messages/common"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type IndicatorType struct {
	common.CommonPropertiesType
	common.DescriptivePropertiesType
	Pattern               string `json:"pattern,omitempty"`
	Valid_from            string `json:"valid_from,omitempty"`
	Valid_from_precision  string `json:"valid_from_precision,omitempty"`
	Valid_until           string `json:"valid_until,omitempty"`
	Valid_until_precision string `json:"valid_until_precision,omitempty"`
	common.KillChainPhasesType
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() IndicatorType {
	var obj IndicatorType
	obj.MessageType = "indicator"
	obj.Id = obj.NewId("indicator")
	obj.Created = obj.GetCurrentTime()
	obj.Modified = obj.Created
	obj.Version = 1
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - IndicatorType
// ----------------------------------------------------------------------

func (this *IndicatorType) SetPattern(s string) {
	this.Pattern = s
}

// SetValidForm takes in two parameters and returns an error if there is one
// param: t a timestamp in either time.Time or string format
// param: s a timestamp precision in string format
func (this *IndicatorType) SetValidFrom(t interface{}, s string) error {

	ts, err := this.VerifyTimestamp(t)
	if err != nil {
		return err
	}
	this.Valid_from = ts

	p, err := this.VerifyPrecision(s)
	if err != nil {
		return err
	}
	this.Valid_from_precision = p

	return nil
}

// SetValidUntil takes in two parameters and returns and error if there is one
// param: t a timestamp in either time.Time or string format
// param: s a timestamp precision in string format
func (this *IndicatorType) SetValidUntil(t interface{}, s string) error {

	ts, err := this.VerifyTimestamp(t)
	if err != nil {
		return err
	}
	this.Valid_until = ts

	p, err := this.VerifyPrecision(s)
	if err != nil {
		return err
	}
	this.Valid_until_precision = p

	return nil
}
