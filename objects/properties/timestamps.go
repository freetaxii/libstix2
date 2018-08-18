// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import "github.com/freetaxii/libstix2/timestamp"

// ----------------------------------------------------------------------
//
// Types
//
// ----------------------------------------------------------------------

/*
BaseTimestampProperties - Properties used by all STIX objects that captures the
time that this object was created and modified in STIX timestamp format.
*/
type BaseTimestampProperties struct {
	Created  string `json:"created,omitempty"`
	Modified string `json:"modified,omitempty"`
}

/*
SeenTimestampProperties - Properties used by one or more STIX objects that captures
the time that this object was first and last seen in STIX timestamp format.
*/
type SeenTimestampProperties struct {
	FirstSeen string `json:"first_seen,omitempty"`
	LastSeen  string `json:"last_seen,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - baseTimestampProperties
// ----------------------------------------------------------------------

/*
SetCreatedToCurrentTime - This methods sets the object created time to the
current time
*/
func (p *BaseTimestampProperties) SetCreatedToCurrentTime() error {
	p.Created = timestamp.CurrentTime("milli")
	return nil
}

/*
SetCreated - This method takes in a timestamp in either time.Time or string
format and updates the created property with it. The value is stored as a
string, so if the value is in time.Time format, it will be converted to the
correct STIX timestamp format.
*/
func (p *BaseTimestampProperties) SetCreated(t interface{}) error {
	ts, _ := timestamp.ToString(t, "milli")
	p.Created = ts
	return nil
}

/*
GetCreated - This method will return the created timestamp as a string.
*/
func (p *BaseTimestampProperties) GetCreated() string {
	return p.Created
}

/*
SetModifiedToCreated sets the object modified time to be the same as the
created time.
*/
func (p *BaseTimestampProperties) SetModifiedToCreated() error {
	p.Modified = p.Created
	return nil
}

/*
SetModifiedToCurrentTime - This methods sets the object modified time to the
current time
*/
func (p *BaseTimestampProperties) SetModifiedToCurrentTime() error {
	p.Modified = timestamp.CurrentTime("milli")
	return nil
}

/*
SetModified - This method takes in a timestamp in either time.Time or string
format and updates the modifed property with it. The value is stored as a
string, so if the value is in time.Time format, it will be converted to the
correct STIX timestamp format.
*/
func (p *BaseTimestampProperties) SetModified(t interface{}) error {
	ts, _ := timestamp.ToString(t, "milli")
	p.Modified = ts
	return nil
}

/*
GetModified - This method will return the modified timestamp as a string. If
the value is the same as the created timestamp, then this object is the
first version of the object.
*/
func (p *BaseTimestampProperties) GetModified() string {
	return p.Modified
}

// ----------------------------------------------------------------------
// Public Methods - SeenTimestampProperties
// ----------------------------------------------------------------------

/*
SetFirstSeenToCurrentTime - This methods sets the first seen time to the
current time
*/
func (p *SeenTimestampProperties) SetFirstSeenToCurrentTime() error {
	p.FirstSeen = timestamp.CurrentTime("micro")
	return nil
}

/*
SetFirstSeen -  This method takes in a timestamp in either time.Time or string
format and updates the first seen property with it. The value is stored as a
string, so if the value is in time.Time format, it will be converted to the
correct STIX timestamp format.
*/
func (p *SeenTimestampProperties) SetFirstSeen(t interface{}) error {
	ts, _ := timestamp.ToString(t, "micro")
	p.FirstSeen = ts
	return nil
}

/*
SetLastSeenToCurrentTime - This methods sets the first seen time to the
current time
*/
func (p *SeenTimestampProperties) SetLastSeenToCurrentTime() error {
	p.LastSeen = timestamp.CurrentTime("micro")
	return nil
}

/*
SetLastSeen -  This method takes in a time stamp in either time.Time or string
format and updates the last seen property with it. The value is stored as a
string, so if the value is in time.Time format, it will be converted to the
correct STIX time stamp format.
*/
func (p *SeenTimestampProperties) SetLastSeen(t interface{}) error {
	ts, _ := timestamp.ToString(t, "micro")
	p.LastSeen = ts
	return nil
}
