// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

/*
Package grouping implements the STIX 2.1 Grouping object.

The following information comes directly from the STIX 2.1 specification.

This implements STIX 2.1 specification section 4.4.
Reference: https://docs.oasis-open.org/cti/stix/v2.1/csprd01/stix-v2.1-csprd01.html#_Toc16070636

A Grouping object explicitly asserts that the referenced STIX Objects have a
shared context, unlike a STIX Bundle (which explicitly conveys no context). A
Grouping object should not be confused with an intelligence product, which
should be conveyed via a STIX Report.

A STIX Grouping object might represent a set of data that, in time, given
sufficient analysis, would mature to convey an incident or threat report as a
STIX Report object. For example, a Grouping could be used to characterize an
ongoing investigation into a security event or incident. A Grouping object could
also be used to assert that the referenced STIX Objects are related to an
ongoing analysis process, such as when a threat analyst is collaborating with
others in their trust community to examine a series of Campaigns and Indicators.
The Grouping SDO contains a list of references to SDOs, SCOs, and SROs, along
with an explicit statement of the context shared by the content, a textual
description, and the name of the grouping.
*/
package grouping
