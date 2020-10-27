// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package infrastructure

import "github.com/wxj95/libstix2/resources"

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/*
AddTypes - This method takes in a string value, a comma separated list of
string values, or a slice of string values that represents an infrastructure
type and adds it to the infrastructure types property. The values SHOULD come
from the infrastructure-type-ov open vocabulary.
*/
func (o *Infrastructure) AddTypes(values interface{}) error {
	return resources.AddValuesToList(&o.InfrastructureTypes, values)
}
