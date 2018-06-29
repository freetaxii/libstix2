// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import (
	"github.com/freetaxii/libstix2/common/timestamp"
)

// ----------------------------------------------------------------------
//
// Types
//
// ----------------------------------------------------------------------

/*
LastSeenProperty - A property used by one or more STIX objects that
captures the time that this object was last seen in STIX timestamp format,
which is an RFC3339 format.
*/
type LastSeenProperty struct {
	LastSeen string `json:"last_seen,omitempty"`
}

// ----------------------------------------------------------------------
//
// Public Methods - LastSeenProperty
//
// ----------------------------------------------------------------------

/*
SetLastSeenToCurrentTime - This methods sets the first seen time to the
current time
*/
func (p *LastSeenProperty) SetLastSeenToCurrentTime() error {
	p.LastSeen = timestamp.GetCurrentTime("micro")
	return nil
}

/*
SetLastSeen -  This method takes in a time stamp in either time.Time or string
format and updates the last seen property with it. The value is stored as a
string, so if the value is in time.Time format, it will be converted to the
correct STIX time stamp format.
*/
func (p *LastSeenProperty) SetLastSeen(t interface{}) error {
	ts, _ := timestamp.ToString(t, "micro")
	p.LastSeen = ts
	return nil
}
