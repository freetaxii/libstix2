// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

// KillChainPhasesPropertyType - A property used by one or more STIX objects
// that captures a list of kll chain phases as defined by STIX.
type KillChainPhasesPropertyType struct {
	KillChainPhases []KillChainPhaseType `json:"kill_chain_phases,omitempty"`
}

// KillChainPhaseType -
// This type defines all of the properties associated with the STIX Kill Chain Phase type.
type KillChainPhaseType struct {
	KillChainName string `json:"kill_chain_name,omitempty"`
	PhaseName     string `json:"phase_name,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - KillChainPhasesPropertyType
// ----------------------------------------------------------------------

// AddKillChainPhase - This method takes in two parameters and creates a adds
// a new kill chain phase to the list. The first value is a string value
// representing the name of the kill chain being used. The second value is a
// string value representing the phase name from that kill chain.
func (p *KillChainPhasesPropertyType) AddKillChainPhase(name, phase string) {
	k := p.newKillChainPhase()
	k.SetName(name)
	k.SetPhase(phase)
}

// ----------------------------------------------------------------------
// Private Methods - KillChainPhasesPropertyType
// ----------------------------------------------------------------------

// newKillChainPhase - This method returns a reference to a slice location. This
// will enable the code to update an object located at that slice location.
func (p *KillChainPhasesPropertyType) newKillChainPhase() *KillChainPhaseType {
	var s KillChainPhaseType

	if p.KillChainPhases == nil {
		a := make([]KillChainPhaseType, 0)
		p.KillChainPhases = a
	}

	positionThatAppendWillUse := len(p.KillChainPhases)
	p.KillChainPhases = append(p.KillChainPhases, s)
	return &p.KillChainPhases[positionThatAppendWillUse]
}

// ----------------------------------------------------------------------
// Public Methods - KillChainPhaseType
// ----------------------------------------------------------------------

// SetName - This method takes in a string value representing the name of a kill
// chain and updates the kill chain name property.
func (p *KillChainPhaseType) SetName(s string) {
	p.KillChainName = s
}

// SetPhase - This method takes in a string value representing the phase of a
// kill chain and updates the phase name property.
func (p *KillChainPhaseType) SetPhase(s string) {
	p.PhaseName = s
}
