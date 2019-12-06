// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package infrastructure

import (
	"errors"

	"github.com/freetaxii/libstix2/resources/helpers"
	"github.com/freetaxii/libstix2/vocabs"
)

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/*
Valid - This method will verify and test all of the properties on the object to
make sure they are valid per the specification.
*/
func (o *Infrastructure) Valid() (bool, error) {

	// Check common base properties first
	if valid, err := o.CommonObjectProperties.Valid(); valid != true {
		return false, err
	}

	// Check Infrastructure Specific Properties
	if valid, err := o.validInfrastructureTypes(); valid == false {
		return valid, err
	}

	return true, nil
}

// ----------------------------------------------------------------------
// Private Methods
// ----------------------------------------------------------------------

func (o *Infrastructure) validInfrastructureTypes() (bool, error) {
	if len(o.InfrastructureTypes) == 0 {
		return false, errors.New("the InfrastructureTypes property is required, but missing")
	}

	return helpers.ValidSlice("InfrastructureTypes", o.InfrastructureTypes, vocabs.InfrastructureType)
}
