// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package baseobject

import "errors"

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
SpecVersionProperty - A property used by all STIX objects that captures the
STIX specification version.
*/
type SpecVersionProperty struct {
	SpecVersion string `json:"spec_version,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - SpecVersionProperty
// ----------------------------------------------------------------------

/*
Valid - This method will ensure that the spec version property is populated and valid.
It will return a true / false and any error information.
*/
func (o *SpecVersionProperty) Valid() (bool, error) {
	if o.SpecVersion == "" {
		return false, errors.New("the spec version property is required, but missing")
	}
	return true, nil
}

/*
SetSpecVersion20 - This method will set the specification version to 2.0.
*/
func (o *SpecVersionProperty) SetSpecVersion20() error {
	o.SpecVersion = "2.0"
	return nil
}

/*
SetSpecVersion21 - This method will set the specification version to 2.1.
*/
func (o *SpecVersionProperty) SetSpecVersion21() error {
	o.SpecVersion = "2.1"
	return nil
}

/*
SetSpecVersion - This method takes in a string representing a STIX specification
version and updates the Version property.
*/
func (o *SpecVersionProperty) SetSpecVersion(s string) error {
	o.SpecVersion = s
	return nil
}

/*
GetSpecVersion - This method returns the version value as a string.
*/
func (o *SpecVersionProperty) GetSpecVersion() string {
	return o.SpecVersion
}
