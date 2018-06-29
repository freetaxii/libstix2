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
SecondaryMotivationsProperty - A property used by one or more STIX objects
that captures a list of motivations.
*/
type SecondaryMotivationsProperty struct {
	SecondaryMotivations []string `json:"secondary_motivations,omitempty"`
}

// ----------------------------------------------------------------------
//
// Public Methods - SecondaryMotivationPropertyType
//
// ----------------------------------------------------------------------

/*
AddSecondaryMotivation - This method takes in a string value that represents
a motivation from the attack-motivation-ov vocab and adds it to the list of
motivations in the secondary motivations property.
*/
func (p *SecondaryMotivationsProperty) AddSecondaryMotivation(s string) error {
	p.SecondaryMotivations = append(p.SecondaryMotivations, s)
	return nil
}
