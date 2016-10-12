// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package report

import (
	"errors"
	"github.com/freetaxii/libstix2/messages/defs"
	"github.com/freetaxii/libstix2/messages/stix"
	"time"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type ReportType struct {
	stix.CommonProperties
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Published   string   `json:"published,omitempty"`
	Object_refs []string `json:"object_refs,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() ReportType {
	var obj ReportType
	obj.MessageType = "report"
	obj.Id = stix.NewId("report")
	obj.Created = stix.GetCurrentTime().UTC().Format(defs.TIME_RFC_3339)
	obj.Modified = obj.Created
	obj.Version = 1
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - Common Properties
// ----------------------------------------------------------------------

func (this *ReportType) SetCreatedBy(s string) {
	this.Created_by_ref = s
}

func (this *ReportType) SetModified(d time.Time) {
	this.Modified = d.UTC().Format(defs.TIME_RFC_3339)
}

func (this *ReportType) SetVersion(i int) error {
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

func (this *ReportType) SetRevoked() {
	this.Revoked = true
}

func (this *ReportType) AddLabel(value string) {
	if this.Labels == nil {
		a := make([]string, 0)
		this.Labels = a
	}
	this.Labels = append(this.Labels, value)
}

func (this *ReportType) GetId() string {
	return this.Id
}

// ----------------------------------------------------------------------
// Public Methods - ReportType
// ----------------------------------------------------------------------

func (this *ReportType) SetName(s string) {
	this.Name = s
}

func (this *ReportType) SetDescription(s string) {
	this.Description = s
}

func (this *ReportType) SetPublished(d time.Time) {
	this.Published = d.UTC().Format(defs.TIME_RFC_3339)
}

// This function will allow you to assign the time as a string instead of using
// a time.Time object
func (this *ReportType) SetPublishedText(s string) {
	this.Published = s
}

func (this *ReportType) AddObject(value string) {
	if this.Object_refs == nil {
		a := make([]string, 0)
		this.Object_refs = a
	}
	this.Object_refs = append(this.Object_refs, value)
}

// ----------------------------------------------------------------------
// Private Methods - ReportType
// ----------------------------------------------------------------------
