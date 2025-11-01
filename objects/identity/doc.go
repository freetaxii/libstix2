// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

/*
Package identity implements the STIX 2.1 Identity object.

The following information comes directly from the STIX 2.1 specification.

This implements STIX 2.1 specification section 4.5.
Reference: https://docs.oasis-open.org/cti/stix/v2.1/csprd01/stix-v2.1-csprd01.html#_Toc16070637

Identities can represent actual individuals, organizations, or groups (e.g.,
ACME, Inc.) as well as classes of individuals, organizations, systems or groups
(e.g., the finance sector).

The Identity SDO can capture basic identifying information, contact information,
and the sectors that the Identity belongs to. Identity is used in STIX to
represent, among other things, targets of attacks, information sources, object
creators, and threat actor identities.
*/
package identity
