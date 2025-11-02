// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package ipv6addr

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
func Decode(data []byte) (*IPv6Addr, error) {
	var o IPv6Addr
	err := json.Unmarshal(data, &o)
	if err != nil {
		return nil, err
	}
	return &o, nil
}

/*
UnmarshalJSON - This method will over write the default UnmarshalJSON method
to enable custom property support. This is needed because there can be
custom properties that are not defined in the specification. This is handled
through the use of a RawMessage field called Custom in the CommonObjectProperties.
*/
func (o *IPv6Addr) UnmarshalJSON(b []byte) error {
	type alias IPv6Addr
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
