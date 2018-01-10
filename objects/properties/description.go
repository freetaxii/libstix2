// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
DescriptionPropertyType - A property used by one or more STIX objects that
captures the description for the object as a string.
*/
type DescriptionPropertyType struct {
	Description string `json:"description,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - DescriptionPropertyType
// ----------------------------------------------------------------------

/*
SetDescription - This method takes in a string value representing a text
description and updates the description property.
*/
func (ezt *DescriptionPropertyType) SetDescription(s string) error {
	ezt.Description = s
	return nil
}

/*
GetDescription - This method returns the description for an object as a string.
*/
func (ezt *DescriptionPropertyType) GetDescription() (string, error) {
	return ezt.Description, nil
}
