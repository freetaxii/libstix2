// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/freetaxii/libstix2/objects/properties"
)

func VerifyCommonProperties(o properties.CommonObjectProperties) error {

	if err := o.TypeProperty.Verify(); err != nil {
		return err
	}
	if err := o.IDProperty.Verify(); err != nil {
		return err
	}
	return nil
}

/*
DecodeObjectType - This function will take in a slice of bytes representing a
random STIX object encoded as JSON and return the STIX object type as a string.
*/
func DecodeObjectType(data []byte) (string, error) {
	var o properties.CommonObjectProperties
	err := json.Unmarshal(data, &o)
	if err != nil {
		return "", err
	}

	if o.ObjectType == "" {
		return "", errors.New("Invalid STIX object: object type is missing")
	}

	return o.ObjectType, nil
}

/*
DecodeObject - This function will take in a slice of bytes representing a
random STIX object encoded as JSON, decode it to the appropriate STIX object
struct, and return object itself, its STIX ID, and any possible errors.
*/
func DecodeObject(data []byte) (interface{}, string, error) {
	var err error

	// Make a first pass to decode just the object type value. Once we have this
	// value we can easily make a second pass and decode the rest of the object.
	stixtype, err := DecodeObjectType(data)
	if err != nil {
		return nil, "", err
	}

	switch stixtype {
	case "campaign":
		o, err := DecodeCampaign(data)
		if err != nil {
			return nil, "", err
		}
		return o, o.ID, nil
	case "indicator":
		o, err := DecodeIndicator(data)
		if err != nil {
			return nil, "", err
		}
		return o, o.ID, nil
	case "infrastructure":
		var o Infrastructure
		err = json.Unmarshal(data, &o)
		fmt.Println(o.ID)
		return o, "", nil
	case "malware":
		var o Malware
		err = json.Unmarshal(data, &o)
		fmt.Println(o.ID)
		return o, "", nil
	case "observed-data":
		var o ObservedData
		err = json.Unmarshal(data, &o)
		fmt.Println(o.ID)
		return o, "", nil
	case "relationship":
		var o Relationship
		err = json.Unmarshal(data, &o)
		fmt.Println(o.ID)
		return o, "", nil
	case "sighting":
		var o Sighting
		err = json.Unmarshal(data, &o)
		fmt.Println(o.ID)
		return o, "", nil
	}
	//TODO add a default that just stores the custom object
	return nil, "", nil
}
