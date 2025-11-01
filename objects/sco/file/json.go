// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package file

import (
	"encoding/json"
)

// ----------------------------------------------------------------------
// Public Methods - JSON Decoder
// ----------------------------------------------------------------------

/*
Decode - This function is a simple wrapper for decoding JSON data. It will
decode a slice of bytes into an actual struct and return a pointer to that
object along with any errors.
*/
func Decode(data []byte) (*File, error) {
	var o File
	err := json.Unmarshal(data, &o)
	if err != nil {
		return nil, err
	}
	return &o, nil
}

/*
UnmarshalJSON - This method will over write the default UnmarshalJSON method
to enable custom property support.
*/
func (o *File) UnmarshalJSON(b []byte) error {
	type alias File
	temp := &struct {
		*alias
	}{
		alias: (*alias)(o),
	}
	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}
	return o.FindCustomProperties(b, o.GetPropertyList())
}
