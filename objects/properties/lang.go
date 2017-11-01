// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

// LangPropertyType - A property used by one or more STIX objects that
// captures the lang string as defined in RFC 5646. This is used to record the
// language that a given object is using.
type LangPropertyType struct {
	Lang string `json:"lang,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - LangPropertyType
// ----------------------------------------------------------------------

// SetLang - This method takes in a string value representing an ISO 639-2
// encoded language code as defined in RFC 5646 and updates the lang property.
func (ezt *LangPropertyType) SetLang(s string) {
	ezt.Lang = s
}

// GetLang - This method returns the current language code for a given object.
func (ezt *LangPropertyType) GetLang() string {
	return ezt.Lang
}
