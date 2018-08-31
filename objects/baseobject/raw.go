// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package baseobject

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
RawProperty - A property used to store the raw bytes of the JSON object.
*/
type RawProperty struct {
	Raw []byte `json:"-"`
}

// ----------------------------------------------------------------------
// Public Methods - RawProperty
// ----------------------------------------------------------------------

/*
SetRaw - This method takes in a slice of bytes representing a full JSON object
and updates the raw property for the object.
*/
func (o *RawProperty) SetRawData(data []byte) error {
	o.Raw = data
	return nil
}

/*
GetRaw - This method will return the raw bytes for a given STIX object.
*/
func (o *RawProperty) GetRawData() []byte {
	return o.Raw
}
