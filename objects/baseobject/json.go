// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package baseobject

import "encoding/json"

// ----------------------------------------------------------------------
// Public Functions - JSON Encoders
// ----------------------------------------------------------------------

/*
Encode - This method is a simple wrapper for encoding the object in to JSON.
*/
func Encode(o interface{}) ([]byte, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
EncodeToString - This method is a simple wrapper for encoding the object in to
JSON.
*/
func EncodeToString(o interface{}) (string, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

/*
Decode - This function will decode some JSON data encoded as a slice of bytes
into an actual struct. It will return the object as a pointer and any errors found.
This function will handle Custom STIX objects by decoding all of the common properties.
*/
func Decode(data []byte) (*CommonObjectProperties, error) {
	var o CommonObjectProperties
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
