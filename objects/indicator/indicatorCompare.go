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
are the same. The indicator receiver is the master and represent the correct
data, the indicator passed in as i represents the one we need to test.
*/
func (o *Indicator) Compare(i *Indicator) (bool, int, []string) {
	problemsFound := 0
	details := make([]string, 0)

	// Check common properties
	if valid, problems, d := o.CommonObjectProperties.Compare(&i.CommonObjectProperties); valid != true {
		problemsFound += problems
		for _, v := range d {
			details = append(details, v)
		}
	} else {
		for _, v := range d {
			details = append(details, v)
		}
	}

	// Check Name Value
	if i.Name != o.Name {
		problemsFound++
		str := fmt.Sprintf("-- Names Do Not Match: %s | %s", o.Name, i.Name)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Names Match: %s | %s", o.Name, i.Name)
		details = append(details, str)
	}

	// Check Description Value
	if i.Description != o.Description {
		problemsFound++
		str := fmt.Sprintf("-- Descriptions Do Not Match: %s | %s", o.Description, i.Description)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Descriptions Match: %s | %s", o.Description, i.Description)
		details = append(details, str)
	}

	// Check Indicator Types Property Length
	if len(i.IndicatorTypes) != len(o.IndicatorTypes) {
		problemsFound++
		str := fmt.Sprintf("-- Indicator Types Length Do Not Match: %d | %d", len(o.IndicatorTypes), len(i.IndicatorTypes))
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Indicator Types Length Match: %d | %d", len(o.IndicatorTypes), len(i.IndicatorTypes))
		details = append(details, str)

		// If lengths are the same, then check each value
		for index, _ := range o.IndicatorTypes {
			if i.IndicatorTypes[index] != o.IndicatorTypes[index] {
				problemsFound++
				str := fmt.Sprintf("-- Indicator Types Do Not Match: %s | %s", o.IndicatorTypes[index], i.IndicatorTypes[index])
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Indicator Types Match: %s | %s", o.IndicatorTypes[index], i.IndicatorTypes[index])
				details = append(details, str)
			}
		}
	}

	// Check Pattern Value
	if i.Pattern != o.Pattern {
		problemsFound++
		str := fmt.Sprintf("-- Patterns Do Not Match: %s | %s", o.Pattern, i.Pattern)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Patterns Match: %s | %s", o.Pattern, i.Pattern)
		details = append(details, str)
	}

	// Check ValidFrom Value
	if i.ValidFrom != o.ValidFrom {
		problemsFound++
		str := fmt.Sprintf("-- ValidFrom Values Do Not Match: %s | %s", o.ValidFrom, i.ValidFrom)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ ValidFrom Values Match: %s | %s", o.ValidFrom, i.ValidFrom)
		details = append(details, str)
	}

	// Check ValidUntil Value
	if i.ValidUntil != o.ValidUntil {
		problemsFound++
		str := fmt.Sprintf("-- ValidUntil Values Do Not Match: %s | %s", o.ValidUntil, i.ValidUntil)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ ValidUntil Values Match: %s | %s", o.ValidUntil, i.ValidUntil)
		details = append(details, str)
	}

	// Check Kill Chain Phases Property Length
	if len(i.KillChainPhases) != len(o.KillChainPhases) {
		problemsFound++
		str := fmt.Sprintf("-- Kill Chain Phases Length Do Not Match: %d | %d", len(o.KillChainPhases), len(i.KillChainPhases))
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Kill Chain Phases Length Match: %d | %d", len(o.KillChainPhases), len(i.KillChainPhases))
		details = append(details, str)
		for index, _ := range o.KillChainPhases {
			// Check Kill Chain Phases values
			if i.KillChainPhases[index].KillChainName != o.KillChainPhases[index].KillChainName {
				problemsFound++
				str := fmt.Sprintf("-- Kill Chain Names Do Not Match: %s | %s", o.KillChainPhases[index].KillChainName, i.KillChainPhases[index].KillChainName)
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Kill Chain Names Match: %s | %s", o.KillChainPhases[index].KillChainName, i.KillChainPhases[index].KillChainName)
				details = append(details, str)
			}

			// Check Kill Chain Phases values
			if i.KillChainPhases[index].PhaseName != o.KillChainPhases[index].PhaseName {
				problemsFound++
				str := fmt.Sprintf("-- Kill Chain Phases Do Not Match: %s | %s", o.KillChainPhases[index].PhaseName, i.KillChainPhases[index].PhaseName)
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Kill Chain Phases Match: %s | %s", o.KillChainPhases[index].PhaseName, i.KillChainPhases[index].PhaseName)
				details = append(details, str)
			}
		}
	}

	if problemsFound > 0 {
		return false, problemsFound, details
	}

	return true, 0, details
}
