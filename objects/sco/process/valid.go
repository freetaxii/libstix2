// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package process

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
func (o *Process) Valid(debug bool) (bool, int, []string) {
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

	// Validate pid if present (should be non-negative)
	if o.Pid < 0 {
		problemsFound++
		resultDetails = append(resultDetails, "-- the pid property cannot be negative")
	} else {
		resultDetails = append(resultDetails, "++ the pid property is non-negative")
	}

	// Validate created_time if present
	if o.CreatedTime != "" {
		if valid := objects.IsTimestampValid(o.CreatedTime); !valid {
			problemsFound++
			resultDetails = append(resultDetails, "-- the created_time property does not contain a valid timestamp")
		} else {
			resultDetails = append(resultDetails, "++ the created_time property contains a valid timestamp")
		}
	}

	// Validate opened_connection_refs if present (should contain valid STIX IDs)
	for _, ref := range o.OpenedConnectionRefs {
		if !objects.IsIDValid(ref) {
			problemsFound++
			resultDetails = append(resultDetails, "-- opened_connection_refs contains an invalid STIX ID: "+ref)
		}
	}
	if len(o.OpenedConnectionRefs) > 0 {
		resultDetails = append(resultDetails, "++ opened_connection_refs contains valid STIX IDs")
	}

	// Validate creator_user_ref if present
	if o.CreatorUserRef != "" {
		if valid := objects.IsIDValid(o.CreatorUserRef); !valid {
			problemsFound++
			resultDetails = append(resultDetails, "-- the creator_user_ref property does not contain a valid identifier")
		} else {
			resultDetails = append(resultDetails, "++ the creator_user_ref property contains a valid identifier")
		}
	}

	// Validate image_ref if present
	if o.ImageRef != "" {
		if valid := objects.IsIDValid(o.ImageRef); !valid {
			problemsFound++
			resultDetails = append(resultDetails, "-- the image_ref property does not contain a valid identifier")
		} else {
			resultDetails = append(resultDetails, "++ the image_ref property contains a valid identifier")
		}
	}

	// Validate parent_ref if present
	if o.ParentRef != "" {
		if valid := objects.IsIDValid(o.ParentRef); !valid {
			problemsFound++
			resultDetails = append(resultDetails, "-- the parent_ref property does not contain a valid identifier")
		} else {
			resultDetails = append(resultDetails, "++ the parent_ref property contains a valid identifier")
		}
	}

	// Validate child_refs if present (should contain valid STIX IDs)
	for _, ref := range o.ChildRefs {
		if !objects.IsIDValid(ref) {
			problemsFound++
			resultDetails = append(resultDetails, "-- child_refs contains an invalid STIX ID: "+ref)
		}
	}
	if len(o.ChildRefs) > 0 {
		resultDetails = append(resultDetails, "++ child_refs contains valid STIX IDs")
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
