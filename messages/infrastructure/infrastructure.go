// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package infrastructure

import (
	"errors"
	"github.com/freetaxii/libstix2/messages/defs"
	"github.com/freetaxii/libstix2/messages/stix"
	"time"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type InfrastructureType struct {
	stix.CommonProperties
	Name                 string                    `json:"name,omitempty"`
	Description          string                    `json:"description,omitempty"`
	Kill_chain_phases    []stix.KillChainPhaseType `json:"kill_chain_phases,omitempty"`
	First_seen           string                    `json:"first_seen,omitempty"`
	First_seen_precision string                    `json:"first_seen_precision,omitempty"`
	Region               string                    `json:"region,omitempty"`
	Country              string                    `json:"country,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() InfrastructureType {
	var obj InfrastructureType
	obj.MessageType = "infrastructure"
	obj.Id = stix.NewId("infrastructure")
	obj.Created = stix.GetCurrentTime().UTC().Format(defs.TIME_RFC_3339)
	obj.Modified = obj.Created
	obj.Version = 1
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - Common Properties
// ----------------------------------------------------------------------

func (this *InfrastructureType) SetCreatedBy(s string) {
	this.Created_by_ref = s
}

func (this *InfrastructureType) SetModified(d time.Time) {
	this.Modified = d.UTC().Format(defs.TIME_RFC_3339)
}

func (this *InfrastructureType) SetVersion(i int) error {
	if i < defs.MIN_VERSION_SIZE {
		return errors.New("No change made, new version is smaller than min size")
	}

	if i > defs.MAX_VERSION_SIZE {
		return errors.New("No change made, new version is larger than max size")
	}

	if i <= this.Version {
		return errors.New("No change made, new version is not larger than original")
	}

	this.Version = i
	return nil
}

func (this *InfrastructureType) SetRevoked() {
	this.Revoked = true
}

func (this *InfrastructureType) AddLabel(value string) {
	if this.Labels == nil {
		a := make([]string, 0)
		this.Labels = a
	}
	this.Labels = append(this.Labels, value)
}

func (this *InfrastructureType) GetId() string {
	return this.Id
}

// ----------------------------------------------------------------------
// Public Methods - InfrastructureType
// ----------------------------------------------------------------------

func (this *InfrastructureType) SetName(s string) {
	this.Name = s
}

func (this *InfrastructureType) SetDescription(s string) {
	this.Description = s
}

func (this *InfrastructureType) AddKillChainPhase(name, phase string) {
	k := this.newKillChainPhase()
	k.AddName(name)
	k.AddPhase(phase)
}

func (this *InfrastructureType) SetFirstSeen(d time.Time) {
	this.First_seen = d.UTC().Format(defs.TIME_RFC_3339)
}

// This function will allow you to assign the time as a string instead of using
// a time.Time object
func (this *InfrastructureType) SetFirstSeenText(s string) {
	this.First_seen = s
}

// TODO Add precision functions

func (this *InfrastructureType) SetRegion(s string) {
	this.Region = s
}

func (this *InfrastructureType) SetCountry(s string) {
	// TODO make sure this is a two digit country code
	this.Country = s
}

// ----------------------------------------------------------------------
// Private Methods - InfrastructureType
// ----------------------------------------------------------------------

// This function will return a reference to a slice location. This
// will enable us to update an object located at that slice location.
func (this *InfrastructureType) newKillChainPhase() *stix.KillChainPhaseType {
	var o stix.KillChainPhaseType

	if this.Kill_chain_phases == nil {
		a := make([]stix.KillChainPhaseType, 0)
		this.Kill_chain_phases = a
	}

	positionThatAppendWillUse := len(this.Kill_chain_phases)
	this.Kill_chain_phases = append(this.Kill_chain_phases, o)
	return &this.Kill_chain_phases[positionThatAppendWillUse]
}
