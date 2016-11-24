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
	First_seen            string   `json:"first_seen,omitempty"`
	First_seen_precision  string   `json:"first_seen_precision,omitempty"`
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

// SetFirstSeen takes in two parameters and returns and error if there is one
// param: t a timestamp in either time.Time or string format
// param: s a timestamp precision in string format
func (this *IntrusionSetType) SetFirstSeen(t interface{}, s string) error {

	ts, err := this.VerifyTimestamp(t)
	if err != nil {
		return err
	}
	this.First_seen = ts

	p, err := this.VerifyPrecision(s)
	if err != nil {
		return err
	}
	this.First_seen_precision = p

	return nil
}

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
