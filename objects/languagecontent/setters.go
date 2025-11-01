// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package languagecontent

import (
	"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
// Public Methods - LanguageContent - Setters
// ----------------------------------------------------------------------

/*
SetObjectRef - This method takes in a string value representing a STIX object
identifier and updates the object_ref property.
*/
func (o *LanguageContent) SetObjectRef(s string) error {
	o.ObjectRef = s
	return nil
}

/*
SetObjectModified - This method takes in a timestamp in either time.Time or
string format and updates the object_modified property.
*/
func (o *LanguageContent) SetObjectModified(t interface{}) error {
	ts, _ := objects.TimeToString(t, "micro")
	o.ObjectModified = ts
	return nil
}

/*
AddContent - This method takes in three parameters and adds content to the
contents property. The first parameter is a string representing the language
code (e.g., "en", "es", "fr"). The second parameter is a string representing
the property selector (e.g., "name", "description"). The third parameter is a
string representing the translated content.
*/
func (o *LanguageContent) AddContent(language, selector, value string) error {
	if o.Contents == nil {
		o.Contents = make(map[string]map[string]string)
	}
	
	if o.Contents[language] == nil {
		o.Contents[language] = make(map[string]string)
	}
	
	o.Contents[language][selector] = value
	return nil
}
