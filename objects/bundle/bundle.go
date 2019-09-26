// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package bundle

import (
	"encoding/json"
	"io"

	"github.com/freetaxii/libstix2/objects/baseobject"
)

// ----------------------------------------------------------------------
//
// Define Object Type
//
// ----------------------------------------------------------------------

/*
Bundle - This type implements the STIX 2 Bundle and defines
all of the properties methods needed to create and work with the STIX Bundle.
All of the methods not defined local to this type are inherited from
the individual properties.
*/
type Bundle struct {
	baseobject.BundleBaseProperties
	Objects []interface{} `json:"objects,omitempty"`
}

/*
BundleRawDecode - This type is used for decoding a STIX bundle since the
Objects property needs special handling.
*/
type BundleRawDecode struct {
	baseobject.BundleBaseProperties
	Objects []json.RawMessage `json:"objects,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Bundle object and return it as
a pointer. This function can not use the InitNewObject() function as a Bundle
does not have all of the fields that are common to a standard object.
*/
func New() *Bundle {
	var obj Bundle
	obj.SetObjectType("bundle")
	obj.SetNewID("bundle")
	return &obj
}

// ----------------------------------------------------------------------
// Public Methods - Bundle - Core Functionality
// ----------------------------------------------------------------------

/*
DecodeRaw - This function will decode the outer layer of a bundle and stop
processing when it gets to the objects. It will leave the objects as a slice of
json.RawMessage objects. This way, later on, we can decode each one individually
*/
func DecodeRaw(r io.Reader) (*BundleRawDecode, error) {
	var b BundleRawDecode
	err := json.NewDecoder(r).Decode(&b)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

/*
Encode - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *Bundle) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
EncodeToString - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *Bundle) EncodeToString() (string, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

/*
EncodeToString - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *BundleRawDecode) EncodeToString() (string, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ----------------------------------------------------------------------
//
// Public Methods - Bundle
//
// ----------------------------------------------------------------------

/*
AddObject - This method will take in an object as an interface and add it to
the list of objects in the bundle.
*/
func (o *Bundle) AddObject(i interface{}) error {
	o.Objects = append(o.Objects, i)
	return nil
}
