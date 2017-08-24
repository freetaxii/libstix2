// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type ObjectMarkingRefsPropertyType struct {
	Object_marking_refs []string `json:"object_marking_refs,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - ObjectMarkingRefsPropertyType
// ----------------------------------------------------------------------

// AddObjectMarkingRef takes in one parameter
// param: s - a string value that represents a STIX identifer
func (this *ObjectMarkingRefsPropertyType) AddObjectMarkingRef(s string) {
	if this.Object_marking_refs == nil {
		a := make([]string, 0)
		this.Object_marking_refs = a
	}
	this.Object_marking_refs = append(this.Object_marking_refs, s)
}
