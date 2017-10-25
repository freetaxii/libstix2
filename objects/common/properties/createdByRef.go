// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

// CreatedByRefPropertyType - A property used by one or more STIX objects that
// captures the STIX identifier of the identity that created this object.
type CreatedByRefPropertyType struct {
	CreatedByRef string `json:"created_by_ref,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - CreatedByRefPropertyType
// ----------------------------------------------------------------------

// SetCreatedByRef - This method takes in a string value representing a STIX
// identifier and updates the Created By Ref property.
func (p *CreatedByRefPropertyType) SetCreatedByRef(s string) {
	p.CreatedByRef = s
}

// GetCreatedByRef - This method returns the STIX identifier for the identity
// that created this object.
func (p *CreatedByRefPropertyType) GetCreatedByRef() string {
	return p.CreatedByRef
}
