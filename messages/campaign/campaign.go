// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package campaign

import (
	"errors"
	"github.com/freetaxii/libstix2/messages/defs"
	"github.com/freetaxii/libstix2/messages/stix"
	"time"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type CampaignType struct {
	stix.CommonProperties
	Name                 string   `json:"name,omitempty"`
	Description          string   `json:"description,omitempty"`
	Aliases              []string `json:"aliases,omitempty"`
	First_seen           string   `json:"first_seen,omitempty"`
	First_seen_precision string   `json:"first_seen_precision,omitempty"`
	Objective            string   `json:"objective,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() CampaignType {
	var obj CampaignType
	obj.MessageType = "campaign"
	obj.Id = stix.NewId("campaign")
	obj.Created = stix.GetCurrentTime().UTC().Format(defs.TIME_RFC_3339)
	obj.Modified = obj.Created
	obj.Version = 1
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - Common Properties
// ----------------------------------------------------------------------

func (this *CampaignType) SetCreatedBy(s string) {
	this.Created_by_ref = s
}

func (this *CampaignType) SetModified(d time.Time) {
	this.Modified = d.UTC().Format(defs.TIME_RFC_3339)
}

func (this *CampaignType) SetVersion(i int) error {
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

func (this *CampaignType) SetRevoked() {
	this.Revoked = true
}

func (this *CampaignType) AddLabel(value string) {
	if this.Labels == nil {
		a := make([]string, 0)
		this.Labels = a
	}
	this.Labels = append(this.Labels, value)
}

func (this *CampaignType) GetId() string {
	return this.Id
}

// ----------------------------------------------------------------------
// Public Methods - CampaignType
// ----------------------------------------------------------------------

func (this *CampaignType) SetName(s string) {
	this.Name = s
}

func (this *CampaignType) SetDescription(s string) {
	this.Description = s
}

func (this *CampaignType) AddAlias(value string) {
	if this.Aliases == nil {
		a := make([]string, 0)
		this.Aliases = a
	}
	this.Aliases = append(this.Aliases, value)
}

func (this *CampaignType) SetFirstSeen(d time.Time) {
	this.First_seen = d.UTC().Format(defs.TIME_RFC_3339)
}

// This function will allow you to assign the time as a string instead of using
// a time.Time object
func (this *CampaignType) SetFirstSeenText(s string) {
	this.First_seen = s
}

// TODO Add precision functions

func (this *CampaignType) SetObjective(s string) {
	this.Objective = s
}
