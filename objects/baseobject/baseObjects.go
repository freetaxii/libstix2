// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package baseobject

import (
	"errors"

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
	ObjectIDProperty
	TypeProperty
	SpecVersionProperty
	IDProperty
	CreatedByRefProperty
	CreatedModifiedProperty
	RevokedProperty
	LabelsProperty
	ConfidenceProperty
	LangProperty
	ExternalReferencesProperty
	ObjectMarkingRefsProperty
	GranularMarkingsProperty
	RawProperty
}

/*
CommonMarkingDefinitionProperties - This type includes all of the common
properties that are used by the STIX Marking Definition object
*/
type CommonMarkingDefinitionProperties struct {
	ObjectIDProperty
	TypeProperty
	SpecVersionProperty
	IDProperty
	CreatedByRefProperty
	CreatedModifiedProperty
	ExternalReferencesProperty
	ObjectMarkingRefsProperty
	GranularMarkingsProperty
	RawProperty
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
func (o *CommonObjectProperties) InitObject(stixType string) error {
	// TODO make sure that the value coming in is a valid STIX object type
	o.SetSpecVersion(defs.STIX_VERSION)
	o.SetObjectType(stixType)
	o.SetNewID(stixType)
	o.SetCreatedToCurrentTime()
	o.SetModifiedToCreated()
	return nil
}

func VerifyCommonProperties(o CommonObjectProperties) error {

	if o.ObjectType == "" {
		return errors.New("The type property is required, but missing")
	}

	if o.SpecVersion == "" {
		return errors.New("The spec version property is required, but missing")
	}

	if o.ID == "" {
		return errors.New("The ID property is required, but missing")
	} else {
		// TOOD check to make sure ID is a valid STIX ID but only if it is defined
	}

	if o.Created == "" {
		return errors.New("The created property is required, but missing")
	} else {
		// TODO check to make sure timestamp is a valid STIX timestamp but only if it is defined
	}

	if o.Modified == "" {
		return errors.New("The modified property is required, but missing")
	} else {
		// TODO check to make sure timestamp is a valid STIX timestamp but only if it is defined
	}

	return nil
}
