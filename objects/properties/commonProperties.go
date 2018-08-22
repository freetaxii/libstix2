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
		str := fmt.Sprintf("++ Types Match: %s | %s", p.ObjectType, b.ObjectType)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Types Match: %s | %s", p.ObjectType, b.ObjectType)
		details = append(details, str)
	}

	// Check Spec Version Value
	if b.SpecVersion != p.SpecVersion {
		problemsFound++
		str := fmt.Sprintf("++ Spec Versions Match: %s | %s", p.SpecVersion, b.SpecVersion)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Spec Versions Match: %s | %s", p.SpecVersion, b.SpecVersion)
		details = append(details, str)
	}

	// Check ID Value
	if b.ID != p.ID {
		problemsFound++
		str := fmt.Sprintf("++ IDs Match: %s | %s", p.ID, b.ID)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ IDs Match: %s | %s", p.ID, b.ID)
		details = append(details, str)
	}

	// Check Created Value
	if b.Created != p.Created {
		problemsFound++
		str := fmt.Sprintf("++ Created Dates Match: %s | %s", p.Created, b.Created)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Created Dates Match: %s | %s", p.Created, b.Created)
		details = append(details, str)
	}

	// Check Modified Value
	if b.Modified != p.Modified {
		problemsFound++
		str := fmt.Sprintf("++ Modified Dates Match: %s | %s", p.Modified, b.Modified)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Modified Dates Match: %s | %s", p.Modified, b.Modified)
		details = append(details, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, details
	}

	return true, 0, details
}
