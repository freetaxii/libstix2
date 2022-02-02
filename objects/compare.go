// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

import "fmt"

// Compare - This function will compare two objects to make sure they are the
// same and will return a boolean, an integer that tracks the number of
// problems found, and a slice of strings that contain the detailed results,
// whether good or bad.
func Compare(o, obj2 *CommonObjectProperties, debug bool) (bool, int, []string) {
	return o.Compare(obj2, debug)
}

// Compare - This method will compare two objects to make sure they are the
// same and will return a boolean, an integer that tracks the number of
// problems found, and a slice of strings that contain the detailed results,
// whether good or bad.
func (o *CommonObjectProperties) Compare(obj2 *CommonObjectProperties, debug bool) (bool, int, []string) {
	var r *results = new(results)
	r.debug = debug

	// Object Type
	if o.ObjectType != obj2.ObjectType {
		str := fmt.Sprintf("-- the type property values do not match: %s | %s", o.ObjectType, obj2.ObjectType)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the type property values match: %s | %s", o.ObjectType, obj2.ObjectType)
		logValid(r, str)
	}

	// Spec Version
	if o.SpecVersion != obj2.SpecVersion {
		str := fmt.Sprintf("-- the spec_version property values do not match: %s | %s", o.SpecVersion, obj2.SpecVersion)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the spec_version property values match: %s | %s", o.SpecVersion, obj2.SpecVersion)
		logValid(r, str)
	}

	// ID
	if o.ID != obj2.ID {
		str := fmt.Sprintf("-- the id property values do not match: %s | %s", o.ID, obj2.ID)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the id property values match: %s | %s", o.ID, obj2.ID)
		logValid(r, str)
	}

	// Check Created By Ref Value
	if o.CreatedByRef != obj2.CreatedByRef {
		str := fmt.Sprintf("-- The created by ref values do not match: %s | %s", o.CreatedByRef, obj2.CreatedByRef)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ The created by ref values match: %s | %s", o.CreatedByRef, obj2.CreatedByRef)
		logValid(r, str)
	}

	// Created
	if o.Created != obj2.Created {
		str := fmt.Sprintf("-- the created dates do not match: %s | %s", o.Created, obj2.Created)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the created dates match: %s | %s", o.Created, obj2.Created)
		logValid(r, str)
	}

	// Modified
	if o.Modified != obj2.Modified {
		str := fmt.Sprintf("-- the modified dates do not match: %s | %s", o.Modified, obj2.Modified)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the modified dates match: %s | %s", o.Modified, obj2.Modified)
		logValid(r, str)
	}

	// Revoked
	if o.Revoked != obj2.Revoked {
		str := fmt.Sprintf("-- the revoked values do not match: %t | %t", o.Revoked, obj2.Revoked)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the revoked values match: %t | %t", o.Revoked, obj2.Revoked)
		logValid(r, str)
	}

	// Labels
	if len(o.Labels) != len(obj2.Labels) {
		str := fmt.Sprintf("-- the number of entries in the labels property do not match: %d | %d", len(o.Labels), len(obj2.Labels))
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ the number of entries in the labels property match: %d | %d", len(o.Labels), len(obj2.Labels))
		logValid(r, str)

		// If lengths are the same, then check each value
		for index := range o.Labels {
			if o.Labels[index] != obj2.Labels[index] {
				str := fmt.Sprintf("-- the label values do not match: %s | %s", o.Labels[index], obj2.Labels[index])
				logProblem(r, str)
			} else {
				str := fmt.Sprintf("++ the label values match: %s | %s", o.Labels[index], obj2.Labels[index])
				logValid(r, str)
			}
		}
	}

	// Confidence
	if o.Confidence != obj2.Confidence {
		str := fmt.Sprintf("-- The confidence values do not match: %d | %d", o.Confidence, obj2.Confidence)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ The confidence values match: %d | %d", o.Confidence, obj2.Confidence)
		logValid(r, str)
	}

	// Lang
	if o.Lang != obj2.Lang {
		str := fmt.Sprintf("-- The lang values do not match: %s | %s", o.Lang, obj2.Lang)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ The lang values match: %s | %s", o.Lang, obj2.Lang)
		logValid(r, str)
	}

	// Check External References
	if len(o.ExternalReferences) != len(obj2.ExternalReferences) {
		str := fmt.Sprintf("-- The number of entries in external references do not match: %d | %d", len(o.ExternalReferences), len(obj2.ExternalReferences))
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ The number of entries in external references match: %d | %d", len(o.ExternalReferences), len(obj2.ExternalReferences))
		logValid(r, str)
		for index := range o.ExternalReferences {

			// Check External Reference Source Name
			if o.ExternalReferences[index].SourceName != obj2.ExternalReferences[index].SourceName {
				str := fmt.Sprintf("-- The source name values do not match: %s | %s", o.ExternalReferences[index].SourceName, obj2.ExternalReferences[index].SourceName)
				logProblem(r, str)
			} else {
				str := fmt.Sprintf("++ The source name values match: %s | %s", o.ExternalReferences[index].SourceName, obj2.ExternalReferences[index].SourceName)
				logValid(r, str)
			}

			// Check External Reference Descriptions
			if o.ExternalReferences[index].Description != obj2.ExternalReferences[index].Description {
				str := fmt.Sprintf("-- The description values do not match: %s | %s", o.ExternalReferences[index].Description, obj2.ExternalReferences[index].Description)
				logProblem(r, str)
			} else {
				str := fmt.Sprintf("++ The description values match: %s | %s", o.ExternalReferences[index].Description, obj2.ExternalReferences[index].Description)
				logValid(r, str)
			}

			// Check External Reference URLs
			if o.ExternalReferences[index].URL != obj2.ExternalReferences[index].URL {
				str := fmt.Sprintf("-- The url values do not match: %s | %s", o.ExternalReferences[index].URL, obj2.ExternalReferences[index].URL)
				logProblem(r, str)
			} else {
				str := fmt.Sprintf("++ The url values match: %s | %s", o.ExternalReferences[index].URL, obj2.ExternalReferences[index].URL)
				logValid(r, str)
			}

			// Check External Reference Hashes
			if len(o.ExternalReferences[index].Hashes) != len(obj2.ExternalReferences[index].Hashes) {
				str := fmt.Sprintf("-- The number of entries in hashes do not match: %d | %d", len(o.ExternalReferences[index].Hashes), len(obj2.ExternalReferences[index].Hashes))
				logProblem(r, str)
			} else {
				str := fmt.Sprintf("++ The number of entries in hashes match: %d | %d", len(o.ExternalReferences[index].Hashes), len(obj2.ExternalReferences[index].Hashes))
				logValid(r, str)

				// If lengths are the same, then check each value
				for key := range o.ExternalReferences[index].Hashes {
					if o.ExternalReferences[index].Hashes[key] != obj2.ExternalReferences[index].Hashes[key] {
						str := fmt.Sprintf("-- The hash values do not match: %s | %s", o.ExternalReferences[index].Hashes[key], obj2.ExternalReferences[index].Hashes[key])
						logProblem(r, str)
					} else {
						str := fmt.Sprintf("++ The hash values match: %s | %s", o.ExternalReferences[index].Hashes[key], obj2.ExternalReferences[index].Hashes[key])
						logValid(r, str)
					}
				}
			}

			// Check External Reference External IDs
			if o.ExternalReferences[index].ExternalID != obj2.ExternalReferences[index].ExternalID {
				str := fmt.Sprintf("-- The external id values do not match: %s | %s", o.ExternalReferences[index].ExternalID, obj2.ExternalReferences[index].ExternalID)
				logProblem(r, str)
			} else {
				str := fmt.Sprintf("++ The external id values match: %s | %s", o.ExternalReferences[index].ExternalID, obj2.ExternalReferences[index].ExternalID)
				logValid(r, str)
			}
		}
	}

	// Check Object Marking Refs
	if len(o.ObjectMarkingRefs) != len(obj2.ObjectMarkingRefs) {
		str := fmt.Sprintf("-- The number of entries in object marking refs do not match: %d | %d", len(o.ObjectMarkingRefs), len(obj2.ObjectMarkingRefs))
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ The number of entries in object marking refs match: %d | %d", len(o.ObjectMarkingRefs), len(obj2.ObjectMarkingRefs))
		logValid(r, str)

		// If lengths are the same, then check each value
		for index := range o.ObjectMarkingRefs {
			if o.ObjectMarkingRefs[index] != obj2.ObjectMarkingRefs[index] {
				str := fmt.Sprintf("-- The object marking ref values do not match: %s | %s", o.ObjectMarkingRefs[index], obj2.ObjectMarkingRefs[index])
				logProblem(r, str)
			} else {
				str := fmt.Sprintf("++ The object marking ref values match: %s | %s", o.ObjectMarkingRefs[index], obj2.ObjectMarkingRefs[index])
				logValid(r, str)
			}
		}
	}

	// Check Granular Markings
	if len(o.GranularMarkings) != len(obj2.GranularMarkings) {
		str := fmt.Sprintf("-- The number of entries in granular markings do not match: %d | %d", len(o.GranularMarkings), len(obj2.GranularMarkings))
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ The number of entries in granular markings match: %d | %d", len(o.GranularMarkings), len(obj2.GranularMarkings))
		logValid(r, str)

		for index := range o.GranularMarkings {

			// Check Granular Marking Languages
			if o.GranularMarkings[index].Lang != obj2.GranularMarkings[index].Lang {
				str := fmt.Sprintf("-- The language values do not match: %s | %s", o.GranularMarkings[index].Lang, obj2.GranularMarkings[index].Lang)
				logProblem(r, str)
			} else {
				str := fmt.Sprintf("++ The languages values match: %s | %s", o.GranularMarkings[index].Lang, obj2.GranularMarkings[index].Lang)
				logValid(r, str)
			}

			// Check Granular Marking Refs
			if o.GranularMarkings[index].MarkingRef != obj2.GranularMarkings[index].MarkingRef {
				str := fmt.Sprintf("-- The marking ref values do not match: %s | %s", o.GranularMarkings[index].MarkingRef, obj2.GranularMarkings[index].MarkingRef)
				logProblem(r, str)
			} else {
				str := fmt.Sprintf("++ The marking ref values match: %s | %s", o.GranularMarkings[index].MarkingRef, obj2.GranularMarkings[index].MarkingRef)
				logValid(r, str)
			}

			// Check Granular Marking Selectors
			if len(o.GranularMarkings[index].Selectors) != len(obj2.GranularMarkings[index].Selectors) {
				str := fmt.Sprintf("-- The number of entries in selectors do not match: %d | %d", len(o.GranularMarkings[index].Selectors), len(obj2.GranularMarkings[index].Selectors))
				logProblem(r, str)
			} else {
				str := fmt.Sprintf("++ The number of entries in selectors match: %d | %d", len(o.GranularMarkings[index].Selectors), len(obj2.GranularMarkings[index].Selectors))
				logValid(r, str)

				// If lengths are the same, then check each value
				for j := range o.GranularMarkings[index].Selectors {
					if o.GranularMarkings[index].Selectors[j] != obj2.GranularMarkings[index].Selectors[j] {
						str := fmt.Sprintf("-- The selector values do not match: %s | %s", o.GranularMarkings[index].Selectors[j], obj2.GranularMarkings[index].Selectors[j])
						logProblem(r, str)
					} else {
						str := fmt.Sprintf("++ The selector values match: %s | %s", o.GranularMarkings[index].Selectors[j], obj2.GranularMarkings[index].Selectors[j])
						logValid(r, str)
					}
				}
			}
		}
	}

	// Return real values not pointers
	if r.problemsFound > 0 {
		return false, r.problemsFound, r.resultDetails
	}
	return true, 0, r.resultDetails
}
