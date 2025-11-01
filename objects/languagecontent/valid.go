// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package languagecontent

import (
	"fmt"
)

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/*
Valid - This method will verify and test all of the properties on an object
to make sure they are valid per the specification. It will return a boolean, an
integer that tracks the number of problems found, and a slice of strings that
contain the detailed results, whether good or bad.
*/
func (o *LanguageContent) Valid(debug bool) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check common base properties first
	_, pBase, dBase := o.CommonObjectProperties.ValidSDO(debug)
	problemsFound += pBase
	resultDetails = append(resultDetails, dBase...)

	// Verify object_ref property is present (required)
	if o.ObjectRef == "" {
		problemsFound++
		resultDetails = append(resultDetails, "-- The object_ref property is required but missing")
	} else {
		resultDetails = append(resultDetails, fmt.Sprintf("++ The object_ref property is present: %s", o.ObjectRef))
	}

	// Verify contents property is present (required)
	if o.Contents == nil || len(o.Contents) == 0 {
		problemsFound++
		resultDetails = append(resultDetails, "-- The contents property is required but missing or empty")
	} else {
		resultDetails = append(resultDetails, fmt.Sprintf("++ The contents property contains %d language(s)", len(o.Contents)))
		
		// Validate each language entry has at least one selector
		for lang, selectors := range o.Contents {
			if len(selectors) == 0 {
				problemsFound++
				resultDetails = append(resultDetails, fmt.Sprintf("-- Language '%s' has no selectors", lang))
			} else {
				resultDetails = append(resultDetails, fmt.Sprintf("++ Language '%s' has %d selector(s)", lang, len(selectors)))
			}
		}
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
