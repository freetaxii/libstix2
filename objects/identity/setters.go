// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package identity

import "github.com/freetaxii/libstix2/resources"

// ----------------------------------------------------------------------
// Public Methods - Identity
// ----------------------------------------------------------------------

/* SetIdentityClass - This method takes in a string value representing a STIX
identity class from the vocab identity-class-ov and updates the identity class
property. */
func (o *Identity) SetIdentityClass(s string) error {
	o.IdentityClass = s
	return nil
}

/* AddSectors - This method takes in a string value, a comma separated list of
string values, or a slice of string values that represents a sector and adds it
to the sectors property. */
func (o *Identity) AddSectors(values interface{}) error {
	return resources.AddValuesToList(&o.Sectors, values)
}

/* SetContactInformation - This method takes in a string value representing
contact information as a text string and updates the contact information
property. */
func (o *Identity) SetContactInformation(s string) error {
	o.ContactInformation = s
	return nil
}
