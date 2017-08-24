// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package infrastructure

import (
	"github.com/freetaxii/libstix2/objects/common"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type InfrastructureType struct {
	common.CommonObjectPropertiesType
	common.DescriptivePropertiesType
	common.KillChainPhasesPropertyType
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() InfrastructureType {
	var obj InfrastructureType
	obj.InitNewObject("infrastructure")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - InfrastructureType
// ----------------------------------------------------------------------
