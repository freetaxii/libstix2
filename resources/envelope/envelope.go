// Copyright 2015-2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package envelope

import (
	"github.com/freetaxii/libstix2/objects/bundle"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
 */
type Envelope struct {
	More bool           `json:"more,omitempty"`
	Data *bundle.Bundle `json:"data,omitempty"`
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
AddBundle - This method takes in an object that represents a STIX Bundle and
adds it to the envelope. This method would be used if the Bundle was created
separately and it just needs to be added in whole to the envelope.
*/
func (r *Envelope) AddBundle(o *bundle.Bundle) error {
	r.Data = o
	return nil
}

/*
NewBundle - This method is used to create a STIX Bundle and automatically
add it to the envelope. It returns a bundle.Bundle which is a pointer to the
actual STIX Bundle that was created in the envelope.
*/
func (r *Envelope) NewBundle() (*bundle.Bundle, error) {
	o := bundle.New()
	r.Data = o
	return r.Data, nil
}

/*
SetMore - This method will set the more property to true
*/
func (r *Envelope) SetMore() error {
	r.More = true
	return nil
}
