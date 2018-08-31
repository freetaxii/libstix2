// Copyright 2018 Bret Jordan, All rights reserved.
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
represent the correct data, the common properties that are passed in as b
represents the one we need to test.
*/
func (o *CommonObjectProperties) Compare(b *CommonObjectProperties) (bool, int, []string) {
	problemsFound := 0
	details := make([]string, 0)

	// Check Type Value
	if b.ObjectType != o.ObjectType {
		problemsFound++
		str := fmt.Sprintf("-- Types Do Not Match: %s | %s", o.ObjectType, b.ObjectType)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Types Match: %s | %s", o.ObjectType, b.ObjectType)
		details = append(details, str)
	}

	// Check Spec Version Value
	if b.SpecVersion != o.SpecVersion {
		problemsFound++
		str := fmt.Sprintf("-- Spec Versions Do Not Match: %s | %s", o.SpecVersion, b.SpecVersion)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Spec Versions Match: %s | %s", o.SpecVersion, b.SpecVersion)
		details = append(details, str)
	}

	// Check ID Value
	if b.ID != o.ID {
		problemsFound++
		str := fmt.Sprintf("-- IDs Do Not Match: %s | %s", o.ID, b.ID)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ IDs Match: %s | %s", o.ID, b.ID)
		details = append(details, str)
	}

	// Check Created By Ref Value
	if b.CreatedByRef != o.CreatedByRef {
		problemsFound++
		str := fmt.Sprintf("-- Created By Refs Do Not Match: %s | %s", o.CreatedByRef, b.CreatedByRef)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Created By Refs Match: %s | %s", o.CreatedByRef, b.CreatedByRef)
		details = append(details, str)
	}

	// Check Created Value
	if b.Created != o.Created {
		problemsFound++
		str := fmt.Sprintf("-- Created Dates Do Not Match: %s | %s", o.Created, b.Created)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Created Dates Match: %s | %s", o.Created, b.Created)
		details = append(details, str)
	}

	// Check Modified Value
	if b.Modified != o.Modified {
		problemsFound++
		str := fmt.Sprintf("-- Modified Dates Do Not Match: %s | %s", o.Modified, b.Modified)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Modified Dates Match: %s | %s", o.Modified, b.Modified)
		details = append(details, str)
	}

	// Check Revoked Value
	if b.Revoked != o.Revoked {
		problemsFound++
		str := fmt.Sprintf("-- Revoked Values Do Not Match: %t | %t", o.Revoked, b.Revoked)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Revoked Values Match: %t | %t", o.Revoked, b.Revoked)
		details = append(details, str)
	}

	// Check Labels Values
	if len(b.Labels) != len(o.Labels) {
		problemsFound++
		str := fmt.Sprintf("-- Labels Length Do Not Match: %d | %d", len(o.Labels), len(b.Labels))
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Labels Length Match: %d | %d", len(o.Labels), len(b.Labels))
		details = append(details, str)

		// If lengths are the same, then check each value
		for index, _ := range o.Labels {
			if b.Labels[index] != o.Labels[index] {
				problemsFound++
				str := fmt.Sprintf("-- Labels Do Not Match: %s | %s", o.Labels[index], b.Labels[index])
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Labels Match: %s | %s", o.Labels[index], b.Labels[index])
				details = append(details, str)
			}
		}
	}

	// Check Confidence Value
	if b.Confidence != o.Confidence {
		problemsFound++
		str := fmt.Sprintf("-- Confidence Values Do Not Match: %d | %d", o.Confidence, b.Confidence)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Confidence Values Match: %d | %d", o.Confidence, b.Confidence)
		details = append(details, str)
	}

	// Check Lang Value
	if b.Lang != o.Lang {
		problemsFound++
		str := fmt.Sprintf("-- Lang Values Do Not Match: %s | %s", o.Lang, b.Lang)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Lang Values Match: %s | %s", o.Lang, b.Lang)
		details = append(details, str)
	}

	// Check External References
	if len(b.ExternalReferences) != len(o.ExternalReferences) {
		problemsFound++
		str := fmt.Sprintf("-- External References Length Do Not Match: %d | %d", len(o.ExternalReferences), len(b.ExternalReferences))
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ External References Length Match: %d | %d", len(o.ExternalReferences), len(b.ExternalReferences))
		details = append(details, str)
		for index, _ := range o.ExternalReferences {

			// Check External Reference Source Name
			if b.ExternalReferences[index].SourceName != o.ExternalReferences[index].SourceName {
				problemsFound++
				str := fmt.Sprintf("-- Source Name Do Not Match: %s | %s", o.ExternalReferences[index].SourceName, b.ExternalReferences[index].SourceName)
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Source Name Match: %s | %s", o.ExternalReferences[index].SourceName, b.ExternalReferences[index].SourceName)
				details = append(details, str)
			}

			// Check External Reference Descriptions
			if b.ExternalReferences[index].Description != o.ExternalReferences[index].Description {
				problemsFound++
				str := fmt.Sprintf("-- Descriptions Do Not Match: %s | %s", o.ExternalReferences[index].Description, b.ExternalReferences[index].Description)
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Descriptions Match: %s | %s", o.ExternalReferences[index].Description, b.ExternalReferences[index].Description)
				details = append(details, str)
			}

			// Check External Reference URLs
			if b.ExternalReferences[index].URL != o.ExternalReferences[index].URL {
				problemsFound++
				str := fmt.Sprintf("-- URLs Do Not Match: %s | %s", o.ExternalReferences[index].URL, b.ExternalReferences[index].URL)
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ URLs Match: %s | %s", o.ExternalReferences[index].URL, b.ExternalReferences[index].URL)
				details = append(details, str)
			}

			// Check External Reference Hashes
			if len(b.ExternalReferences[index].Hashes) != len(o.ExternalReferences[index].Hashes) {
				problemsFound++
				str := fmt.Sprintf("-- Hashes Length Do Not Match: %d | %d", len(o.ExternalReferences[index].Hashes), len(b.ExternalReferences[index].Hashes))
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Hashes Length Match: %d | %d", len(o.ExternalReferences[index].Hashes), len(b.ExternalReferences[index].Hashes))
				details = append(details, str)

				// If lengths are the same, then check each value
				for key, _ := range o.ExternalReferences[index].Hashes {
					if b.ExternalReferences[index].Hashes[key] != o.ExternalReferences[index].Hashes[key] {
						problemsFound++
						str := fmt.Sprintf("-- Hashes Do Not Match: %s | %s", o.ExternalReferences[index].Hashes[key], b.ExternalReferences[index].Hashes[key])
						details = append(details, str)
					} else {
						str := fmt.Sprintf("++ Hashes Match: %s | %s", o.ExternalReferences[index].Hashes[key], b.ExternalReferences[index].Hashes[key])
						details = append(details, str)
					}
				}
			}

			// Check External Reference External IDs
			if b.ExternalReferences[index].ExternalID != o.ExternalReferences[index].ExternalID {
				problemsFound++
				str := fmt.Sprintf("-- External IDs Do Not Match: %s | %s", o.ExternalReferences[index].ExternalID, b.ExternalReferences[index].ExternalID)
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ External IDs Match: %s | %s", o.ExternalReferences[index].ExternalID, b.ExternalReferences[index].ExternalID)
				details = append(details, str)
			}
		}
	}

	// Check Object Marking Refs
	if len(b.ObjectMarkingRefs) != len(o.ObjectMarkingRefs) {
		problemsFound++
		str := fmt.Sprintf("-- Object Marking Refs Length Do Not Match: %d | %d", len(o.ObjectMarkingRefs), len(b.ObjectMarkingRefs))
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Object Marking Refs Length Match: %d | %d", len(o.ObjectMarkingRefs), len(b.ObjectMarkingRefs))
		details = append(details, str)

		// If lengths are the same, then check each value
		for index, _ := range o.ObjectMarkingRefs {
			if b.ObjectMarkingRefs[index] != o.ObjectMarkingRefs[index] {
				problemsFound++
				str := fmt.Sprintf("-- Object Marking Refs Do Not Match: %s | %s", o.ObjectMarkingRefs[index], b.ObjectMarkingRefs[index])
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Object Marking Refs Match: %s | %s", o.ObjectMarkingRefs[index], b.ObjectMarkingRefs[index])
				details = append(details, str)
			}
		}
	}

	// Check Granular Markings
	if len(b.GranularMarkings) != len(o.GranularMarkings) {
		problemsFound++
		str := fmt.Sprintf("-- Granular Markings Length Do Not Match: %d | %d", len(o.GranularMarkings), len(b.GranularMarkings))
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Granular Markings Length Match: %d | %d", len(o.GranularMarkings), len(b.GranularMarkings))
		details = append(details, str)
		for index, _ := range o.GranularMarkings {

			// Check Granular Marking Languages
			if b.GranularMarkings[index].Lang != o.GranularMarkings[index].Lang {
				problemsFound++
				str := fmt.Sprintf("-- Languages Do Not Match: %s | %s", o.GranularMarkings[index].Lang, b.GranularMarkings[index].Lang)
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Languages Match: %s | %s", o.GranularMarkings[index].Lang, b.GranularMarkings[index].Lang)
				details = append(details, str)
			}

			// Check Granular Marking Refs
			if b.GranularMarkings[index].MarkingRef != o.GranularMarkings[index].MarkingRef {
				problemsFound++
				str := fmt.Sprintf("-- Refs Do Not Match: %s | %s", o.GranularMarkings[index].MarkingRef, b.GranularMarkings[index].MarkingRef)
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Refs Match: %s | %s", o.GranularMarkings[index].MarkingRef, b.GranularMarkings[index].MarkingRef)
				details = append(details, str)
			}

			// Check Granular Marking Selectors
			if len(b.GranularMarkings[index].Selectors) != len(o.GranularMarkings[index].Selectors) {
				problemsFound++
				str := fmt.Sprintf("-- Selectors Length Do Not Match: %d | %d", len(o.GranularMarkings[index].Selectors), len(b.GranularMarkings[index].Selectors))
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Selectors Length Match: %d | %d", len(o.GranularMarkings[index].Selectors), len(b.GranularMarkings[index].Selectors))
				details = append(details, str)

				// If lengths are the same, then check each value
				for j, _ := range o.GranularMarkings[index].Selectors {
					if b.GranularMarkings[index].Selectors[j] != o.GranularMarkings[index].Selectors[j] {
						problemsFound++
						str := fmt.Sprintf("-- Selectors Do Not Match: %s | %s", o.GranularMarkings[index].Selectors[j], b.GranularMarkings[index].Selectors[j])
						details = append(details, str)
					} else {
						str := fmt.Sprintf("++ Selectors Match: %s | %s", o.GranularMarkings[index].Selectors[j], b.GranularMarkings[index].Selectors[j])
						details = append(details, str)
					}
				}
			}
		}
	}

	if problemsFound > 0 {
		return false, problemsFound, details
	}

	return true, 0, details
}
