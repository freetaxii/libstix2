// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package threat_actor

import (
	"github.com/freetaxii/libstix2/messages/stix"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type ThreatActorType struct {
	stix.CommonPropertiesType
	stix.DescriptivePropertiesType
	stix.AliasesType
	Roles                 []string `json:"roles,omitempty"`
	Goals                 []string `json:"goals,omitempty"`
	Sophistication        string   `json:"sophistication,omitempty"`
	Resource_level        string   `json:"resource_level,omitempty"`
	Primary_motivation    string   `json:"primary_motivation,omitempty"`
	Secondary_motivations []string `json:"secondary_motivations,omitempty"`
	Personal_motivations  []string `json:"personal_motivations,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() ThreatActorType {
	var obj ThreatActorType
	obj.MessageType = "threat-actor"
	obj.Id = obj.NewId("threat-actor")
	obj.Created = obj.GetCurrentTime()
	obj.Modified = obj.Created
	obj.Version = 1
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - ThreatActorType
// ----------------------------------------------------------------------

func (this *ThreatActorType) AddRole(value string) {
	if this.Roles == nil {
		a := make([]string, 0)
		this.Roles = a
	}
	this.Roles = append(this.Roles, value)
}

func (this *ThreatActorType) AddGoal(value string) {
	if this.Goals == nil {
		a := make([]string, 0)
		this.Goals = a
	}
	this.Goals = append(this.Goals, value)
}

func (this *ThreatActorType) SetSophistication(s string) {
	this.Sophistication = s
}

func (this *ThreatActorType) SetResourceLevel(s string) {
	this.Resource_level = s
}

func (this *ThreatActorType) SetPrimaryMotivation(s string) {
	this.Primary_motivation = s
}

func (this *ThreatActorType) AddSecondaryMotivation(value string) {
	if this.Secondary_motivations == nil {
		a := make([]string, 0)
		this.Secondary_motivations = a
	}
	this.Secondary_motivations = append(this.Secondary_motivations, value)
}

func (this *ThreatActorType) AddPersonalMotivation(value string) {
	if this.Personal_motivations == nil {
		a := make([]string, 0)
		this.Personal_motivations = a
	}
	this.Personal_motivations = append(this.Personal_motivations, value)
}
