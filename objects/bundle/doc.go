// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

/*
Package bundle implements the STIX 2.1 Bundle object.

The following information comes directly from the STIX 2.1 specification.

A Bundle is a collection of arbitrary STIX Objects grouped together in a single
container. A Bundle does not have any semantic meaning and the objects contained
within the Bundle are not considered related by virtue of being in the same
Bundle.

A STIX Bundle Object is not a STIX Object but makes use of the type and id
Common Properties. A Bundle is transient, and implementations SHOULD NOT assume
that other implementations will treat it as a persistent object or keep any
custom properties found on the bundle itself.

The JSON MTI serialization uses the JSON Object type [RFC8259] when representing
bundle.
*/
package bundle
