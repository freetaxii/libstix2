// Copyright 2016 Bret Jordan, All rights reserved.
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
	common.CommonPropertiesType
	common.DescriptivePropertiesType
	common.KillChainPhasesType
	Tool_version string `json:"tool_version,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() ToolType {
	var obj ToolType
	obj.MessageType = "tool"
	obj.Id = obj.NewId("tool")
	obj.Created = obj.GetCurrentTime()
	obj.Modified = obj.Created
	obj.Version = 1
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - ToolType
// ----------------------------------------------------------------------

func (this *ToolType) SetToolVersion(s string) {
	this.Tool_version = s
}
