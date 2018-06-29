// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import (
	"github.com/freetaxii/libstix2/defs"
)

// ----------------------------------------------------------------------
//
// Common Property Types - Used to populate the common object properties
//
// ----------------------------------------------------------------------

/*
CommonBaseProperties - This type includes all of the common properties
that are used by all STIX objects
*/
type CommonBaseProperties struct {
	TypeProperty
	IDProperty
}

/*
CommonObjectProperties - This type includes all of the common properties
that are used by all STIX SDOs and SROs
*/
type CommonObjectProperties struct {
	ObjectIDProperty
	CommonBaseProperties
	SpecVersionProperty
	CreatedByRefProperty
	CreatedProperty
	ModifiedProperty
	RevokedProperty
	LabelsProperty
	ConfidenceProperty
	LangProperty
	ExternalReferencesProperty
	ObjectMarkingRefsProperty
	GranularMarkingsProperty
	RawDataProperty
}

/*
CommonMarkingDefinitionProperties - This type includes all of the common
properties that are used by the STIX Marking Definition object
*/
type CommonMarkingDefinitionProperties struct {
	ObjectIDProperty
	CommonBaseProperties
	SpecVersionProperty
	CreatedByRefProperty
	CreatedProperty
	ExternalReferencesProperty
	ObjectMarkingRefsProperty
	GranularMarkingsProperty
	RawDataProperty
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

// SetModifiedToCreated sets the object modified time to be the same as the
// created time. This has to be done at this level, since at the individual
// properties type say "ModifiedPropertyType" this.Created does not exist.
// But it will exist at this level of inheritance
func (p *CommonObjectProperties) SetModifiedToCreated() error {
	p.Modified = p.Created
	return nil
}
