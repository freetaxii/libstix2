// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package report

import "github.com/freetaxii/libstix2/objects"

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/*
SetPublished - This method takes in a timestamp in either time.Time or string
format and updates the published timestamp property.
*/
func (o *Report) SetPublished(t interface{}) error {
	ts, _ := objects.TimeToString(t, "micro")
	o.Published = ts
	return nil
}
