// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
//
// Types
//
// ----------------------------------------------------------------------

/*
GranularMarkingsProperty - A property used by one or more STIX objects
that captures a list of granular markings as defined by STIX.
*/
type GranularMarkingsProperty struct {
	GranularMarkings []GranularMarking `json:"granular_markings,omitempty"`
}

/*
GranularMarking - This type defines all of the properties associated with
the STIX Granular Marking type. All of the methods not defined local to this
type are inherited from the individual properties.
*/
type GranularMarking struct {
	LangProperty
	MarkingRef string   `json:"marking_ref,omitempty"`
	Selectors  []string `json:"selectors,omitempty"`
}

/*
ObjectMarkingRefsProperty - A property used by one or more STIX objects
that captures a list of STIX identifier that represent marking definition
objects.
*/
type ObjectMarkingRefsProperty struct {
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - GranularMarking
// ----------------------------------------------------------------------

/*
SetMarkingRef - This method takes in a string value representing a STIX
identifier of a marking definition object and sets the marking ref property
to that value.
*/
func (p *GranularMarking) SetMarkingRef(s string) error {
	p.MarkingRef = s
	return nil
}

/*
GetMarkingRef - This method returns the STIX identifier of the marking
definition object that was recorded in this granular marking type.
*/
func (p *GranularMarking) GetMarkingRef() string {
	return p.MarkingRef
}

/*
AddSelector - This method takes in a string value representing a STIX
granular marking selector and adds it to the list of selectors.
*/
func (p *GranularMarking) AddSelector(s string) error {
	p.Selectors = append(p.Selectors, s)
	return nil
}

// ----------------------------------------------------------------------
// Public Methods - ObjectMarkingRefsProperty
// ----------------------------------------------------------------------

/*
AddObjectMarkingRef - This method takes in a string value that represents a
STIX identifer for a marking definition object and adds it to the list of
object marking refs.
*/
func (p *ObjectMarkingRefsProperty) AddObjectMarkingRef(s string) error {
	p.ObjectMarkingRefs = append(p.ObjectMarkingRefs, s)
	return nil
}
