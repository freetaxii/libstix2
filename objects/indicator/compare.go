// Copyright 2015-2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package indicator

import (
	"fmt"
)

/*
Compare - This method will compare two indicators to make sure they
are the same. The indicator receiver is the known good and represent the correct
data, the indicator passed in as i represents the one we need to test/check.
*/
func (o *Indicator) Compare(toTest *Indicator) (bool, int, []string) {
	return Compare(o, toTest)
}

/*
Compare - This function will compare two indicators to make sure they
are the same. Indicator correct is the master and represent the correct
data, indicator toTest represents the one we need to test.
*/
func Compare(correct, toTest *Indicator) (bool, int, []string) {
	problemsFound := 0
	errorDetails := make([]string, 0)

	// Check common properties
	if valid, problems, d := correct.CommonObjectProperties.Compare(&toTest.CommonObjectProperties); valid != true {
		problemsFound += problems
		for _, v := range d {
			errorDetails = append(errorDetails, v)
		}
	} else {
		for _, v := range d {
			errorDetails = append(errorDetails, v)
		}
	}

	// Check Name Value
	if toTest.Name != correct.Name {
		problemsFound++
		str := fmt.Sprintf("-- Names Do Not Match: %s | %s", correct.Name, toTest.Name)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Names Match: %s | %s", correct.Name, toTest.Name)
		errorDetails = append(errorDetails, str)
	}

	// Check Description Value
	if toTest.Description != correct.Description {
		problemsFound++
		str := fmt.Sprintf("-- Descriptions Do Not Match: %s | %s", correct.Description, toTest.Description)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Descriptions Match: %s | %s", correct.Description, toTest.Description)
		errorDetails = append(errorDetails, str)
	}

	// Check Indicator Types Property Length
	if len(toTest.IndicatorTypes) != len(correct.IndicatorTypes) {
		problemsFound++
		str := fmt.Sprintf("-- Indicator Types Length Do Not Match: %d | %d", len(correct.IndicatorTypes), len(toTest.IndicatorTypes))
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Indicator Types Length Match: %d | %d", len(correct.IndicatorTypes), len(toTest.IndicatorTypes))
		errorDetails = append(errorDetails, str)

		// If lengths are the same, then check each value
		for index, _ := range correct.IndicatorTypes {
			if toTest.IndicatorTypes[index] != correct.IndicatorTypes[index] {
				problemsFound++
				str := fmt.Sprintf("-- Indicator Types Do Not Match: %s | %s", correct.IndicatorTypes[index], toTest.IndicatorTypes[index])
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ Indicator Types Match: %s | %s", correct.IndicatorTypes[index], toTest.IndicatorTypes[index])
				errorDetails = append(errorDetails, str)
			}
		}
	}

	// Check Pattern Value
	if toTest.Pattern != correct.Pattern {
		problemsFound++
		str := fmt.Sprintf("-- Patterns Do Not Match: %s | %s", correct.Pattern, toTest.Pattern)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Patterns Match: %s | %s", correct.Pattern, toTest.Pattern)
		errorDetails = append(errorDetails, str)
	}

	// Check ValidFrom Value
	if toTest.ValidFrom != correct.ValidFrom {
		problemsFound++
		str := fmt.Sprintf("-- ValidFrom Values Do Not Match: %s | %s", correct.ValidFrom, toTest.ValidFrom)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ ValidFrom Values Match: %s | %s", correct.ValidFrom, toTest.ValidFrom)
		errorDetails = append(errorDetails, str)
	}

	// Check ValidUntil Value
	if toTest.ValidUntil != correct.ValidUntil {
		problemsFound++
		str := fmt.Sprintf("-- ValidUntil Values Do Not Match: %s | %s", correct.ValidUntil, toTest.ValidUntil)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ ValidUntil Values Match: %s | %s", correct.ValidUntil, toTest.ValidUntil)
		errorDetails = append(errorDetails, str)
	}

	// Check Kill Chain Phases Property Length
	if len(toTest.KillChainPhases) != len(correct.KillChainPhases) {
		problemsFound++
		str := fmt.Sprintf("-- Kill Chain Phases Length Do Not Match: %d | %d", len(correct.KillChainPhases), len(toTest.KillChainPhases))
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Kill Chain Phases Length Match: %d | %d", len(correct.KillChainPhases), len(toTest.KillChainPhases))
		errorDetails = append(errorDetails, str)
		for index, _ := range correct.KillChainPhases {
			// Check Kill Chain Phases values
			if toTest.KillChainPhases[index].KillChainName != correct.KillChainPhases[index].KillChainName {
				problemsFound++
				str := fmt.Sprintf("-- Kill Chain Names Do Not Match: %s | %s", correct.KillChainPhases[index].KillChainName, toTest.KillChainPhases[index].KillChainName)
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ Kill Chain Names Match: %s | %s", correct.KillChainPhases[index].KillChainName, toTest.KillChainPhases[index].KillChainName)
				errorDetails = append(errorDetails, str)
			}

			// Check Kill Chain Phases values
			if toTest.KillChainPhases[index].PhaseName != correct.KillChainPhases[index].PhaseName {
				problemsFound++
				str := fmt.Sprintf("-- Kill Chain Phases Do Not Match: %s | %s", correct.KillChainPhases[index].PhaseName, toTest.KillChainPhases[index].PhaseName)
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ Kill Chain Phases Match: %s | %s", correct.KillChainPhases[index].PhaseName, toTest.KillChainPhases[index].PhaseName)
				errorDetails = append(errorDetails, str)
			}
		}
	}

	if problemsFound > 0 {
		return false, problemsFound, errorDetails
	}

	return true, 0, errorDetails
}
