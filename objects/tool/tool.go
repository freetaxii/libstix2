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

The following information comes directly from the STIX 2 specification documents.

Tools are legitimate software that can be used by threat actors to perform
attacks. Knowing how and when threat actors use such tools can be important for
understanding how campaigns are executed. Unlike malware, these tools or
software packages are often found on a system and have legitimate purposes for
power users, system administrators, network administrators, or even normal
users. Remote access tools (e.g., RDP) and network scanning tools (e.g., Nmap)
are examples of Tools that may be used by a Threat Actor during an attack.

The Tool SDO characterizes the properties of these software tools and can be
used as a basis for making an assertion about how a Threat Actor uses them
during an attack. It contains properties to name and describe the tool, a list
of Kill Chain Phases the tool can be used to carry out, and the version of the
tool.

This SDO MUST NOT be used to characterize malware. Further, Tool MUST NOT be
used to characterize tools used as part of a course of action in response to an
attack. Tools used during response activities can be included directly as part
of a Course of Action SDO.
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
into an actual struct. It will return:
 - the object as a pointer
 - the STIX ID
 - the SITX Version
 - any errors found
*/
func Decode(data []byte) (*Tool, string, string, error) {
	var o Tool
	err := json.Unmarshal(data, &o)
	if err != nil {
		return nil, "", "", err
	}

	if valid, err := o.Valid(); valid != true {
		return nil, "", "", err
	}

	o.SetRawData(data)
	return &o, o.ID, o.Modified, nil
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
