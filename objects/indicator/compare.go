// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package indicator

import (
	"fmt"

	"github.com/freetaxii/libstix2/objects/baseobject"
	"github.com/freetaxii/libstix2/objects/properties"
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
	_, pBase, dBase := baseobject.Compare(&obj1.CommonObjectProperties, &obj2.CommonObjectProperties)
	problemsFound += pBase
	resultDetails = append(resultDetails, dBase...)

	// Check Name Values
	_, pNames, dNames := properties.CompareNameProperties(&obj1.NameProperty, &obj2.NameProperty)
	problemsFound += pNames
	resultDetails = append(resultDetails, dNames...)

	_, pDescriptions, dDescriptions := properties.CompareDescriptionProperties(&obj1.DescriptionProperty, &obj2.DescriptionProperty)
	problemsFound += pDescriptions
	resultDetails = append(resultDetails, dDescriptions...)

	// Check Indicator Types Property Lengths
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

	// Check Pattern Values
	if obj1.Pattern != obj2.Pattern {
		problemsFound++
		str := fmt.Sprintf("-- Patterns Do Not Match: %s | %s", obj1.Pattern, obj2.Pattern)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ Patterns Match: %s | %s", obj1.Pattern, obj2.Pattern)
		resultDetails = append(resultDetails, str)
	}

	// Check PatternType Values
	if obj1.PatternType != obj2.PatternType {
		problemsFound++
		str := fmt.Sprintf("-- Pattern Types Do Not Match: %s | %s", obj1.PatternType, obj2.PatternType)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ Pattern Types Match: %s | %s", obj1.PatternType, obj2.PatternType)
		resultDetails = append(resultDetails, str)
	}

	// Check PatternVersion Values
	if obj1.PatternVersion != obj2.PatternVersion {
		problemsFound++
		str := fmt.Sprintf("-- Pattern Versions Do Not Match: %s | %s", obj1.PatternVersion, obj2.PatternVersion)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ Pattern Versions Match: %s | %s", obj1.PatternVersion, obj2.PatternVersion)
		resultDetails = append(resultDetails, str)
	}

	// Check ValidFrom Values
	if obj1.ValidFrom != obj2.ValidFrom {
		problemsFound++
		str := fmt.Sprintf("-- Valid From Values Do Not Match: %s | %s", obj1.ValidFrom, obj2.ValidFrom)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ Valid From Values Match: %s | %s", obj1.ValidFrom, obj2.ValidFrom)
		resultDetails = append(resultDetails, str)
	}

	// Check ValidUntil Values
	if obj1.ValidUntil != obj2.ValidUntil {
		problemsFound++
		str := fmt.Sprintf("-- Valid Until Values Do Not Match: %s | %s", obj1.ValidUntil, obj2.ValidUntil)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ Valid Until Values Match: %s | %s", obj1.ValidUntil, obj2.ValidUntil)
		resultDetails = append(resultDetails, str)
	}

	// Check Kill Chain Phases
	if valid, problems, details := properties.CompareKillChainPhases(&obj1.KillChainPhasesProperty, &obj2.KillChainPhasesProperty); valid != true {
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
