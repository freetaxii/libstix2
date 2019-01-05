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
make sure they are the same. The common properties receiver is the master and
represent the correct data, the common properties that are passed in as toTest
represents the one we need to test.
*/
func (o *CommonObjectProperties) Compare(toTest *CommonObjectProperties) (bool, int, []string) {
	return Compare(o, toTest)
}

/*
Compare - This method will compare the common properties from two objects to
make sure they are the same. The common properties correct is the master and
represent the correct data, the common properties toTest represents the one we
need to test.
*/
func Compare(correct, toTest *CommonObjectProperties) (bool, int, []string) {
	problemsFound := 0
	errorDetails := make([]string, 0)

	// Check Type Value
	if toTest.ObjectType != correct.ObjectType {
		problemsFound++
		str := fmt.Sprintf("-- Types Do Not Match: %s | %s", correct.ObjectType, toTest.ObjectType)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Types Match: %s | %s", correct.ObjectType, toTest.ObjectType)
		errorDetails = append(errorDetails, str)
	}

	// Check Spec Version Value
	if toTest.SpecVersion != correct.SpecVersion {
		problemsFound++
		str := fmt.Sprintf("-- Spec Versions Do Not Match: %s | %s", correct.SpecVersion, toTest.SpecVersion)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Spec Versions Match: %s | %s", correct.SpecVersion, toTest.SpecVersion)
		errorDetails = append(errorDetails, str)
	}

	// Check ID Value
	if toTest.ID != correct.ID {
		problemsFound++
		str := fmt.Sprintf("-- IDs Do Not Match: %s | %s", correct.ID, toTest.ID)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ IDs Match: %s | %s", correct.ID, toTest.ID)
		errorDetails = append(errorDetails, str)
	}

	// Check Created By Ref Value
	if toTest.CreatedByRef != correct.CreatedByRef {
		problemsFound++
		str := fmt.Sprintf("-- Created By Refs Do Not Match: %s | %s", correct.CreatedByRef, toTest.CreatedByRef)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Created By Refs Match: %s | %s", correct.CreatedByRef, toTest.CreatedByRef)
		errorDetails = append(errorDetails, str)
	}

	// Check Created Value
	if toTest.Created != correct.Created {
		problemsFound++
		str := fmt.Sprintf("-- Created Dates Do Not Match: %s | %s", correct.Created, toTest.Created)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Created Dates Match: %s | %s", correct.Created, toTest.Created)
		errorDetails = append(errorDetails, str)
	}

	// Check Modified Value
	if toTest.Modified != correct.Modified {
		problemsFound++
		str := fmt.Sprintf("-- Modified Dates Do Not Match: %s | %s", correct.Modified, toTest.Modified)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Modified Dates Match: %s | %s", correct.Modified, toTest.Modified)
		errorDetails = append(errorDetails, str)
	}

	// Check Revoked Value
	if toTest.Revoked != correct.Revoked {
		problemsFound++
		str := fmt.Sprintf("-- Revoked Values Do Not Match: %t | %t", correct.Revoked, toTest.Revoked)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Revoked Values Match: %t | %t", correct.Revoked, toTest.Revoked)
		errorDetails = append(errorDetails, str)
	}

	// Check Labels Values
	if len(toTest.Labels) != len(correct.Labels) {
		problemsFound++
		str := fmt.Sprintf("-- Labels Length Do Not Match: %d | %d", len(correct.Labels), len(toTest.Labels))
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Labels Length Match: %d | %d", len(correct.Labels), len(toTest.Labels))
		errorDetails = append(errorDetails, str)

		// If lengths are the same, then check each value
		for index, _ := range correct.Labels {
			if toTest.Labels[index] != correct.Labels[index] {
				problemsFound++
				str := fmt.Sprintf("-- Labels Do Not Match: %s | %s", correct.Labels[index], toTest.Labels[index])
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ Labels Match: %s | %s", correct.Labels[index], toTest.Labels[index])
				errorDetails = append(errorDetails, str)
			}
		}
	}

	// Check Confidence Value
	if toTest.Confidence != correct.Confidence {
		problemsFound++
		str := fmt.Sprintf("-- Confidence Values Do Not Match: %d | %d", correct.Confidence, toTest.Confidence)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Confidence Values Match: %d | %d", correct.Confidence, toTest.Confidence)
		errorDetails = append(errorDetails, str)
	}

	// Check Lang Value
	if toTest.Lang != correct.Lang {
		problemsFound++
		str := fmt.Sprintf("-- Lang Values Do Not Match: %s | %s", correct.Lang, toTest.Lang)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Lang Values Match: %s | %s", correct.Lang, toTest.Lang)
		errorDetails = append(errorDetails, str)
	}

	// Check External References
	if len(toTest.ExternalReferences) != len(correct.ExternalReferences) {
		problemsFound++
		str := fmt.Sprintf("-- External References Length Do Not Match: %d | %d", len(correct.ExternalReferences), len(toTest.ExternalReferences))
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ External References Length Match: %d | %d", len(correct.ExternalReferences), len(toTest.ExternalReferences))
		errorDetails = append(errorDetails, str)
		for index, _ := range correct.ExternalReferences {

			// Check External Reference Source Name
			if toTest.ExternalReferences[index].SourceName != correct.ExternalReferences[index].SourceName {
				problemsFound++
				str := fmt.Sprintf("-- Source Name Do Not Match: %s | %s", correct.ExternalReferences[index].SourceName, toTest.ExternalReferences[index].SourceName)
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ Source Name Match: %s | %s", correct.ExternalReferences[index].SourceName, toTest.ExternalReferences[index].SourceName)
				errorDetails = append(errorDetails, str)
			}

			// Check External Reference Descriptions
			if toTest.ExternalReferences[index].Description != correct.ExternalReferences[index].Description {
				problemsFound++
				str := fmt.Sprintf("-- Descriptions Do Not Match: %s | %s", correct.ExternalReferences[index].Description, toTest.ExternalReferences[index].Description)
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ Descriptions Match: %s | %s", correct.ExternalReferences[index].Description, toTest.ExternalReferences[index].Description)
				errorDetails = append(errorDetails, str)
			}

			// Check External Reference URLs
			if toTest.ExternalReferences[index].URL != correct.ExternalReferences[index].URL {
				problemsFound++
				str := fmt.Sprintf("-- URLs Do Not Match: %s | %s", correct.ExternalReferences[index].URL, toTest.ExternalReferences[index].URL)
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ URLs Match: %s | %s", correct.ExternalReferences[index].URL, toTest.ExternalReferences[index].URL)
				errorDetails = append(errorDetails, str)
			}

			// Check External Reference Hashes
			if len(toTest.ExternalReferences[index].Hashes) != len(correct.ExternalReferences[index].Hashes) {
				problemsFound++
				str := fmt.Sprintf("-- Hashes Length Do Not Match: %d | %d", len(correct.ExternalReferences[index].Hashes), len(toTest.ExternalReferences[index].Hashes))
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ Hashes Length Match: %d | %d", len(correct.ExternalReferences[index].Hashes), len(toTest.ExternalReferences[index].Hashes))
				errorDetails = append(errorDetails, str)

				// If lengths are the same, then check each value
				for key, _ := range correct.ExternalReferences[index].Hashes {
					if toTest.ExternalReferences[index].Hashes[key] != correct.ExternalReferences[index].Hashes[key] {
						problemsFound++
						str := fmt.Sprintf("-- Hashes Do Not Match: %s | %s", correct.ExternalReferences[index].Hashes[key], toTest.ExternalReferences[index].Hashes[key])
						errorDetails = append(errorDetails, str)
					} else {
						str := fmt.Sprintf("++ Hashes Match: %s | %s", correct.ExternalReferences[index].Hashes[key], toTest.ExternalReferences[index].Hashes[key])
						errorDetails = append(errorDetails, str)
					}
				}
			}

			// Check External Reference External IDs
			if toTest.ExternalReferences[index].ExternalID != correct.ExternalReferences[index].ExternalID {
				problemsFound++
				str := fmt.Sprintf("-- External IDs Do Not Match: %s | %s", correct.ExternalReferences[index].ExternalID, toTest.ExternalReferences[index].ExternalID)
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ External IDs Match: %s | %s", correct.ExternalReferences[index].ExternalID, toTest.ExternalReferences[index].ExternalID)
				errorDetails = append(errorDetails, str)
			}
		}
	}

	// Check Object Marking Refs
	if len(toTest.ObjectMarkingRefs) != len(correct.ObjectMarkingRefs) {
		problemsFound++
		str := fmt.Sprintf("-- Object Marking Refs Length Do Not Match: %d | %d", len(correct.ObjectMarkingRefs), len(toTest.ObjectMarkingRefs))
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Object Marking Refs Length Match: %d | %d", len(correct.ObjectMarkingRefs), len(toTest.ObjectMarkingRefs))
		errorDetails = append(errorDetails, str)

		// If lengths are the same, then check each value
		for index, _ := range correct.ObjectMarkingRefs {
			if toTest.ObjectMarkingRefs[index] != correct.ObjectMarkingRefs[index] {
				problemsFound++
				str := fmt.Sprintf("-- Object Marking Refs Do Not Match: %s | %s", correct.ObjectMarkingRefs[index], toTest.ObjectMarkingRefs[index])
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ Object Marking Refs Match: %s | %s", correct.ObjectMarkingRefs[index], toTest.ObjectMarkingRefs[index])
				errorDetails = append(errorDetails, str)
			}
		}
	}

	// Check Granular Markings
	if len(toTest.GranularMarkings) != len(correct.GranularMarkings) {
		problemsFound++
		str := fmt.Sprintf("-- Granular Markings Length Do Not Match: %d | %d", len(correct.GranularMarkings), len(toTest.GranularMarkings))
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Granular Markings Length Match: %d | %d", len(correct.GranularMarkings), len(toTest.GranularMarkings))
		errorDetails = append(errorDetails, str)
		for index, _ := range correct.GranularMarkings {

			// Check Granular Marking Languages
			if toTest.GranularMarkings[index].Lang != correct.GranularMarkings[index].Lang {
				problemsFound++
				str := fmt.Sprintf("-- Languages Do Not Match: %s | %s", correct.GranularMarkings[index].Lang, toTest.GranularMarkings[index].Lang)
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ Languages Match: %s | %s", correct.GranularMarkings[index].Lang, toTest.GranularMarkings[index].Lang)
				errorDetails = append(errorDetails, str)
			}

			// Check Granular Marking Refs
			if toTest.GranularMarkings[index].MarkingRef != correct.GranularMarkings[index].MarkingRef {
				problemsFound++
				str := fmt.Sprintf("-- Refs Do Not Match: %s | %s", correct.GranularMarkings[index].MarkingRef, toTest.GranularMarkings[index].MarkingRef)
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ Refs Match: %s | %s", correct.GranularMarkings[index].MarkingRef, toTest.GranularMarkings[index].MarkingRef)
				errorDetails = append(errorDetails, str)
			}

			// Check Granular Marking Selectors
			if len(toTest.GranularMarkings[index].Selectors) != len(correct.GranularMarkings[index].Selectors) {
				problemsFound++
				str := fmt.Sprintf("-- Selectors Length Do Not Match: %d | %d", len(correct.GranularMarkings[index].Selectors), len(toTest.GranularMarkings[index].Selectors))
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ Selectors Length Match: %d | %d", len(correct.GranularMarkings[index].Selectors), len(toTest.GranularMarkings[index].Selectors))
				errorDetails = append(errorDetails, str)

				// If lengths are the same, then check each value
				for j, _ := range correct.GranularMarkings[index].Selectors {
					if toTest.GranularMarkings[index].Selectors[j] != correct.GranularMarkings[index].Selectors[j] {
						problemsFound++
						str := fmt.Sprintf("-- Selectors Do Not Match: %s | %s", correct.GranularMarkings[index].Selectors[j], toTest.GranularMarkings[index].Selectors[j])
						errorDetails = append(errorDetails, str)
					} else {
						str := fmt.Sprintf("++ Selectors Match: %s | %s", correct.GranularMarkings[index].Selectors[j], toTest.GranularMarkings[index].Selectors[j])
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
