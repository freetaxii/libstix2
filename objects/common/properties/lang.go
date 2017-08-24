// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type LangPropertyType struct {
	Lang string `json:"lang,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - LangPropertyType
// ----------------------------------------------------------------------

// SetLang takes in one parameter
// param: s - a string value representing an ISO 639-2 encoded language code
func (this *LangPropertyType) SetLang(s string) {
	this.Lang = s
}

func (this *LangPropertyType) GetLang() string {
	return this.Lang
}
