// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

/*
Package report implements the STIX 2.1 Report object.

The following information comes directly from the STIX 2.1 specification.

Reports are collections of threat intelligence focused on one or more topics,
such as a description of a threat actor, malware, or attack technique, including
context and related details. They are used to group related threat intelligence
together so that it can be published as a comprehensive cyber threat story.

The Report SDO contains a list of references to STIX Objects (the CTI objects
included in the report) along with a textual description and the name of the
report.

For example, a threat report produced by ACME Defense Corp. discussing the Glass
Gazelle campaign should be represented using Report. The Report itself would
contain the narrative of the report while the Campaign SDO and any related SDOs
(e.g., Indicators for the Campaign, Malware it uses, and the associated
Relationships) would be referenced in the report contents.
*/
package report
