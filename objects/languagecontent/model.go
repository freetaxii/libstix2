// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package languagecontent

import (
	"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
LanguageContent - This type implements the STIX 2.1 Language Content object and
defines all of the properties and methods needed to create and work with this
object. All of the methods not defined local to this type are inherited from the
individual properties.

Reference: STIX 2.1 specification section 7.2

The Language Content object represents text content for STIX Objects represented
in languages other than that of the original object.
*/
type LanguageContent struct {
	objects.CommonObjectProperties
	ObjectRef     string                       `json:"object_ref" bson:"object_ref"`
	ObjectModified string                      `json:"object_modified,omitempty" bson:"object_modified,omitempty"`
	Contents      map[string]map[string]string `json:"contents" bson:"contents"`
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *LanguageContent) GetPropertyList() []string {
	return []string{"object_ref", "object_modified", "contents"}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Language Content object and return
it as a pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *LanguageContent {
	var obj LanguageContent
	obj.InitSDO("language-content")
	return &obj
}
