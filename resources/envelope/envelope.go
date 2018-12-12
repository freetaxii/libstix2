// Copyright 2015-2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package envelope

import (
	"encoding/json"
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
