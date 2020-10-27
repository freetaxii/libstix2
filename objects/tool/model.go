// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package tool

import (
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Type
// ----------------------------------------------------------------------

/*
Tool - This type implements the STIX 2 Tool SDO and defines all of the
properties and methods needed to create and work with this object. All of the
methods not defined local to this type are inherited from the individual
properties.
*/
type Tool struct {
	objects.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	ToolTypes []string `json:"tool_types,omitempty"`
	properties.AliasesProperty
	properties.KillChainPhasesProperty
	ToolVersion string `json:"tool_version,omitempty"`
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *Tool) GetPropertyList() []string {
	return []string{"name", "description", "tool_types", "aliases", "kill_chain_phases", "tool_version"}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Tool object and return
it as a pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *Tool {
	var obj Tool
	obj.InitSDO("tool")
	return &obj
}
