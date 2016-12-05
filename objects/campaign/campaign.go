// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package campaign

import (
	"github.com/freetaxii/libstix2/objects/common"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type CampaignType struct {
	common.CommonPropertiesType
	common.DescriptivePropertiesType
	common.AliasesType
	common.FirstLastSeenType
	Objective string `json:"objective,omitempty"`
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

func (this *CampaignType) SetObjective(s string) {
	this.Objective = s
}
