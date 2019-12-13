// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package intrusionset

import (
	"encoding/json"

	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
//
// Define Object Type
//
// ----------------------------------------------------------------------

/*
IntrusionSet - This type implements the STIX 2 Intrusion Set SDO and defines
all of the properties methods needed to create and work with the STIX Intrusion Set
SDO. All of the methods not defined local to this type are inherited from
the individual properties.
*/
type IntrusionSet struct {
	objects.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	properties.AliasesProperty
	properties.SeenTimestampProperties
	properties.GoalsProperty
	properties.ResourceLevelProperty
	properties.MotivationProperties
}

// ----------------------------------------------------------------------
//
// Initialization Functions
//
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Intrusion Set object and return it as
a pointer.
*/
func New() *IntrusionSet {
	var obj IntrusionSet
	obj.InitObject("intrusion-set")
	return &obj
}

// ----------------------------------------------------------------------
// Public Methods - Intrusion Set - Core Functionality
// ----------------------------------------------------------------------

/*
Decode - This function will decode some JSON data encoded as a slice of bytes
into an actual struct. It will return the object as a pointer and any errors found.
*/
func Decode(data []byte) (*IntrusionSet, error) {
	var o IntrusionSet
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
func (o *IntrusionSet) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
EncodeToString - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *IntrusionSet) EncodeToString() (string, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

/*
Valid - This method will verify all of the properties on the object.
*/
func (o *IntrusionSet) Valid() (bool, error) {

	// Check common base properties first
	if valid, err := o.CommonObjectProperties.Valid(); valid != true {
		return false, err
	}

	return true, nil
}

// ----------------------------------------------------------------------
// Public Methods - IntrusionSet
// ----------------------------------------------------------------------
