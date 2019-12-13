// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import "github.com/freetaxii/libstix2/resources/helpers"

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
ResolvesToRefsProperty -
*/
type ResolvesToRefsProperty struct {
	ResolvesToRefs []string `json:"resolves_to_refs,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - ResolvesToRefsProperty
// ----------------------------------------------------------------------

/*
AddResolvesToRefs -
*/
func (o *ResolvesToRefsProperty) AddResolvesToRefs(ids []string) error {
	arr, err := helpers.AddToList(o.ResolvesToRefs, ids)

	if err != nil {
		return err
	}

	o.ResolvesToRefs = arr

	return nil
}

// AddResolvesToRef -
func (o *ResolvesToRefsProperty) AddResolvesToRef(id string) error {
	o.ResolvesToRefs = append(o.ResolvesToRefs, id)

	return nil
}
