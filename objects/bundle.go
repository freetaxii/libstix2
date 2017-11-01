// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package objects

import (
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
BundleType - This type implements the STIX 2 Bundle and defines
all of the properties methods needed to create and work with the STIX Bundle.
All of the methods not defined local to this type are inherited from
the individual properties.

The following information comes directly from the STIX 2 specification documents.

A Bundle is a collection of arbitrary STIX Objects and Marking Definitions
grouped together in a single container. A Bundle does not have any semantic
meaning and Objects are not considered related by virtue of being in the same
Bundle.

Bundle is not STIX Object, so it does not have any of the Common Properties
other than the type and id properties. Bundle is transient and implementations
should not assume that other implementations will treat it as a persistent
object.
*/
type BundleType struct {
	properties.CommonBundlePropertiesType
	SpecVersion string        `json:"spec_version,omitempty"`
	Objects     []interface{} `json:"objects,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - BundleType
// ----------------------------------------------------------------------

// SetSpecVersion20 - This method will set the specification version to 2.1.
func (ezt *BundleType) SetSpecVersion20() {
	ezt.SpecVersion = "2.0"
}

// AddObject - This method will take in an object as an interface and add it to
// the list of objects in the bundle.
func (ezt *BundleType) AddObject(i interface{}) {
	ezt.Objects = append(ezt.Objects, i)
}
