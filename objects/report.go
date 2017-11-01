// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package objects

import (
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
ReportType - This type implements the STIX 2 Report SDO and defines
all of the properties methods needed to create and work with the STIX Report
SDO. All of the methods not defined local to this type are inherited from
the individual properties.

The following information comes directly from the STIX 2 specification documents.

Reports are collections of threat intelligence focused on one or more topics,
such as a description of a threat actor, malware, or attack technique, including
context and related details. They are used to group related threat intelligence
together so that it can be published as a comprehensive cyber threat story.

The Report SDO contains a list of references to SDOs and SROs (the CTI objects
included in the report) along with a textual description and the name of the
report.

For example, a threat report produced by ACME Defense Corp. discussing the Glass
Gazelle campaign should be represented using Report. The Report itself would
contain the narrative of the report while the Campaign SDO and any related SDOs
(e.g., Indicators for the Campaign, Malware it uses, and the associated
Relationships) would be referenced in the report contents.
*/
type ReportType struct {
	properties.CommonObjectPropertiesType
	properties.NamePropertyType
	properties.DescriptionPropertyType
	Published  string   `json:"published,omitempty"`
	ObjectRefs []string `json:"object_refs,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - ReportType
// ----------------------------------------------------------------------

// SetPublished - This method takes in a timestamp in either time.Time or string
// format and updates the published timestamp property.
func (ezt *ReportType) SetPublished(t interface{}) {
	ts := ezt.VerifyTimestamp(t)
	ezt.Published = ts
}

// AddObject - This methods takes in a string value that represents a STIX
// identifier and adds it to the objects ref property.
func (ezt *ReportType) AddObject(s string) {
	ezt.ObjectRefs = append(ezt.ObjectRefs, s)
}
