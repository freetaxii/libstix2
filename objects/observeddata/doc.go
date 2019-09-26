// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

/*
Package observeddata implements the STIX 2.1 Observed Data object.

The following information comes directly from the STIX 2.1 specification.

Observed Data conveys information about cyber security related entities such as
files, systems, and networks using the STIX Cyber-observable Objects (SCOs). For
example, Observed Data can capture information about an IP address, a network
connection, a file, or a registry key. Observed Data is not an intelligence
assertion, it is simply the raw information without any context for what it
means.

Observed Data can capture that a piece of information was seen one or more
times. Meaning, it can capture both a single observation of a single entity
(file, network connection) as well as the aggregation of multiple observations
of an entity. When the number_observed property is 1 the Observed Data
represents a single entity. When the number_observed property is greater than 1,
the Observed Data represents several instances of an entity potentially
collected over a period of time. If a time window is known, that can be captured
using the first_observed and last_observed properties. When used to collect
aggregate data, it is likely that some properties in the SCO (e.g., timestamp
properties) will be omitted because they would differ for each of the individual
observations.

Observed Data may be used by itself (without relationships) to convey raw data
collected from any source including analyst reports, sandboxes, and network and
host-based detection tools. An intelligence producer conveying Observed Data
SHOULD include as much context (e.g. SCOs) as possible that supports the use of
the observed data set in systems expecting to utilize the Observed Data for
improved security. This includes all SCOs that matched on an Indicator pattern
and are represented in the collected observed event (or events) being conveyed
in the Observed Data object. For example, a firewall could emit a single
Observed Data instance containing a single Network Traffic object for each
connection it sees. The firewall could also aggregate data and instead send out
an Observed Data instance every ten minutes with an IP address and an
appropriate number_observed value to indicate the number of times that IP
address was observed in that window. A sandbox could emit an Observed Data
instance containing a file hash that it discovered.

Observed Data may also be related to other SDOs to represent raw data that is
relevant to those objects. For example, the Sighting Relationship object, can
relate an Indicator, Malware, or other SDO to a specific Observed Data to
represent the raw information that led to the creation of the Sighting (e.g.,
what was actually seen that suggested that a particular instance of malware was
active).

To support backwards compatibility, related SCOs can still be specified using
the objects properties, Either the objects property or the object_refs property
MUST be provided, but both MUST NOT be present at the same time.
*/
package observeddata
