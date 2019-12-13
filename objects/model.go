// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

import "github.com/freetaxii/libstix2/objects/properties"

// ----------------------------------------------------------------------
// Define Internal Object Model
// ----------------------------------------------------------------------

/*
baseProperties - This type includes the Base properties that are common to
nearly every STIX Object. If an object does not define one of these properties
then it will be removed / zeroed out during encoding.
*/
type baseProperties struct {
	properties.TypeProperty
	properties.SpecVersionProperty
	properties.IDProperty
}

/*
versioningProperties - This type includes the versioning properties that are
common on many STIX Objects. If an object only defines some of these properties,
then the ones that are not defined will be removed / zeroed out during encoding.
*/
type versioningProperties struct {
	properties.CreatedByRefProperty
	properties.CreatedModifiedProperty
	properties.RevokedProperty
}

/*
extendedProperties - This type includes all of the extended properties that are
used on many STIX Objects. If an object only defines some of these properties,
then the ones that are not defined will be removed / zeroed out during encoding.
*/
type extendedProperties struct {
	properties.LabelsProperty
	properties.ConfidenceProperty
	properties.LangProperty
}

/*
markingProperties - This type includes all of the marking properties that are
used on many STIX Objects. If an object only defines some of these properties,
then the ones that are not defined will be removed / zeroed out during encoding.
*/
type markingProperties struct {
	properties.ExternalReferencesProperty
	properties.MarkingProperty
}

// ----------------------------------------------------------------------
// Define Exported Object Model
// ----------------------------------------------------------------------

type CommonObjectProperties struct {
	properties.DatastoreIDProperty
	baseProperties
	versioningProperties
	extendedProperties
	markingProperties
	properties.RawProperty
}

type CommonBundleProperties struct {
	baseProperties
}

type CommonScoProperties struct {
	properties.DatastoreIDProperty
	baseProperties
	markingProperties
	properties.RawProperty
}
