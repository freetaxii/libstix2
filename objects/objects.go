// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

/*
Package objects implements the STIX 2 object model.
*/
package objects

import (
	"encoding/json"
	"fmt"

	"github.com/freetaxii/libstix2/objects/attackpattern"
	"github.com/freetaxii/libstix2/objects/baseobject"
	"github.com/freetaxii/libstix2/objects/campaign"
	"github.com/freetaxii/libstix2/objects/courseofaction"
	"github.com/freetaxii/libstix2/objects/identity"
	"github.com/freetaxii/libstix2/objects/indicator"
	"github.com/freetaxii/libstix2/objects/infrastructure"
	"github.com/freetaxii/libstix2/objects/intrusionset"
	"github.com/freetaxii/libstix2/objects/malware"
	"github.com/freetaxii/libstix2/objects/observeddata"
	"github.com/freetaxii/libstix2/objects/relationship"
	"github.com/freetaxii/libstix2/objects/report"
	"github.com/freetaxii/libstix2/objects/sighting"
	"github.com/freetaxii/libstix2/objects/threatactor"
	"github.com/freetaxii/libstix2/objects/tool"
	"github.com/freetaxii/libstix2/objects/vulnerability"
)

type STIXObject interface {
	GetObjectType() string
	GetID() string
	GetModified() string
	GetCommonProperties() *baseobject.CommonObjectProperties
}

/*
DecodeType - This function will take in a slice of bytes representing a
random STIX object encoded as JSON and return the STIX object type as a string.
*/
func DecodeType(data []byte) (string, error) {
	var o baseobject.TypeProperty
	err := json.Unmarshal(data, &o)
	if err != nil {
		return "", err
	}

	if valid, err := o.Valid(); valid != true {
		return "", fmt.Errorf("invalid STIX object: %s", err)
	}

	return o.ObjectType, nil
}

/*
Decode - This function will take in a slice of bytes representing a
random STIX object encoded as JSON, decode it to the appropriate STIX object
struct, and return
 - object itself as an interface
 - its STIX ID
 - its Modified time stamp
 - and any possible errors
*/
func Decode(data []byte) (STIXObject, error) {
	var err error

	// Make a first pass to decode just the object type value. Once we have this
	// value we can easily make a second pass and decode the rest of the object.
	stixtype, err := DecodeType(data)
	if err != nil {
		return nil, err
	}

	switch stixtype {
	case "attack-pattern":
		return attackpattern.Decode(data)
	case "campaign":
		return campaign.Decode(data)
	case "course-of-action":
		return courseofaction.Decode(data)
	case "identity":
		return identity.Decode(data)
	case "indicator":
		return indicator.Decode(data)
	case "infrastructure":
		return infrastructure.Decode(data)
	case "intrusion-set":
		return intrusionset.Decode(data)
	case "malware":
		return malware.Decode(data)
	case "observed-data":
		return observeddata.Decode(data)
	case "relationship":
		return relationship.Decode(data)
	case "report":
		return report.Decode(data)
	case "sighting":
		return sighting.Decode(data)
	case "threat-actor":
		return threatactor.Decode(data)
	case "tool":
		return tool.Decode(data)
	case "vulnerability":
		return vulnerability.Decode(data)
	}
	//TODO add a default that just stores the custom object
	// probably just decode this to common properties and leave everything else
	// as RAW data.  This is also how I could add support for all types before
	// I get around to coding support for the actual writing to the tables.
	return nil, nil
}
