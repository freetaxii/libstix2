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
