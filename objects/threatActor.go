// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

import (
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
//
// Define Message Type
//
// ----------------------------------------------------------------------

/*
ThreatActorType - This type implements the STIX 2 Threat Actor SDO and defines
all of the properties methods needed to create and work with the STIX Threat Actor
SDO. All of the methods not defined local to this type are inherited from
the individual properties.

The following information comes directly from the STIX 2 specification documents.

Threat Actors are actual individuals, groups, or organizations believed to be
operating with malicious intent. A Threat Actor is not an Intrusion Set but may
support or be affiliated with various Intrusion Sets, groups, or organizations
over time.

Threat Actors leverage their resources, and possibly the resources of an
Intrusion Set, to conduct attacks and run Campaigns against targets.

Threat Actors can be characterized by their motives, capabilities, goals,
sophistication level, past activities, resources they have access to, and their
role in the organization.
*/
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
//
// Initialization Functions
//
// ----------------------------------------------------------------------

/*
NewThreatActor - This function will create a new STIX Threat Actor object
and return it as a pointer.
*/
func NewThreatActor(ver string) *ThreatActorType {
	var obj ThreatActorType
	obj.InitObjectProperties("threat-actor", ver)
	return &obj
}

// ----------------------------------------------------------------------
//
// Public Methods - ThreatActorType
//
// ----------------------------------------------------------------------

/*
AddRole - This method takes in a string value representing a threat actor
role from the threat-actor-role-ov and adds it to the role property.
*/
func (o *ThreatActorType) AddRole(s string) error {
	o.Roles = append(o.Roles, s)
	return nil
}

/*
SetSophistication - This method takes in a string value representing the
sophistication level of a threat actor from the threat-actor-sophistication-ov
and adds it to the sophistication property.
*/
func (o *ThreatActorType) SetSophistication(s string) error {
	o.Sophistication = s
	return nil
}

/*
AddPersonalMotivation - This method takes in a string value representing the
motivation of a threat actor from the threat-actor-motivation-ov and adds it
to the personal motivations property.
*/
func (o *ThreatActorType) AddPersonalMotivation(s string) error {
	o.PersonalMotivations = append(o.PersonalMotivations, s)
	return nil
}
