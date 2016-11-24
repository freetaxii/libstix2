// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package campaign

import (
	"github.com/freetaxii/libstix2/messages/stix"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type CampaignType struct {
	stix.CommonPropertiesType
	stix.DescriptivePropertiesType
	stix.AliasesType
	First_seen           string `json:"first_seen,omitempty"`
	First_seen_precision string `json:"first_seen_precision,omitempty"`
	Objective            string `json:"objective,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() CampaignType {
	var obj CampaignType
	obj.MessageType = "campaign"
	obj.Id = obj.NewId("campaign")
	obj.Created = obj.GetCurrentTime()
	obj.Modified = obj.Created
	obj.Version = 1
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - CampaignType
// ----------------------------------------------------------------------

func (this *CampaignType) SetFirstSeen(t interface{}, s string) error {

	ts, err := this.VerifyTimestamp(t)
	if err != nil {
		return err
	}
	this.First_seen = ts

	p, err := this.VerifyPrecision(s)
	if err != nil {
		return err
	}
	this.First_seen_precision = p

	return nil
}

func (this *CampaignType) SetObjective(s string) {
	this.Objective = s
}
