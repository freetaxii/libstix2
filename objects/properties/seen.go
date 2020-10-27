// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import "github.com/wxj95/libstix2/timestamp"

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/*
SeenProperties - Properties used by one or more STIX objects that
captures the time that this object was first and last seen in STIX timestamp
format.
*/
type SeenProperties struct {
	FirstSeen string `json:"first_seen,omitempty"`
	LastSeen  string `json:"last_seen,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - SeenProperties
// ----------------------------------------------------------------------

/*
SetFirstSeenToCurrentTime - This methods sets the first seen time to the
current time
*/
func (o *SeenProperties) SetFirstSeenToCurrentTime() error {
	o.FirstSeen = timestamp.CurrentTime("micro")
	return nil
}

/*
SetFirstSeen -  This method takes in a timestamp in either time.Time or
string format and updates the first seen property with it. The value is stored
as a string, so if the value is in time.Time format, it will be converted to the
correct STIX timestamp format.
*/
func (o *SeenProperties) SetFirstSeen(t interface{}) error {
	ts, _ := timestamp.ToString(t, "micro")
	o.FirstSeen = ts
	return nil
}

/*
SetLastSeenToCurrentTime - This methods sets the first seen time to the
current time
*/
func (o *SeenProperties) SetLastSeenToCurrentTime() error {
	o.LastSeen = timestamp.CurrentTime("micro")
	return nil
}

/*
SetLastSeen -  This method takes in a time stamp in either time.Time or
string format and updates the last seen property with it. The value is stored as
a string, so if the value is in time.Time format, it will be converted to the
correct STIX time stamp format.
*/
func (o *SeenProperties) SetLastSeen(t interface{}) error {
	ts, _ := timestamp.ToString(t, "micro")
	o.LastSeen = ts
	return nil
}
