// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package threatactor

import (
	"github.com/freetaxii/libstix2/objects"
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
	objects.NameProperty
	objects.DescriptionProperty
	ThreatActorTypes []string `json:"threat_actor_types,omitempty"`
	objects.AliasesProperty
	objects.SeenProperties
	objects.RolesProperty
	objects.GoalsProperty
	Sophistication string `json:"sophistication,omitempty"`
	objects.ResourceLevelProperty
	objects.MotivationProperties
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
