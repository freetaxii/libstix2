// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

// Package bundle - This package defines the properties and methods needed to
// create a work with the STIX Bundle.
//
// The following information comes directly from the STIX 2 specificaton documents:
//
// A Bundle is a collection of arbitrary STIX Objects and Marking Definitions grouped together in a single container. A Bundle does not have any semantic meaning and Objects are not considered related by virtue of being in the same Bundle.
//
// Bundle is not STIX Object, so it does not have any of the Common Properties other than the type and id properties. Bundle is transient and implementations should not assume that other implementations will treat it as a persistent object.
package bundle
