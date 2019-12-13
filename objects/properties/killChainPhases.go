// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import "fmt"

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
KillChainPhasesProperty - A property used by one or more STIX objects
that captures a list of kll chain phases as defined by STIX.
*/
type KillChainPhasesProperty struct {
	KillChainPhases []KillChainPhase `json:"kill_chain_phases,omitempty"`
}

/*
KillChainPhase - This type defines all of the properties associated with
the STIX Kill Chain Phase type.
*/
type KillChainPhase struct {
	KillChainName string `json:"kill_chain_name,omitempty"`
	PhaseName     string `json:"phase_name,omitempty"`
}

// ----------------------------------------------------------------------
// Public Functions - KillChainPhasesProperty
// ----------------------------------------------------------------------

/*
CompareKillChainPhases - This function will compare two kill chain phases
(object 1 and object 2) to make sure they are the same. This function will
return an integer that tracks the number of problems and a slice of strings that
contain the detailed results, whether good or bad.
*/
func CompareKillChainPhases(obj1, obj2 *KillChainPhasesProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

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

// ----------------------------------------------------------------------
// Public Methods - KillChainPhasesProperty
// ----------------------------------------------------------------------

/*
CreateKillChainPhase - This method takes in two parameters and creates and adds
a new kill chain phase to the list. The first value is a string value
representing the name of the kill chain being used. The second value is a
string value representing the phase name from that kill chain.
*/
func (o *KillChainPhasesProperty) CreateKillChainPhase(name, phase string) error {
	k, _ := o.newKillChainPhase()
	k.SetName(name)
	k.SetPhase(phase)
	return nil
}

// ----------------------------------------------------------------------
// Private Methods - KillChainPhasesProperty
// ----------------------------------------------------------------------

/*
newKillChainPhase - This method returns a reference to a slice location. This
will enable the code to update an object located at that slice location.
*/
func (o *KillChainPhasesProperty) newKillChainPhase() (*KillChainPhase, error) {
	var s KillChainPhase

	// if o.KillChainPhases == nil {
	// 	a := make([]KillChainPhase, 0)
	// 	o.KillChainPhases = a
	// }

	positionThatAppendWillUse := len(o.KillChainPhases)
	o.KillChainPhases = append(o.KillChainPhases, s)
	return &o.KillChainPhases[positionThatAppendWillUse], nil
}

// ----------------------------------------------------------------------
// Public Methods - KillChainPhase
// ----------------------------------------------------------------------

/*
SetName - This method takes in a string value representing the name of a kill
chain and updates the kill chain name property.
*/
func (o *KillChainPhase) SetName(s string) error {
	o.KillChainName = s
	return nil
}

/*
SetPhase - This method takes in a string value representing the phase of a
kill chain and updates the phase name property.
*/
func (o *KillChainPhase) SetPhase(s string) error {
	o.PhaseName = s
	return nil
}
