// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package tool

import (
	"github.com/freetaxii/libstix2/objects/common/properties"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
ToolType defines all of the properties associated with the STIX Tool
SDO. All of the methods not defined local to this type are inherited from the
individual properties.
*/
type ToolType struct {
	properties.CommonObjectPropertiesType
	properties.NamePropertyType
	properties.DescriptionPropertyType
	properties.KillChainPhasesPropertyType
	ToolVersion string `json:"tool_version,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// New - This function will create a new tool object.
func New() ToolType {
	var obj ToolType
	obj.InitNewObject("tool")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - ToolType
// ----------------------------------------------------------------------

// SetToolVersion - This method takes in a string value representing the version
// of the tool and updates the tool version property.
func (ezt *ToolType) SetToolVersion(s string) {
	ezt.ToolVersion = s
}
