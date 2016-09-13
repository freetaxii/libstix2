// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package bundle

import (
	"github.com/freetaxii/libstix2/messages/indicator"
	"github.com/freetaxii/libstix2/messages/relationship"
	"github.com/freetaxii/libstix2/messages/stix"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type BundleType struct {
	MessageType   string                          `json:"type,omitempty"`
	Id            string                          `json:"id,omitempty"`
	Spec_version  string                          `json:"spec_version,omitempty"`
	Indicators    []indicator.IndicatorType       `json:"indicators,omitempty"`
	Relationships []relationship.RelationshipType `json:"relationships,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() BundleType {
	var obj BundleType
	obj.MessageType = "bundle"
	obj.Id = stix.NewId("bundle")
	obj.SetSpecVersion20()
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - Common Properties
// ----------------------------------------------------------------------
func (this *BundleType) SetSpecVersion20() {
	this.Spec_version = "stix-2.0"
}

// ----------------------------------------------------------------------
// Public Methods - BundleType
// ----------------------------------------------------------------------

func (this *BundleType) NewIndicator() *indicator.IndicatorType {
	i := indicator.New()
	slicePosition := this.addIndicator(i)
	return &this.Indicators[slicePosition]
}

func (this *BundleType) NewRelationship() *relationship.RelationshipType {
	i := relationship.New()
	slicePosition := this.addRelationship(i)
	return &this.Relationships[slicePosition]
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

func (this *BundleType) addRelationship(c relationship.RelationshipType) int {
	if this.Relationships == nil {
		a := make([]relationship.RelationshipType, 0)
		this.Relationships = a
	}
	positionThatAppendWillUse := len(this.Relationships)
	this.Relationships = append(this.Relationships, c)
	return positionThatAppendWillUse
}
