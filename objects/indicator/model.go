// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package indicator

import (
	"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
Indicator - This type implements the STIX 2 Indicator SDO and defines all of
the properties and methods needed to create and work with this object. All of
the methods not defined local to this type are inherited from the individual
properties.
*/
type Indicator struct {
	objects.CommonObjectProperties
	objects.NameProperty
	objects.DescriptionProperty
	IndicatorTypes []string `json:"indicator_types,omitempty" bson:"indicator_types,omitempty"`
	Pattern        string   `json:"pattern,omitempty" bson:"pattern,omitempty"`
	PatternType    string   `json:"pattern_type,omitempty" bson:"pattern_type,omitempty"`
	PatternVersion string   `json:"pattern_version,omitempty" bson:"pattern_version,omitempty"`
	ValidFrom      string   `json:"valid_from,omitempty" bson:"valid_from,omitempty"`
	ValidUntil     string   `json:"valid_until,omitempty" bson:"valid_until,omitempty"`
	objects.KillChainPhasesProperty
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *Indicator) GetPropertyList() []string {
	return []string{"name", "description", "indicator_types", "pattern", "pattern_type", "pattern_version", "valid_from", "valid_until", "kill_chain_phases"}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Indicator object and return it as
a pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *Indicator {
	var obj Indicator
	obj.InitSDO("indicator")
	return &obj
}
