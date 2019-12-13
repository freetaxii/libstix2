// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package attackpattern

import (
	"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
// Public Functions - JSON Decoder
// ----------------------------------------------------------------------

/*
Decode - This function is a simple wrapper for decoding JSON data. It will
decode a slice of bytes into an actual struct and return a pointer to that
object along with any errors.
*/
func Decode(data []byte) (*AttackPattern, error) {
	var o AttackPattern
	err := objects.Decode(data, &o)
	if err != nil {
		return nil, err
	}
	return &o, nil
}

// ----------------------------------------------------------------------
// Public Functions / Methods - JSON Encoders
// ----------------------------------------------------------------------

/*
Encode - This function is a simple wrapper for encoding an object into JSON
*/
func Encode(o *AttackPattern) ([]byte, error) {
	// Use this object's method to make use of preprocessing rules
	return o.Encode()
}

/*
EncodeToString - This function is a simple wrapper for encoding an object into JSON
*/
func EncodeToString(o *AttackPattern) (string, error) {
	// Use this object's method to make use of preprocessing rules
	return o.EncodeToString()
}

/*
Encode - This method is a simple wrapper for encoding an object into JSON
*/
func (o *AttackPattern) Encode() ([]byte, error) {
	// Any needed preprocessing would be done here
	return objects.Encode(o)
}

/*
EncodeToString - This method is a simple wrapper for encoding an object into JSON
*/
func (o *AttackPattern) EncodeToString() (string, error) {
	// Any needed preprocessing would be done here
	return objects.EncodeToString(o)
}
