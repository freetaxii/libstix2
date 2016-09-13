// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package indicator

import (
	"errors"
	"github.com/freetaxii/libstix2/messages/defs"
	"github.com/freetaxii/libstix2/messages/stix"
	"time"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type IndicatorType struct {
	stix.CommonProperties
	Name                  string                    `json:"name,omitempty"`
	Description           string                    `json:"description,omitempty"`
	Pattern_lang          string                    `json:"pattern_lang,omitempty"`
	Pattern_lang_version  string                    `json:"pattern_lang_version,omitempty"`
	Pattern               string                    `json:"pattern,omitempty"`
	Valid_from            string                    `json:"valid_from,omitempty"`
	Valid_from_precision  string                    `json:"valid_from_precision,omitempty"`
	Valid_until           string                    `json:"valid_until,omitempty"`
	Valid_until_precision string                    `json:"valid_until_precision,omitempty"`
	Kill_chain_phases     []stix.KillChainPhaseType `json:"kill_chain_phases,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() IndicatorType {
	var obj IndicatorType
	obj.MessageType = "indicator"
	obj.Id = stix.NewId("indicator")
	obj.Created = stix.GetCurrentTime().UTC().Format(defs.TIME_RFC_3339)
	obj.Modified = obj.Created
	obj.Version = 1
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - Common Properties
// ----------------------------------------------------------------------

func (this *IndicatorType) SetCreatedBy(s string) {
	this.Created_by_ref = s
}

func (this *IndicatorType) SetModified(d time.Time) {
	this.Modified = d.UTC().Format(defs.TIME_RFC_3339)
}

func (this *IndicatorType) SetVersion(i int) error {
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

func (this *IndicatorType) SetRevoked() {
	this.Revoked = true
}

// ----------------------------------------------------------------------
// Public Methods - IndicatorType
// ----------------------------------------------------------------------

func (this *IndicatorType) SetName(s string) {
	this.Name = s
}

func (this *IndicatorType) SetDescription(s string) {
	this.Description = s
}

func (this *IndicatorType) SetPatternLang(s string) {
	this.Pattern_lang = s
}

func (this *IndicatorType) SetPatternLangVersion(s string) {
	this.Pattern_lang_version = s
}

func (this *IndicatorType) SetValidFrom(d time.Time) {
	this.Valid_from = d.UTC().Format(defs.TIME_RFC_3339)
}

func (this *IndicatorType) SetValidFromText(s string) {
	this.Valid_from = s
}

func (this *IndicatorType) SetValidUntil(d time.Time) {
	this.Valid_until = d.UTC().Format(defs.TIME_RFC_3339)
}

func (this *IndicatorType) SetValidUntilText(s string) {
	this.Valid_until = s
}

func (this *IndicatorType) NewKillChainPhase() *stix.KillChainPhaseType {
	var s stix.KillChainPhaseType
	slicePosition := this.addKillChainPhase(s)
	return &this.Kill_chain_phases[slicePosition]
}

func (this *IndicatorType) AddKillChainPhase(name, phase string) {
	k := this.NewKillChainPhase()
	k.AddName(name)
	k.AddPhase(phase)
}

// ----------------------------------------------------------------------

func (this *IndicatorType) SetPrecisionYear(s string) {
	if s == "valid_from" {
		this.Valid_from_precision = "year"
	} else if s == "valid_until" {
		this.Valid_until_precision = "year"
	}
}

func (this *IndicatorType) SetPrecisionMonth(s string) {
	if s == "valid_from" {
		this.Valid_from_precision = "month"
	} else if s == "valid_until" {
		this.Valid_until_precision = "month"
	}
}

func (this *IndicatorType) SetPrecisionDay(s string) {
	if s == "valid_from" {
		this.Valid_from_precision = "day"
	} else if s == "valid_until" {
		this.Valid_until_precision = "day"
	}
}

func (this *IndicatorType) SetPrecisionHour(s string) {
	if s == "valid_from" {
		this.Valid_from_precision = "hour"
	} else if s == "valid_until" {
		this.Valid_until_precision = "hour"
	}
}

func (this *IndicatorType) SetPrecisionMinute(s string) {
	if s == "valid_from" {
		this.Valid_from_precision = "minute"
	} else if s == "valid_until" {
		this.Valid_until_precision = "minute"
	}
}

func (this *IndicatorType) SetPrecisionFull(s string) {
	if s == "valid_from" {
		this.Valid_from_precision = "full"
	} else if s == "valid_until" {
		this.Valid_until_precision = "full"
	}
}

// ----------------------------------------------------------------------
// Private Methods - IndicatorType
// ----------------------------------------------------------------------

func (this *IndicatorType) addKillChainPhase(s stix.KillChainPhaseType) int {
	if this.Kill_chain_phases == nil {
		a := make([]stix.KillChainPhaseType, 0)
		this.Kill_chain_phases = a
	}
	positionThatAppendWillUse := len(this.Kill_chain_phases)
	this.Kill_chain_phases = append(this.Kill_chain_phases, s)
	return positionThatAppendWillUse
}
