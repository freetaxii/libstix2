// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package tool

import "github.com/freetaxii/libstix2/objects"

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/*
AddTypes - This method takes in a string value, a comma separated list of
string values, or a slice of string values that represents an tool type and
adds it to the tool types property. The values SHOULD come from the
tool-type-ov open vocabulary.
*/
func (o *Tool) AddTypes(values interface{}) error {
	return objects.AddValuesToList(&o.ToolTypes, values)
}

/*
SetToolVersion - This method takes in a string value representing the version
of the tool and updates the tool version property.
*/
func (o *Tool) SetToolVersion(s string) error {
	o.ToolVersion = s
	return nil
}
