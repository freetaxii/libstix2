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
ObjectIDProperty - A property used by one or more STIX objects that captures
the unique object ID. This is not included in the JSON serialization, but is
used for writing to the database.
*/
type ObjectIDProperty struct {
	ObjectID int `json:"-"`
}

/*
RawDataProperty - A property used to store the raw bytes of the JSON object.
*/
type RawDataProperty struct {
	Raw []byte `json:"-"`
}

// ----------------------------------------------------------------------
// Public Methods - ObjectIDProperty
// ----------------------------------------------------------------------

/*
SetObjectID - This method takes in a int64 representing an object ID and
updates the Version property.
*/
func (p *ObjectIDProperty) SetObjectID(i int) error {
	p.ObjectID = i
	return nil
}

/*
GetObjectID - This method returns the object ID value as a int64.
*/
func (p *ObjectIDProperty) GetObjectID() int {
	return p.ObjectID
}

// ----------------------------------------------------------------------
// Public Methods - IdPropertyType
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
