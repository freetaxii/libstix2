// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

/*
Package indicator implements the STIX 2.1 Indicator object.

The following information comes directly from the STIX 2.1 specification.

Indicators contain a pattern that can be used to detect suspicious or malicious
cyber activity. For example, an Indicator may be used to represent a set of
malicious domains and use the STIX Patterning Language (see section 9) to
specify these domains.

The Indicator SDO contains a simple textual description, the Kill Chain Phases
that it detects behavior in, a time window for when the Indicator is valid or
useful, and a required pattern property to capture a structured detection
pattern. Conforming STIX implementations MUST support the STIX Patterning
Language as defined in section 9.

Relationships from the Indicator can describe the malicious or suspicious
behavior that it directly detects (Malware, Tool, and Attack Pattern). In
addition, it may also imply the presence of a Campaigns, Intrusion Sets, and
Threat Actors, etc.
*/
package indicator
