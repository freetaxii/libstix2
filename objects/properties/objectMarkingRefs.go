// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

// ObjectMarkingRefsPropertyType - A property used by one or more STIX objects
// that captures a list of STIX identifier that represent marking definition
// objects.
type ObjectMarkingRefsPropertyType struct {
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - ObjectMarkingRefsPropertyType
// ----------------------------------------------------------------------

// AddObjectMarkingRef - This method takes in a string value that represents a
// STIX identifer for a marking definition object and adds it to the list of
// object marking refs.
func (ezt *ObjectMarkingRefsPropertyType) AddObjectMarkingRef(s string) {
	ezt.ObjectMarkingRefs = append(ezt.ObjectMarkingRefs, s)
}
