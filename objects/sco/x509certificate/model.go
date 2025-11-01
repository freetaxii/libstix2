// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package x509certificate

import (
"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
X509Certificate - This type implements the STIX 2.1 X509Certificate SCO and defines all of the
properties and methods needed to create and work with this object. All of the
methods not defined local to this type are inherited from the individual
properties.

Reference: STIX 2.1 specification section 6.17

TODO: Complete implementation of all properties per specification
*/
type X509Certificate struct {
objects.CommonObjectProperties
// TODO: Add specific properties for X509Certificate based on STIX 2.1 spec section 6.17
// TODO: Add is_self_signed, hashes, version, serial_number, signature_algorithm, issuer, validity_not_before, validity_not_after, subject, subject_public_key_algorithm, subject_public_key_modulus, subject_public_key_exponent, x509_v3_extensions
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *X509Certificate) GetPropertyList() []string {
// TODO: Update with actual property names
return []string{}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX X509Certificate SCO and return it as a
pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *X509Certificate {
var obj X509Certificate
obj.InitSCO("x509certificate")
return &obj
}
