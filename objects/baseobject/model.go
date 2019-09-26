// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package baseobject

import (
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Common Property Types - Used to populate the common object properties
// ----------------------------------------------------------------------

/*
BundleBaseProperties - This type includes all of the common properties
that are used by by the STIX Bundle. It is done here to make it similar to
all other STIX object definitions. Meaning, that they all use this baseobject
package.
*/
type BundleBaseProperties struct {
	properties.TypeProperty
	properties.IDProperty
}

/*
CommonObjectProperties - This type includes all of the common properties that
are used by all STIX SDOs, SROs, Marking Definition Objects, and the Language
object.  For objects where some of these properties are not defined, they will
be removed / zeroed out in their respective encoding methods.
*/
type CommonObjectProperties struct {
	properties.DatastoreIDProperty
	properties.TypeProperty
	properties.SpecVersionProperty
	properties.IDProperty
	properties.CreatedByRefProperty
	properties.CreatedModifiedProperty
	properties.RevokedProperty
	properties.LabelsProperty
	properties.ConfidenceProperty
	properties.LangProperty
	properties.ExternalReferencesProperty
	properties.MarkingProperty
	properties.RawProperty
}
