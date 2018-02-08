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
RevokedPropertyType - A property used by one or more STIX objects that
captures whether or not this STIX object has been revoked by the object
// creator.
*/
type RevokedPropertyType struct {
	Revoked bool `json:"revoked,omitempty"`
}

// ----------------------------------------------------------------------
//
// Public Methods - RevokedPropertyType
//
// ----------------------------------------------------------------------

/*
SetRevoked - This method sets the revoked boolean to true
*/
func (p *RevokedPropertyType) SetRevoked() error {
	p.Revoked = true
	return nil
}

/*
GetRevoked - This method returns the current value of the revoked property.
*/
func (p *RevokedPropertyType) GetRevoked() bool {
	return p.Revoked
}
