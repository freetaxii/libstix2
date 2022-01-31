// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import "fmt"

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/*
MarkingProperties - Properties used by one or more STIX objects that captures
the data markings for an object. These can be in the form of object markings or
granular markings. The object markings is a list of STIX identifier that
represent marking definition objects. The granular markings is a list of
granular markings.
*/
type MarkingProperties struct {
	ObjectMarkingRefs []string          `json:"object_marking_refs,omitempty" bson:"object_marking_refs,omitempty"`
	GranularMarkings  []GranularMarking `json:"granular_markings,omitempty" bson:"granular_markings,omitempty"`
}

/*
GranularMarking - This type defines all of the properties associated with the
STIX Granular Marking type. All of the methods not defined local to this type
are inherited from the individual properties.
*/
type GranularMarking struct {
	LangProperty
	MarkingRef string   `json:"marking_ref,omitempty" bson:"marking_ref,omitempty"`
	Selectors  []string `json:"selectors,omitempty" bson:"selectors,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - MarkingProperty - Setters
// ----------------------------------------------------------------------

/*
AddObjectMarkingRef - This method takes in a string value that represents a
STIX identifier for a marking definition object and adds it to the list of object
marking refs.
*/
func (o *MarkingProperties) AddObjectMarkingRef(s string) error {
	o.ObjectMarkingRefs = append(o.ObjectMarkingRefs, s)
	return nil
}

// ----------------------------------------------------------------------
// Public Methods - GranularMarking - Setters
// ----------------------------------------------------------------------

/*
SetMarkingRef - This method takes in a string value representing a STIX
identifier of a marking definition object and sets the marking ref property to
that value.
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

// ----------------------------------------------------------------------
// Public Methods - MarkingProperties - Checks
// ----------------------------------------------------------------------

/*
Compare - This method will compare two properties to make sure they are the
same and will return a boolean, an integer that tracks the number of problems
found, and a slice of strings that contain the detailed results, whether good or
bad.
*/
func (o *MarkingProperties) Compare(obj2 *MarkingProperties) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Object Marking Refs
	if len(o.ObjectMarkingRefs) != len(obj2.ObjectMarkingRefs) {
		problemsFound++
		str := fmt.Sprintf("-- The number of entries in object marking refs do not match: %d | %d", len(o.ObjectMarkingRefs), len(obj2.ObjectMarkingRefs))
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The number of entries in object marking refs match: %d | %d", len(o.ObjectMarkingRefs), len(obj2.ObjectMarkingRefs))
		resultDetails = append(resultDetails, str)

		// If lengths are the same, then check each value
		for index := range o.ObjectMarkingRefs {
			if o.ObjectMarkingRefs[index] != obj2.ObjectMarkingRefs[index] {
				problemsFound++
				str := fmt.Sprintf("-- The object marking ref values do not match: %s | %s", o.ObjectMarkingRefs[index], obj2.ObjectMarkingRefs[index])
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ The object marking ref values match: %s | %s", o.ObjectMarkingRefs[index], obj2.ObjectMarkingRefs[index])
				resultDetails = append(resultDetails, str)
			}
		}
	}

	// Check Granular Markings
	if len(o.GranularMarkings) != len(obj2.GranularMarkings) {
		problemsFound++
		str := fmt.Sprintf("-- The number of entries in granular markings do not match: %d | %d", len(o.GranularMarkings), len(obj2.GranularMarkings))
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The number of entries in granular markings match: %d | %d", len(o.GranularMarkings), len(obj2.GranularMarkings))
		resultDetails = append(resultDetails, str)
		for index := range o.GranularMarkings {

			// Check Granular Marking Languages
			if o.GranularMarkings[index].Lang != obj2.GranularMarkings[index].Lang {
				problemsFound++
				str := fmt.Sprintf("-- The language values do not match: %s | %s", o.GranularMarkings[index].Lang, obj2.GranularMarkings[index].Lang)
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ The languages values match: %s | %s", o.GranularMarkings[index].Lang, obj2.GranularMarkings[index].Lang)
				resultDetails = append(resultDetails, str)
			}

			// Check Granular Marking Refs
			if o.GranularMarkings[index].MarkingRef != obj2.GranularMarkings[index].MarkingRef {
				problemsFound++
				str := fmt.Sprintf("-- The marking ref values do not match: %s | %s", o.GranularMarkings[index].MarkingRef, obj2.GranularMarkings[index].MarkingRef)
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ The marking ref values match: %s | %s", o.GranularMarkings[index].MarkingRef, obj2.GranularMarkings[index].MarkingRef)
				resultDetails = append(resultDetails, str)
			}

			// Check Granular Marking Selectors
			if len(o.GranularMarkings[index].Selectors) != len(obj2.GranularMarkings[index].Selectors) {
				problemsFound++
				str := fmt.Sprintf("-- The number of entries in selectors do not match: %d | %d", len(o.GranularMarkings[index].Selectors), len(obj2.GranularMarkings[index].Selectors))
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ The number of entries in selectors match: %d | %d", len(o.GranularMarkings[index].Selectors), len(obj2.GranularMarkings[index].Selectors))
				resultDetails = append(resultDetails, str)

				// If lengths are the same, then check each value
				for j := range o.GranularMarkings[index].Selectors {
					if o.GranularMarkings[index].Selectors[j] != obj2.GranularMarkings[index].Selectors[j] {
						problemsFound++
						str := fmt.Sprintf("-- The selector values do not match: %s | %s", o.GranularMarkings[index].Selectors[j], obj2.GranularMarkings[index].Selectors[j])
						resultDetails = append(resultDetails, str)
					} else {
						str := fmt.Sprintf("++ The selector values match: %s | %s", o.GranularMarkings[index].Selectors[j], obj2.GranularMarkings[index].Selectors[j])
						resultDetails = append(resultDetails, str)
					}
				}
			}
		}
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
