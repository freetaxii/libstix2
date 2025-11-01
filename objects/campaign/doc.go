// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

/*
Package campaign implements the STIX 2.1 Campaign object.

The following information comes directly from the STIX 2.1 specification.

This implements STIX 2.1 specification section 4.2.
Reference: https://docs.oasis-open.org/cti/stix/v2.1/csprd01/stix-v2.1-csprd01.html#_Toc16070634

A Campaign is a grouping of adversarial behaviors that describes a set of
malicious activities or attacks (sometimes called waves) that occur over a
period of time against a specific set of targets. Campaigns usually have well
defined objectives and may be part of an Intrusion Set.

Campaigns are often attributed to an intrusion set and threat actors. The threat
actors may reuse known infrastructure from the intrusion set or may set up new
infrastructure specific for conducting that campaign.

Campaigns can be characterized by their objectives and the incidents they cause,
people or resources they target, and the resources (infrastructure,
intelligence, Malware, Tools, etc.) they use.

For example, a Campaign could be used to describe a crime syndicate's attack
using a specific variant of malware and new C2 servers against the executives of
ACME Bank during the summer of 2016 in order to gain secret information about an
upcoming merger with another bank.
*/
package campaign
