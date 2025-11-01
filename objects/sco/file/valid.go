// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package file

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
func (o *File) Valid(debug bool) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check common SCO properties (type, spec_version, id)
if o.ObjectType == "" {
problemsFound++
resultDetails = append(resultDetails, "-- the type property is required but missing")
}

if o.SpecVersion == "" {
problemsFound++
resultDetails = append(resultDetails, "-- the spec_version property is required but missing")
}

if o.ID == "" {
problemsFound++
resultDetails = append(resultDetails, "-- the id property is required but missing")
}

	// File object MUST contain at least one of hashes or name
	if (o.Hashes == nil || len(o.Hashes) == 0) && o.Name == "" {
		problemsFound++
		resultDetails = append(resultDetails, "-- The file object MUST contain at least one of hashes or name")
	} else {
		resultDetails = append(resultDetails, "++ The file object contains at least one of hashes or name")
	}

	// Validate size if present
	if o.Size < 0 {
		problemsFound++
		resultDetails = append(resultDetails, fmt.Sprintf("-- The size property cannot be negative: %d", o.Size))
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
