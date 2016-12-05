// Copyright 2016 Bret Jordan, All rights reserved.
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
	common.CommonPropertiesType
	common.FirstLastSeenType
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
	obj.MessageType = "sighting"
	obj.Id = obj.NewId("sighting")
	obj.Created = obj.GetCurrentTime()
	obj.Modified = obj.Created
	obj.Version = 1
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - SightingType
// ----------------------------------------------------------------------

func (this *SightingType) SetCount(i int) {
	this.Count = i
}

func (this *SightingType) SetSightingOfRef(s string) {
	this.Sighting_of_ref = s
}

func (this *SightingType) AddObservedDataRef(s string) {
	if this.Observed_data_refs == nil {
		a := make([]string, 0)
		this.Observed_data_refs = a
	}
	this.Observed_data_refs = append(this.Observed_data_refs, s)
}

func (this *SightingType) AddWhereSightedRef(s string) {
	if this.Where_sighted_refs == nil {
		a := make([]string, 0)
		this.Where_sighted_refs = a
	}
	this.Where_sighted_refs = append(this.Where_sighted_refs, s)
}

func (this *SightingType) SetSummary() {
	this.Summary = true
}
