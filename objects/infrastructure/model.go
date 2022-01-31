// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package infrastructure

import (
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Type
// ----------------------------------------------------------------------

/*
Infrastructure - This type implements the STIX 2 Infrastructure SDO and
defines all of the properties and methods needed to create and work with this
object. All of the methods not defined local to this type are inherited from the
individual properties.
*/
type Infrastructure struct {
	objects.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	InfrastructureTypes []string `json:"infrastructure_types,omitempty" bson:"infrastructure_types,omitempty"`
	properties.AliasesProperty
	properties.KillChainPhasesProperty
	properties.SeenProperties
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *Infrastructure) GetPropertyList() []string {
	return []string{"name", "description", "infrastructure_types", "aliases", "kill_chain_phases", "first_seen", "last_seen"}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Infrastructure object and return
it as a pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *Infrastructure {
	var obj Infrastructure
	obj.InitSDO("infrastructure")
	return &obj
}
