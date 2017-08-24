// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package threat_actor

import (
	"github.com/freetaxii/libstix2/objects/common"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type ThreatActorType struct {
	common.CommonObjectPropertiesType
	common.DescriptivePropertiesType
	common.AliasesPropertiesType
	Roles []string `json:"roles,omitempty"`
	common.GoalsPropertiesType
	Sophistication string `json:"sophistication,omitempty"`
	common.ResourceLevelPropertyType
	common.PrimaryMotivationPropertyType
	common.SecondaryMotivationsPropertyType
	Personal_motivations []string `json:"personal_motivations,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() ThreatActorType {
	var obj ThreatActorType
	obj.InitNewObject("threat-actor")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - ThreatActorType
// ----------------------------------------------------------------------

// AddRole takes in one parameter
// param: s - a string value representing a threat actor role from the threat-actor-role-ov
func (this *ThreatActorType) AddRole(s string) {
	if this.Roles == nil {
		a := make([]string, 0)
		this.Roles = a
	}
	this.Roles = append(this.Roles, s)
}

// SetSophistication takes in one parameter
// param: s - a string value representing the sophistication level of a threat actor from the threat-actor-sophistication-ov
func (this *ThreatActorType) SetSophistication(s string) {
	this.Sophistication = s
}

// AddPersonalMotivation takes in one parameter
// param: s - a string value representing the motivation of a threat actor from the threat-actor-motivation-ov
func (this *ThreatActorType) AddPersonalMotivation(s string) {
	if this.Personal_motivations == nil {
		a := make([]string, 0)
		this.Personal_motivations = a
	}
	this.Personal_motivations = append(this.Personal_motivations, s)
}
