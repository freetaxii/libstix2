// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package domainnameobject

import (
	"errors"
)

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/*
Valid - This method will verify and test all of the properties on the object to
make sure they are valid per the specification.
*/
func (o *DomainNameObject) Valid() (bool, error) {

	// Check common base properties first
	if valid, err := o.CommonObjectProperties.Valid(); valid != true {
		return false, err
	}

	// Check DomainNameObject Specific Properties
	if valid, err := o.validValue(); valid == false {
		return valid, err
	}

	return true, nil
}

// ----------------------------------------------------------------------
// Private Methods
// ----------------------------------------------------------------------

func (o *DomainNameObject) validValue() (bool, error) {
	if o.Value == "" {
		return false, errors.New("the Value property is required, but missing")
	}

	return true, nil
}
