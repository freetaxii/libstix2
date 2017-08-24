// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type ResourceLevelPropertyType struct {
	Resource_level string `json:"resource_level,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - ResourceLevelPropertyType
// ----------------------------------------------------------------------

// SetResourceLevel takes in one parameter
// param: s - a string value representing a resource level from the attack-resrouce-level-ov vocab
func (this *ResourceLevelPropertyType) SetResourceLevel(s string) {
	this.Resource_level = s
}

func (this *ResourceLevelPropertyType) GetResourceLevel() string {
	return this.Resource_level
}
