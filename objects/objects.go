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

	"github.com/freetaxii/libstix2/objects/baseobject"
	"github.com/freetaxii/libstix2/objects/indicator"
)

type STIXObject interface {
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
func Decode(data []byte) (interface{}, string, string, error) {
	var err error

	// Make a first pass to decode just the object type value. Once we have this
	// value we can easily make a second pass and decode the rest of the object.
	stixtype, err := DecodeType(data)
	if err != nil {
		return nil, "", "", err
	}

	switch stixtype {
	// case "campaign":
	// 	return campaign.Decode(data)
	case "indicator":
		return indicator.Decode(data)
		// case "infrastructure":
		// 	var o infrastructure.Infrastructure
		// 	err = json.Unmarshal(data, &o)
		// 	return o, o.ID, nil
		// case "malware":
		// 	var o malware.Malware
		// 	err = json.Unmarshal(data, &o)
		// 	return o, o.ID, nil
		// case "observed-data":
		// 	var o observeddata.ObservedData
		// 	err = json.Unmarshal(data, &o)
		// 	return o, o.ID, nil
		// case "relationship":
		// 	var o relationship.Relationship
		// 	err = json.Unmarshal(data, &o)
		// 	return o, o.ID, nil
		// case "sighting":
		// 	var o sighting.Sighting
		// 	err = json.Unmarshal(data, &o)
		// 	return o, o.ID, nil
	}
	//TODO add a default that just stores the custom object
	return nil, "", "", nil
}
