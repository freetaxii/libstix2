// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package urlobject

import (
	"encoding/json"
)

// ----------------------------------------------------------------------
// Public Functions - JSON Decoder
// ----------------------------------------------------------------------

/* Decode - This function is a simple wrapper for decoding JSON data. It will
decode a slice of bytes into an actual struct and return a pointer to that
object along with any errors. */
func Decode(data []byte) (*URLObject, error) {
	var o URLObject

	err := json.Unmarshal(data, &o)
	if err != nil {
		return nil, err
	}

	o.SetRawData(data)

	return &o, nil
}

// ----------------------------------------------------------------------
// Public Methods JSON Encoders
// The encoding is done here at the individual object level instead of at
// the STIX Object level so that individual pre/post processing rules can
// be applied. Since some of the STIX Objects do not follow a universal
// model, we need to cleanup some things that were inherited but not valid
// for the object.
// ----------------------------------------------------------------------

/* Encode - This method is a simple wrapper for encoding an object into JSON */
func (o *URLObject) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return nil, err
	}

	// Any needed preprocessing would be done here
	return data, nil
}

/* EncodeToString - This method is a simple wrapper for encoding an object into
JSON */
func (o *URLObject) EncodeToString() (string, error) {
	data, err := o.Encode()
	if err != nil {
		return "", err
	}
	return string(data), nil
}
