// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package grouping

import (
	"github.com/wxj95/libstix2/objects"
	"github.com/wxj95/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
Grouping - This type implements the STIX 2 Grouping SDO and defines all of
the properties and methods needed to create and work with this object. All of
the methods not defined local to this type are inherited from the individual
properties.
*/
type Grouping struct {
	objects.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	Context string `json:"context,omitempty"`
	properties.ObjectRefsProperty
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *Grouping) GetPropertyList() []string {
	return []string{"name", "description", "context", "object_refs"}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Grouping object and return
it as a pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *Grouping {
	var obj Grouping
	obj.InitSDO("grouping")
	return &obj
}
