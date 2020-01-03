// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

/*
Package objects implements the STIX 2.1 object model.

The following information comes directly from the STIX 2.1 specification.

This specification defines the set of STIX Domain Objects (SDOs), each of which
corresponds to a unique concept commonly represented in CTI. Using SDOs, STIX
Cyber-observable Objects (SCOs), and STIX Relationship Objects (SROs) as
building blocks, individuals can create and share broad and comprehensive cyber
threat intelligence.

Property information, relationship information, and examples are provided for
each SDO defined below. Property information includes common properties as well
as properties that are specific to each SDO. Relationship information includes
embedded relationships (e.g., created_by_ref), common relationships (e.g.,
related-to), and SDO-specific relationships. Forward relationships (i.e.,
relationships from the SDO to other SDOs or SCOs) are fully defined, while
reverse relationships (i.e., relationships to the SDO from other SDOs or SCOs)
are duplicated for convenience.

Some SDOs are similar and can be grouped together into categories. Attack
Pattern, Malware, and Tool can all be considered types of tactics, techniques,
and procedures (TTPs): they describe behaviors and resources that attackers use
to carry out their attacks. Similarly, Campaign, Intrusion Set, and Threat Actor
all describe information about why adversaries carry out attacks and how they
organize themselves.
*/
package objects
