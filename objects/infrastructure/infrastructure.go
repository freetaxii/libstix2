// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package infrastructure

import (
	"github.com/freetaxii/libstix2/objects/common/properties"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
InfrastructureType defines all of the properties associated with the STIX
Infrastructure SDO. All of the methods not defined local to this type are
inherited from the individual properties.
*/
type InfrastructureType struct {
	properties.CommonObjectPropertiesType
	properties.NamePropertyType
	properties.DescriptionPropertyType
	properties.KillChainPhasesPropertyType
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// New - This function will create a new infrastructure object.
func New() InfrastructureType {
	var obj InfrastructureType
	obj.InitNewObject("infrastructure")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - InfrastructureType
// ----------------------------------------------------------------------
