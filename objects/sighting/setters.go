// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sighting

import (
	"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
// Public Methods
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
AddObservedDataRefs - This method takes in a string value, a comma separated
list of string values, or a slice of string values that represents an id of an
observed data object that identifies what was sighted and adds it to the
observed data refs property.
*/
func (o *Sighting) AddObservedDataRefs(values interface{}) error {
	return objects.AddValuesToList(&o.ObservedDataRefs, values)
}

/*
AddWhereSightedRefs - This method takes in a string value, a comma separated
list of string values, or a slice of string values that represents an id of a
location object that identifies where this was sighted (location, sector, etc)
and adds it to the where sighted refs property.
*/
func (o *Sighting) AddWhereSightedRefs(values interface{}) error {
	return objects.AddValuesToList(&o.WhereSightedRefs, values)
}

/*
SetSummary - This method set the boolean value of the summary to true.
*/
func (o *Sighting) SetSummary() error {
	o.Summary = true
	return nil
}
