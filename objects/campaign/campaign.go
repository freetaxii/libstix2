// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package campaign

import (
	"encoding/json"
	"errors"

	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
//
// Define Object Type
//
// ----------------------------------------------------------------------

/*
Campaign - This type implements the STIX 2 Campaign SDO and defines
all of the properties methods needed to create and work with the STIX Campaign
SDO. All of the methods not defined local to this type are inherited from
the individual properties.
*/
type Campaign struct {
	objects.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	properties.AliasesProperty
	properties.SeenTimestampProperties
	Objective string `json:"objective,omitempty"`
}

// ----------------------------------------------------------------------
//
// Initialization Functions
//
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Campaign object and return it as a
pointer. It will also initialize the object by setting all of the basic properties.
*/
func New() *Campaign {
	var obj Campaign
	obj.InitObject("campaign")
	return &obj
}

// ----------------------------------------------------------------------
// Public Methods - Campaign - Core Functionality
// ----------------------------------------------------------------------

/*
Decode - This function will decode some JSON data encoded as a slice of bytes
into an actual struct. It will return the object as a pointer and any errors found.
*/
func Decode(data []byte) (*Campaign, error) {
	var o Campaign
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
func (o *Campaign) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
EncodeToString - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *Campaign) EncodeToString() (string, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

/*
Valid - This method will verify all of the properties on the object.
*/
func (o *Campaign) Valid() (bool, error) {

	// Check common base properties first
	if valid, err := o.CommonObjectProperties.Valid(); valid != true {
		return false, err
	}

	if o.Name == "" {
		return false, errors.New("the name property is required, but missing")
	}

	return true, nil
}

// ----------------------------------------------------------------------
// Public Methods - Campaign
// ----------------------------------------------------------------------

/*
SetObjective - This method will take in a string representing an objective,
goal, desired outcome, or intended effect and update the objective property.
*/
func (o *Campaign) SetObjective(s string) error {
	o.Objective = s
	return nil
}
