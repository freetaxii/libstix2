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
NameProperty - A property used by one or more STIX objects that captures a
vanity name for the STIX object in string format.
*/
type NameProperty struct {
	Name string `json:"name,omitempty"`
}

// ----------------------------------------------------------------------
//
// Public Methods - NameProperty
//
// ----------------------------------------------------------------------

/*
SetName - This method takes in a string value representing a name of the object
and updates the name property.
*/
func (p *NameProperty) SetName(s string) error {
	p.Name = s
	return nil
}

/*
GetName - This method returns the current name of the object.
*/
func (p *NameProperty) GetName() string {
	return p.Name
}
