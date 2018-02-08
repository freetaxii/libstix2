// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
//
// Types
//
// ----------------------------------------------------------------------

/*
DescriptionPropertyType - A property used by one or more TAXII resources.
*/
type DescriptionPropertyType struct {
	Description string `json:"description,omitempty"`
}

// ----------------------------------------------------------------------
//
// Public Methods - DescriptionPropertyType
//
// ----------------------------------------------------------------------

/*
SetDescription - This method takes in a string value representing a text
description and udpates the description property.
*/
func (p *DescriptionPropertyType) SetDescription(s string) error {
	p.Description = s
	return nil
}

/*
GetDescription - This method returns the description.
*/
func (p *DescriptionPropertyType) GetDescription() string {
	return p.Description
}
