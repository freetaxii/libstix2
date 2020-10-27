// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sighting

import (
	"github.com/wxj95/libstix2/objects"
	"github.com/wxj95/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Type
// ----------------------------------------------------------------------

/*
Sighting - This type implements the STIX 2 Sighting SRO and defines all of
the properties and methods needed to create and work with this object. All of
the methods not defined local to this type are inherited from the individual
properties.
*/
type Sighting struct {
	objects.CommonObjectProperties
	properties.DescriptionProperty
	properties.SeenProperties
	Count            int      `json:"count,omitempty"`
	SightingOfRef    string   `json:"sighting_of_ref,omitempty"`
	ObservedDataRefs []string `json:"observed_data_refs,omitempty"`
	WhereSightedRefs []string `json:"where_sighted_refs,omitempty"`
	Summary          bool     `json:"summary,omitempty"`
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *Sighting) GetPropertyList() []string {
	return []string{"description", "first_seen", "last_seen", "count", "sighting_of_ref", "observed_data_refs", "where_sighted_refs", "summary"}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Sighting object and return
it as a pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *Sighting {
	var obj Sighting
	obj.InitSRO("sighting")
	return &obj
}
