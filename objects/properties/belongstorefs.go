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
BelongsToRefsProperty -
*/
type BelongsToRefsProperty struct {
	BelongsToRefs []string `json:"belongs_to_refs,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - BelongsToRefsProperty
// ----------------------------------------------------------------------

/*
AddBelongsToRefs - This method takes in a string value, a comma separated
list of string values, or a slice of string values that represents an id that
this object belongs to and adds it to the belongs to refs property.
*/
func (o *BelongsToRefsProperty) AddBelongsToRefs(values interface{}) error {
	return resources.AddValuesToList(&o.BelongsToRefs, values)
}
