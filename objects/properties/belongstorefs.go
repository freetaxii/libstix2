// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import "github.com/freetaxii/libstix2/resources/helpers"

// ----------------------------------------------------------------------
// Types
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
AddBelongsToRefs -
*/
func (o *BelongsToRefsProperty) AddBelongsToRefs(ids []string) error {
	arr, err := helpers.AddToList(o.BelongsToRefs, ids)

	if err != nil {
		return err
	}

	o.BelongsToRefs = arr
	return nil
}

// AddResolvesToRef -
func (o *BelongsToRefsProperty) AddBelongsToRef(id string) error {
	o.BelongsToRefs = append(o.BelongsToRefs, id)

	return nil
}
