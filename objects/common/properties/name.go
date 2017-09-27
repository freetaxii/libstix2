// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

// NamePropertyType - A property used by one or more STIX objects that
// captures a vanity name for the STIX object in string format.
type NamePropertyType struct {
	Name string `json:"name,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - NamePropertyType
// ----------------------------------------------------------------------

// SetName - This method takes in a string value representing a name of the
// object and updates the name property.
func (this *NamePropertyType) SetName(s string) {
	this.Name = s
}

// GetName - This method returns the current name of the object.
func (this *NamePropertyType) GetName() string {
	return this.Name
}
