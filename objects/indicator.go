// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

import (
	"encoding/json"

	"github.com/freetaxii/libstix2/common/timestamp"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
//
// Define Object Type
//
// ----------------------------------------------------------------------

/*
Indicator - This type implements the STIX 2 Indicator SDO and defines
all of the properties methods needed to create and work with the STIX Indicator
SDO. All of the methods not defined local to this type are inherited from
the individual properties.

The following information comes directly from the STIX 2 specification documents.

Indicators contain a pattern that can be used to detect suspicious or malicious
cyber activity. For example, an Indicator may be used to represent a set of
malicious domains and use the STIX Patterning Language (STIX™ Version 2.0.
Part 5: STIX Patterning) to specify these domains.

The Indicator SDO contains a simple textual description, the Kill Chain Phases
that it detects behavior in, a time window for when the Indicator is valid or
useful, and a required pattern property to capture a structured detection
pattern. Conforming STIX implementations MUST support the STIX Patterning
Language as defined in STIX™ Version 2.0. Part 5: STIX Patterning. While each
structured pattern language has different syntax and potentially different
semantics, in general an Indicator is considered to have "matched" (or been
"sighted") when the conditions specified in the structured pattern are satisfied
in whatever context they are evaluated in.

Relationships from the Indicator can describe the malicious or suspicious
behavior that it directly detects (Malware, Tool, and Attack Pattern) as well
as the Campaigns, Intrusion Sets, and Threat Actors that it might indicate the
presence of.
*/
type Indicator struct {
	properties.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	Pattern    string `json:"pattern,omitempty"`
	ValidFrom  string `json:"valid_from,omitempty"`
	ValidUntil string `json:"valid_until,omitempty"`
	properties.KillChainPhasesProperty
}

// ----------------------------------------------------------------------
//
// Initialization Functions
//
// ----------------------------------------------------------------------

/*
NewIndicator - This function will create a new STIX Indicator object and return
it as a pointer. It will also initialize the object by setting all of the basic
properties.
*/
func NewIndicator() *Indicator {
	var obj Indicator
	obj.InitObject("indicator")
	return &obj
}

/*
DecodeIndicator - This function will decode some JSON data encoded as a slice
of bytes into an actual struct. It will return the object as a pointer.
*/
func DecodeIndicator(data []byte) (*Indicator, error) {
	var o Indicator
	err := json.Unmarshal(data, &o)
	if err != nil {
		return nil, err
	}

	if err := VerifyCommonProperties(o.CommonObjectProperties); err != nil {
		return nil, err
	}

	o.SetRawData(data)
	return &o, nil
}

// ----------------------------------------------------------------------
//
// Public Methods - Indicator
//
// ----------------------------------------------------------------------

/*
New - This method will initialize the object by setting all of the basic properties.
*/
func (o *Indicator) New() {
	o.InitObject("indicator")
}

/*
Encode - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *Indicator) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
SetPattern - This method will take in a string value representing a complete
and valid STIX pattern and set the pattern property to that value.
*/
func (o *Indicator) SetPattern(s string) error {
	o.Pattern = s
	return nil
}

/*
SetValidFromToCurrentTime - This methods sets the valid from time to the
current time
*/
func (o *Indicator) SetValidFromToCurrentTime() error {
	o.ValidFrom = timestamp.GetCurrentTime("micro")
	return nil
}

/*
SetValidFrom - This method will take in a timestamp in either time.Time or
string format and will set the valid from property to that value.
*/
func (o *Indicator) SetValidFrom(t interface{}) error {
	ts, _ := timestamp.ToString(t, "micro")
	o.ValidFrom = ts
	return nil
}

/*
SetValidUntilToCurrentTime - This methods sets the valid until time to the
current time
*/
func (o *Indicator) SetValidUntilToCurrentTime() error {
	o.ValidUntil = timestamp.GetCurrentTime("micro")
	return nil
}

/*
SetValidUntil - This method will take in a timestamp in either time.Time or
string format and will set the valid until property to that value.
*/
func (o *Indicator) SetValidUntil(t interface{}) error {
	ts, _ := timestamp.ToString(t, "micro")

	// TODO check to make sure this is later than the vaild_from
	o.ValidUntil = ts
	return nil
}
