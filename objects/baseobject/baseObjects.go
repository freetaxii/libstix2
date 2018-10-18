// Copyright 2015-2018 Bret Jordan, All rights reserved.
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
CommonBaseProperties - This type includes all of the common properties
that are used by all STIX objects
*/
type CommonBaseProperties struct {
	TypeProperty
	SpecVersionProperty
	IDProperty
}

/*
CommonObjectProperties - This type includes all of the common properties
that are used by all STIX SDOs and SROs
*/
type CommonObjectProperties struct {
	DatastoreIDProperty
	CommonBaseProperties
	CreatedByRefProperty
	CreatedModifiedProperty
	RevokedProperty
	LabelsProperty
	ConfidenceProperty
	LangProperty
	ExternalReferencesProperty
	MarkingProperties
	RawProperty
}

type CommonLanguageContentProperties struct {
	DatastoreIDProperty
	CommonBaseProperties
	CreatedByRefProperty
	CreatedModifiedProperty
	RevokedProperty
	LabelsProperty
	ConfidenceProperty
	ExternalReferencesProperty
	MarkingProperties
	RawProperty
}

/*
CommonMarkingDefinitionProperties - This type includes all of the common
properties that are used by the STIX Marking Definition object. This inherits
the CommonBaseProperties even though the Modified and Revoked properties are
not technically valid for this object.
*/
type CommonMarkingDefinitionProperties struct {
	DatastoreIDProperty
	CommonBaseProperties
	CreatedByRefProperty
	CreatedModifiedProperty
	ExternalReferencesProperty
	MarkingProperties
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

/*
Verify - This method will ensure that all of the required properties are
populated and try to ensure all of values are valid. The difference between this
method and the one for Common Object Properties is this one is missing the check
for the modified property.
*/
func (o *CommonBaseProperties) Verify() error {

	if o.ObjectType == "" {
		return errors.New("the type property is required, but missing")
	}

	if o.SpecVersion == "" {
		return errors.New("the spec version property is required, but missing")
	}

	if o.ID == "" {
		return errors.New("the ID property is required, but missing")
	} else {
		// TOOD check to make sure ID is a valid STIX ID but only if it is defined
	}
	return nil
}

/*
Verify - This method will ensure that all of the required
properties are populated and try to ensure all of values are valid.
*/
func (o *CommonObjectProperties) Verify() error {

	if o.ObjectType == "" {
		return errors.New("the type property is required, but missing")
	}

	if o.SpecVersion == "" {
		return errors.New("the spec version property is required, but missing")
	}

	if o.ID == "" {
		return errors.New("the ID property is required, but missing")
	} else {
		// TOOD check to make sure ID is a valid STIX ID but only if it is defined
	}

	if o.Created == "" {
		return errors.New("the created property is required, but missing")
	} else {
		// TODO check to make sure timestamp is a valid STIX timestamp but only if it is defined
	}

	if o.Modified == "" {
		return errors.New("the modified property is required, but missing")
	} else {
		// TODO check to make sure timestamp is a valid STIX timestamp but only if it is defined
	}

	return nil
}
