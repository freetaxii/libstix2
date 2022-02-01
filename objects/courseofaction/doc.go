// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

/*
Package courseofaction implements the STIX 2.1 Course of Action object.

The following information comes directly from the STIX 2.1 specification.

A Course of Action (CoA) is a recommendation from a producer of intelligence to
a consumer on the actions that they might take in response to that intelligence.
The CoA may be preventative to deter exploitation or corrective to counter its
potential impact. The CoA may describe automatable actions (applying patches,
configuring firewalls, etc.), manual processes, or a combination of the two. For
example, a CoA that describes how to remediate a vulnerability could describe
how to apply the patch that removes that vulnerability.

The CoA includes the encoded content of an action or a reference to an
externally defined action identified by the action_type property.
*/
package courseofaction
