// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package directory

import (
	"github.com/freetaxii/libstix2/objects"
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
func (o *Directory) Valid(debug bool) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check common SCO properties (type, spec_version, id) - these are required for SCOs
	if o.ObjectType == "" {
		problemsFound++
		resultDetails = append(resultDetails, "-- the type property is required but missing")
	} else {
		resultDetails = append(resultDetails, "++ the type property is present")
	}

	if o.SpecVersion == "" {
		problemsFound++
		resultDetails = append(resultDetails, "-- the spec_version property is required but missing")
	} else {
		resultDetails = append(resultDetails, "++ the spec_version property is present")
	}

	if o.ID == "" {
		problemsFound++
		resultDetails = append(resultDetails, "-- the id property is required but missing")
	} else {
		resultDetails = append(resultDetails, "++ the id property is present")
	}

	// Directory specific validations
	if o.Path == "" {
		problemsFound++
		resultDetails = append(resultDetails, "-- the path property is required but missing")
	} else {
		resultDetails = append(resultDetails, "++ the path property is present")
	}

	// Validate timestamps if present
	if o.Ctime != "" {
		if valid := objects.IsTimestampValid(o.Ctime); !valid {
			problemsFound++
			resultDetails = append(resultDetails, "-- the ctime property does not contain a valid timestamp")
		} else {
			resultDetails = append(resultDetails, "++ the ctime property contains a valid timestamp")
		}
	}

	if o.Mtime != "" {
		if valid := objects.IsTimestampValid(o.Mtime); !valid {
			problemsFound++
			resultDetails = append(resultDetails, "-- the mtime property does not contain a valid timestamp")
		} else {
			resultDetails = append(resultDetails, "++ the mtime property contains a valid timestamp")
		}
	}

	if o.Atime != "" {
		if valid := objects.IsTimestampValid(o.Atime); !valid {
			problemsFound++
			resultDetails = append(resultDetails, "-- the atime property does not contain a valid timestamp")
		} else {
			resultDetails = append(resultDetails, "++ the atime property contains a valid timestamp")
		}
	}

	// Validate contains_refs if present (should contain valid STIX IDs)
	for _, ref := range o.ContainsRefs {
		if !objects.IsIDValid(ref) {
			problemsFound++
			resultDetails = append(resultDetails, "-- contains_refs contains an invalid STIX ID: "+ref)
		}
	}
	if len(o.ContainsRefs) > 0 {
		resultDetails = append(resultDetails, "++ contains_refs contains valid STIX IDs")
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
