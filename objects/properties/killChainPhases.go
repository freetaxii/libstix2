// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

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
// Public Methods - KillChainPhasesProperty
// ----------------------------------------------------------------------

/*
CreateKillChainPhase - This method takes in two parameters and creates and adds
a new kill chain phase to the list. The first value is a string value
representing the name of the kill chain being used. The second value is a
string value representing the phase name from that kill chain.
*/
func (o *KillChainPhasesProperty) CreateKillChainPhase(name, phase string) error {
	k, _ := o.NewKillChainPhase()
	k.SetName(name)
	k.SetPhase(phase)
	return nil
}

// ----------------------------------------------------------------------
// Private Methods - KillChainPhasesProperty
// ----------------------------------------------------------------------

/*
NewKillChainPhase - This method returns a reference to a slice location. This
will enable the code to update an object located at that slice location.
*/
func (o *KillChainPhasesProperty) NewKillChainPhase() (*KillChainPhase, error) {
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
