// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package indicator

import "encoding/json"

// ----------------------------------------------------------------------
// Public Functions - JSON Decoder
// ----------------------------------------------------------------------

/* Decode - This function is a simple wrapper for decoding JSON data. It will
decode a slice of bytes into an actual struct and return a pointer to that
object along with any errors. */
func Decode(data []byte) (*Indicator, error) {
	var o Indicator

	err := json.Unmarshal(data, o)
	if err != nil {
		return nil, err
	}

	if valid, err := o.Valid(); valid != true {
		return nil, err
	}

	o.SetRawData(data)

	return &o, nil
}

// ----------------------------------------------------------------------
// Public Methods JSON Encoders
// ----------------------------------------------------------------------

/* Encode - This method is a simple wrapper for encoding an object into JSON */
func (o *Indicator) Encode() ([]byte, error) {
	return Encode(o)
}

/* EncodeToString - This method is a simple wrapper for encoding an object into
JSON */
func (o *Indicator) EncodeToString() (string, error) {
	return EncodeToString(o)
}

// ----------------------------------------------------------------------
// Public Functions JSON Encoders
// ----------------------------------------------------------------------

/* Encode - This function is a simple wrapper for encoding an object into JSON
 */
func Encode(o *Indicator) ([]byte, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return nil, err
	}

	// Any needed preprocessing would be done here
	return data, nil
}

/* EncodeToString - This function is a simple wrapper for encoding an object
into JSON */
func EncodeToString(o *Indicator) (string, error) {
	data, err := Encode(o)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
