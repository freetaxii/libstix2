// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sighting

import (
	"encoding/json"

	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Type
// ----------------------------------------------------------------------

/* Sighting - This type implements the STIX 2 Sighting SRO and defines all of
the properties and methods needed to create and work with this object. All of
the methods not defined local to this type are inherited from the individual
properties. */
type Sighting struct {
	objects.CommonObjectProperties
	properties.SeenProperties
	Count            int      `json:"count,omitempty"`
	SightingOfRef    string   `json:"sighting_of_ref,omitempty"`
	ObservedDataRefs []string `json:"observed_data_refs,omitempty"`
	WhereSightedRefs []string `json:"where_sighted_refs,omitempty"`
	Summary          bool     `json:"summary,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/* New - This function will create a new STIX Sighting object and return
it as a pointer. It will also initialize the object by setting all of the basic
properties. */
func New() *Sighting {
	var obj Sighting
	obj.InitSRO("sighting")
	return &obj
}

// ----------------------------------------------------------------------
// Public Methods - Sighting - Core Functionality
// ----------------------------------------------------------------------

/*
Decode - This function will decode some JSON data encoded as a slice of bytes
into an actual struct. It will return the object as a pointer and any errors found.
*/
func Decode(data []byte) (*Sighting, error) {
	var o Sighting
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
func (o *Sighting) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
EncodeToString - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *Sighting) EncodeToString() (string, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

/*
Valid - This method will verify all of the properties on the object.
*/
func (o *Sighting) Valid() (bool, error) {

	// Check common base properties first
	if valid, err := o.CommonObjectProperties.Valid(); valid != true {
		return false, err
	}

	return true, nil
}

// ----------------------------------------------------------------------
// Public Methods - Sighting
// ----------------------------------------------------------------------

/*
SetCount - This method takes in an integer that represents the number of
sightings and upates the count properties.
*/
func (o *Sighting) SetCount(i int) error {
	o.Count = i
	return nil
}

/*
SetSightingOfRef - This method takes in a string value that represents a STIX
identifier of the object that was sighted and updates the sighting of ref
property.
*/
func (o *Sighting) SetSightingOfRef(s string) error {
	o.SightingOfRef = s
	return nil
}

/*
AddObservedDataRef - This method takes in a string value that represents a
STIX identifier of the STIX Observed Data object that identifies what was
sighted and adds it to the observed data refs property.
*/
func (o *Sighting) AddObservedDataRef(s string) error {
	o.ObservedDataRefs = append(o.ObservedDataRefs, s)
	return nil
}

/*
AddWhereSightedRef - This method takes in a string value that represents a
STIX identifier of the STIX Sighting object that identifies where this was
sighted (location, sector, etc) and adds it to the where sighted ref property.
*/
func (o *Sighting) AddWhereSightedRef(s string) error {
	o.WhereSightedRefs = append(o.WhereSightedRefs, s)
	return nil
}

/*
SetSummary - This method set the boolean value of the summary to true.
*/
func (o *Sighting) SetSummary() error {
	o.Summary = true
	return nil
}
