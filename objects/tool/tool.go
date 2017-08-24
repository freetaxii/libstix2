// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package tool

import (
	"github.com/freetaxii/libstix2/objects/common"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type ToolType struct {
	common.CommonObjectPropertiesType
	common.DescriptivePropertiesType
	common.KillChainPhasesPropertyType
	Tool_version string `json:"tool_version,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() ToolType {
	var obj ToolType
	obj.InitNewObject("tool")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - ToolType
// ----------------------------------------------------------------------

// SetToolVersion takes in one parameter
// param: s - a string value representing the version of the tool
func (this *ToolType) SetToolVersion(s string) {
	this.Tool_version = s
}
