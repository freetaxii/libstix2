// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

import (
	"encoding/json"
	"fmt"

	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Public Functions - JSON Decoders
// ----------------------------------------------------------------------

/*
DecodeType - This function will take in a slice of bytes representing a
random STIX object encoded as JSON and return the STIX object type as a string.
*/
func DecodeType(data []byte) (string, error) {
	var o properties.TypeProperty
	err := json.Unmarshal(data, &o)
	if err != nil {
		return "", err
	}

	// This will call the Valid function on the TypeProperty type
	if valid, err := o.Valid(); valid != true {
		return "", fmt.Errorf("invalid STIX object: %s", err)
	}

	return o.ObjectType, nil
}

/*
Decode - This function takes in some JSON data encoded as a slice of bytes
and a STIXObject interface. It will decode the JSON data into an actual struct
that is provided as a STIXObject interface. This function is used by all of the
STIX objects in this library. The function will return any errors found.
*/
func Decode(data []byte, o STIXObject) error {
	err := json.Unmarshal(data, o)
	if err != nil {
		return err
	}

	if valid, err := o.Valid(); valid != true {
		return err
	}

	o.SetRawData(data)

	return nil
}

// ----------------------------------------------------------------------
// Public Functions / Methods - JSON Encoders
// ----------------------------------------------------------------------

/*
Encode - This function is a simple wrapper for encoding an object in to JSON
*/
func Encode(o STIXObject) ([]byte, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
EncodeToString - This function is a simple wrapper for encoding an object in
to JSON
*/
func EncodeToString(o STIXObject) (string, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
