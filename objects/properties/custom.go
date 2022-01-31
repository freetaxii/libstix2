// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/*
CustomProperties - A property used by all STIX objects that captures any
custom properties. These are all stored in a map.
*/
type CustomProperties struct {
	Custom map[string][]byte `json:"custom,omitempty" bson:"custom,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - CreatedProperty - Setters
// ----------------------------------------------------------------------
