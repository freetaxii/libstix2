// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package observeddata

import "fmt"

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/*
Valid - This method will verify and test all of the properties on an object
to make sure they are valid per the specification. It will return a boolean, an
integer that tracks the number of problems found, and a slice of strings that
contain the detailed results, whether good or bad.
*/
func (o *ObservedData) Valid(debug bool) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check common base properties first
	_, pBase, dBase := o.CommonObjectProperties.ValidSDO(debug)
	problemsFound += pBase
	resultDetails = append(resultDetails, dBase...)

	// Verify First Observed property is present
	if o.FirstObserved == "" {
		problemsFound++
		str := fmt.Sprintf("-- The first observed property is required but missing")
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The first observed property is required and is present")
		resultDetails = append(resultDetails, str)
	}

	// Verify Last Observed property is present
	if o.LastObserved == "" {
		problemsFound++
		str := fmt.Sprintf("-- The last observed property is required but missing")
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The last observed property is required and is present")
		resultDetails = append(resultDetails, str)
	}

	// Verify Number Observed property is present
	if o.NumberObserved == 0 {
		problemsFound++
		str := fmt.Sprintf("-- The number observed property is required and is missing or set to zero")
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The number observed property is required and is present")
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
