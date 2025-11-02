// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package artifact

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/*
Valid - This method will verify and test all of the properties on an object
to make sure they are valid per the specification. It will return a boolean, an
integer that tracks the number of problems found, and a slice of strings that
contain the detailed results, whether good or bad.
*/
func (o *Artifact) Valid(debug bool) (bool, int, []string) {
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

	// Artifact MUST contain at least one of payload_bin or url
	if o.PayloadBin == "" && o.URL == "" {
		problemsFound++
		resultDetails = append(resultDetails, "-- The artifact object MUST contain at least one of payload_bin or url")
	} else {
		resultDetails = append(resultDetails, "++ The artifact object contains at least one of payload_bin or url")
	}

	// If hashes are provided when payload_bin is used, validate
	if o.PayloadBin != "" && (o.Hashes == nil || len(o.Hashes) == 0) {
		resultDetails = append(resultDetails, "-- Warning: hashes SHOULD be provided when payload_bin is specified")
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
