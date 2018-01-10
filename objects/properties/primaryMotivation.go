// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
PrimaryMotivationPropertyType - A property used by one or more STIX objects
that captures the primary motivation.
*/
type PrimaryMotivationPropertyType struct {
	PrimaryMotivation string `json:"primary_motivation,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - PrimaryMotivationPropertyType
// ----------------------------------------------------------------------

/*
SetPrimaryMotivation - This methods takes in a string value representing a
motivation from the attack-motivation-ov vocab and updates the primary
motivation property.
*/
func (ezt *PrimaryMotivationPropertyType) SetPrimaryMotivation(s string) error {
	ezt.PrimaryMotivation = s
	return nil
}

/*
GetPrimaryMotivation - This method returns the primary motivation.
*/
func (ezt *PrimaryMotivationPropertyType) GetPrimaryMotivation() string {
	return ezt.PrimaryMotivation
}
