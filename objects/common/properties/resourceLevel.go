// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

// ResourceLevelPropertyType - A property used by one or more STIX objects
// that captures the resource level.
type ResourceLevelPropertyType struct {
	ResourceLevel string `json:"resource_level,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - ResourceLevelPropertyType
// ----------------------------------------------------------------------

// SetResourceLevel - This method takes in a string value representing a
// resource level from the attack-resrouce-level-ov vocab and updates the
// resource level property.
func (p *ResourceLevelPropertyType) SetResourceLevel(s string) {
	p.ResourceLevel = s
}

// GetResourceLevel - This method returns the resource level.
func (p *ResourceLevelPropertyType) GetResourceLevel() string {
	return p.ResourceLevel
}
