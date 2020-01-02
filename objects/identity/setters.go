// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package identity

// ----------------------------------------------------------------------
// Public Methods - Identity
// ----------------------------------------------------------------------

/* AddRole - This method takes in a string value that represents an Identity
/* role and adds it to the identity object. */
func (o *Identity) AddRole(s string) error {
	o.Roles = append(o.Roles, s)
	return nil
}

/* SetIdentityClass - This method takes in a string value representing a STIX
identity class from the vocab identity-class-ov and updates the identity class
property. */
func (o *Identity) SetIdentityClass(s string) error {
	o.IdentityClass = s
	return nil
}

/* AddSector - This method takes in a string value that represents a STIX sector
from the vocab industry-sector-ov and adds it to the identity object. */
func (o *Identity) AddSector(s string) error {
	o.Sectors = append(o.Sectors, s)
	return nil
}

/* SetContactInformation - This method takes in a string value representing
contact information as a text string and updates the contact information
property. */
func (o *Identity) SetContactInformation(s string) error {
	o.ContactInformation = s
	return nil
}
