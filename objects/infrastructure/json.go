// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package infrastructure

import (
	"encoding/json"

	"github.com/freetaxii/libstix2/defs"
)

// ----------------------------------------------------------------------
// Public Functions - JSON Decoder
// ----------------------------------------------------------------------

/* Decode - This function is a simple wrapper for decoding JSON data. It will
decode a slice of bytes into an actual struct and return a pointer to that
object along with any errors. */
func Decode(data []byte) (*Infrastructure, error) {
	var o Infrastructure

	if err := json.Unmarshal(data, &o); err != nil {
		return nil, err
	}

	return &o, nil
}

/* UnmarshalJSON - This method will over write the default UnmarshalJSON method
to enable custom properties that this library does not know about. It will store
them as map where the value of each key is a byte arrays. This way a tool that
does know how to deal with them can then further process them after this is
done. This will also allow the storage of the raw JSON data. */
func (o *Infrastructure) UnmarshalJSON(b []byte) error {

	type alias Infrastructure
	temp := &struct {
		*alias
	}{
		alias: (*alias)(o),
	}
	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}

	// This will create a map of all of the custom properties and store them in a
	// property called o.Custom
	if err := o.FindCustomProperties(b, o.GetPropertyList()); err != nil {
		return err
	}

	// This will store a complete copy of the original JSON in a byte array called
	// o.Raw. This could be useful if you need to digitally sign the JSON or do
	// verification on what was actually received.
	if defs.KEEP_RAW_DATA == true {
		o.SetRawData(b)
	}

	return nil
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
func (o *Infrastructure) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return nil, err
	}

	// Any needed preprocessing would be done here
	return data, nil
}

/* EncodeToString - This method is a simple wrapper for encoding an object into
JSON */
func (o *Infrastructure) EncodeToString() (string, error) {
	data, err := o.Encode()
	if err != nil {
		return "", err
	}
	return string(data), nil
}
