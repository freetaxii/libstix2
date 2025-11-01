// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

/*
Package location implements the STIX 2.1 Attack Pattern object.

The following information comes directly from the STIX 2.1 specification.

This implements STIX 2.1 specification section 4.9.
Reference: https://docs.oasis-open.org/cti/stix/v2.1/csprd01/stix-v2.1-csprd01.html#_Toc16070641

A Location represents a geographic location. The location may be described as
any, some or all of the following: region (e.g., North America), civic address
(e.g. New York, US), latitude and longitude.

Locations are primarily used to give context to other SDOs. For example, a
Location could be used in a relationship to describe that the Bourgeois Swallow
intrusion set originates from Eastern Europe.  The Location SDO can be related
to an Identity or Intrusion Set to indicate that the identity or intrusion set
is located in that location. It can also be related from a malware or attack
pattern to indicate that they target victims in that location. The Location
object describes geographic areas, not governments, even in cases where that
area might have a government. For example, a Location representing the United
States describes the United States as a geographic area, not the federal
government of the United States.

At least one of the following properties/sets of properties MUST be provided:
region, country, latitude and longitude

When a combination of properties is provided (e.g. a region and a latitude and
longitude) the more precise properties are what the location describes. In other
words, if a location contains both a region of northern-america and a country of
us, then the location describes the United States, not all of North America. In
cases where a latitude and longitude are specified without a precision, the
location describes the most precise other value.

If precision is specified, then the datum for latitude and longitude MUST be WGS
84 [WGS84]. Organizations specifying a designated location using latitude and
longitude SHOULD specify the precision which is appropriate for the scope of the
location being identified. The scope is defined by the boundary as outlined by
the precision around the coordinates.
*/
package location
