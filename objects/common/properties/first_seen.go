// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

import (
	"github.com/freetaxii/libstix2/objects/common/timestamp"
)

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type FirstSeenPropertyType struct {
	First_seen string `json:"first_seen,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - FirstSeenPropertyType
// ----------------------------------------------------------------------

// SetFirstSeen takes in one parameter
// param: t - a timestamp in either time.Time or string format
func (this *FirstSeenPropertyType) SetFirstSeen(t interface{}) {

	ts := timestamp.Verify(t)
	this.First_seen = ts
}
