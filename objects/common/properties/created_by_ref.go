// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type CreatedByRefPropertyType struct {
	Created_by_ref string `json:"created_by_ref,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - CreatedByRefPropertyType
// ----------------------------------------------------------------------

// SetCreatedBy takes in one parameter
// param: s - a string value representing a STIX Identifier
func (this *CreatedByRefPropertyType) SetCreatedByRef(s string) {
	this.Created_by_ref = s
}

func (this *CreatedByRefPropertyType) GetCreatedByRef() string {
	return this.Created_by_ref
}
