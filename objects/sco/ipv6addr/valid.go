// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package ipv6addr

import (
	"fmt"
	"net"
	"strings"
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
func (o *IPv6Addr) Valid(debug bool) (bool, int, []string) {
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

	// Validate IPv6 address format
	if o.Value != "" {
		if valid := isValidIPv6(o.Value); valid {
			resultDetails = append(resultDetails, fmt.Sprintf("++ The value property is a valid IPv6 address: %s", o.Value))
		} else {
			problemsFound++
			resultDetails = append(resultDetails, fmt.Sprintf("-- The value property is not a valid IPv6 address: %s", o.Value))
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

// isValidIPv6 validates an IPv6 address or CIDR notation
func isValidIPv6(value string) bool {
	// Check if it's CIDR notation
	if strings.Contains(value, "/") {
		_, _, err := net.ParseCIDR(value)
		return err == nil
	}

	// Check if it's a plain IPv6 address
	ip := net.ParseIP(value)
	return ip != nil && ip.To4() == nil
}
