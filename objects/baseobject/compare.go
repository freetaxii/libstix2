// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package baseobject

import (
	"fmt"
)

/*
Compare - This method will compare the common properties from two objects to
make sure they are the same. The common properties receiver is object 1 and the
common properties passed in is object 2. This method will return an integer that
tracks the number of problems and a slice of strings that contain the detailed
results, whether good or bad.
*/
func (o *CommonObjectProperties) Compare(obj2 *CommonObjectProperties) (bool, int, []string) {
	return Compare(o, obj2)
}

/*
Compare - This function will compare the common properties from two objects
(object 1 and object 2) to make sure they are the same. This function will
return an integer that tracks the number of problems and a slice of strings that
contain the detailed results, whether good or bad.
*/
func Compare(obj1, obj2 *CommonObjectProperties) (bool, int, []string) {
	problemsFound := 0
	errorDetails := make([]string, 0)

	// Check Type Value
	if obj1.ObjectType != obj2.ObjectType {
		problemsFound++
		str := fmt.Sprintf("-- Types Do Not Match: %s | %s", obj1.ObjectType, obj2.ObjectType)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Types Match: %s | %s", obj1.ObjectType, obj2.ObjectType)
		errorDetails = append(errorDetails, str)
	}

	// Check Spec Version Value
	if obj1.SpecVersion != obj2.SpecVersion {
		problemsFound++
		str := fmt.Sprintf("-- Spec Versions Do Not Match: %s | %s", obj1.SpecVersion, obj2.SpecVersion)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Spec Versions Match: %s | %s", obj1.SpecVersion, obj2.SpecVersion)
		errorDetails = append(errorDetails, str)
	}

	// Check ID Value
	if obj1.ID != obj2.ID {
		problemsFound++
		str := fmt.Sprintf("-- IDs Do Not Match: %s | %s", obj1.ID, obj2.ID)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ IDs Match: %s | %s", obj1.ID, obj2.ID)
		errorDetails = append(errorDetails, str)
	}

	// Check Created By Ref Value
	if obj1.CreatedByRef != obj2.CreatedByRef {
		problemsFound++
		str := fmt.Sprintf("-- Created By Refs Do Not Match: %s | %s", obj1.CreatedByRef, obj2.CreatedByRef)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Created By Refs Match: %s | %s", obj1.CreatedByRef, obj2.CreatedByRef)
		errorDetails = append(errorDetails, str)
	}

	// Check Created Value
	if obj1.Created != obj2.Created {
		problemsFound++
		str := fmt.Sprintf("-- Created Dates Do Not Match: %s | %s", obj1.Created, obj2.Created)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Created Dates Match: %s | %s", obj1.Created, obj2.Created)
		errorDetails = append(errorDetails, str)
	}

	// Check Modified Value
	if obj1.Modified != obj2.Modified {
		problemsFound++
		str := fmt.Sprintf("-- Modified Dates Do Not Match: %s | %s", obj1.Modified, obj2.Modified)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Modified Dates Match: %s | %s", obj1.Modified, obj2.Modified)
		errorDetails = append(errorDetails, str)
	}

	// Check Revoked Value
	if obj1.Revoked != obj2.Revoked {
		problemsFound++
		str := fmt.Sprintf("-- Revoked Values Do Not Match: %t | %t", obj1.Revoked, obj2.Revoked)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Revoked Values Match: %t | %t", obj1.Revoked, obj2.Revoked)
		errorDetails = append(errorDetails, str)
	}

	// Check Labels Values
	if len(obj1.Labels) != len(obj2.Labels) {
		problemsFound++
		str := fmt.Sprintf("-- Labels Length Do Not Match: %d | %d", len(obj1.Labels), len(obj2.Labels))
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Labels Length Match: %d | %d", len(obj1.Labels), len(obj2.Labels))
		errorDetails = append(errorDetails, str)

		// If lengths are the same, then check each value
		for index := range obj1.Labels {
			if obj1.Labels[index] != obj2.Labels[index] {
				problemsFound++
				str := fmt.Sprintf("-- Labels Do Not Match: %s | %s", obj1.Labels[index], obj2.Labels[index])
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ Labels Match: %s | %s", obj1.Labels[index], obj2.Labels[index])
				errorDetails = append(errorDetails, str)
			}
		}
	}

	// Check Confidence Value
	if obj1.Confidence != obj2.Confidence {
		problemsFound++
		str := fmt.Sprintf("-- Confidence Values Do Not Match: %d | %d", obj1.Confidence, obj2.Confidence)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Confidence Values Match: %d | %d", obj1.Confidence, obj2.Confidence)
		errorDetails = append(errorDetails, str)
	}

	// Check Lang Value
	if obj1.Lang != obj2.Lang {
		problemsFound++
		str := fmt.Sprintf("-- Lang Values Do Not Match: %s | %s", obj1.Lang, obj2.Lang)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Lang Values Match: %s | %s", obj1.Lang, obj2.Lang)
		errorDetails = append(errorDetails, str)
	}

	// Check External References
	if len(obj1.ExternalReferences) != len(obj2.ExternalReferences) {
		problemsFound++
		str := fmt.Sprintf("-- External References Length Do Not Match: %d | %d", len(obj1.ExternalReferences), len(obj2.ExternalReferences))
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ External References Length Match: %d | %d", len(obj1.ExternalReferences), len(obj2.ExternalReferences))
		errorDetails = append(errorDetails, str)
		for index := range obj1.ExternalReferences {

			// Check External Reference Source Name
			if obj1.ExternalReferences[index].SourceName != obj2.ExternalReferences[index].SourceName {
				problemsFound++
				str := fmt.Sprintf("-- Source Name Do Not Match: %s | %s", obj1.ExternalReferences[index].SourceName, obj2.ExternalReferences[index].SourceName)
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ Source Name Match: %s | %s", obj1.ExternalReferences[index].SourceName, obj2.ExternalReferences[index].SourceName)
				errorDetails = append(errorDetails, str)
			}

			// Check External Reference Descriptions
			if obj1.ExternalReferences[index].Description != obj2.ExternalReferences[index].Description {
				problemsFound++
				str := fmt.Sprintf("-- Descriptions Do Not Match: %s | %s", obj1.ExternalReferences[index].Description, obj2.ExternalReferences[index].Description)
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ Descriptions Match: %s | %s", obj1.ExternalReferences[index].Description, obj2.ExternalReferences[index].Description)
				errorDetails = append(errorDetails, str)
			}

			// Check External Reference URLs
			if obj1.ExternalReferences[index].URL != obj2.ExternalReferences[index].URL {
				problemsFound++
				str := fmt.Sprintf("-- URLs Do Not Match: %s | %s", obj1.ExternalReferences[index].URL, obj2.ExternalReferences[index].URL)
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ URLs Match: %s | %s", obj1.ExternalReferences[index].URL, obj2.ExternalReferences[index].URL)
				errorDetails = append(errorDetails, str)
			}

			// Check External Reference Hashes
			if len(obj1.ExternalReferences[index].Hashes) != len(obj2.ExternalReferences[index].Hashes) {
				problemsFound++
				str := fmt.Sprintf("-- Hashes Length Do Not Match: %d | %d", len(obj1.ExternalReferences[index].Hashes), len(obj2.ExternalReferences[index].Hashes))
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ Hashes Length Match: %d | %d", len(obj1.ExternalReferences[index].Hashes), len(obj2.ExternalReferences[index].Hashes))
				errorDetails = append(errorDetails, str)

				// If lengths are the same, then check each value
				for key := range obj1.ExternalReferences[index].Hashes {
					if obj1.ExternalReferences[index].Hashes[key] != obj2.ExternalReferences[index].Hashes[key] {
						problemsFound++
						str := fmt.Sprintf("-- Hashes Do Not Match: %s | %s", obj1.ExternalReferences[index].Hashes[key], obj2.ExternalReferences[index].Hashes[key])
						errorDetails = append(errorDetails, str)
					} else {
						str := fmt.Sprintf("++ Hashes Match: %s | %s", obj1.ExternalReferences[index].Hashes[key], obj2.ExternalReferences[index].Hashes[key])
						errorDetails = append(errorDetails, str)
					}
				}
			}

			// Check External Reference External IDs
			if obj1.ExternalReferences[index].ExternalID != obj2.ExternalReferences[index].ExternalID {
				problemsFound++
				str := fmt.Sprintf("-- External IDs Do Not Match: %s | %s", obj1.ExternalReferences[index].ExternalID, obj2.ExternalReferences[index].ExternalID)
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ External IDs Match: %s | %s", obj1.ExternalReferences[index].ExternalID, obj2.ExternalReferences[index].ExternalID)
				errorDetails = append(errorDetails, str)
			}
		}
	}

	// Check Object Marking Refs
	if len(obj1.ObjectMarkingRefs) != len(obj2.ObjectMarkingRefs) {
		problemsFound++
		str := fmt.Sprintf("-- Object Marking Refs Length Do Not Match: %d | %d", len(obj1.ObjectMarkingRefs), len(obj2.ObjectMarkingRefs))
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Object Marking Refs Length Match: %d | %d", len(obj1.ObjectMarkingRefs), len(obj2.ObjectMarkingRefs))
		errorDetails = append(errorDetails, str)

		// If lengths are the same, then check each value
		for index := range obj1.ObjectMarkingRefs {
			if obj1.ObjectMarkingRefs[index] != obj2.ObjectMarkingRefs[index] {
				problemsFound++
				str := fmt.Sprintf("-- Object Marking Refs Do Not Match: %s | %s", obj1.ObjectMarkingRefs[index], obj2.ObjectMarkingRefs[index])
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ Object Marking Refs Match: %s | %s", obj1.ObjectMarkingRefs[index], obj2.ObjectMarkingRefs[index])
				errorDetails = append(errorDetails, str)
			}
		}
	}

	// Check Granular Markings
	if len(obj1.GranularMarkings) != len(obj2.GranularMarkings) {
		problemsFound++
		str := fmt.Sprintf("-- Granular Markings Length Do Not Match: %d | %d", len(obj1.GranularMarkings), len(obj2.GranularMarkings))
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Granular Markings Length Match: %d | %d", len(obj1.GranularMarkings), len(obj2.GranularMarkings))
		errorDetails = append(errorDetails, str)
		for index := range obj1.GranularMarkings {

			// Check Granular Marking Languages
			if obj1.GranularMarkings[index].Lang != obj2.GranularMarkings[index].Lang {
				problemsFound++
				str := fmt.Sprintf("-- Languages Do Not Match: %s | %s", obj1.GranularMarkings[index].Lang, obj2.GranularMarkings[index].Lang)
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ Languages Match: %s | %s", obj1.GranularMarkings[index].Lang, obj2.GranularMarkings[index].Lang)
				errorDetails = append(errorDetails, str)
			}

			// Check Granular Marking Refs
			if obj1.GranularMarkings[index].MarkingRef != obj2.GranularMarkings[index].MarkingRef {
				problemsFound++
				str := fmt.Sprintf("-- Refs Do Not Match: %s | %s", obj1.GranularMarkings[index].MarkingRef, obj2.GranularMarkings[index].MarkingRef)
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ Refs Match: %s | %s", obj1.GranularMarkings[index].MarkingRef, obj2.GranularMarkings[index].MarkingRef)
				errorDetails = append(errorDetails, str)
			}

			// Check Granular Marking Selectors
			if len(obj1.GranularMarkings[index].Selectors) != len(obj2.GranularMarkings[index].Selectors) {
				problemsFound++
				str := fmt.Sprintf("-- Selectors Length Do Not Match: %d | %d", len(obj1.GranularMarkings[index].Selectors), len(obj2.GranularMarkings[index].Selectors))
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ Selectors Length Match: %d | %d", len(obj1.GranularMarkings[index].Selectors), len(obj2.GranularMarkings[index].Selectors))
				errorDetails = append(errorDetails, str)

				// If lengths are the same, then check each value
				for j := range obj1.GranularMarkings[index].Selectors {
					if obj1.GranularMarkings[index].Selectors[j] != obj2.GranularMarkings[index].Selectors[j] {
						problemsFound++
						str := fmt.Sprintf("-- Selectors Do Not Match: %s | %s", obj1.GranularMarkings[index].Selectors[j], obj2.GranularMarkings[index].Selectors[j])
						errorDetails = append(errorDetails, str)
					} else {
						str := fmt.Sprintf("++ Selectors Match: %s | %s", obj1.GranularMarkings[index].Selectors[j], obj2.GranularMarkings[index].Selectors[j])
						errorDetails = append(errorDetails, str)
					}
				}
			}
		}
	}

	if problemsFound > 0 {
		return false, problemsFound, errorDetails
	}

	return true, 0, errorDetails
}
