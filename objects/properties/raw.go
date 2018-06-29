// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
//
// Types
//
// ----------------------------------------------------------------------

/*
RawDataProperty - A property used to store the raw bytes of the JSON object.
*/
type RawDataProperty struct {
	Raw []byte `json:"-"`
}

// ----------------------------------------------------------------------
//
// Public Methods - IdPropertyType
//
// ----------------------------------------------------------------------

/*
SetRaw - This method takes in a slice of bytes representing a full JSON object
and updates the raw property for the object.
*/
func (p *RawDataProperty) SetRawData(data []byte) error {
	p.Raw = data
	return nil
}

/*
GetRaw - This method will return the raw bytes for a given STIX object.
*/
func (p *RawDataProperty) GetRawData() []byte {
	return p.Raw
}
