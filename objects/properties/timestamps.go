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
CreatedProperty - A property used by one or more STIX objects that
captures the time that this object was created in STIX timestamp format.
*/
type CreatedProperty struct {
	Created string `json:"created,omitempty"`
}

/*
ModifiedProperty - A property used by one or more STIX objects that
captures the time that this object was modified in STIX timestamp format.
*/
type ModifiedProperty struct {
	Modified string `json:"modified,omitempty"`
}

/*
SeenProperties - Properties used by one or more STIX objects that captures
the time that this object was first and last seen in STIX timestamp format.
*/
type SeenProperties struct {
	FirstSeen string `json:"first_seen,omitempty"`
	LastSeen  string `json:"last_seen,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - CommonObjectProperties
// ----------------------------------------------------------------------

/*
SetModifiedToCreated sets the object modified time to be the same as the
created time. This has to be done at this level, since at the individual
properties type say "ModifiedProperty" p.Created does not exist.
But it will exist at this level of inheritance
*/
func (p *CommonObjectProperties) SetModifiedToCreated() error {
	p.Modified = p.Created
	return nil
}

// ----------------------------------------------------------------------
// Public Methods - CreatedProperty
// ----------------------------------------------------------------------

/*
SetCreatedToCurrentTime - This methods sets the object created time to the
current time
*/
func (p *CreatedProperty) SetCreatedToCurrentTime() error {
	p.Created = timestamp.CurrentTime("milli")
	return nil
}

/*
SetCreated - This method takes in a timestamp in either time.Time or string
format and updates the created property with it. The value is stored as a
string, so if the value is in time.Time format, it will be converted to the
correct STIX timestamp format.
*/
func (p *CreatedProperty) SetCreated(t interface{}) error {
	ts, _ := timestamp.ToString(t, "milli")
	p.Created = ts
	return nil
}

/*
GetCreated - This method will return the created timestamp as a string.
*/
func (p *CreatedProperty) GetCreated() string {
	return p.Created
}

// ----------------------------------------------------------------------
// Public Methods - ModifiedProperty
// ----------------------------------------------------------------------

/*
SetModifiedToCurrentTime - This methods sets the object modified time to the
current time
*/
func (p *ModifiedProperty) SetModifiedToCurrentTime() error {
	p.Modified = timestamp.CurrentTime("milli")
	return nil
}

/*
SetModified - This method takes in a timestamp in either time.Time or string
format and updates the modifed property with it. The value is stored as a
string, so if the value is in time.Time format, it will be converted to the
correct STIX timestamp format.
*/
func (p *ModifiedProperty) SetModified(t interface{}) error {
	ts, _ := timestamp.ToString(t, "milli")
	p.Modified = ts
	return nil
}

/*
GetModified - This method will return the modified timestamp as a string. If
the value is the same as the created timestamp, then this object is the
first version of the object.
*/
func (p *ModifiedProperty) GetModified() string {
	return p.Modified
}

// ----------------------------------------------------------------------
// Public Methods - SeenProperties
// ----------------------------------------------------------------------

/*
SetFirstSeenToCurrentTime - This methods sets the first seen time to the
current time
*/
func (p *SeenProperties) SetFirstSeenToCurrentTime() error {
	p.FirstSeen = timestamp.CurrentTime("micro")
	return nil
}

/*
SetFirstSeen -  This method takes in a timestamp in either time.Time or string
format and updates the first seen property with it. The value is stored as a
string, so if the value is in time.Time format, it will be converted to the
correct STIX timestamp format.
*/
func (p *SeenProperties) SetFirstSeen(t interface{}) error {
	ts, _ := timestamp.ToString(t, "micro")
	p.FirstSeen = ts
	return nil
}

/*
SetLastSeenToCurrentTime - This methods sets the first seen time to the
current time
*/
func (p *SeenProperties) SetLastSeenToCurrentTime() error {
	p.LastSeen = timestamp.CurrentTime("micro")
	return nil
}

/*
SetLastSeen -  This method takes in a time stamp in either time.Time or string
format and updates the last seen property with it. The value is stored as a
string, so if the value is in time.Time format, it will be converted to the
correct STIX time stamp format.
*/
func (p *SeenProperties) SetLastSeen(t interface{}) error {
	ts, _ := timestamp.ToString(t, "micro")
	p.LastSeen = ts
	return nil
}
