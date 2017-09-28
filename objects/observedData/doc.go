// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

/*
Package observedData implements the STIX 2 Observed Data Domain Object.
This package defines the properties and methods needed to create and work with
the STIX Observed Data SDO.

STIX 2 Specification Text

The following information comes directly from the STIX 2 specificaton documents.

Observed Data conveys information that was observed on systems and networks
using the Cyber Observable specification defined in parts 3 and 4 of this
specification. For example, Observed Data can capture the observation of an IP
address, a network connection, a file, or a registry key. Observed Data is not
an intelligence assertion, it is simply information: this file was seen, without
any context for what it means.

Observed Data captures both a single observation of a single entity (file,
network connection) as well as the aggregation of multiple observations of an
entity. When the number_observed property is 1 the Observed Data is of a single
entity. When the number_observed property is greater than 1, the observed data
consists of several instances of an entity collected over the time window
specified by the first_observed and last_observed properties. When used to
collect aggregate data, it is likely that some fields in the Cyber Observable
Object (e.g., timestamp fields) will be omitted because they would differ for
each of the individual observations.

Observed Data may be used by itself (without relationships) to convey raw data
collected from network and host-based detection tools. A firewall could emit a
single Observed Data instance containing a single Network Traffic object for
each connection it sees. The firewall could also aggregate data and instead send
out an Observed Data instance every ten minutes with an IP address and an
appropriate number_observed value to indicate the number of times that IP
address was observed in that window.

Observed Data may also be related to other SDOs to represent raw data that is
relevant to those objects. The Sighting object, which captures the sighting of
an Indicator, Malware, or other SDO, uses Observed Data to represent the raw
information that led to the creation of the Sighting (e.g., what was actually
seen that suggested that a particular instance of malware was active).
*/
package observedData
