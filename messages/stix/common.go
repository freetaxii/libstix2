// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package stix

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/freetaxii/libstix2/messages/defs"
	"time"
)

// ----------------------------------------------------------------------
// Common Types and Properties
// ----------------------------------------------------------------------

type ExteralReferenceType struct {
	Source_name string `json:"source_name,omitempty"`
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
	External_id string `json:"url,omitempty"`
}

type KillChainPhaseType struct {
	Kill_chain_name string `json:"kill_chain_name,omitempty"`
	Phase_name      string `json:"phase_name,omitempty"`
}

type CommonProperties struct {
	MessageType         string                 `json:"type,omitempty"`
	Id                  string                 `json:"id,omitempty"`
	Created_by_ref      string                 `json:"created_by_ref,omitempty"`
	Created             string                 `json:"created,omitempty"`
	Modified            string                 `json:"modified,omitempty"`
	Version             int                    `json:"version,omitempty"`
	Revoked             bool                   `json:"revoked,omitempty"`
	Labels              []string               `json:"labels,omitempty"`
	External_references []ExteralReferenceType `json:"external_references,omitempty"`
	Object_marking_refs []string               `json:"object_marking_refs,omitempty"`
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

// ----------------------------------------------------------------------
// Public Methods - CommonProperties
// ----------------------------------------------------------------------

func NewId(s string) string {
	// TODO Add check to validate input value
	id := s + "--" + uuid.New()
	return id
}

func GetCurrentTime() time.Time {
	return time.Now()
}

// This function will the time as a string instead of a time.Time object
func GetCurrentTimeText() string {
	t := time.Now().UTC().Format(defs.TIME_RFC_3339)
	return t
}
