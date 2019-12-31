// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

import (
	"encoding/json"
	"errors"

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
	if valid, _, details := o.VerifyExists(); valid != true {
		return "", errors.New(details[0])
	}

	return o.ObjectType, nil
}

/* Decode - This function is a simple wrapper for decoding JSON data. It will
decode a slice of bytes into an actual struct and return a pointer to that
object along with any errors. */
func Decode(data []byte) (*CommonObjectProperties, error) {
	var o CommonObjectProperties

	err := json.Unmarshal(data, o)
	if err != nil {
		return nil, err
	}

	o.SetRawData(data)

	return &o, nil
}

// ----------------------------------------------------------------------
// Public Functions / Methods - JSON Encoders
// ----------------------------------------------------------------------

/*
Encode - This function is a simple wrapper for encoding an object in to JSON
*/
// func Encode(o STIXObject) ([]byte, error) {
// 	data, err := json.MarshalIndent(o, "", "    ")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return data, nil
// }

/*
EncodeToString - This function is a simple wrapper for encoding an object in
to JSON
*/
// func EncodeToString(o STIXObject) (string, error) {
// 	data, err := json.MarshalIndent(o, "", "    ")
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(data), nil
// }
