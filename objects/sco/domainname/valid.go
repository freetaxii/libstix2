// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package domainname

import "fmt"

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/* Valid - This method will verify and test all of the properties on an object
to make sure they are valid per the specification. It will return a boolean, an
integer that tracks the number of problems found, and a slice of strings that
contain the detailed results, whether good or bad. */
func (o *DomainName) Valid() (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check common base properties first
	_, pBase, dBase := o.CommonObjectProperties.ValidSDO()
	problemsFound += pBase
	resultDetails = append(resultDetails, dBase...)

	// Verify the value property exists
	if o.Value == "" {
		problemsFound++
		str := fmt.Sprintf("-- The value property is required but missing")
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The value property is required and is present")
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
