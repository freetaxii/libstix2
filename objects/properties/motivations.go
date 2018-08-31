// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
MotivationProperties - Properties used by one or more STIX objects that capture
the primary and secondary motivations.
*/
type MotivationProperties struct {
	PrimaryMotivation    string   `json:"primary_motivation,omitempty"`
	SecondaryMotivations []string `json:"secondary_motivations,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - MotivationProperties
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
AddSecondaryMotivation - This method takes in a string value that represents
a motivation from the attack-motivation-ov vocab and adds it to the list of
motivations in the secondary motivations property.
*/
func (o *MotivationProperties) AddSecondaryMotivation(s string) error {
	o.SecondaryMotivations = append(o.SecondaryMotivations, s)
	return nil
}
