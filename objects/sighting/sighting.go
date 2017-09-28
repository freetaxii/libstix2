// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sighting

import (
	"github.com/freetaxii/libstix2/objects/common/properties"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
SightingType defines all of the properties associated with the STIX Sighting
SRO. All of the methods not defined local to this type are inherited from the
individual properties.
*/
type SightingType struct {
	properties.CommonObjectPropertiesType
	properties.FirstSeenPropertyType
	properties.LastSeenPropertyType
	Count            int      `json:"count,omitempty"`
	SightingOfRef    string   `json:"sighting_of_ref,omitempty"`
	ObservedDataRefs []string `json:"observed_data_refs,omitempty"`
	WhereSightedRefs []string `json:"where_sighted_refs,omitempty"`
	Summary          bool     `json:"summary,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// New - This function will create a new sighting object.
func New() SightingType {
	var obj SightingType
	obj.InitNewObject("sighting")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - SightingType
// ----------------------------------------------------------------------

// SetCount - This method takes in an integer that represents the number of
// sightings and upates the count properties.
func (ezt *SightingType) SetCount(i int) {
	ezt.Count = i
}

// SetSightingOfRef - This method takes in a string value that represents a STIX
// identifier of the object that was sighted and updates the sighting of ref
// property.
func (ezt *SightingType) SetSightingOfRef(s string) {
	ezt.SightingOfRef = s
}

// AddObservedDataRef - This method takes in a string value that represents a
// STIX identifier of the STIX Observed Data object that identifies what was
// sighted and adds it to the observed data refs property.
func (ezt *SightingType) AddObservedDataRef(s string) {
	ezt.ObservedDataRefs = append(ezt.ObservedDataRefs, s)
}

// AddWhereSightedRef - This method takes in a string value that represents a
// STIX identifier of the STIX Identity object that identifies where this was
// sighted (location, sector, etc) and adds it to the where sighted ref property.
func (ezt *SightingType) AddWhereSightedRef(s string) {
	ezt.WhereSightedRefs = append(ezt.WhereSightedRefs, s)
}

// SetSummary - This method set the boolean value of the summary to true.
func (ezt *SightingType) SetSummary() {
	ezt.Summary = true
}
