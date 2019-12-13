// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package infrastructure

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
Infrastructure - This type implements the STIX 2 Infrastructure SDO and defines
all of the properties methods needed to create and work with the STIX Infrastructure
SDO. All of the methods not defined local to this type are inherited from
the individual properties.
*/
type Infrastructure struct {
	objects.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	properties.KillChainPhasesProperty
	properties.AliasesProperty
	properties.SeenProperties
	InfrastructureTypes []string `json:"infrastructure_types,omitempty"`
}

// ----------------------------------------------------------------------
//
// Initialization Functions
//
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Infrastructure object and return it
as a pointer.
*/
func New() *Infrastructure {
	var obj Infrastructure
	obj.InitObject("infrastructure")
	return &obj
}

// ----------------------------------------------------------------------
// Public Methods - Infrastructure - Core Functionality
// ----------------------------------------------------------------------

/*
Decode - This function will decode some JSON data encoded as a slice of bytes
into an actual struct. It will return the object as a pointer and any errors found.
*/
func Decode(data []byte) (*Infrastructure, error) {
	var o Infrastructure
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
func (o *Infrastructure) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
EncodeToString - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *Infrastructure) EncodeToString() (string, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
