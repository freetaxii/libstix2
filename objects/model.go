// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

import "github.com/freetaxii/libstix2/objects/properties"

// ----------------------------------------------------------------------
// Define Exported Object Model
// ----------------------------------------------------------------------

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
	properties.MarkingProperties
	properties.RawProperty
}

type CommonBundleProperties struct {
	properties.TypeProperty
	properties.IDProperty
}

type CommonScoProperties struct {
	properties.DatastoreIDProperty
	properties.TypeProperty
	properties.SpecVersionProperty
	properties.IDProperty
	properties.MarkingProperties
	properties.RawProperty
}
