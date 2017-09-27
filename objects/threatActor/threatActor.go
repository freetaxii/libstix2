// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package threatActor

import (
	"github.com/freetaxii/libstix2/objects/common/properties"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

// ThreatActorType -
// This type defines all of the properties associated with the STIX Threat Actor SDO.
// All of the methods not defined local to this type are inherited from the individual properties.
type ThreatActorType struct {
	properties.CommonObjectPropertiesType
	properties.NamePropertyType
	properties.DescriptionPropertyType
	properties.AliasesPropertyType
	Roles []string `json:"roles,omitempty"`
	properties.GoalsPropertyType
	Sophistication string `json:"sophistication,omitempty"`
	properties.ResourceLevelPropertyType
	properties.PrimaryMotivationPropertyType
	properties.SecondaryMotivationsPropertyType
	PersonalMotivations []string `json:"personal_motivations,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// New - This function will create a new threat actor object.
func New() ThreatActorType {
	var obj ThreatActorType
	obj.InitNewObject("threat-actor")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - ThreatActorType
// ----------------------------------------------------------------------

// AddRole - This method takes in a string value representing a threat actor
// role from the threat-actor-role-ov and adds it to the role property.
func (this *ThreatActorType) AddRole(s string) {
	this.Roles = append(this.Roles, s)
}

// SetSophistication - This method takes in a string value representing the
// sophistication level of a threat actor from the threat-actor-sophistication-ov
// and adds it to the sophistication property.
func (this *ThreatActorType) SetSophistication(s string) {
	this.Sophistication = s
}

// AddPersonalMotivation - This method takes in a string value representing the
// motivation of a threat actor from the threat-actor-motivation-ov and adds it
// to the personal motivations property.
func (this *ThreatActorType) AddPersonalMotivation(s string) {
	this.PersonalMotivations = append(this.PersonalMotivations, s)
}
