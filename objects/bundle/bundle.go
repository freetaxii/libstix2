// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package bundle

import (
	"encoding/json"

	"github.com/freetaxii/libstix2/objects"
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
	objects.CommonBundleProperties
	Objects []objects.STIXObject `json:"objects,omitempty"`
}

/*
bundleRawDecode - This type is used for decoding a STIX bundle since the
Objects property needs special handling.
*/
type bundleRawDecode struct {
	objects.CommonBundleProperties
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
//
// Public Methods - Bundle
//
// ----------------------------------------------------------------------

/*
AddObject - This method will take in an object as an interface and add it to
the list of objects in the bundle.
*/
func (o *Bundle) AddObject(i objects.STIXObject) error {
	o.Objects = append(o.Objects, i)
	return nil
}
