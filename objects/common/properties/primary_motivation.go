// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type PrimaryMotivationPropertyType struct {
	Primary_motivation string `json:"primary_motivation,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - PrimaryMotivationPropertyType
// ----------------------------------------------------------------------

// SetPrimaryMotivation takes in one parameter
// param: s - a string value representing a motivation from the attack-motivation-ov vocab
func (this *PrimaryMotivationPropertyType) SetPrimaryMotivation(s string) {
	this.Primary_motivation = s
}

func (this *PrimaryMotivationPropertyType) GetPrimaryMotivation() string {
	return this.Primary_motivation
}
