// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package directory

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/*
Valid - This method will verify and test all of the properties on an object
to make sure they are valid per the specification. It will return a boolean, an
integer that tracks the number of problems found, and a slice of strings that
contain the detailed results, whether good or bad.

TODO: Implement full validation per STIX 2.1 specification section 6.3
*/
func (o *Directory) Valid(debug bool) (bool, int, []string) {
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

	// TODO: Add specific validation rules for Directory

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
