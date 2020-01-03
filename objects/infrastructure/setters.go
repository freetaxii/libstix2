// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package infrastructure

import (
	"errors"
	"strings"
)

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/* AddTypes - This method takes in a string value, a comma separated list of
string values, or a slice of string values that all representing a
categorization for this infrastructure. The values SHOULD come from the
infrastructure-type-ov open vocabulary. */
func (o *Infrastructure) AddTypes(data interface{}) error {

	switch data.(type) {
	case string:
		types := strings.Split(data.(string), ",")
		o.InfrastructureTypes = append(o.InfrastructureTypes, types...)
	case []string:
		o.InfrastructureTypes = append(o.InfrastructureTypes, data.([]string)...)
	default:
		return errors.New("wrong data type passed in to AddTypes()")
	}

	return nil
}
