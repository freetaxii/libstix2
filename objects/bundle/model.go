// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package bundle

import (
	"encoding/json"

	"github.com/wxj95/libstix2/objects"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
Bundle - This type implements the STIX 2 Bundle SDO and defines all of the
properties and methods needed to create and work with this object. All of the
methods not defined local to this type are inherited from the individual
properties.
*/
type Bundle struct {
	objects.CommonObjectProperties
	Objects []objects.STIXObject `json:"objects,omitempty"`
}

/*
bundleRawDecode - This type is used for decoding a STIX bundle since the
Objects property needs special handling.
*/
type bundleRawDecode struct {
	objects.CommonObjectProperties
	Objects []json.RawMessage `json:"objects,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Bundle object and return it as a
pointer. This function can not use the InitNewObject() function as a Bundle does
not have all of the fields that are common to a standard object.
*/
func New() *Bundle {
	var obj Bundle
	obj.InitBundle()
	return &obj
}
