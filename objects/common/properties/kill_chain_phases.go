// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type KillChainPhasesPropertyType struct {
	Kill_chain_phases []KillChainPhaseType `json:"kill_chain_phases,omitempty"`
}

type KillChainPhaseType struct {
	Kill_chain_name string `json:"kill_chain_name,omitempty"`
	Phase_name      string `json:"phase_name,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - KillChainPhasesPropertyType
// ----------------------------------------------------------------------

// AddKillChainPhase takes in two parameters
// param: name - a string value representing the name of a kill chain
// param: phase - a string value representing the phase of the kill chain
func (this *KillChainPhasesPropertyType) AddKillChainPhase(name, phase string) {
	k := this.newKillChainPhase()
	k.SetName(name)
	k.SetPhase(phase)
}

// ----------------------------------------------------------------------
// Private Methods - KillChainPhasesPropertyType
// ----------------------------------------------------------------------

// newKillChainPhase returns a reference to a slice location. This
// will enable the code to update an object located at that slice location.
func (this *KillChainPhasesPropertyType) newKillChainPhase() *KillChainPhaseType {
	var s KillChainPhaseType

	if this.Kill_chain_phases == nil {
		a := make([]KillChainPhaseType, 0)
		this.Kill_chain_phases = a
	}

	positionThatAppendWillUse := len(this.Kill_chain_phases)
	this.Kill_chain_phases = append(this.Kill_chain_phases, s)
	return &this.Kill_chain_phases[positionThatAppendWillUse]
}

// ----------------------------------------------------------------------
// Public Methods - KillChainPhaseType
// ----------------------------------------------------------------------

// SetName takes in one parameter
// param: s - a string value representing the name of a kill chain
func (this *KillChainPhaseType) SetName(s string) {
	this.Kill_chain_name = s
}

// SetPhase takes in one parameter
// param: s - a string value representing the phase of a kill chain
func (this *KillChainPhaseType) SetPhase(s string) {
	this.Phase_name = s
}
