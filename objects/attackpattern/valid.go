// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package attackpattern

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/* Valid - This method will verify and test all of the properties on an object
to make sure they are valid per the specification. It will return a boolean, an
integer that tracks the number of problems found, and a slice of strings that
contain the detailed results, whether good or bad. */
func (o *AttackPattern) Valid() (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check common base properties first
	_, pBase, dBase := o.CommonObjectProperties.Valid()
	problemsFound += pBase
	resultDetails = append(resultDetails, dBase...)

	// Check attack pattern specific properties
	_, pSpecific, dSpecific := o.validSpecificProperties()
	problemsFound += pSpecific
	resultDetails = append(resultDetails, dSpecific...)

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}

// ----------------------------------------------------------------------
// Private Methods
// ----------------------------------------------------------------------

// TODO This needs to be moved to the name property
func (o *AttackPattern) validSpecificProperties() (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	if o.Name == "" {
		problemsFound++
		str := fmtSprintf("-- The Name property is required on Attack Pattern but is missing")
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
