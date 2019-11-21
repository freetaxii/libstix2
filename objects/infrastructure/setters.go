// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package infrastructure

import (
	"github.com/freetaxii/libstix2/resources/helpers"
)

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/*
AddInfrastructureTypes - add malware types.

The type of infrastructure being described.

This is an open vocabulary and values SHOULD come from the infrastructure-type-ov vocabulary.
*/
func (o *Infrastructure) AddInfrastructureTypes(data interface{}) error {
	arr, err := helpers.AddToList(o.InfrastructureTypes, data)

	if err != nil {
		return err
	}

	o.InfrastructureTypes = arr
	return nil
}

/*
AddInfrastructureType - add malware type.

The type of infrastructure being described.

This is an open vocabulary and values SHOULD come from the infrastructure-type-ov vocabulary.
*/
func (o *Infrastructure) AddInfrastructureType(s string) error {
	o.InfrastructureTypes = append(o.InfrastructureTypes, s)

	return nil
}
