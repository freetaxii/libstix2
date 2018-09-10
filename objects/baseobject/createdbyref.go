// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package baseobject

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
CreatedByRefProperty - A property used by all STIX objects that captures the STIX
identifier of the identity that created this object.
*/
type CreatedByRefProperty struct {
	CreatedByRef string `json:"created_by_ref,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - CreatedByRefProperty
// ----------------------------------------------------------------------

/*
SetCreatedByRef - This method takes in a string value representing a STIX
identifier and updates the Created By Ref property.
*/
func (o *CreatedByRefProperty) SetCreatedByRef(s string) error {
	o.CreatedByRef = s
	return nil
}

/*
GetCreatedByRef - This method returns the STIX identifier for the identity
that created this object.
*/
func (o *CreatedByRefProperty) GetCreatedByRef() string {
	return o.CreatedByRef
}
