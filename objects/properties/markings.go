// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import "fmt"

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
// Public Methods - MarkingProperty
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

// ----------------------------------------------------------------------
// Public Functions - MarkingProperties
// ----------------------------------------------------------------------

/* CompareMarkingProperties - This function will compare two properties to make
sure they are the same and will return a boolean, an integer that tracks the
number of problems found, and a slice of strings that contain the detailed
results, whether good or bad. */
func CompareMarkingProperties(obj1, obj2 *MarkingProperties) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Object Marking Refs
	if len(obj1.ObjectMarkingRefs) != len(obj2.ObjectMarkingRefs) {
		problemsFound++
		str := fmt.Sprintf("-- The number of entries in Object Marking Refs do not match: %d | %d", len(obj1.ObjectMarkingRefs), len(obj2.ObjectMarkingRefs))
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The number of entries in Object Marking Refs match: %d | %d", len(obj1.ObjectMarkingRefs), len(obj2.ObjectMarkingRefs))
		resultDetails = append(resultDetails, str)

		// If lengths are the same, then check each value
		for index := range obj1.ObjectMarkingRefs {
			if obj1.ObjectMarkingRefs[index] != obj2.ObjectMarkingRefs[index] {
				problemsFound++
				str := fmt.Sprintf("-- The Object Marking Ref values do not match: %s | %s", obj1.ObjectMarkingRefs[index], obj2.ObjectMarkingRefs[index])
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ The Object Marking Ref values match: %s | %s", obj1.ObjectMarkingRefs[index], obj2.ObjectMarkingRefs[index])
				resultDetails = append(resultDetails, str)
			}
		}
	}

	// Check Granular Markings
	if len(obj1.GranularMarkings) != len(obj2.GranularMarkings) {
		problemsFound++
		str := fmt.Sprintf("-- The number of entries in Granular Markings do not match: %d | %d", len(obj1.GranularMarkings), len(obj2.GranularMarkings))
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The number of entries in Granular Markings match: %d | %d", len(obj1.GranularMarkings), len(obj2.GranularMarkings))
		resultDetails = append(resultDetails, str)
		for index := range obj1.GranularMarkings {

			// Check Granular Marking Languages
			if obj1.GranularMarkings[index].Lang != obj2.GranularMarkings[index].Lang {
				problemsFound++
				str := fmt.Sprintf("-- The Language values do not match: %s | %s", obj1.GranularMarkings[index].Lang, obj2.GranularMarkings[index].Lang)
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ The Languages values match: %s | %s", obj1.GranularMarkings[index].Lang, obj2.GranularMarkings[index].Lang)
				resultDetails = append(resultDetails, str)
			}

			// Check Granular Marking Refs
			if obj1.GranularMarkings[index].MarkingRef != obj2.GranularMarkings[index].MarkingRef {
				problemsFound++
				str := fmt.Sprintf("-- The Ref values do not match: %s | %s", obj1.GranularMarkings[index].MarkingRef, obj2.GranularMarkings[index].MarkingRef)
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ The Ref values match: %s | %s", obj1.GranularMarkings[index].MarkingRef, obj2.GranularMarkings[index].MarkingRef)
				resultDetails = append(resultDetails, str)
			}

			// Check Granular Marking Selectors
			if len(obj1.GranularMarkings[index].Selectors) != len(obj2.GranularMarkings[index].Selectors) {
				problemsFound++
				str := fmt.Sprintf("-- The number of entries in Selectors do not match: %d | %d", len(obj1.GranularMarkings[index].Selectors), len(obj2.GranularMarkings[index].Selectors))
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ The number of entries in Selectors match: %d | %d", len(obj1.GranularMarkings[index].Selectors), len(obj2.GranularMarkings[index].Selectors))
				resultDetails = append(resultDetails, str)

				// If lengths are the same, then check each value
				for j := range obj1.GranularMarkings[index].Selectors {
					if obj1.GranularMarkings[index].Selectors[j] != obj2.GranularMarkings[index].Selectors[j] {
						problemsFound++
						str := fmt.Sprintf("-- The Selector values do not match: %s | %s", obj1.GranularMarkings[index].Selectors[j], obj2.GranularMarkings[index].Selectors[j])
						resultDetails = append(resultDetails, str)
					} else {
						str := fmt.Sprintf("++ The Selector values match: %s | %s", obj1.GranularMarkings[index].Selectors[j], obj2.GranularMarkings[index].Selectors[j])
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
