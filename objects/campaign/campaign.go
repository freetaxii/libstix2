// Copyright 2017 Bret Jordan, All rights reserved.
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
	common.CommonObjectPropertiesType
	common.DescriptivePropertiesType
	common.AliasesPropertiesType
	common.FirstLastSeenPropertiesType
	Objective string `json:"objective,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() CampaignType {
	var obj CampaignType
	obj.InitNewObject("campaign")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - CampaignType
// ----------------------------------------------------------------------

// SetObjective takes in one parameter
// param: s - a string value representing an objective, goal, desired outcome, or indended effect
func (this *CampaignType) SetObjective(s string) {
	this.Objective = s
}
