// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package attack_pattern

import (
	"github.com/freetaxii/libstix2/objects/common"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type AttackPatternType struct {
	common.CommonPropertiesType
	common.DescriptivePropertiesType
	common.KillChainPhasesType
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() AttackPatternType {
	var obj AttackPatternType
	obj.MessageType = "attack-pattern"
	obj.Id = obj.NewId("attack-pattern")
	obj.Created = obj.GetCurrentTime()
	obj.Modified = obj.Created
	obj.Version = 1
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - AttackPatternType
// ----------------------------------------------------------------------
