// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import "github.com/freetaxii/libstix2/resources"

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/* GoalsProperty - A property used by one or more STIX objects that captures a
list of goals that are part of the STIX object. */
type GoalsProperty struct {
	Goals []string `json:"goals,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - GoalsType - Setters
// ----------------------------------------------------------------------

/* AddGoals - This method takes in a string value, a comma separated list of
string values, or a slice of string values that represents a goal and adds it to
the goals property. */
func (o *GoalsProperty) AddGoals(values interface{}) error {
	return resources.AddValuesToList(&o.Goals, values)
}
