// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package indicator

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/freetaxii/libstix2/objects/baseobject"
	"github.com/freetaxii/libstix2/objects/properties"
	"github.com/freetaxii/libstix2/timestamp"
)

// ----------------------------------------------------------------------
// Define Object Type
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
	baseobject.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	IndicatorTypes []string `json:"indicator_types,omitempty"`
	Pattern        string   `json:"pattern,omitempty"`
	ValidFrom      string   `json:"valid_from,omitempty"`
	ValidUntil     string   `json:"valid_until,omitempty"`
	properties.KillChainPhasesProperty
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Indicator object and return it as a
pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *Indicator {
	var obj Indicator
	obj.InitObject("indicator")
	return &obj
}

// ----------------------------------------------------------------------
// Public Methods - Indicator - Core Functionality
// ----------------------------------------------------------------------

/*
Decode - This function will decode some JSON data encoded as a slice of bytes
into an actual struct. It will return the object as a pointer, the STIX ID, and
any errors.
*/
func Decode(data []byte) (*Indicator, string, string, error) {
	var o Indicator
	err := json.Unmarshal(data, &o)
	if err != nil {
		return nil, "", "", err
	}

	if valid, err := o.Valid(); valid != true {
		return nil, "", "", err
	}

	o.SetRawData(data)
	return &o, o.ID, o.Modified, nil
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
EncodeToString - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *Indicator) EncodeToString() (string, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

/*
Valid - This method will verify all of the properties on the object.
*/
func (o *Indicator) Valid() (bool, error) {

	// Check common base properties first
	if valid, err := o.CommonObjectProperties.Valid(); valid != true {
		return false, err
	}

	if len(o.IndicatorTypes) == 0 {
		return false, errors.New("the indicator types property is required, but missing")
	}

	if o.Pattern == "" {
		return false, errors.New("the pattern property is required, but missing")
	} else {
		// TODO verify the pattern is correct
	}

	if o.ValidFrom == "" {
		return false, errors.New("the valid from property is required, but missing")
	} else {
		// TODO check to make sure timestamp is a valid STIX timestamp but only if it is defined
	}

	return true, nil
}

// ----------------------------------------------------------------------
// Public Methods - Indicator
// ----------------------------------------------------------------------

/*
AddType - This method takes in a string value representing an indicator
type from the indicator-type-ov and adds it to the indicator type property.
*/
func (o *Indicator) AddType(s string) error {

	indicatorTypes := strings.Split(s, ",")
	for _, iType := range indicatorTypes {
		o.IndicatorTypes = append(o.IndicatorTypes, iType)
	}
	return nil
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
	o.ValidFrom = timestamp.CurrentTime("micro")
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
	o.ValidUntil = timestamp.CurrentTime("micro")
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
