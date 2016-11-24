// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package common

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type KillChainPhasesType struct {
	Kill_chain_phases []KillChainPhaseType `json:"kill_chain_phases,omitempty"`
}

type KillChainPhaseType struct {
	Kill_chain_name string `json:"kill_chain_name,omitempty"`
	Phase_name      string `json:"phase_name,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - KillChainPhasesType
// ----------------------------------------------------------------------

func (this *KillChainPhasesType) AddKillChainPhase(name, phase string) {
	k := this.newKillChainPhase()
	k.AddName(name)
	k.AddPhase(phase)
}

// ----------------------------------------------------------------------
// Private Methods - KillChainPhaseType
// ----------------------------------------------------------------------

// This method will return a reference to a slice location. This
// will enable us to update an object located at that slice location.
func (this *KillChainPhasesType) newKillChainPhase() *KillChainPhaseType {
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

func (this *KillChainPhaseType) AddName(s string) {
	this.Kill_chain_name = s
}

func (this *KillChainPhaseType) AddPhase(s string) {
	this.Phase_name = s
}
