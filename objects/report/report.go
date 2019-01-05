// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package report

import (
	"encoding/json"

	"github.com/freetaxii/libstix2/objects/baseobject"
	"github.com/freetaxii/libstix2/objects/properties"
	"github.com/freetaxii/libstix2/timestamp"
)

// ----------------------------------------------------------------------
//
// Define Object Type
//
// ----------------------------------------------------------------------

/*
Report - This type implements the STIX 2 Report SDO and defines
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
type Report struct {
	baseobject.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	Published  string   `json:"published,omitempty"`
	ObjectRefs []string `json:"object_refs,omitempty"`
}

// ----------------------------------------------------------------------
//
// Initialization Functions
//
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Report object and return it as a
pointer.
*/
func New() *Report {
	var obj Report
	obj.InitObject("report")
	return &obj
}

// ----------------------------------------------------------------------
// Public Methods - Report - Core Functionality
// ----------------------------------------------------------------------

/*
Decode - This function will decode some JSON data encoded as a slice of bytes
into an actual struct. It will return:
 - the object as a pointer
 - the STIX ID
 - the SITX Version
 - any errors found
*/
func Decode(data []byte) (*Report, string, string, error) {
	var o Report
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
func (o *Report) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
EncodeToString - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *Report) EncodeToString() (string, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

/*
Valid - This method will verify all of the properties on the object.
*/
func (o *Report) Valid() (bool, error) {

	// Check common base properties first
	if valid, err := o.CommonObjectProperties.Valid(); valid != true {
		return false, err
	}

	return true, nil
}

// ----------------------------------------------------------------------
// Public Methods - Report
// ----------------------------------------------------------------------

/*
SetPublished - This method takes in a timestamp in either time.Time or string
format and updates the published timestamp property.
*/
func (o *Report) SetPublished(t interface{}) error {
	ts, _ := timestamp.ToString(t, "micro")
	o.Published = ts
	return nil
}

/*
AddObject - This methods takes in a string value that represents a STIX
identifier and adds it to the objects ref property.
*/
func (o *Report) AddObject(s string) error {
	o.ObjectRefs = append(o.ObjectRefs, s)
	return nil
}
