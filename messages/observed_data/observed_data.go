// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package observed_data

import (
	"errors"
	"github.com/freetaxii/libstix2/messages/defs"
	"github.com/freetaxii/libstix2/messages/stix"
	"time"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type ObservedDataType struct {
	stix.CommonProperties
	First_observed  string `json:"first_observed,omitempty"`
	Last_observed   string `json:"last_observed,omitempty"`
	Number_observed int    `json:"number_observed,omitempty"`
	Cybox           string `json:"cybox,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() ObservedDataType {
	var obj ObservedDataType
	obj.MessageType = "observed-data"
	obj.Id = stix.NewId("observed-data")
	obj.Created = stix.GetCurrentTime().UTC().Format(defs.TIME_RFC_3339)
	obj.Modified = obj.Created
	obj.Version = 1
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - Common Properties
// ----------------------------------------------------------------------

func (this *ObservedDataType) SetCreatedBy(s string) {
	this.Created_by_ref = s
}

func (this *ObservedDataType) SetModified(d time.Time) {
	this.Modified = d.UTC().Format(defs.TIME_RFC_3339)
}

func (this *ObservedDataType) SetVersion(i int) error {
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

func (this *ObservedDataType) SetRevoked() {
	this.Revoked = true
}

func (this *ObservedDataType) AddLabel(value string) {
	if this.Labels == nil {
		a := make([]string, 0)
		this.Labels = a
	}
	this.Labels = append(this.Labels, value)
}

func (this *ObservedDataType) GetId() string {
	return this.Id
}

// ----------------------------------------------------------------------
// Public Methods - ObservedDataType
// ----------------------------------------------------------------------

func (this *ObservedDataType) SetFirstObserved(d time.Time) {
	this.First_observed = d.UTC().Format(defs.TIME_RFC_3339)
}

// This function will allow you to assign the time as a string instead of using
// a time.Time object
func (this *ObservedDataType) SetFirstObservedText(s string) {
	this.First_observed = s
}

func (this *ObservedDataType) SetLastObserved(d time.Time) {
	this.Last_observed = d.UTC().Format(defs.TIME_RFC_3339)
}

// This function will allow you to assign the time as a string instead of using
// a time.Time object
func (this *ObservedDataType) SetLastObservedText(s string) {
	this.Last_observed = s
}

func (this *ObservedDataType) SetNumberObserved(i int) {
	this.Number_observed = i
}

func (this *ObservedDataType) SetCybox(s string) {
	this.Cybox = s
}

// ----------------------------------------------------------------------
// Private Methods - ObservedDataType
// ----------------------------------------------------------------------
