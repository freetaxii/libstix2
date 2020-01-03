// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package relationship

import (
	"encoding/json"

	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Type
// ----------------------------------------------------------------------

/* Relationship - This type implements the STIX 2 Relationship SRO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties. */
type Relationship struct {
	objects.CommonObjectProperties
	RelationshipType string `json:"relationship_type,omitempty"`
	properties.DescriptionProperty
	SourceRef string `json:"source_ref,omitempty"`
	TargetRef string `json:"target_ref,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/* New - This function will create a new STIX Relationship object and return
it as a pointer. It will also initialize the object by setting all of the basic
properties. */
func New() *Relationship {
	var obj Relationship
	obj.InitSRO("relationship")
	return &obj
}

// ----------------------------------------------------------------------
// Public Methods - Relationship - Core Functionality
// ----------------------------------------------------------------------

/*
Decode - This function will decode some JSON data encoded as a slice of bytes
into an actual struct. It will return the object as a pointer and any errors found.
*/
func Decode(data []byte) (*Relationship, error) {
	var o Relationship
	err := json.Unmarshal(data, &o)
	if err != nil {
		return nil, err
	}

	if valid, err := o.Valid(); valid != true {
		return nil, err
	}

	o.SetRawData(data)
	return &o, nil
}

/*
Encode - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *Relationship) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
EncodeToString - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *Relationship) EncodeToString() (string, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

/*
Valid - This method will verify all of the properties on the object.
*/
func (o *Relationship) Valid() (bool, error) {

	// Check common base properties first
	if valid, err := o.CommonObjectProperties.Valid(); valid != true {
		return false, err
	}

	return true, nil
}

// ----------------------------------------------------------------------
// Public Methods - Relationship
// ----------------------------------------------------------------------

/*
SetRelationshipType - This method takes in a string value that represents the
type name of the relationship and updates the relationship type property.
*/
func (o *Relationship) SetRelationshipType(s string) error {
	o.RelationshipType = s
	return nil
}

/*
SetSourceRef - This method takes in a string value that represents a STIX
identifier of the source STIX object in the relationship and updates the
source ref property.
*/
func (o *Relationship) SetSourceRef(s string) error {
	o.SourceRef = s
	return nil
}

/*
SetTargetRef - This method takes in a string value that represents a STIX
identifier of the target STIX object in the relationship and updates the
target ref property.
*/
func (o *Relationship) SetTargetRef(s string) error {
	o.TargetRef = s
	return nil
}

/*
SetSourceTarget - This methods takes in two string values where both
represent a STIX identifier. This is a convenience function for setting both
ends of the relationship at the same time. The first identifier is for the
source and the second is for the target.
*/
func (o *Relationship) SetSourceTarget(s, t string) error {
	o.SourceRef = s
	o.TargetRef = t
	return nil
}
