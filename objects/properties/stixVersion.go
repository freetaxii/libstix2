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
STIXVersionPropertyType - A property used by one or more STIX objects that
captures the STIX specification version.
*/
type STIXVersionPropertyType struct {
	SpecVersion string `json:"spec_version,omitempty"`
}

// ----------------------------------------------------------------------
//
// Public Methods - STIXVersionPropertyType
//
// ----------------------------------------------------------------------

/*
SetSpecVersion20 - This method will set the specification version to 2.0.
*/
func (p *STIXVersionPropertyType) SetSpecVersion20() error {
	p.SpecVersion = "2.0"
	return nil
}

/*
SetSpecVersion20 - This method will set the specification version to 2.1.
*/
func (p *STIXVersionPropertyType) SetSpecVersion21() error {
	p.SpecVersion = "2.1"
	return nil
}

/*
SetSpecVersion - This method takes in a string representing a STIX specification
version and updates the Version property.
*/
func (p *STIXVersionPropertyType) SetSpecVersion(s string) error {
	p.SpecVersion = s
	return nil
}

/*
GetSpecVersion - This method returns the version value as a string.
*/
func (p *STIXVersionPropertyType) GetSpecVersion() string {
	return p.SpecVersion
}
