// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package campaign

// ----------------------------------------------------------------------
// Public Methods - Campaign
// ----------------------------------------------------------------------

/* SetObjective - This method will take in a string representing an objective,
goal, desired outcome, or intended effect and update the objective property. */
func (o *Campaign) SetObjective(s string) error {
	o.Objective = s
	return nil
}
