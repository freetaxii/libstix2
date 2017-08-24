// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type NamePropertyType struct {
	Name string `json:"name,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - NamePropertyType
// ----------------------------------------------------------------------

// SetName takes in one parameter
// param: s - a string value representing a name of the object
func (this *NamePropertyType) SetName(s string) {
	this.Name = s
}

func (this *NamePropertyType) GetName() string {
	return this.Name
}
