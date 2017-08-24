// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type DescriptionPropertyType struct {
	Description string `json:"description,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - DescriptionPropertyType
// ----------------------------------------------------------------------

// SetDescription takes in one parameter
// param: s - a string value representing a text description
func (this *DescriptionPropertyType) SetDescription(s string) {
	this.Description = s
}

func (this *DescriptionPropertyType) GetDescription() string {
	return this.Description
}
