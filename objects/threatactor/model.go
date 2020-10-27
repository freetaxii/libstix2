// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package threatactor

import (
	"github.com/wxj95/libstix2/objects"
	"github.com/wxj95/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Type
// ----------------------------------------------------------------------

/*
ThreatActor - This type implements the STIX 2 Threat Actor SDO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties.
*/
type ThreatActor struct {
	objects.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	ThreatActorTypes []string `json:"threat_actor_types,omitempty"`
	properties.AliasesProperty
	properties.SeenProperties
	properties.RolesProperty
	properties.GoalsProperty
	Sophistication string `json:"sophistication,omitempty"`
	properties.ResourceLevelProperty
	properties.MotivationProperties
	PersonalMotivations []string `json:"personal_motivations,omitempty"`
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *ThreatActor) GetPropertyList() []string {
	return []string{"name", "description", "threat_actor_types", "aliases", "first_seen", "last_seen", "roles", "goals", "sophistication", "resource_level", "primary_motivation", "secondary_motivations", "personal_motivations"}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Threat Actor object and return it
as a pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *ThreatActor {
	var obj ThreatActor
	obj.InitSDO("threat-actor")
	return &obj
}
