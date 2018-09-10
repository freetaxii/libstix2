// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package campaign

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
Campaign - This type implements the STIX 2 Campaign SDO and defines
all of the properties methods needed to create and work with the STIX Campaign
SDO. All of the methods not defined local to this type are inherited from
the individual properties.

The following information comes directly from the STIX 2 specification documents.

A Campaign is a grouping of adversarial behaviors that describes a set of
malicious activities or attacks (sometimes called waves) that occur over a
period of time against a specific set of targets. Campaigns usually have well
defined objectives and may be part of an Intrusion Set.

Campaigns are often attributed to an intrusion set and threat actors. The threat
actors may reuse known infrastructure from the intrusion set or may set up new
infrastructure specific for conducting that campaign.

Campaigns can be characterized by their objectives and the incidents they
cause, people or resources they target, and the resources (infrastructure,
intelligence, Malware, Tools, etc.) they use.

For example, a Campaign could be used to describe a crime syndicate's attack
using a specific variant of malware and new C2 servers against the executives
of ACME Bank during the summer of 2016 in order to gain secret information
about an upcoming merger with another bank.
*/
type Campaign struct {
	baseobject.CommonObjectProperties
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

/*
Decode - This function will decode some JSON data encoded as a slice of bytes
into an actual struct. It will return the object as a pointer.
*/
func Decode(data []byte) (*Campaign, error) {
	var o Campaign
	err := json.Unmarshal(data, &o)
	if err != nil {
		return nil, err
	}

	if err := o.CommonObjectProperties.Verify(); err != nil {
		return nil, err
	}

	o.SetRawData(data)
	return &o, nil
}

// ----------------------------------------------------------------------
//
// Public Methods - Campaign
//
// ----------------------------------------------------------------------

/*
SetObjective - This method will take in a string representing an objective,
goal, desired outcome, or intended effect and update the objective property.
*/
func (o *Campaign) SetObjective(s string) error {
	o.Objective = s
	return nil
}
