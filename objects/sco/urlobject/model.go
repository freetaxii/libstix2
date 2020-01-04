// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package urlobject

import (
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/* URLObject - This type implements the STIX 2 URL SCO and defines all of the
properties and methods needed to create and work with this object. All of the
methods not defined local to this type are inherited from the individual
properties. */
type URLObject struct {
	objects.CommonObjectProperties
	properties.ValueProperty
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/* New - This function will create a new STIX URL SCO and return it as a
pointer. It will also initialize the object by setting all of the basic
properties. */
func New() *URLObject {
	var obj URLObject
	obj.InitSCO("url")
	return &obj
}
