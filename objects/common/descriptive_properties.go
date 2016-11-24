// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package common

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type DescriptivePropertiesType struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - KillChainPhaseType
// ----------------------------------------------------------------------

func (this *DescriptivePropertiesType) SetName(s string) {
	this.Name = s
}

func (this *DescriptivePropertiesType) SetDescription(s string) {
	this.Description = s
}
