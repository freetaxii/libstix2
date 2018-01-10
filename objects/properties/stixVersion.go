// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
STIXVersionPropertyType - A property used by one or more STIX objects that
captures the STIX specification version. This is not included in the JSON
serialization, but is used for writing to the database.
*/
type STIXVersionPropertyType struct {
	SpecVersion string `json:"-"`
}

// ----------------------------------------------------------------------
// Public Methods - STIXVersionPropertyType
// ----------------------------------------------------------------------

/*
SetSpecVersion - This method takes in a string representing a STIX specification
version and updates the Version property.
*/
func (ezt *STIXVersionPropertyType) SetSpecVersion(s string) error {
	ezt.SpecVersion = s
	return nil
}

/*
GetSpecVersion - This method returns the version value as a string.
*/
func (ezt *STIXVersionPropertyType) GetSpecVersion() string {
	return ezt.SpecVersion
}
