// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/* ResourceLevelProperty - A property used by one or more STIX objects that
captures the resource level. */
type ResourceLevelProperty struct {
	ResourceLevel string `json:"resource_level,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - ResourceLevelProperty - Setters
// ----------------------------------------------------------------------

/* SetResourceLevel - This method takes in a string value representing a
resource level from the attack-resrouce-level-ov vocab and updates the resource
level property. */
func (o *ResourceLevelProperty) SetResourceLevel(s string) error {
	o.ResourceLevel = s
	return nil
}

/* GetResourceLevel - This method returns the resource level. */
func (o *ResourceLevelProperty) GetResourceLevel() string {
	return o.ResourceLevel
}
