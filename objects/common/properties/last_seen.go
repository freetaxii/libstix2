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

type LastSeenPropertyType struct {
	Last_seen string `json:"last_seen,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - LastSeenPropertyType
// ----------------------------------------------------------------------

// SetLastSeen takes in one parameter
// param: t - a timestamp in either time.Time or string format
func (this *LastSeenPropertyType) SetLastSeen(t interface{}) {

	ts := timestamp.Verify(t)
	this.Last_seen = ts
}
