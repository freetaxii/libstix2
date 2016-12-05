// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package intrusion_set

import (
	"github.com/freetaxii/libstix2/objects/common"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type IntrusionSetType struct {
	common.CommonPropertiesType
	common.DescriptivePropertiesType
	common.AliasesType
	common.FirstLastSeenType
	Goals                 []string `json:"goals,omitempty"`
	Resource_level        string   `json:"resource_level,omitempty"`
	Primary_motivation    string   `json:"primary_motivation,omitempty"`
	Secondary_motivations []string `json:"secondary_motivations,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() IntrusionSetType {
	var obj IntrusionSetType
	obj.MessageType = "intrusion-set"
	obj.Id = obj.NewId("intrusion-set")
	obj.Created = obj.GetCurrentTime()
	obj.Modified = obj.Created
	obj.Version = 1
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - IntrusionSetType
// ----------------------------------------------------------------------

func (this *IntrusionSetType) AddGoal(value string) {
	if this.Goals == nil {
		a := make([]string, 0)
		this.Goals = a
	}
	this.Goals = append(this.Goals, value)
}

func (this *IntrusionSetType) SetResourceLevel(s string) {
	this.Resource_level = s
}

func (this *IntrusionSetType) SetPrimaryMotivation(s string) {
	this.Primary_motivation = s
}

func (this *IntrusionSetType) AddSecondaryMotivation(value string) {
	if this.Secondary_motivations == nil {
		a := make([]string, 0)
		this.Secondary_motivations = a
	}
	this.Secondary_motivations = append(this.Secondary_motivations, value)
}
