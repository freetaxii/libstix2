// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package attackpattern

import (
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
AttackPattern - This type implements the STIX 2 Attack Pattern SDO and defines
all of the properties methods needed to create and work with the STIX Attack Pattern
SDO. All of the methods not defined local to this type are inherited from
the individual properties.
*/
type AttackPattern struct {
	objects.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	properties.AliasesProperty
	properties.KillChainPhasesProperty
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Attack Pattern object and return it
as a pointer.
*/
func New() *AttackPattern {
	var obj AttackPattern
	obj.InitObject("attack-pattern")
	return &obj
}

// ----------------------------------------------------------------------
// Public Methods - Attack Pattern
// ----------------------------------------------------------------------

/*
Valid - This method will verify all of the properties on the object.
*/
func (o *AttackPattern) Valid() (bool, error) {

	// Check common base properties first
	if valid, err := o.CommonObjectProperties.Valid(); valid != true {
		return false, err
	}

	return true, nil
}
