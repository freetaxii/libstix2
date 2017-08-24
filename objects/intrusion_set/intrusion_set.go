// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package intrusion_set

import (
	"github.com/freetaxii/libstix2/objects/common"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type IntrusionSetType struct {
	common.CommonObjectPropertiesType
	common.DescriptivePropertiesType
	common.AliasesPropertiesType
	common.FirstLastSeenPropertiesType
	common.GoalsPropertiesType
	common.ResourceLevelPropertyType
	common.PrimaryMotivationPropertyType
	common.SecondaryMotivationsPropertyType
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() IntrusionSetType {
	var obj IntrusionSetType
	obj.InitNewObject("intrusion-set")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - IntrusionSetType
// ----------------------------------------------------------------------
