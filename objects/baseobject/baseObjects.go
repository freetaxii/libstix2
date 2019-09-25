// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package baseobject

import (
	"github.com/freetaxii/libstix2/defs"
)

// ----------------------------------------------------------------------
//
// Common Property Types - Used to populate the common object properties
//
// ----------------------------------------------------------------------

/*
BundleBaseProperties - This type includes all of the common properties
that are used by by the STIX Bundle. It is done here to make it similar to
all other STIX object definitions. Meaning, that they all use this baseobject
package.
*/
type BundleBaseProperties struct {
	TypeProperty
	IDProperty
}

/*
CommonObjectProperties - This type includes all of the common properties that
are used by all STIX SDOs, SROs, Marking Definition Objects, and the Language
object.  For objects where some of these properties are not defined, they will
be removed / zeroed out in their respective encoding methods.
*/
type CommonObjectProperties struct {
	DatastoreIDProperty
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
	MarkingProperties
	RawProperty
}

// ----------------------------------------------------------------------
// Public Methods - CommonObjectProperties
// ----------------------------------------------------------------------

/*
InitObject - This method will initialize the object by setting all of the basic
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
GetCommonProperties - This method will return a pointer to the common properties
of this object.
*/
func (o *CommonObjectProperties) GetCommonProperties() *CommonObjectProperties {
	return o
}

/*
Valid - This method will ensure that all of the required properties are
populated and try to ensure all of values are valid.
*/
func (o *CommonObjectProperties) Valid() (bool, error) {

	if valid, err := o.TypeProperty.Valid(); valid != true {
		return valid, err
	}

	if valid, err := o.SpecVersionProperty.Valid(); valid != true {
		return valid, err
	}

	if valid, err := o.IDProperty.Valid(); valid != true {
		return valid, err
	}

	if valid, err := o.CreatedModifiedProperty.Valid(); valid != true {
		return valid, err
	}

	return true, nil
}
