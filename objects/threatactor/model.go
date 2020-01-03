// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package threatactor

import (
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Type
// ----------------------------------------------------------------------

/* ThreatActor - This type implements the STIX 2 Threat Actor SDO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties. */
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

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/* New - This function will create a new STIX Threat Actor object and return it
as a pointer. It will also initialize the object by setting all of the basic
properties. */
func New() *ThreatActor {
	var obj ThreatActor
	obj.InitSDO("threat-actor")
	return &obj
}
