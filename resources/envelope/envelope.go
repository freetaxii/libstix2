// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package envelope

import (
	"encoding/json"
	"io"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
Envelope - This type implements the TAXII 2 Envelope Resource and defines
all of the properties and methods needed to create and work with the TAXII
Envelope Resource.

The following information comes directly from the TAXII 2 specification documents.

The envelope is a simple wrapper for STIX 2 content. When returning STIX 2
content in a TAXII response the HTTP root object payload MUST be an envelope.
This specification does not define any other form of content wrapper for objects
outside of STIX content.

For example:
*/
type Envelope struct {
	More    bool          `json:"more,omitempty"`
	Objects []interface{} `json:"objects,omitempty"`
}

type EnvelopeRawDecode struct {
	More    bool              `json:"more,omitempty"`
	Objects []json.RawMessage `json:"objects,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new TAXII Envelope object and return it as a
pointer.
*/
func New() *Envelope {
	var obj Envelope
	return &obj
}

// ----------------------------------------------------------------------
// Public Methods - Envelope - Core Functionality
// ----------------------------------------------------------------------

/*
DecodeRaw - This function will decode the outer layer of an envelope and stop
processing when it gets to the objects. It will leave the objects as a slice of
json.RawMessage objects. This way, later on, we can decode each one individually
*/
func DecodeRaw(r io.Reader) (*EnvelopeRawDecode, error) {
	var b EnvelopeRawDecode
	err := json.NewDecoder(r).Decode(&b)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

/*
Encode - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *Envelope) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
EncodeToString - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *Envelope) EncodeToString() (string, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

/*
EncodeToString - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *EnvelopeRawDecode) EncodeToString() (string, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ----------------------------------------------------------------------
// Public Methods - Envelope
// ----------------------------------------------------------------------

/*
AddObject - This method will take in an object as an interface and add it to
the list of objects in the envelope
*/
func (r *Envelope) AddObject(o interface{}) error {
	r.Objects = append(r.Objects, o)
	return nil
}

/*
GetMore - This method will return the more property
*/
func (r *Envelope) GetMore() bool {
	return r.More
}

/*
SetMore - This method will set the more property to true
*/
func (r *Envelope) SetMore() error {
	r.More = true
	return nil
}
