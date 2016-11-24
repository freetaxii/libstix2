// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package bundle

import (
	"github.com/freetaxii/libstix2/objects/campaign"
	"github.com/freetaxii/libstix2/objects/common"
	"github.com/freetaxii/libstix2/objects/indicator"
	"github.com/freetaxii/libstix2/objects/infrastructure"
	"github.com/freetaxii/libstix2/objects/malware"
	"github.com/freetaxii/libstix2/objects/observed_data"
	"github.com/freetaxii/libstix2/objects/relationship"
	"github.com/freetaxii/libstix2/objects/sighting"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type BundleType struct {
	MessageType     string                              `json:"type,omitempty"`
	Id              string                              `json:"id,omitempty"`
	Spec_version    string                              `json:"spec_version,omitempty"`
	Campaigns       []campaign.CampaignType             `json:"campaigns,omitempty"`
	Indicators      []indicator.IndicatorType           `json:"indicators,omitempty"`
	Infrastructures []infrastructure.InfrastructureType `json:"infrastructures,omitempty"`
	Malware         []malware.MalwareType               `json:"malware,omitempty"`
	ObservedData    []observed_data.ObservedDataType    `json:"observed-data,omitempty"`
	Relationships   []relationship.RelationshipType     `json:"relationships,omitempty"`
	Sightings       []sighting.SightingType             `json:"sightings,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() BundleType {
	var obj BundleType
	obj.MessageType = "bundle"
	obj.Id = common.NewId("bundle")
	obj.SetSpecVersion20()
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - Common Properties
// ----------------------------------------------------------------------
func (this *BundleType) SetSpecVersion20() {
	this.Spec_version = "2.0"
}

// ----------------------------------------------------------------------
// Public Methods - BundleType
// ----------------------------------------------------------------------

func (this *BundleType) NewCampaign() *campaign.CampaignType {
	i := campaign.New()
	slicePosition := this.addCampaign(i)
	return &this.Campaigns[slicePosition]
}

func (this *BundleType) NewIndicator() *indicator.IndicatorType {
	i := indicator.New()
	slicePosition := this.addIndicator(i)
	return &this.Indicators[slicePosition]
}

func (this *BundleType) NewInfrastructure() *infrastructure.InfrastructureType {
	i := infrastructure.New()
	slicePosition := this.addInfrastructure(i)
	return &this.Infrastructures[slicePosition]
}

func (this *BundleType) NewMalware() *malware.MalwareType {
	i := malware.New()
	slicePosition := this.addMalware(i)
	return &this.Malware[slicePosition]
}

func (this *BundleType) NewObservedData() *observed_data.ObservedDataType {
	i := observed_data.New()
	slicePosition := this.addObservedData(i)
	return &this.ObservedData[slicePosition]
}

func (this *BundleType) NewRelationship() *relationship.RelationshipType {
	i := relationship.New()
	slicePosition := this.addRelationship(i)
	return &this.Relationships[slicePosition]
}

func (this *BundleType) NewSighting() *sighting.SightingType {
	i := sighting.New()
	slicePosition := this.addSighting(i)
	return &this.Sightings[slicePosition]
}

// ----------------------------------------------------------------------
// Private Methods - IndicatorType
// ----------------------------------------------------------------------

func (this *BundleType) addCampaign(o campaign.CampaignType) int {
	if this.Campaigns == nil {
		a := make([]campaign.CampaignType, 0)
		this.Campaigns = a
	}
	positionThatAppendWillUse := len(this.Campaigns)
	this.Campaigns = append(this.Campaigns, o)
	return positionThatAppendWillUse
}

func (this *BundleType) addIndicator(o indicator.IndicatorType) int {
	if this.Indicators == nil {
		a := make([]indicator.IndicatorType, 0)
		this.Indicators = a
	}
	positionThatAppendWillUse := len(this.Indicators)
	this.Indicators = append(this.Indicators, o)
	return positionThatAppendWillUse
}

func (this *BundleType) addInfrastructure(o infrastructure.InfrastructureType) int {
	if this.Infrastructures == nil {
		a := make([]infrastructure.InfrastructureType, 0)
		this.Infrastructures = a
	}
	positionThatAppendWillUse := len(this.Infrastructures)
	this.Infrastructures = append(this.Infrastructures, o)
	return positionThatAppendWillUse
}

func (this *BundleType) addMalware(o malware.MalwareType) int {
	if this.Malware == nil {
		a := make([]malware.MalwareType, 0)
		this.Malware = a
	}
	positionThatAppendWillUse := len(this.Malware)
	this.Malware = append(this.Malware, o)
	return positionThatAppendWillUse
}

func (this *BundleType) addObservedData(o observed_data.ObservedDataType) int {
	if this.ObservedData == nil {
		a := make([]observed_data.ObservedDataType, 0)
		this.ObservedData = a
	}
	positionThatAppendWillUse := len(this.ObservedData)
	this.ObservedData = append(this.ObservedData, o)
	return positionThatAppendWillUse
}

func (this *BundleType) addRelationship(o relationship.RelationshipType) int {
	if this.Relationships == nil {
		a := make([]relationship.RelationshipType, 0)
		this.Relationships = a
	}
	positionThatAppendWillUse := len(this.Relationships)
	this.Relationships = append(this.Relationships, o)
	return positionThatAppendWillUse
}

func (this *BundleType) addSighting(o sighting.SightingType) int {
	if this.Sightings == nil {
		a := make([]sighting.SightingType, 0)
		this.Sightings = a
	}
	positionThatAppendWillUse := len(this.Sightings)
	this.Sightings = append(this.Sightings, o)
	return positionThatAppendWillUse
}
