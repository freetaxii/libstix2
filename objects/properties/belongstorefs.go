// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import (
	"errors"
	"strings"
)

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

/* AddBelongsToRefs - This method takes in a string value, a comma separated
list of string values, or a slice of string values that all representing an id
of an objects that this object belongs / is related to. */
func (o *BelongsToRefsProperty) AddBelongsToRefs(data interface{}) error {

	switch data.(type) {
	case string:
		types := strings.Split(data.(string), ",")
		o.BelongsToRefs = append(o.BelongsToRefs, types...)
	case []string:
		o.BelongsToRefs = append(o.BelongsToRefs, data.([]string)...)
	default:
		return errors.New("wrong data type passed in to AddBelongsToRefs()")
	}

	return nil
}
