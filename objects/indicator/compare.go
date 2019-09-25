// Copyright 2015-2019 Bret Jordan, All rights reserved.
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
Compare - This method will compare two indicators to make sure they are the
same. The indicator receiver is object 1 and the indicator passed in is object
2. This method will return an integer that tracks the number of problems and a
slice of strings that contain the detailed results, whether good or bad.
*/
func (o *Indicator) Compare(obj2 *Indicator) (bool, int, []string) {
	return Compare(o, obj2)
}

// ----------------------------------------------------------------------
// Public Functions
// ----------------------------------------------------------------------

/*
Compare - This function will compare two indicators (object 1 and object 2) to
make sure they are the same. This function will return an integer that tracks
the number of problems and a slice of strings that contain the detailed results,
whether good or bad.
*/
func Compare(obj1, obj2 *Indicator) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check common properties
	if valid, problems, d := obj1.CommonObjectProperties.Compare(&obj2.CommonObjectProperties); valid != true {
		problemsFound += problems
		for _, v := range d {
			resultDetails = append(resultDetails, v)
		}
	} else {
		// The Common Properties were good, so lets just capture any details
		// that were returned.
		for _, v := range d {
			resultDetails = append(resultDetails, v)
		}
	}

	// Check Name Value
	if obj1.Name != obj2.Name {
		problemsFound++
		str := fmt.Sprintf("-- Names Do Not Match: %s | %s", obj1.Name, obj2.Name)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ Names Match: %s | %s", obj1.Name, obj2.Name)
		resultDetails = append(resultDetails, str)
	}

	// Check Description Value
	if obj1.Description != obj2.Description {
		problemsFound++
		str := fmt.Sprintf("-- Descriptions Do Not Match: %s | %s", obj1.Description, obj2.Description)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ Descriptions Match: %s | %s", obj1.Description, obj2.Description)
		resultDetails = append(resultDetails, str)
	}

	// Check Indicator Types Property Length
	if len(obj1.IndicatorTypes) != len(obj2.IndicatorTypes) {
		problemsFound++
		str := fmt.Sprintf("-- Indicator Types Length Do Not Match: %d | %d", len(obj1.IndicatorTypes), len(obj2.IndicatorTypes))
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ Indicator Types Length Match: %d | %d", len(obj1.IndicatorTypes), len(obj2.IndicatorTypes))
		resultDetails = append(resultDetails, str)

		// If lengths are the same, then check each value
		for index := range obj1.IndicatorTypes {
			if obj1.IndicatorTypes[index] != obj2.IndicatorTypes[index] {
				problemsFound++
				str := fmt.Sprintf("-- Indicator Types Do Not Match: %s | %s", obj1.IndicatorTypes[index], obj2.IndicatorTypes[index])
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ Indicator Types Match: %s | %s", obj1.IndicatorTypes[index], obj2.IndicatorTypes[index])
				resultDetails = append(resultDetails, str)
			}
		}
	}

	// Check Pattern Value
	if obj1.Pattern != obj2.Pattern {
		problemsFound++
		str := fmt.Sprintf("-- Patterns Do Not Match: %s | %s", obj1.Pattern, obj2.Pattern)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ Patterns Match: %s | %s", obj1.Pattern, obj2.Pattern)
		resultDetails = append(resultDetails, str)
	}

	// Check PatternType Value
	if obj1.PatternType != obj2.PatternType {
		problemsFound++
		str := fmt.Sprintf("-- Pattern Types Do Not Match: %s | %s", obj1.PatternType, obj2.PatternType)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ Pattern Types Match: %s | %s", obj1.PatternType, obj2.PatternType)
		resultDetails = append(resultDetails, str)
	}

	// Check PatternVersion Value
	if obj1.PatternVersion != obj2.PatternVersion {
		problemsFound++
		str := fmt.Sprintf("-- Pattern Versions Do Not Match: %s | %s", obj1.PatternVersion, obj2.PatternVersion)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ Pattern Versions Match: %s | %s", obj1.PatternVersion, obj2.PatternVersion)
		resultDetails = append(resultDetails, str)
	}

	// Check ValidFrom Value
	if obj1.ValidFrom != obj2.ValidFrom {
		problemsFound++
		str := fmt.Sprintf("-- ValidFrom Values Do Not Match: %s | %s", obj1.ValidFrom, obj2.ValidFrom)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ ValidFrom Values Match: %s | %s", obj1.ValidFrom, obj2.ValidFrom)
		resultDetails = append(resultDetails, str)
	}

	// Check ValidUntil Value
	if obj1.ValidUntil != obj2.ValidUntil {
		problemsFound++
		str := fmt.Sprintf("-- ValidUntil Values Do Not Match: %s | %s", obj1.ValidUntil, obj2.ValidUntil)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ ValidUntil Values Match: %s | %s", obj1.ValidUntil, obj2.ValidUntil)
		resultDetails = append(resultDetails, str)
	}

	// Check Kill Chain Phases Property Length
	if len(obj1.KillChainPhases) != len(obj2.KillChainPhases) {
		problemsFound++
		str := fmt.Sprintf("-- Kill Chain Phases Length Do Not Match: %d | %d", len(obj1.KillChainPhases), len(obj2.KillChainPhases))
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ Kill Chain Phases Length Match: %d | %d", len(obj1.KillChainPhases), len(obj2.KillChainPhases))
		resultDetails = append(resultDetails, str)
		for index := range obj1.KillChainPhases {
			// Check Kill Chain Phases values
			if obj1.KillChainPhases[index].KillChainName != obj2.KillChainPhases[index].KillChainName {
				problemsFound++
				str := fmt.Sprintf("-- Kill Chain Names Do Not Match: %s | %s", obj1.KillChainPhases[index].KillChainName, obj2.KillChainPhases[index].KillChainName)
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ Kill Chain Names Match: %s | %s", obj1.KillChainPhases[index].KillChainName, obj2.KillChainPhases[index].KillChainName)
				resultDetails = append(resultDetails, str)
			}

			// Check Kill Chain Phases values
			if obj1.KillChainPhases[index].PhaseName != obj2.KillChainPhases[index].PhaseName {
				problemsFound++
				str := fmt.Sprintf("-- Kill Chain Phases Do Not Match: %s | %s", obj1.KillChainPhases[index].PhaseName, obj2.KillChainPhases[index].PhaseName)
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ Kill Chain Phases Match: %s | %s", obj1.KillChainPhases[index].PhaseName, obj2.KillChainPhases[index].PhaseName)
				resultDetails = append(resultDetails, str)
			}
		}
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
