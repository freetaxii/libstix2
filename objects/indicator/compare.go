// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package indicator

import (
	"fmt"
)

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/*
Compare - This method will compare two objects to make sure they are the
same. The receiver is object 1 and the object passed in is object 2. This method
will return a boolean, an integer that tracks the number of problems found, and
a slice of strings that contain the detailed results, whether good or bad.
*/
func (o *Indicator) Compare(obj2 *Indicator) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check common properties
	_, pBase, dBase := o.CommonObjectProperties.Compare(&obj2.CommonObjectProperties, false)
	problemsFound += pBase
	resultDetails = append(resultDetails, dBase...)

	// Check Name Values
	_, pNames, dNames := o.NameProperty.Compare(&obj2.NameProperty, false)
	problemsFound += pNames
	resultDetails = append(resultDetails, dNames...)

	_, pDescriptions, dDescriptions := o.DescriptionProperty.Compare(&obj2.DescriptionProperty, false)
	problemsFound += pDescriptions
	resultDetails = append(resultDetails, dDescriptions...)

	// Check Indicator Types Property Lengths
	if len(o.IndicatorTypes) != len(obj2.IndicatorTypes) {
		problemsFound++
		str := fmt.Sprintf("-- The indicator types length do not match: %d | %d", len(o.IndicatorTypes), len(obj2.IndicatorTypes))
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The indicator types length match: %d | %d", len(o.IndicatorTypes), len(obj2.IndicatorTypes))
		resultDetails = append(resultDetails, str)

		// If lengths are the same, then check each value
		for index := range o.IndicatorTypes {
			if o.IndicatorTypes[index] != obj2.IndicatorTypes[index] {
				problemsFound++
				str := fmt.Sprintf("-- The indicator types do not match: %s | %s", o.IndicatorTypes[index], obj2.IndicatorTypes[index])
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ The indicator types match: %s | %s", o.IndicatorTypes[index], obj2.IndicatorTypes[index])
				resultDetails = append(resultDetails, str)
			}
		}
	}

	// Check Pattern Values
	if o.Pattern != obj2.Pattern {
		problemsFound++
		str := fmt.Sprintf("-- The patterns do not match: %s | %s", o.Pattern, obj2.Pattern)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The patterns match: %s | %s", o.Pattern, obj2.Pattern)
		resultDetails = append(resultDetails, str)
	}

	// Check PatternType Values
	if o.PatternType != obj2.PatternType {
		problemsFound++
		str := fmt.Sprintf("-- The pattern types do not match: %s | %s", o.PatternType, obj2.PatternType)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The pattern types match: %s | %s", o.PatternType, obj2.PatternType)
		resultDetails = append(resultDetails, str)
	}

	// Check PatternVersion Values
	if o.PatternVersion != obj2.PatternVersion {
		problemsFound++
		str := fmt.Sprintf("-- The pattern versions do not match: %s | %s", o.PatternVersion, obj2.PatternVersion)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The pattern versions match: %s | %s", o.PatternVersion, obj2.PatternVersion)
		resultDetails = append(resultDetails, str)
	}

	// Check ValidFrom Values
	if o.ValidFrom != obj2.ValidFrom {
		problemsFound++
		str := fmt.Sprintf("-- The valid from values do not match: %s | %s", o.ValidFrom, obj2.ValidFrom)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The valid from values match: %s | %s", o.ValidFrom, obj2.ValidFrom)
		resultDetails = append(resultDetails, str)
	}

	// Check ValidUntil Values
	if o.ValidUntil != obj2.ValidUntil {
		problemsFound++
		str := fmt.Sprintf("-- The valid until values do not match: %s | %s", o.ValidUntil, obj2.ValidUntil)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The valid until values match: %s | %s", o.ValidUntil, obj2.ValidUntil)
		resultDetails = append(resultDetails, str)
	}

	// Check Kill Chain Phases
	if valid, problems, details := o.KillChainPhasesProperty.Compare(&obj2.KillChainPhasesProperty, false); valid != true {
		problemsFound += problems
		for _, v := range details {
			resultDetails = append(resultDetails, v)
		}
	} else {
		// Everything was good, let's just capture any details that were returned.
		for _, v := range details {
			resultDetails = append(resultDetails, v)
		}
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}

// ----------------------------------------------------------------------
// Public Functions
// ----------------------------------------------------------------------

/*
Compare - This function will compare two objects to make sure they are the
same and will return a boolean, an integer that tracks the number of problems
found, and a slice of strings that contain the detailed results, whether good or
bad.
*/
func Compare(obj1, obj2 *Indicator) (bool, int, []string) {
	return obj1.Compare(obj2)
}
