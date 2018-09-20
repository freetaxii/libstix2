// Copyright 2018 Bret Jordan, All rights reserved.
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
	problemsFound := 0
	errorDetails := make([]string, 0)

	// Check common properties
	if valid, problems, d := o.CommonObjectProperties.Compare(&toTest.CommonObjectProperties); valid != true {
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
	if toTest.Name != o.Name {
		problemsFound++
		str := fmt.Sprintf("-- Names Do Not Match: %s | %s", o.Name, toTest.Name)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Names Match: %s | %s", o.Name, toTest.Name)
		errorDetails = append(errorDetails, str)
	}

	// Check Description Value
	if toTest.Description != o.Description {
		problemsFound++
		str := fmt.Sprintf("-- Descriptions Do Not Match: %s | %s", o.Description, toTest.Description)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Descriptions Match: %s | %s", o.Description, toTest.Description)
		errorDetails = append(errorDetails, str)
	}

	// Check Indicator Types Property Length
	if len(toTest.IndicatorTypes) != len(o.IndicatorTypes) {
		problemsFound++
		str := fmt.Sprintf("-- Indicator Types Length Do Not Match: %d | %d", len(o.IndicatorTypes), len(toTest.IndicatorTypes))
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Indicator Types Length Match: %d | %d", len(o.IndicatorTypes), len(toTest.IndicatorTypes))
		errorDetails = append(errorDetails, str)

		// If lengths are the same, then check each value
		for index, _ := range o.IndicatorTypes {
			if toTest.IndicatorTypes[index] != o.IndicatorTypes[index] {
				problemsFound++
				str := fmt.Sprintf("-- Indicator Types Do Not Match: %s | %s", o.IndicatorTypes[index], toTest.IndicatorTypes[index])
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ Indicator Types Match: %s | %s", o.IndicatorTypes[index], toTest.IndicatorTypes[index])
				errorDetails = append(errorDetails, str)
			}
		}
	}

	// Check Pattern Value
	if toTest.Pattern != o.Pattern {
		problemsFound++
		str := fmt.Sprintf("-- Patterns Do Not Match: %s | %s", o.Pattern, toTest.Pattern)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Patterns Match: %s | %s", o.Pattern, toTest.Pattern)
		errorDetails = append(errorDetails, str)
	}

	// Check ValidFrom Value
	if toTest.ValidFrom != o.ValidFrom {
		problemsFound++
		str := fmt.Sprintf("-- ValidFrom Values Do Not Match: %s | %s", o.ValidFrom, toTest.ValidFrom)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ ValidFrom Values Match: %s | %s", o.ValidFrom, toTest.ValidFrom)
		errorDetails = append(errorDetails, str)
	}

	// Check ValidUntil Value
	if toTest.ValidUntil != o.ValidUntil {
		problemsFound++
		str := fmt.Sprintf("-- ValidUntil Values Do Not Match: %s | %s", o.ValidUntil, toTest.ValidUntil)
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ ValidUntil Values Match: %s | %s", o.ValidUntil, toTest.ValidUntil)
		errorDetails = append(errorDetails, str)
	}

	// Check Kill Chain Phases Property Length
	if len(toTest.KillChainPhases) != len(o.KillChainPhases) {
		problemsFound++
		str := fmt.Sprintf("-- Kill Chain Phases Length Do Not Match: %d | %d", len(o.KillChainPhases), len(toTest.KillChainPhases))
		errorDetails = append(errorDetails, str)
	} else {
		str := fmt.Sprintf("++ Kill Chain Phases Length Match: %d | %d", len(o.KillChainPhases), len(toTest.KillChainPhases))
		errorDetails = append(errorDetails, str)
		for index, _ := range o.KillChainPhases {
			// Check Kill Chain Phases values
			if toTest.KillChainPhases[index].KillChainName != o.KillChainPhases[index].KillChainName {
				problemsFound++
				str := fmt.Sprintf("-- Kill Chain Names Do Not Match: %s | %s", o.KillChainPhases[index].KillChainName, toTest.KillChainPhases[index].KillChainName)
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ Kill Chain Names Match: %s | %s", o.KillChainPhases[index].KillChainName, toTest.KillChainPhases[index].KillChainName)
				errorDetails = append(errorDetails, str)
			}

			// Check Kill Chain Phases values
			if toTest.KillChainPhases[index].PhaseName != o.KillChainPhases[index].PhaseName {
				problemsFound++
				str := fmt.Sprintf("-- Kill Chain Phases Do Not Match: %s | %s", o.KillChainPhases[index].PhaseName, toTest.KillChainPhases[index].PhaseName)
				errorDetails = append(errorDetails, str)
			} else {
				str := fmt.Sprintf("++ Kill Chain Phases Match: %s | %s", o.KillChainPhases[index].PhaseName, toTest.KillChainPhases[index].PhaseName)
				errorDetails = append(errorDetails, str)
			}
		}
	}

	if problemsFound > 0 {
		return false, problemsFound, errorDetails
	}

	return true, 0, errorDetails
}
