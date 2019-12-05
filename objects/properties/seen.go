// Copyright 2015-2019 Bret Jordan, All rights reserved.
// Copyright 2019 Oleksii Morozov, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import (
	"github.com/freetaxii/libstix2/timestamp"
)

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type SeenProperties struct {
	FirstSeen string `json:"first_seen,omitempty"`
	LastSeen  string `json:"last_seen,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - SeenProperties
// ----------------------------------------------------------------------

/*
SetFirstSeen - sets first seen date.

The time that this Infrastructure was first seen performing malicious activities.
*/
func (o *SeenProperties) SetFirstSeen(t interface{}) error {
	ts, _ := timestamp.ToString(t, "milli")
	o.FirstSeen = ts

	return nil
}

/*
SetLastSeen - sets last seen date.

The time that this Infrastructure was last seen performing malicious activities.

If this property and the first_seen property are both defined, then this property MUST be greater than or equal to the timestamp in the first_seen property.
*/
func (o *SeenProperties) SetLastSeen(t interface{}) error {
	ts, _ := timestamp.ToString(t, "milli")
	o.LastSeen = ts

	return nil
}
