// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package campaign

import (
	"github.com/freetaxii/libstix2/objects/common/properties"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
CampaignType defines all of the properties associated with the STIX Campaign
SDO. All of the methods not defined local to this type are inherited from the
individual properties.
*/
type CampaignType struct {
	properties.CommonObjectPropertiesType
	properties.NamePropertyType
	properties.DescriptionPropertyType
	properties.AliasesPropertyType
	properties.FirstSeenPropertyType
	properties.LastSeenPropertyType
	Objective string `json:"objective,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// New - This function will create a new campaign object.
func New() CampaignType {
	var obj CampaignType
	obj.InitNewObject("campaign")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - CampaignType
// ----------------------------------------------------------------------

// SetObjective - This method will take in a string representing an objective,
// goal, desired outcome, or intended effect and update the objective property.
func (ezt *CampaignType) SetObjective(s string) {
	ezt.Objective = s
}
