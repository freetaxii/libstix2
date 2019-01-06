// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package intrusionset

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
IntrusionSet - This type implements the STIX 2 Intrusion Set SDO and defines
all of the properties methods needed to create and work with the STIX Intrusion Set
SDO. All of the methods not defined local to this type are inherited from
the individual properties.

The following information comes directly from the STIX 2 specification documents.

An Intrusion Set is a grouped set of adversarial behaviors and resources with
common properties that is believed to be orchestrated by a single organization.
An Intrusion Set may capture multiple Campaigns or other activities that are all
tied together by shared attributes indicating a common known or unknown Threat
Actor. New activity can be attributed to an Intrusion Set even if the Threat
Actors behind the attack are not known. Threat Actors can move from supporting
one Intrusion Set to supporting another, or they may support multiple Intrusion
Sets.

Where a Campaign is a set of attacks over a period of time against a specific
set of targets to achieve some objective, an Intrusion Set is the entire attack
package and may be used over a very long period of time in multiple Campaigns to
achieve potentially multiple purposes.

While sometimes an Intrusion Set is not active, or changes focus, it is usually
difficult to know if it has truly disappeared or ended. Analysts may have
varying level of fidelity on attributing an Intrusion Set back to Threat Actors
and may be able to only attribute it back to a nation state or perhaps back to
an organization within that nation state.
*/
type IntrusionSet struct {
	baseobject.CommonObjectProperties
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
