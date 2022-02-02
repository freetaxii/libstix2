// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package attackpattern

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/*
Compare - This method will compare two objects to make sure they are the
same. The receiver is object 1 and the object passed in is object 2. This method
will return a boolean, an integer that tracks the number of problems found, and
a slice of strings that contain the detailed results, whether good or bad.
*/
func (o *AttackPattern) Compare(obj2 *AttackPattern) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check common properties
	_, pBase, dBase := o.CommonObjectProperties.Compare(&obj2.CommonObjectProperties, false)
	problemsFound += pBase
	resultDetails = append(resultDetails, dBase...)

	// Check Name Property
	_, pNames, dNames := o.NameProperty.Compare(&obj2.NameProperty, false)
	problemsFound += pNames
	resultDetails = append(resultDetails, dNames...)

	// Check Description Property
	_, pDescriptions, dDescriptions := o.DescriptionProperty.Compare(&obj2.DescriptionProperty, false)
	problemsFound += pDescriptions
	resultDetails = append(resultDetails, dDescriptions...)

	// Check Aliases Property
	_, pAliases, dAliases := o.AliasesProperty.Compare(&obj2.AliasesProperty, false)
	problemsFound += pAliases
	resultDetails = append(resultDetails, dAliases...)

	// Check KillChainPhases Property
	_, pKillChainPhases, dKillChainPhases := o.KillChainPhasesProperty.Compare(&obj2.KillChainPhasesProperty, false)
	problemsFound += pKillChainPhases
	resultDetails = append(resultDetails, dKillChainPhases...)

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
func Compare(obj1, obj2 *AttackPattern) (bool, int, []string) {
	return obj1.Compare(obj2)
}
