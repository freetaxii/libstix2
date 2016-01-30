// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package stix

import (
	"github.com/freetaxii/libstix2/messages/common"
	"github.com/freetaxii/libstix2/messages/indicator"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type PackageType struct {
	MessageType string                    `json:"type,omitempty"`
	Id          string                    `json:"id,omitempty"`
	CreatedAt   string                    `json:"created_at,omitempty"`
	Indicators  []indicator.IndicatorType `json:"indicators,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() PackageType {
	var obj PackageType
	obj.MessageType = "package"
	obj.Id = common.CreateId("package")
	obj.CreatedAt = common.GetCurrentTime()
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - IndicatorType
// ----------------------------------------------------------------------

func (this *PackageType) NewIndicator() *indicator.IndicatorType {
	i := indicator.New()
	slicePosition := this.addIndicator(i)
	return &this.Indicators[slicePosition]
}

// ----------------------------------------------------------------------
// Private Methods - IndicatorType
// ----------------------------------------------------------------------

func (this *PackageType) addIndicator(c indicator.IndicatorType) int {
	if this.Indicators == nil {
		a := make([]indicator.IndicatorType, 0)
		this.Indicators = a
	}
	positionThatAppendWillUse := len(this.Indicators)
	this.Indicators = append(this.Indicators, c)
	return positionThatAppendWillUse
}
