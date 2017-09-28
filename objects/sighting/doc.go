// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

/*
Package sighting implements the STIX 2 Sighting Relationship Object.
This package defines the properties and methods needed to create and work with
the STIX Sighting SRO.

STIX 2 Specification Text

The following information comes directly from the STIX 2 specificaton documents.

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
package sighting
