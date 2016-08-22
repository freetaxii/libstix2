// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package bundle

import (
	"github.com/freetaxii/libstix2/messages/indicator"
	"github.com/freetaxii/libstix2/messages/stix"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type BundleType struct {
	MessageType string                    `json:"type,omitempty"`
	Id          string                    `json:"id,omitempty"`
	Created     string                    `json:"created,omitempty"`
	Indicators  []indicator.IndicatorType `json:"indicators,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() BundleType {
	var obj BundleType
	obj.MessageType = "bundle"
	obj.Id = stix.CreateId("bundle")
	obj.Created = stix.GetCurrentTime()
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - IndicatorType
// ----------------------------------------------------------------------

func (this *BundleType) NewIndicator() *indicator.IndicatorType {
	i := indicator.New()
	slicePosition := this.addIndicator(i)
	return &this.Indicators[slicePosition]
}

// ----------------------------------------------------------------------
// Private Methods - IndicatorType
// ----------------------------------------------------------------------

func (this *BundleType) addIndicator(c indicator.IndicatorType) int {
	if this.Indicators == nil {
		a := make([]indicator.IndicatorType, 0)
		this.Indicators = a
	}
	positionThatAppendWillUse := len(this.Indicators)
	this.Indicators = append(this.Indicators, c)
	return positionThatAppendWillUse
}
