// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package intrusionSet

import (
	"github.com/freetaxii/libstix2/objects/common/properties"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

// IntrusionSetType -
// This type defines all of the properties associated with the STIX Intrusion Set SDO.
// All of the methods not defined local to this type are inherited from the individual properties.
type IntrusionSetType struct {
	properties.CommonObjectPropertiesType
	properties.NamePropertyType
	properties.DescriptionPropertyType
	properties.AliasesPropertyType
	properties.FirstSeenPropertyType
	properties.LastSeenPropertyType
	properties.GoalsPropertyType
	properties.ResourceLevelPropertyType
	properties.PrimaryMotivationPropertyType
	properties.SecondaryMotivationsPropertyType
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// New - This function will create a new intrusion set object.
func New() IntrusionSetType {
	var obj IntrusionSetType
	obj.InitNewObject("intrusion-set")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - IntrusionSetType
// ----------------------------------------------------------------------
