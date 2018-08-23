// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import (
	"fmt"

	"github.com/freetaxii/libstix2/defs"
)

// ----------------------------------------------------------------------
//
// Common Property Types - Used to populate the common object properties
//
// ----------------------------------------------------------------------

/*
CommonBundleProperties - This type includes all of the common properties
that are used by all STIX objects
*/
type CommonBundleProperties struct {
	TypeProperty
	IDProperty
}

/*
CommonObjectProperties - This type includes all of the common properties
that are used by all STIX SDOs and SROs
*/
type CommonObjectProperties struct {
	BaseProperties
	RevokedProperty
	LabelsProperty
	ConfidenceProperty
	LangProperty
	BaseExtendedProperties
}

/*
CommonMarkingDefinitionProperties - This type includes all of the common
properties that are used by the STIX Marking Definition object
*/
type CommonMarkingDefinitionProperties struct {
	BaseProperties
	BaseExtendedProperties
}

// ----------------------------------------------------------------------
//
// Public Methods - CommonObjectProperties
//
// ----------------------------------------------------------------------

/*
InitObject- This method will initialize the object by setting all of the basic
properties.
*/
func (p *CommonObjectProperties) InitObject(stixType string) error {
	// TODO make sure that the value coming in is a valid STIX object type
	p.SetSpecVersion(defs.STIX_VERSION)
	p.SetObjectType(stixType)
	p.SetNewID(stixType)
	p.SetCreatedToCurrentTime()
	p.SetModifiedToCreated()
	return nil
}

/*
Compare - This method will compare the common properties from two objects to
make sure they are the same. The common properties receiver is the master and
represent the correct data, the common properties that are passed in as b
represents the one we need to test.
*/
func (p *CommonObjectProperties) Compare(b *CommonObjectProperties) (bool, int, []string) {
	problemsFound := 0
	details := make([]string, 0)

	// Check Type Value
	if b.ObjectType != p.ObjectType {
		problemsFound++
		str := fmt.Sprintf("-- Types Do Not Match: %s | %s", p.ObjectType, b.ObjectType)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Types Match: %s | %s", p.ObjectType, b.ObjectType)
		details = append(details, str)
	}

	// Check Spec Version Value
	if b.SpecVersion != p.SpecVersion {
		problemsFound++
		str := fmt.Sprintf("-- Spec Versions Do Not Match: %s | %s", p.SpecVersion, b.SpecVersion)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Spec Versions Match: %s | %s", p.SpecVersion, b.SpecVersion)
		details = append(details, str)
	}

	// Check ID Value
	if b.ID != p.ID {
		problemsFound++
		str := fmt.Sprintf("-- IDs Do Not Match: %s | %s", p.ID, b.ID)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ IDs Match: %s | %s", p.ID, b.ID)
		details = append(details, str)
	}

	// Check Created By Ref Value
	if b.CreatedByRef != p.CreatedByRef {
		problemsFound++
		str := fmt.Sprintf("-- Created By Refs Do Not Match: %s | %s", p.CreatedByRef, b.CreatedByRef)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Created By Refs Match: %s | %s", p.CreatedByRef, b.CreatedByRef)
		details = append(details, str)
	}

	// Check Created Value
	if b.Created != p.Created {
		problemsFound++
		str := fmt.Sprintf("-- Created Dates Do Not Match: %s | %s", p.Created, b.Created)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Created Dates Match: %s | %s", p.Created, b.Created)
		details = append(details, str)
	}

	// Check Modified Value
	if b.Modified != p.Modified {
		problemsFound++
		str := fmt.Sprintf("-- Modified Dates Do Not Match: %s | %s", p.Modified, b.Modified)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Modified Dates Match: %s | %s", p.Modified, b.Modified)
		details = append(details, str)
	}

	// Check Revoked Value
	if b.Revoked != p.Revoked {
		problemsFound++
		str := fmt.Sprintf("-- Revoked Values Do Not Match: %t | %t", p.Revoked, b.Revoked)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Revoked Values Match: %t | %t", p.Revoked, b.Revoked)
		details = append(details, str)
	}

	// Check Labels Values
	if len(b.Labels) != len(p.Labels) {
		problemsFound++
		str := fmt.Sprintf("-- Labels Length Do Not Match: %d | %d", len(p.Labels), len(b.Labels))
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Labels Length Match: %d | %d", len(p.Labels), len(b.Labels))
		details = append(details, str)

		// If lengths are the same, then check each value
		for index, _ := range p.Labels {
			if b.Labels[index] != p.Labels[index] {
				problemsFound++
				str := fmt.Sprintf("-- Labels Do Not Match: %s | %s", p.Labels[index], b.Labels[index])
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Labels Match: %s | %s", p.Labels[index], b.Labels[index])
				details = append(details, str)
			}
		}
	}

	// Check Confidence Value
	if b.Confidence != p.Confidence {
		problemsFound++
		str := fmt.Sprintf("-- Confidence Values Do Not Match: %d | %d", p.Confidence, b.Confidence)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Confidence Values Match: %d | %d", p.Confidence, b.Confidence)
		details = append(details, str)
	}

	// Check Lang Value
	if b.Lang != p.Lang {
		problemsFound++
		str := fmt.Sprintf("-- Lang Values Do Not Match: %s | %s", p.Lang, b.Lang)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Lang Values Match: %s | %s", p.Lang, b.Lang)
		details = append(details, str)
	}

	// Check External References
	if len(b.ExternalReferences) != len(p.ExternalReferences) {
		problemsFound++
		str := fmt.Sprintf("-- External References Length Do Not Match: %d | %d", len(p.ExternalReferences), len(b.ExternalReferences))
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ External References Length Match: %d | %d", len(p.ExternalReferences), len(b.ExternalReferences))
		details = append(details, str)
		for index, _ := range p.ExternalReferences {

			// Check External Reference Source Name
			if b.ExternalReferences[index].SourceName != p.ExternalReferences[index].SourceName {
				problemsFound++
				str := fmt.Sprintf("-- Source Name Do Not Match: %s | %s", p.ExternalReferences[index].SourceName, b.ExternalReferences[index].SourceName)
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Source Name Match: %s | %s", p.ExternalReferences[index].SourceName, b.ExternalReferences[index].SourceName)
				details = append(details, str)
			}

			// Check External Reference Descriptions
			if b.ExternalReferences[index].Description != p.ExternalReferences[index].Description {
				problemsFound++
				str := fmt.Sprintf("-- Descriptions Do Not Match: %s | %s", p.ExternalReferences[index].Description, b.ExternalReferences[index].Description)
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Descriptions Match: %s | %s", p.ExternalReferences[index].Description, b.ExternalReferences[index].Description)
				details = append(details, str)
			}

			// Check External Reference URLs
			if b.ExternalReferences[index].URL != p.ExternalReferences[index].URL {
				problemsFound++
				str := fmt.Sprintf("-- URLs Do Not Match: %s | %s", p.ExternalReferences[index].URL, b.ExternalReferences[index].URL)
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ URLs Match: %s | %s", p.ExternalReferences[index].URL, b.ExternalReferences[index].URL)
				details = append(details, str)
			}

			// Check External Reference Hashes
			if len(b.ExternalReferences[index].Hashes) != len(p.ExternalReferences[index].Hashes) {
				problemsFound++
				str := fmt.Sprintf("-- Hashes Length Do Not Match: %d | %d", len(p.ExternalReferences[index].Hashes), len(b.ExternalReferences[index].Hashes))
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Hashes Length Match: %d | %d", len(p.ExternalReferences[index].Hashes), len(b.ExternalReferences[index].Hashes))
				details = append(details, str)

				// If lengths are the same, then check each value
				for key, _ := range p.ExternalReferences[index].Hashes {
					if b.ExternalReferences[index].Hashes[key] != p.ExternalReferences[index].Hashes[key] {
						problemsFound++
						str := fmt.Sprintf("-- Hashes Do Not Match: %s | %s", p.ExternalReferences[index].Hashes[key], b.ExternalReferences[index].Hashes[key])
						details = append(details, str)
					} else {
						str := fmt.Sprintf("++ Hashes Match: %s | %s", p.ExternalReferences[index].Hashes[key], b.ExternalReferences[index].Hashes[key])
						details = append(details, str)
					}
				}
			}

			// Check External Reference External IDs
			if b.ExternalReferences[index].ExternalID != p.ExternalReferences[index].ExternalID {
				problemsFound++
				str := fmt.Sprintf("-- External IDs Do Not Match: %s | %s", p.ExternalReferences[index].ExternalID, b.ExternalReferences[index].ExternalID)
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ External IDs Match: %s | %s", p.ExternalReferences[index].ExternalID, b.ExternalReferences[index].ExternalID)
				details = append(details, str)
			}
		}
	}

	// Check Object Marking Refs
	if len(b.ObjectMarkingRefs) != len(p.ObjectMarkingRefs) {
		problemsFound++
		str := fmt.Sprintf("-- Object Marking Refs Length Do Not Match: %d | %d", len(p.ObjectMarkingRefs), len(b.ObjectMarkingRefs))
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Object Marking Refs Length Match: %d | %d", len(p.ObjectMarkingRefs), len(b.ObjectMarkingRefs))
		details = append(details, str)

		// If lengths are the same, then check each value
		for index, _ := range p.ObjectMarkingRefs {
			if b.ObjectMarkingRefs[index] != p.ObjectMarkingRefs[index] {
				problemsFound++
				str := fmt.Sprintf("-- Object Marking Refs Do Not Match: %s | %s", p.ObjectMarkingRefs[index], b.ObjectMarkingRefs[index])
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Object Marking Refs Match: %s | %s", p.ObjectMarkingRefs[index], b.ObjectMarkingRefs[index])
				details = append(details, str)
			}
		}
	}

	// Check Granular Markings
	if len(b.GranularMarkings) != len(p.GranularMarkings) {
		problemsFound++
		str := fmt.Sprintf("-- Granular Markings Length Do Not Match: %d | %d", len(p.GranularMarkings), len(b.GranularMarkings))
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Granular Markings Length Match: %d | %d", len(p.GranularMarkings), len(b.GranularMarkings))
		details = append(details, str)
		for index, _ := range p.GranularMarkings {

			// Check Granular Marking Languages
			if b.GranularMarkings[index].Lang != p.GranularMarkings[index].Lang {
				problemsFound++
				str := fmt.Sprintf("-- Languages Do Not Match: %s | %s", p.GranularMarkings[index].Lang, b.GranularMarkings[index].Lang)
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Languages Match: %s | %s", p.GranularMarkings[index].Lang, b.GranularMarkings[index].Lang)
				details = append(details, str)
			}

			// Check Granular Marking Refs
			if b.GranularMarkings[index].MarkingRef != p.GranularMarkings[index].MarkingRef {
				problemsFound++
				str := fmt.Sprintf("-- Refs Do Not Match: %s | %s", p.GranularMarkings[index].MarkingRef, b.GranularMarkings[index].MarkingRef)
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Refs Match: %s | %s", p.GranularMarkings[index].MarkingRef, b.GranularMarkings[index].MarkingRef)
				details = append(details, str)
			}

			// Check Granular Marking Selectors
			if len(b.GranularMarkings[index].Selectors) != len(p.GranularMarkings[index].Selectors) {
				problemsFound++
				str := fmt.Sprintf("-- Selectors Length Do Not Match: %d | %d", len(p.GranularMarkings[index].Selectors), len(b.GranularMarkings[index].Selectors))
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Selectors Length Match: %d | %d", len(p.GranularMarkings[index].Selectors), len(b.GranularMarkings[index].Selectors))
				details = append(details, str)

				// If lengths are the same, then check each value
				for j, _ := range p.GranularMarkings[index].Selectors {
					if b.GranularMarkings[index].Selectors[j] != p.GranularMarkings[index].Selectors[j] {
						problemsFound++
						str := fmt.Sprintf("-- Selectors Do Not Match: %s | %s", p.GranularMarkings[index].Selectors[j], b.GranularMarkings[index].Selectors[j])
						details = append(details, str)
					} else {
						str := fmt.Sprintf("++ Selectors Match: %s | %s", p.GranularMarkings[index].Selectors[j], b.GranularMarkings[index].Selectors[j])
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
