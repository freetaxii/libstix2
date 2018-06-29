// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

import (
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
//
// Define Object Type
//
// ----------------------------------------------------------------------

/*
Sighting - This type implements the STIX 2 Sighting SRO and defines
all of the properties methods needed to create and work with the STIX Sighting
SRO. All of the methods not defined local to this type are inherited from
the individual properties.

The following information comes directly from the STIX 2 specification documents.

A Sighting denotes the belief that something in CTI (e.g., an indicator,
malware, tool, threat actor, etc.) was seen. Sightings are used to track who and
what are being targeted, how attacks are carried out, and to track trends in
attack behavior.

The Sighting relationship object is a special type of SRO; it is a relationship
that contains extra properties not present on the generic Relationship object.
These extra properties are included to represent data specific to sighting
relationships (e.g., count, representing how many times something was seen), but
for other purposes a Sighting can be thought of as a Relationship with a name of
"sighting-of". Sighting is captured as a relationship because you cannot have a
sighting unless you have something that has been sighted. Sighting does not make
sense without the relationship to what was sighted.

Sighting relationships relate three aspects of the sighting:
  * What was sighted, such as the Indicator, Malware, Campaign, or other SDO (sighting_of_ref)
  * Who sighted it and/or where it was sighted, represented as an Identity (where_sighted_refs) and
  * What was actually seen on systems and networks, represented as Observed Data (observed_data_refs)

What was sighted is required; a sighting does not make sense unless you say what
you saw. Who sighted it, where it was sighted, and what was actually seen are
optional. In many cases it is not necessary to provide that level of detail in
order to provide value.

Sightings are used whenever any SDO has been "seen". In some cases, the object
creator wishes to convey very little information about the sighting; the details
might be sensitive, but the fact that they saw a malware instance or threat actor
could still be very useful. In other cases, providing the details may be helpful
or even necessary; saying exactly which of the 1000 IP addresses in an indicator
were sighted is helpful when tracking which of those IPs is still malicious.

Sighting is distinct from Observed Data in that Sighting is an intelligence
assertion ("I saw this threat actor") while Observed Data is simply information
("I saw this file"). When you combine them by including the linked Observed Data
(observed_data_refs) from a Sighting, you can say "I saw this file, and that
makes me think I saw this threat actor". Although confidence is currently
reserved, notionally confidence would be added to Sighting (the intelligence
relationship) but not to Observed Data (the raw information).
*/
type Sighting struct {
	properties.CommonObjectProperties
	properties.FirstSeenProperty
	properties.LastSeenProperty
	Count            int      `json:"count,omitempty"`
	SightingOfRef    string   `json:"sighting_of_ref,omitempty"`
	ObservedDataRefs []string `json:"observed_data_refs,omitempty"`
	WhereSightedRefs []string `json:"where_sighted_refs,omitempty"`
	Summary          bool     `json:"summary,omitempty"`
}

// ----------------------------------------------------------------------
//
// Initialization Functions
//
// ----------------------------------------------------------------------

/*
NewSighting - This function will create a new STIX Sighting object and return
it as a pointer.
*/
func NewSighting() *Sighting {
	var obj Sighting
	obj.InitObject("sighting")
	return &obj
}

// ----------------------------------------------------------------------
//
// Public Methods - Sighting
//
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
STIX identifier of the STIX Identity object that identifies where this was
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
