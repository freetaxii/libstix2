// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import (
	"fmt"

	"github.com/freetaxii/libstix2/resources"
)

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/*
AliasesProperty - A property used by one or more STIX objects.
*/
type AliasesProperty struct {
	Aliases []string `json:"aliases,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - AliasesProperty - Setters
// ----------------------------------------------------------------------

/*
AddAliases - This method takes in a string value, a comma separated list of
string values, or a slice of string values that represents an alias and adds it
to the aliases property.
*/
func (o *AliasesProperty) AddAliases(values interface{}) error {
	return resources.AddValuesToList(&o.Aliases, values)
}

// ----------------------------------------------------------------------
// Public Methods - AliasesProperty - Checks
// ----------------------------------------------------------------------

/*
Compare - This method will compare two properties to make sure they are the
same and will return a boolean, an integer that tracks the number of problems
found, and a slice of strings that contain the detailed results, whether good or
bad.
*/
func (o *AliasesProperty) Compare(obj2 *AliasesProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	if len(o.Aliases) != len(obj2.Aliases) {
		problemsFound++
		str := fmt.Sprintf("-- The number of entries in aliases do not match: %d | %d", len(o.Aliases), len(obj2.Aliases))
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The number of entries in aliases match: %d | %d", len(o.Aliases), len(obj2.Aliases))
		resultDetails = append(resultDetails, str)

		// If lengths are the same, then check each value
		for index := range o.Aliases {
			if o.Aliases[index] != obj2.Aliases[index] {
				problemsFound++
				str := fmt.Sprintf("-- The alias values do not match: %s | %s", o.Aliases[index], obj2.Aliases[index])
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ The alias values match: %s | %s", o.Aliases[index], obj2.Aliases[index])
				resultDetails = append(resultDetails, str)
			}
		}
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
