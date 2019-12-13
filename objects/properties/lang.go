// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import "fmt"

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
LangProperty - A property used by one or more STIX objects that
captures the lang string as defined in RFC 5646. This is used to record the
language that a given object is using.
*/
type LangProperty struct {
	Lang string `json:"lang,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - LangProperty
// ----------------------------------------------------------------------

/*
SetLang - This method takes in a string value representing an ISO 639-2
encoded language code as defined in RFC 5646 and updates the lang property.
*/
func (o *LangProperty) SetLang(s string) error {
	o.Lang = s
	return nil
}

/*
GetLang - This method returns the current language code for a given object.
*/
func (o *LangProperty) GetLang() string {
	return o.Lang
}

/*
CompareLangProperties - This function will compare two lang properties (object 1
and object 2) to make sure they are the same. This function will return an
integer that tracks the number of problems and a slice of strings that contain
the detailed results, whether good or bad.
*/
func CompareLangProperties(obj1, obj2 *LangProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Lang Value
	if obj1.Lang != obj2.Lang {
		problemsFound++
		str := fmt.Sprintf("-- Lang Values Do Not Match: %s | %s", obj1.Lang, obj2.Lang)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ Lang Values Match: %s | %s", obj1.Lang, obj2.Lang)
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
