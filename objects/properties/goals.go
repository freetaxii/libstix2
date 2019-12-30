// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

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

/* AddGoal - This method takes in a string value that represents a goal and adds
it to the list of goals in the goals property. */
func (o *GoalsProperty) AddGoal(s string) error {
	// if this.Goals == nil {
	// 	a := make([]string, 0)
	// 	this.Goals = a
	// }
	o.Goals = append(o.Goals, s)
	return nil
}
