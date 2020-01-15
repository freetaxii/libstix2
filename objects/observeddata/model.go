// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package observeddata

import (
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Type
// ----------------------------------------------------------------------

/*
ObservedData - This type implements the STIX 2 Observed Data SDO and
defines all of the properties and methods needed to create and work with this
object. All of the methods not defined local to this type are inherited from the
individual properties.
*/
type ObservedData struct {
	objects.CommonObjectProperties
	FirstObserved  string `json:"first_observed,omitempty"`
	LastObserved   string `json:"last_observed,omitempty"`
	NumberObserved int    `json:"number_observed,omitempty"`
	Objects        string `json:"objects,omitempty"`
	properties.ObjectRefsProperty
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *ObservedData) GetPropertyList() []string {
	return []string{"first_observed", "last_observed", "number_observed", "objects", "object_refs"}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Observed Data object and return
it as a pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *ObservedData {
	var obj ObservedData
	obj.InitSDO("observed-data")
	return &obj
}
