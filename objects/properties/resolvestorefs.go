// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import "github.com/freetaxii/libstix2/resources"

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/*
ResolvesToRefsProperty -
*/
type ResolvesToRefsProperty struct {
	ResolvesToRefs []string `json:"resolves_to_refs,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - ResolvesToRefsProperty - Setters
// ----------------------------------------------------------------------

/*
AddResolvesToRefs - This method takes in a string value, a comma separated
list of string values, or a slice of string values that represents an id of an
object that this resolves to and adds it to the resolves to refs property.
*/
func (o *ResolvesToRefsProperty) AddResolvesToRefs(values interface{}) error {
	return resources.AddValuesToList(&o.ResolvesToRefs, values)
}
