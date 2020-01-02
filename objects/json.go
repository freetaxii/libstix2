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

/* DecodeType - This function will take in a slice of bytes representing a
random STIX object encoded as JSON and return the STIX object type as a string.
This is called from the Bundle Decode() to determine which type of STIX object
the data represents, so that the data can be dispatched to the right object
decoder. */
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
object along with any errors. This is called from the Bundle Decode() if the
object type can not be determined. So for custom objects, it will at least
decode any of the common object properties that might be found.*/
func Decode(data []byte) (*CommonObjectProperties, error) {
	var o CommonObjectProperties

	err := json.Unmarshal(data, o)
	if err != nil {
		return nil, err
	}

	o.SetRawData(data)

	return &o, nil
}
