// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import "github.com/wxj95/libstix2/resources"

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/*
MotivationProperties - Properties used by one or more STIX objects that
capture the primary and secondary motivations.
*/
type MotivationProperties struct {
	PrimaryMotivation    string   `json:"primary_motivation,omitempty"`
	SecondaryMotivations []string `json:"secondary_motivations,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - MotivationProperties - Setters
// ----------------------------------------------------------------------

/*
SetPrimaryMotivation - This methods takes in a string value representing a
motivation from the attack-motivation-ov vocab and updates the primary
motivation property.
*/
func (o *MotivationProperties) SetPrimaryMotivation(s string) error {
	o.PrimaryMotivation = s
	return nil
}

/*
GetPrimaryMotivation - This method returns the primary motivation.
*/
func (o *MotivationProperties) GetPrimaryMotivation() string {
	return o.PrimaryMotivation
}

/*
AddSecondaryMotivations - This method takes in a string value, a comma
separated list of string values, or a slice of string values that represents a
secondary motivation and adds it to the secondary motivations property.
*/
func (o *MotivationProperties) AddSecondaryMotivations(values interface{}) error {
	return resources.AddValuesToList(&o.SecondaryMotivations, values)
}
