// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package threatactor

import (
	"encoding/json"

	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Type
// ----------------------------------------------------------------------

/* ThreatActor - This type implements the STIX 2 Threat Actor SDO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties. */
type ThreatActor struct {
	objects.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	ThreatActorType []string `json:"threat_actor_types,omitempty"`
	properties.AliasesProperty
	Roles []string `json:"roles,omitempty"`
	properties.GoalsProperty
	Sophistication string `json:"sophistication,omitempty"`
	properties.ResourceLevelProperty
	properties.MotivationProperties
	PersonalMotivations []string `json:"personal_motivations,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/* New - This function will create a new STIX Threat Actor object and return it
as a pointer. It will also initialize the object by setting all of the basic
properties. */
func New() *ThreatActor {
	var obj ThreatActor
	obj.InitSDO("threat-actor")
	return &obj
}

// ----------------------------------------------------------------------
// Public Methods - Threat Actor - Core Functionality
// ----------------------------------------------------------------------

/*
Decode - This function will decode some JSON data encoded as a slice of bytes
into an actual struct. It will return the object as a pointer and any errors found.
*/
func Decode(data []byte) (*ThreatActor, error) {
	var o ThreatActor
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
func (o *ThreatActor) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
EncodeToString - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *ThreatActor) EncodeToString() (string, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

/*
Valid - This method will verify all of the properties on the object.
*/
func (o *ThreatActor) Valid() (bool, error) {

	// Check common base properties first
	if valid, err := o.CommonObjectProperties.Valid(); valid != true {
		return false, err
	}

	return true, nil
}

// ----------------------------------------------------------------------
// Public Methods - ThreatActor
// ----------------------------------------------------------------------

/*
AddType - This method takes in a string value representing a threat actor
type from the threat-actor-type-ov and adds it to the threat actor type property.
*/
func (o *ThreatActor) AddType(s string) error {
	o.ThreatActorType = append(o.ThreatActorType, s)
	return nil
}

/*
AddRole - This method takes in a string value representing a threat actor
role from the threat-actor-role-ov and adds it to the role property.
*/
func (o *ThreatActor) AddRole(s string) error {
	o.Roles = append(o.Roles, s)
	return nil
}

/*
SetSophistication - This method takes in a string value representing the
sophistication level of a threat actor from the threat-actor-sophistication-ov
and adds it to the sophistication property.
*/
func (o *ThreatActor) SetSophistication(s string) error {
	o.Sophistication = s
	return nil
}

/*
AddPersonalMotivation - This method takes in a string value representing the
motivation of a threat actor from the threat-actor-motivation-ov and adds it
to the personal motivations property.
*/
func (o *ThreatActor) AddPersonalMotivation(s string) error {
	o.PersonalMotivations = append(o.PersonalMotivations, s)
	return nil
}
