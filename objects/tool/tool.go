// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package tool

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
Tool - This type implements the STIX 2 Tool SDO and defines
all of the properties methods needed to create and work with the STIX Tool
SDO. All of the methods not defined local to this type are inherited from
the individual properties.
*/
type Tool struct {
	baseobject.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	properties.KillChainPhasesProperty
	ToolVersion string `json:"tool_version,omitempty"`
}

// ----------------------------------------------------------------------
//
// Initialization Functions
//
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Tool object and return it as a
pointer.
*/
func New() *Tool {
	var obj Tool
	obj.InitObject("tool")
	return &obj
}

// ----------------------------------------------------------------------
// Public Methods - Tool - Core Functionality
// ----------------------------------------------------------------------

/*
Decode - This function will decode some JSON data encoded as a slice of bytes
into an actual struct. It will return the object as a pointer and any errors found.
*/
func Decode(data []byte) (*Tool, error) {
	var o Tool
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
func (o *Tool) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
EncodeToString - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *Tool) EncodeToString() (string, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

/*
Valid - This method will verify all of the properties on the object.
*/
func (o *Tool) Valid() (bool, error) {

	// Check common base properties first
	if valid, err := o.CommonObjectProperties.Valid(); valid != true {
		return false, err
	}

	return true, nil
}

// ----------------------------------------------------------------------
// Public Methods - Tool
// ----------------------------------------------------------------------

/*
SetToolVersion - This method takes in a string value representing the version
of the tool and updates the tool version property.
*/
func (o *Tool) SetToolVersion(s string) error {
	o.ToolVersion = s
	return nil
}
