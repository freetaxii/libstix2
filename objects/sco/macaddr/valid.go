// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package macaddr

import (
	"fmt"
	"regexp"
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
func (o *MACAddr) Valid(debug bool) (bool, int, []string) {
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

	// Verify value property is present
	_, p, d := o.ValueProperty.VerifyExists()
	problemsFound += p
	resultDetails = append(resultDetails, d...)

	// Validate MAC address format
	if o.Value != "" {
		if valid := isValidMAC(o.Value); valid {
			resultDetails = append(resultDetails, fmt.Sprintf("++ The value property is a valid MAC address: %s", o.Value))
		} else {
			problemsFound++
			resultDetails = append(resultDetails, fmt.Sprintf("-- The value property is not a valid MAC address: %s", o.Value))
		}
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}

// ----------------------------------------------------------------------
// Private Methods
// ----------------------------------------------------------------------

// isValidMAC validates a MAC address format
func isValidMAC(value string) bool {
	// MAC address patterns:
	// - 00:00:00:00:00:00 (colon-separated)
	// - 00-00-00-00-00-00 (dash-separated)
	// - 0000.0000.0000 (dot-separated, Cisco style)
	patterns := []string{
		`^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`,
		`^([0-9A-Fa-f]{4}\.){2}([0-9A-Fa-f]{4})$`,
	}

	for _, pattern := range patterns {
		matched, _ := regexp.MatchString(pattern, value)
		if matched {
			return true
		}
	}
	return false
}
