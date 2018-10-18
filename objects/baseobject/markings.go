// Copyright 2015-2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package baseobject

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
MarkingProperties - Properties used by one or more STIX objects that captures
the data markings for an object. These can be in the form of object markings
or granular markings. The object markings is a list of STIX identifier that
represent marking definition objects. The granular markings is a list of
granular markings.
*/
type MarkingProperties struct {
	ObjectMarkingRefs []string          `json:"object_marking_refs,omitempty"`
	GranularMarkings  []GranularMarking `json:"granular_markings,omitempty"`
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

// ----------------------------------------------------------------------
// Public Methods - MarkingProperties
// ----------------------------------------------------------------------

/*
AddObjectMarkingRef - This method takes in a string value that represents a
STIX identifer for a marking definition object and adds it to the list of
object marking refs.
*/
func (o *MarkingProperties) AddObjectMarkingRef(s string) error {
	o.ObjectMarkingRefs = append(o.ObjectMarkingRefs, s)
	return nil
}

// ----------------------------------------------------------------------
// Public Methods - GranularMarking
// ----------------------------------------------------------------------

/*
SetMarkingRef - This method takes in a string value representing a STIX
identifier of a marking definition object and sets the marking ref property
to that value.
*/
func (o *GranularMarking) SetMarkingRef(s string) error {
	o.MarkingRef = s
	return nil
}

/*
GetMarkingRef - This method returns the STIX identifier of the marking
definition object that was recorded in this granular marking type.
*/
func (o *GranularMarking) GetMarkingRef() string {
	return o.MarkingRef
}

/*
AddSelector - This method takes in a string value representing a STIX
granular marking selector and adds it to the list of selectors.
*/
func (o *GranularMarking) AddSelector(s string) error {
	o.Selectors = append(o.Selectors, s)
	return nil
}
