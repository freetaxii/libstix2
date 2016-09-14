// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sighting

import (
	"errors"
	"github.com/freetaxii/libstix2/messages/defs"
	"github.com/freetaxii/libstix2/messages/stix"
	"time"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type SightingType struct {
	stix.CommonProperties
	First_seen           string   `json:"first_seen,omitempty"`
	First_seen_precision string   `json:"first_seen_precision,omitempty"`
	Last_seen            string   `json:"last_seen,omitempty"`
	Last_seen_precision  string   `json:"last_seen_precision,omitempty"`
	Count                int      `json:"count,omitempty"`
	Sighting_of_ref      string   `json:"sighting_of_ref,omitempty"`
	Observed_data_refs   []string `json:"observed_data_ref,omitempty"`
	Where_sighted_refs   []string `json:"where_sighted_ref,omitempty"`
	Summary              bool     `json:"summary,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() SightingType {
	var obj SightingType
	obj.MessageType = "sighting"
	obj.Id = stix.NewId("sighting")
	obj.Created = stix.GetCurrentTime().UTC().Format(defs.TIME_RFC_3339)
	obj.Modified = obj.Created
	obj.Version = 1
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - Common Properties
// ----------------------------------------------------------------------

func (this *SightingType) SetCreatedBy(s string) {
	this.Created_by_ref = s
}

func (this *SightingType) SetModified(d time.Time) {
	this.Modified = d.UTC().Format(defs.TIME_RFC_3339)
}

func (this *SightingType) SetVersion(i int) error {
	if i < defs.MIN_VERSION_SIZE {
		return errors.New("No change made, new version is smaller than min size")
	}

	if i > defs.MAX_VERSION_SIZE {
		return errors.New("No change made, new version is larger than max size")
	}

	if i <= this.Version {
		return errors.New("No change made, new version is not larger than original")
	}

	this.Version = i
	return nil
}

func (this *SightingType) SetRevoked() {
	this.Revoked = true
}

func (this *SightingType) GetId() string {
	return this.Id
}

// ----------------------------------------------------------------------
// Public Methods - SightingType
// ----------------------------------------------------------------------

// TODO add precision functions

func (this *SightingType) SetFirstSeen(d time.Time) {
	this.First_seen = d.UTC().Format(defs.TIME_RFC_3339)
}

// This function will allow you to assign the time as a string instead of using
// a time.Time object
func (this *SightingType) SetFirstSeenText(s string) {
	this.First_seen = s
}

func (this *SightingType) SetLastSeen(d time.Time) {
	this.Last_seen = d.UTC().Format(defs.TIME_RFC_3339)
}

// This function will allow you to assign the time as a string instead of using
// a time.Time object
func (this *SightingType) SetLastSeenText(s string) {
	this.Last_seen = s
}

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
