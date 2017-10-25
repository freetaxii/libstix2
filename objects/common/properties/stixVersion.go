// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

// STIXVersionPropertyType - A property used by one or more STIX objects that
// captures the STIX specification version. This is not included in the JSON
// serialization, but is used for writing to the database.
type STIXVersionPropertyType struct {
	Version string `json:"-"`
}

// ----------------------------------------------------------------------
// Public Methods - STIXVersionPropertyType
// ----------------------------------------------------------------------

// SetVersion - This method takes in a string representing a STIX specification
// version and updates the Version property.
func (p *STIXVersionPropertyType) SetSpecVersion(s string) {
	p.Version = s
}

// GetVersion - This method returns the version value as a string.
func (p *STIXVersionPropertyType) GetSpecVersion() string {
	return p.Version
}
