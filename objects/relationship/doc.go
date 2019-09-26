// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

/*
Package relationship implements the STIX 2.1 Relationship object.

The following information comes directly from the STIX 2.1 specification.

The Relationship object is used to link together two SDOs or SCOs in order to
describe how they are related to each other. If SDOs and SCOs are considered
"nodes" or "vertices" in the graph, the Relationship Objects (SROs) represent
"edges".

STIX defines many relationship types to link together SDOs and SCOs. These
relationships are contained in the "Relationships" table under each SDO and SCO
definition. Relationship types defined in the specification SHOULD be used to
ensure consistency. An example of a specification-defined relationship is that
an indicator indicates a campaign. That relationship type is listed in the
Relationships section of the Indicator SDO definition.

STIX also allows relationships from any SDO or SCO to any SDO or SCO that have
not been defined in this specification. These relationships MAY use the
related-to relationship type or MAY use a user-defined relationship type. As an
example, a user might want to link malware directly to a tool. They can do so
using related-to to say that the Malware is related to the Tool but not describe
how, or they could use delivered-by (a user-defined name they determined) to
indicate more detail.

Note that some relationships in STIX may seem like "shortcuts". For example, an
Indicator doesn't really detect a Campaign: it detects activity (Attack
Patterns, Malware, Infrastructure, etc.) that are often used by that campaign.
While some analysts might want all of the source data and think that shortcuts
are misleading, in many cases it's helpful to provide just the key points
(shortcuts) and leave out the low-level details. In other cases, the low-level
analysis may not be known or sharable, while the high-level analysis is. For
these reasons, relationships that might appear to be "shortcuts" are not
excluded from STIX.
*/
package relationship
