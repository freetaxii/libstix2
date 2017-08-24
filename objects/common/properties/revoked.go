// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type RevokedPropertyType struct {
	Revoked bool `json:"revoked,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - RevokedPropertyType
// ----------------------------------------------------------------------

// SetRevoked set the revoked boolean to true
func (this *RevokedPropertyType) SetRevoked() {
	this.Revoked = true
}

func (this *RevokedPropertyType) GetRevoked() bool {
	return this.Revoked
}
