// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package attackPattern

import (
	"github.com/freetaxii/libstix2/objects/common/properties"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

// AttackPatternType -
// This type defines all of the properties associated with the STIX Attack Pattern SDO.
// All of the methods not defined local to this type are inherited from the individual properties.
type AttackPatternType struct {
	properties.CommonObjectPropertiesType
	properties.NamePropertyType
	properties.DescriptionPropertyType
	properties.KillChainPhasesPropertyType
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// New - This function will create a new attack pattern object
func New() AttackPatternType {
	var obj AttackPatternType
	obj.InitNewObject("atttack-pattern")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - AttackPatternType
// ----------------------------------------------------------------------
