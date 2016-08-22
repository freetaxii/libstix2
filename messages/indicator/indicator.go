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
	obj.Id = stix.CreateId("indicator")
	obj.Created = stix.GetCurrentTime()
	obj.Modified = obj.Created
	obj.Version = 1
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - Common Properties
// ----------------------------------------------------------------------

func (this *IndicatorType) SetModified(d time.Time) {
	this.Modified = d.UTC().Format(defs.TIME_RFC_3339)
}

func (this *IndicatorType) SetVersion(i int) error {
	if i <= this.Version {
		return errors.New("No change made, new version is not larger than original")
	}

	if i > defs.MAX_VERSION_SIZE {
		return errors.New("No change made, new version is larger than max size")
	}

	this.Version = i
	return nil
}

// ----------------------------------------------------------------------
// Public Methods - IndicatorType
// ----------------------------------------------------------------------

func (this *IndicatorType) SetName(t string) {
	this.Name = t
}

func (this *IndicatorType) SetDescription(t string) {
	this.Description = t
}

func (this *IndicatorType) SetPatternLang(t string) {
	this.Pattern_lang = t
}

// ----------------------------------------------------------------------
// Private Methods - IndicatorType
// ----------------------------------------------------------------------
