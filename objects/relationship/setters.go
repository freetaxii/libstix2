// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package relationship

import "github.com/freetaxii/libstix2/objects"

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/*
SetType - This method takes in a string value that represents the type name
of the relationship and updates the relationship type property.
*/
func (o *Relationship) SetType(s string) error {
	o.RelationshipType = s
	return nil
}

/*
SetSourceRef - This method takes in a string value that represents a STIX
identifier of the source STIX object in the relationship and updates the source
ref property.
*/
func (o *Relationship) SetSourceRef(s string) error {
	o.SourceRef = s
	return nil
}

/*
SetTargetRef - This method takes in a string value that represents a STIX
identifier of the target STIX object in the relationship and updates the target
ref property.
*/
func (o *Relationship) SetTargetRef(s string) error {
	o.TargetRef = s
	return nil
}

/*
SetSourceTarget - This methods takes in two string values where both
represent a STIX identifier. This is a convenience function for setting both
ends of the relationship at the same time. The first identifier is for the
source and the second is for the target.
*/
func (o *Relationship) SetSourceTarget(s, t string) error {
	o.SourceRef = s
	o.TargetRef = t
	return nil
}

/*
SetStartTime - This method will take in a timestamp in either time.Time or
string format and will set the valid_from property to that value.
*/
func (o *Relationship) SetStartTime(t interface{}) error {
	ts, _ := objects.TimeToString(t, "micro")
	o.StartTime = ts
	return nil
}

/*
SetStopTime - This method will take in a timestamp in either time.Time or
string format and will set the valid_from property to that value.
*/
func (o *Relationship) SetStopTime(t interface{}) error {
	ts, _ := objects.TimeToString(t, "micro")
	o.StopTime = ts
	return nil
}
