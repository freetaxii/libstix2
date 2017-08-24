// Copyright 2017 Bret Jordan, All rights reserved.
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
	common.CommonObjectPropertiesType
	common.DescriptivePropertiesType
	common.KillChainPhasesPropertyType
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() AttackPatternType {
	var obj AttackPatternType
	obj.InitNewObject("atttack-pattern")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - AttackPatternType
// ----------------------------------------------------------------------
