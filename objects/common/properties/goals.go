// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type GoalsPropertyType struct {
	Goals []string `json:"goals,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - GoalsType
// ----------------------------------------------------------------------

// AddGoal takes in one parameter
// param: s - a string value that represents a goal
func (this *GoalsPropertyType) AddGoal(s string) {
	if this.Goals == nil {
		a := make([]string, 0)
		this.Goals = a
	}
	this.Goals = append(this.Goals, s)
}
