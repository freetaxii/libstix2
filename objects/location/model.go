// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package location

import (
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
Location - This type implements the STIX 2 Location SDO and
defines all of the properties and methods needed to create and work with this
object. All of the methods not defined local to this type are inherited from the
individual properties.
*/
type Location struct {
	objects.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	Latitude           float64 `json:"latitude,omitempty" bson:"latitude,omitempty"`
	Longitude          float64 `json:"longitude,omitempty" bson:"longitude,omitempty"`
	Precision          float64 `json:"precision,omitempty" bson:"precision,omitempty"`
	Region             string  `json:"region,omitempty" bson:"region,omitempty"`
	Country            string  `json:"country,omitempty" bson:"country,omitempty"`
	AdministrativeArea string  `json:"administrative_area,omitempty" bson:"administrative_area,omitempty"`
	City               string  `json:"city,omitempty" bson:"city,omitempty"`
	StreetAddress      string  `json:"street_address,omitempty" bson:"street_address,omitempty"`
	PostalCode         string  `json:"postal_code,omitempty" bson:"postal_code,omitempty"`
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *Location) GetPropertyList() []string {
	return []string{"name", "description", "latitude", "longitude", "precision", "region", "country", "administrative_area", "city", "street_address", "postal_code"}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Location object and return
it as a pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *Location {
	var obj Location
	obj.InitSDO("location")
	return &obj
}
