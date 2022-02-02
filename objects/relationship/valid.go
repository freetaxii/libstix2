// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package relationship

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
func (o *Relationship) Valid(debug bool) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check common base properties first
	_, pBase, dBase := o.CommonObjectProperties.ValidSDO(debug)
	problemsFound += pBase
	resultDetails = append(resultDetails, dBase...)

	// Verify relationship type property is present
	if o.RelationshipType == "" {
		problemsFound++
		str := fmt.Sprintf("-- The relationship type property is required but missing")
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The relationship type property is required and is present")
		resultDetails = append(resultDetails, str)
	}

	// Verify source ref property is present
	if o.SourceRef == "" {
		problemsFound++
		str := fmt.Sprintf("-- The source ref property is required but missing")
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The source ref property is required and is present")
		resultDetails = append(resultDetails, str)
	}

	// Verify target ref property is present
	if o.TargetRef == "" {
		problemsFound++
		str := fmt.Sprintf("-- The target ref property is required but missing")
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The target ref property is required and is present")
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
