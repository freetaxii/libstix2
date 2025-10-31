// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package markingdefinition

import (
	"github.com/freetaxii/libstix2/objects/properties"
)

/*
Valid - This method will verify and test all of the properties on an object
to make sure they are valid per the specification. It will return a boolean, an
integer that tracks the number of problems found, and a slice of strings that
contain the detailed results, whether good or bad.
*/
func (o *MarkingDefinition) Valid(debug bool) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check common base properties first
	_, pBase, dBase := o.CommonObjectProperties.ValidSDO(debug)
	problemsFound += pBase
	resultDetails = append(resultDetails, dBase...)

	// Verify object Name property is present
	if o.GetName() == "" {
		problemsFound++
		str := "-- The markingDefinition name property is required but missing"
		resultDetails = append(resultDetails, str)
	}

	// Verify object DefinitionType property is present
	if o.DefinitionType != "tlp" && o.DefinitionType != "statement" {
		problemsFound++
		str := "-- The markingDefinition definition type property is neither tlp nor statement"
		resultDetails = append(resultDetails, str)
	}

	// Verify object Definition property is present
	if t, ok := o.Definition.(properties.TlpDefinition); ok {
		if t.Tlp == "" {
			problemsFound++
			str := "-- The markingDefinition definition tlp property is required but missing"
			resultDetails = append(resultDetails, str)
		}
	} else if t, ok := o.Definition.(properties.StatementDefinition); ok {
		if t.Statement == "" {
			problemsFound++
			str := "-- The markingDefinition definition statement property is required but missing"
			resultDetails = append(resultDetails, str)
		}
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
