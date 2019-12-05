// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
ValueProperty -
*/
type ValueProperty struct {
	Value string `json:"value,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - ValueProperty
// ----------------------------------------------------------------------

/*
SetValue -
*/
func (o *ValueProperty) SetValue(val string) error {
	o.Value = val
	return nil
}
