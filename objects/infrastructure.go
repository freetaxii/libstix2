// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

import (
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
//
// Define Object Type
//
// ----------------------------------------------------------------------

/*
Infrastructure - This type implements the STIX 2 Infrastructure SDO and defines
all of the properties methods needed to create and work with the STIX Infrastructure
SDO. All of the methods not defined local to this type are inherited from
the individual properties.

The following information comes directly from the STIX 2 specification documents.
*/
type Infrastructure struct {
	properties.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	properties.KillChainPhasesProperty
}

// ----------------------------------------------------------------------
//
// Initialization Functions
//
// ----------------------------------------------------------------------

/*
NewInfrastructure - This function will create a new STIX Infrastructure object
and return it as a pointer.
*/
func NewInfrastructure() *Infrastructure {
	var obj Infrastructure
	obj.InitObject("infrastructure")
	return &obj
}

// ----------------------------------------------------------------------
//
// Public Methods - Infrastructure
//
// ----------------------------------------------------------------------
