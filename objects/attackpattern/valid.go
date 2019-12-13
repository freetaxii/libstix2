// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package attackpattern

import "errors"

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/* Valid - This method will verify and test all of the properties on the object
to make sure they are valid per the specification. */
func (o *AttackPattern) Valid() (bool, error) {
	return Valid(o)
}

// ----------------------------------------------------------------------
// Public Functions
// ----------------------------------------------------------------------
/* Valid - This method will verify and test all of the properties on the object
to make sure they are valid per the specification. */
func Valid(o *AttackPattern) (bool, error) {
	// Check common base properties first
	if valid, err := o.CommonObjectProperties.Valid(); valid != true {
		return false, err
	}

	// Check attack pattern specific properties
	if valid, err := validAttackPatternName(o); valid != true {
		return valid, err
	}

	return true, nil
}

// ----------------------------------------------------------------------
// Private Functions
// ----------------------------------------------------------------------

func validAttackPatternName(o *AttackPattern) (bool, error) {
	if o.Name == "" {
		return false, errors.New("the name property is required but missing")
	}
	return true, nil
}
