// Copyright 2015-2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package threatactor

import (
	"github.com/freetaxii/libstix2/objects/baseobject"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
//
// Define Object Type
//
// ----------------------------------------------------------------------

/*
ThreatActor - This type implements the STIX 2 Threat Actor SDO and defines
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
type ThreatActor struct {
	baseobject.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	ThreatActorType []string `json:"threat_actor_types,omitempty"`
	properties.AliasesProperty
	Roles []string `json:"roles,omitempty"`
	properties.GoalsProperty
	Sophistication string `json:"sophistication,omitempty"`
	properties.ResourceLevelProperty
	properties.MotivationProperties
	PersonalMotivations []string `json:"personal_motivations,omitempty"`
}

// ----------------------------------------------------------------------
//
// Initialization Functions
//
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Threat Actor object and return it as
a pointer.
*/
func New() *ThreatActor {
	var obj ThreatActor
	obj.InitObject("threat-actor")
	return &obj
}

// ----------------------------------------------------------------------
//
// Public Methods - ThreatActor
//
// ----------------------------------------------------------------------

/*
AddType - This method takes in a string value representing a threat actor
type from the threat-actor-type-ov and adds it to the threat actor type property.
*/
func (o *ThreatActor) AddType(s string) error {
	o.ThreatActorType = append(o.ThreatActorType, s)
	return nil
}

/*
AddRole - This method takes in a string value representing a threat actor
role from the threat-actor-role-ov and adds it to the role property.
*/
func (o *ThreatActor) AddRole(s string) error {
	o.Roles = append(o.Roles, s)
	return nil
}

/*
SetSophistication - This method takes in a string value representing the
sophistication level of a threat actor from the threat-actor-sophistication-ov
and adds it to the sophistication property.
*/
func (o *ThreatActor) SetSophistication(s string) error {
	o.Sophistication = s
	return nil
}

/*
AddPersonalMotivation - This method takes in a string value representing the
motivation of a threat actor from the threat-actor-motivation-ov and adds it
to the personal motivations property.
*/
func (o *ThreatActor) AddPersonalMotivation(s string) error {
	o.PersonalMotivations = append(o.PersonalMotivations, s)
	return nil
}
