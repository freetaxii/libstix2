// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sighting

import (
	"github.com/freetaxii/libstix2/objects/common"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type SightingType struct {
	common.CommonObjectPropertiesType
	common.FirstLastSeenPropertiesType
	Count              int      `json:"count,omitempty"`
	Sighting_of_ref    string   `json:"sighting_of_ref,omitempty"`
	Observed_data_refs []string `json:"observed_data_refs,omitempty"`
	Where_sighted_refs []string `json:"where_sighted_refs,omitempty"`
	Summary            bool     `json:"summary,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() SightingType {
	var obj SightingType
	obj.InitNewObject("sighting")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - SightingType
// ----------------------------------------------------------------------

// SetCount takes in one parameter
// param: i - an integer that represents the number of sightings
func (this *SightingType) SetCount(i int) {
	this.Count = i
}

// SetSightingOfRef takes in one parameter
// param: s - a string value that represents a STIX identifier of the object that was sighted
func (this *SightingType) SetSightingOfRef(s string) {
	this.Sighting_of_ref = s
}

// AddObservedDataRef takes in one parameter
// param: s - a string value that represents a STIX identifier of the STIX Observed Data object that identifies what was sighted
func (this *SightingType) AddObservedDataRef(s string) {
	if this.Observed_data_refs == nil {
		a := make([]string, 0)
		this.Observed_data_refs = a
	}
	this.Observed_data_refs = append(this.Observed_data_refs, s)
}

// AddObservedDataRef takes in one parameter
// param: s - a string value that represents a STIX identifier of the STIX Identity object that identifies where this was sighted
func (this *SightingType) AddWhereSightedRef(s string) {
	if this.Where_sighted_refs == nil {
		a := make([]string, 0)
		this.Where_sighted_refs = a
	}
	this.Where_sighted_refs = append(this.Where_sighted_refs, s)
}

// SetSummary set the boolean value of the summary to true
// param: i - an integer that represents the number of sightings
func (this *SightingType) SetSummary() {
	this.Summary = true
}
