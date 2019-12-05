// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package domainnameobject

import (
	"encoding/json"

	"github.com/freetaxii/libstix2/objects/baseobject"
)

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX DomainNameObject object and return it as a
pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *DomainNameObject {
	var obj DomainNameObject
	obj.InitObject("domain-name")
	return &obj
}

// ----------------------------------------------------------------------
// Public Functions - JSON Decoder
// ----------------------------------------------------------------------

/*
Decode - This function will take in some JSON data encoded as a slice of bytes
and decode it into an actual struct. It will return the object as a pointer
along with any errors found.
*/
func Decode(data []byte) (*DomainNameObject, error) {
	var o DomainNameObject
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

// ----------------------------------------------------------------------
// Public Methods - JSON Encoders
// ----------------------------------------------------------------------

/*
Encode - This method is a simple wrapper to JSON encode the object.
*/
func (o *DomainNameObject) Encode() ([]byte, error) {
	return baseobject.Encode(o)
}

/*
EncodeToString - This method is a simple wrapper to JSON encode the object.
*/
func (o *DomainNameObject) EncodeToString() (string, error) {
	return baseobject.EncodeToString(o)
}
