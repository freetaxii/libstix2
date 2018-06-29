// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
//
// Types
//
// ----------------------------------------------------------------------

/*
PrimaryMotivationProperty - A property used by one or more STIX objects
that captures the primary motivation.
*/
type PrimaryMotivationProperty struct {
	PrimaryMotivation string `json:"primary_motivation,omitempty"`
}

// ----------------------------------------------------------------------
//
// Public Methods - PrimaryMotivationProperty
//
// ----------------------------------------------------------------------

/*
SetPrimaryMotivation - This methods takes in a string value representing a
motivation from the attack-motivation-ov vocab and updates the primary
motivation property.
*/
func (p *PrimaryMotivationProperty) SetPrimaryMotivation(s string) error {
	p.PrimaryMotivation = s
	return nil
}

/*
GetPrimaryMotivation - This method returns the primary motivation.
*/
func (p *PrimaryMotivationProperty) GetPrimaryMotivation() string {
	return p.PrimaryMotivation
}
