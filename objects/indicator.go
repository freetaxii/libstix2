// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package objects

import (
	"github.com/freetaxii/libstix2/common/timestamp"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Indicator Type
// ----------------------------------------------------------------------

/*
IndicatorType - This type implements the STIX 2 Indicator SDO and defines
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
type IndicatorType struct {
	properties.CommonObjectPropertiesType
	properties.NamePropertyType
	properties.DescriptionPropertyType
	Pattern    string `json:"pattern,omitempty"`
	ValidFrom  string `json:"valid_from,omitempty"`
	ValidUntil string `json:"valid_until,omitempty"`
	properties.KillChainPhasesPropertyType
}

// ----------------------------------------------------------------------
// Public Methods - IndicatorType
// ----------------------------------------------------------------------

// SetPattern - This method will take in a string value representing a complete
// and valid STIX pattern and set the pattern property to that value.
func (ezt *IndicatorType) SetPattern(s string) {
	ezt.Pattern = s
}

// SetValidFromToCurrentTime - This methods sets the valid from time to the
// current time
func (ezt *IndicatorType) SetValidFromToCurrentTime() {
	ezt.ValidFrom = timestamp.GetCurrentTime("micro")
}

// SetValidFrom - This method will take in a timestamp in either time.Time or
// string format and will set the valid from property to that value.
func (ezt *IndicatorType) SetValidFrom(t interface{}) {
	ts, _ := timestamp.ToString(t, "micro")
	ezt.ValidFrom = ts
}

// SetValidUntilToCurrentTime - This methods sets the valid until time to the
// current time
func (ezt *IndicatorType) SetValidUntilToCurrentTime() {
	ezt.ValidUntil = timestamp.GetCurrentTime("micro")
}

// SetValidUntil - This method will take in a timestamp in either time.Time or
// string format and will set the valid until property to that value.
func (ezt *IndicatorType) SetValidUntil(t interface{}) {
	ts, _ := timestamp.ToString(t, "micro")

	// TODO check to make sure this is later than the vaild_from
	ezt.ValidUntil = ts
}
