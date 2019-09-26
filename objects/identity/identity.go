// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package identity

import (
	"encoding/json"

	"github.com/freetaxii/libstix2/objects/baseobject"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
//
// Define Object Type
//
// ----------------------------------------------------------------------

/*
Identity - This type implements the STIX 2 Identity SDO and defines
all of the properties methods needed to create and work with the STIX Identity
SDO. All of the methods not defined local to this type are inherited from
the individual properties.
*/
type Identity struct {
	baseobject.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	IdentityClass      string   `json:"identity_class,omitempty"`
	Sectors            []string `json:"sectors,omitempty"`
	ContactInformation string   `json:"contact_information,omitempty"`
}

// ----------------------------------------------------------------------
//
// Initialization Functions
//
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Identity object and return it as a
pointer.
*/
func New() *Identity {
	var obj Identity
	obj.InitObject("identity")
	return &obj
}

// ----------------------------------------------------------------------
// Public Methods - Identity - Core Functionality
// ----------------------------------------------------------------------

/*
Decode - This function will decode some JSON data encoded as a slice of bytes
into an actual struct. It will return the object as a pointer and any errors found.
*/
func Decode(data []byte) (*Identity, error) {
	var o Identity
	err := json.Unmarshal(data, &o)
	if err != nil {
		return nil, err
	}

	if valid, err := o.Valid(); valid != true {
		return nil, err
	}

	o.SetRawData(data)
	return &o, nil
}

/*
Encode - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *Identity) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
EncodeToString - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *Identity) EncodeToString() (string, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

/*
Valid - This method will verify all of the properties on the object.
*/
func (o *Identity) Valid() (bool, error) {

	// Check common base properties first
	if valid, err := o.CommonObjectProperties.Valid(); valid != true {
		return false, err
	}

	return true, nil
}

// ----------------------------------------------------------------------
// Public Methods - Identity
// ----------------------------------------------------------------------

/*
SetIdentityClass - This method takes in a string value representing a STIX
identity class from the vocab identity-class-ov and updates the identity class
property.
*/
func (o *Identity) SetIdentityClass(s string) error {
	o.IdentityClass = s
	return nil
}

/*
AddSector - This method takes in a string value that represents a STIX sector
from the vocab industry-sector-ov and adds it to the identity object.
*/
func (o *Identity) AddSector(s string) error {
	o.Sectors = append(o.Sectors, s)
	return nil
}

/*
SetContactInformation - This method takes in a string value representing
contact information as a text string and updates the contact information
property.
*/
func (o *Identity) SetContactInformation(s string) error {
	o.ContactInformation = s
	return nil
}
