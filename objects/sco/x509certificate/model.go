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
X509Certificate - This type implements the STIX 2 Domain Name SCO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties.
*/
type X509Certificate struct {
	objects.CommonObjectProperties
	objects.ValueProperty
	IsSelfSigned              bool              `json:"is_self_signed,omitempty" bson:"is_self_signed,omitempty"`
	Hashes                    map[string]string `json:"hashes,omitempty" bson:"hashes,omitempty"`
	Version                   string            `json:"version,omitempty" bson:"version,omitempty"`
	SerialNumber              string            `json:"serial_number,omitempty" bson:"serial_number,omitempty"`
	SignatureAlgorithm        string            `json:"signature_algorithm,omitempty" bson:"signature_algorithm,omitempty"`
	Issuer                    string            `json:"issuer,omitempty" bson:"issuer,omitempty"`
	ValidityNotBefore         string            `json:"validity_not_before,omitempty" bson:"validity_not_before,omitempty"`
	ValidityNotAfter          string            `json:"validity_not_after,omitempty" bson:"validity_not_after,omitempty"`
	Subject                   string            `json:"subject,omitempty" bson:"subject,omitempty"`
	SubjectPublicKeyAlgorithm string            `json:"subject_public_key_algorithm,omitempty" bson:"subject_public_key_algorithm,omitempty"`
	SubjectPublicKeyModulus   string            `json:"subject_public_key_modulus,omitempty" bson:"subject_public_key_modulus,omitempty"`
	SubjectPublicKeyExponent  int               `json:"subject_public_key_exponent,omitempty" bson:"subject_public_key_exponent,omitempty"`
	//TODO add X.509 V3 extension
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *X509Certificate) GetPropertyList() []string {
	return []string{
		"value",
		"is_self_signed",
		"hashes",
		"version",
		"serial_number",
		"signature_algorithm",
		"issuer",
		"validity_not_before",
		"validity_not_after",
		"subject",
		"subject_public_key_algorithm",
		"subject_public_key_modulus",
		"subject_public_key_exponent",
	}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Domain Name SCO and return it as a
pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *X509Certificate {
	var obj X509Certificate
	obj.InitSCO("x509-certificate")
	return &obj
}
