// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/freetaxii/libstix2/objects/properties"
	"github.com/freetaxii/libstix2/timestamp"
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
	IndicatorTypes []string `json:"indicator_types,omitempty"`
	Pattern        string   `json:"pattern,omitempty"`
	ValidFrom      string   `json:"valid_from,omitempty"`
	ValidUntil     string   `json:"valid_until,omitempty"`
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

/*
Compare - This method will compare two indicators to make sure they
are the same. The indicator receiver is the master and represent the correct
data, the indicator passed in as i represents the one we need to test.
*/
func (o *Indicator) Compare(i *Indicator) (bool, int, []string) {
	problemsFound := 0
	details := make([]string, 0)

	// Check common properties
	if valid, problems, d := o.CommonObjectProperties.Compare(&i.CommonObjectProperties); valid != true {
		problemsFound += problems
		for _, v := range d {
			details = append(details, v)
		}
	} else {
		for _, v := range d {
			details = append(details, v)
		}
	}

	// Check Name Value
	if i.Name != o.Name {
		problemsFound++
		str := fmt.Sprintf("-- Names Do Not Match: %s | %s", o.Name, i.Name)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Names Match: %s | %s", o.Name, i.Name)
		details = append(details, str)
	}

	// Check Description Value
	if i.Description != o.Description {
		problemsFound++
		str := fmt.Sprintf("-- Descriptions Do Not Match: %s | %s", o.Description, i.Description)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Descriptions Match: %s | %s", o.Description, i.Description)
		details = append(details, str)
	}

	// Check Indicator Types Property Length
	if len(i.IndicatorTypes) != len(o.IndicatorTypes) {
		problemsFound++
		str := fmt.Sprintf("-- Indicator Types Length Do Not Match: %d | %d", len(o.IndicatorTypes), len(i.IndicatorTypes))
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Indicator Types Length Match: %d | %d", len(o.IndicatorTypes), len(i.IndicatorTypes))
		details = append(details, str)

		// If lengths are the same, then check each value
		for index, _ := range o.IndicatorTypes {
			if i.IndicatorTypes[index] != o.IndicatorTypes[index] {
				problemsFound++
				str := fmt.Sprintf("-- Indicator Types Do Not Match: %s | %s", o.IndicatorTypes[index], i.IndicatorTypes[index])
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Indicator Types Match: %s | %s", o.IndicatorTypes[index], i.IndicatorTypes[index])
				details = append(details, str)
			}
		}
	}

	// Check Pattern Value
	if i.Pattern != o.Pattern {
		problemsFound++
		str := fmt.Sprintf("-- Patterns Do Not Match: %s | %s", o.Pattern, i.Pattern)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Patterns Match: %s | %s", o.Pattern, i.Pattern)
		details = append(details, str)
	}

	// Check ValidFrom Value
	if i.ValidFrom != o.ValidFrom {
		problemsFound++
		str := fmt.Sprintf("-- ValidFrom Values Do Not Match: %s | %s", o.ValidFrom, i.ValidFrom)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ ValidFrom Values Match: %s | %s", o.ValidFrom, i.ValidFrom)
		details = append(details, str)
	}

	// Check ValidUntil Value
	if i.ValidUntil != o.ValidUntil {
		problemsFound++
		str := fmt.Sprintf("-- ValidUntil Values Do Not Match: %s | %s", o.ValidUntil, i.ValidUntil)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ ValidUntil Values Match: %s | %s", o.ValidUntil, i.ValidUntil)
		details = append(details, str)
	}

	// Check Kill Chain Phases Property Length
	if len(i.KillChainPhases) != len(o.KillChainPhases) {
		problemsFound++
		str := fmt.Sprintf("-- Kill Chain Phases Length Do Not Match: %d | %d", len(o.KillChainPhases), len(i.KillChainPhases))
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Kill Chain Phases Length Match: %d | %d", len(o.KillChainPhases), len(i.KillChainPhases))
		details = append(details, str)
		for index, _ := range o.KillChainPhases {
			// Check Kill Chain Phases values
			if i.KillChainPhases[index].KillChainName != o.KillChainPhases[index].KillChainName {
				problemsFound++
				str := fmt.Sprintf("-- Kill Chain Names Do Not Match: %s | %s", o.KillChainPhases[index].KillChainName, i.KillChainPhases[index].KillChainName)
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Kill Chain Names Match: %s | %s", o.KillChainPhases[index].KillChainName, i.KillChainPhases[index].KillChainName)
				details = append(details, str)
			}

			// Check Kill Chain Phases values
			if i.KillChainPhases[index].PhaseName != o.KillChainPhases[index].PhaseName {
				problemsFound++
				str := fmt.Sprintf("-- Kill Chain Phases Do Not Match: %s | %s", o.KillChainPhases[index].PhaseName, i.KillChainPhases[index].PhaseName)
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Kill Chain Phases Match: %s | %s", o.KillChainPhases[index].PhaseName, i.KillChainPhases[index].PhaseName)
				details = append(details, str)
			}
		}
	}

	if problemsFound > 0 {
		return false, problemsFound, details
	}

	return true, 0, details
}
