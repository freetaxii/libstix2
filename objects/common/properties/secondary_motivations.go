// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type SecondaryMotivationsPropertyType struct {
	Secondary_motivations []string `json:"secondary_motivations,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - SecondaryMotivationPropertyType
// ----------------------------------------------------------------------

// AddSecondaryMotivation takes in one parameter
// param: s - a string value that represents a motivation from the attack-motivation-ov vocab
func (this *SecondaryMotivationsPropertyType) AddSecondaryMotivation(s string) {
	if this.Secondary_motivations == nil {
		a := make([]string, 0)
		this.Secondary_motivations = a
	}
	this.Secondary_motivations = append(this.Secondary_motivations, s)
}
