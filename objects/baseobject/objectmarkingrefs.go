// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package baseobject

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
ObjectMarkingRefsProperty - A property used by one or more STIX objects
that captures a list of STIX identifier that represent marking definition
objects.
*/
type ObjectMarkingRefsProperty struct {
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - ObjectMarkingRefsProperty
// ----------------------------------------------------------------------

/*
AddObjectMarkingRef - This method takes in a string value that represents a
STIX identifer for a marking definition object and adds it to the list of
object marking refs.
*/
func (o *ObjectMarkingRefsProperty) AddObjectMarkingRef(s string) error {
	o.ObjectMarkingRefs = append(o.ObjectMarkingRefs, s)
	return nil
}
