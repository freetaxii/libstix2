// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
DescriptionProperty - A property used by one or more STIX objects that
captures the description for the object as a string.
*/
type DescriptionProperty struct {
	Description string `json:"description,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - DescriptionProperty
// ----------------------------------------------------------------------

/*
SetDescription - This method takes in a string value representing a text
description and updates the description property.
*/
func (o *DescriptionProperty) SetDescription(s string) error {
	o.Description = s
	return nil
}

/*
GetDescription - This method returns the description for an object as a string.
*/
func (o *DescriptionProperty) GetDescription() string {
	return o.Description
}
