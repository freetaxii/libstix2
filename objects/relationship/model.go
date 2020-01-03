// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package relationship

import (
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Type
// ----------------------------------------------------------------------

/* Relationship - This type implements the STIX 2 Relationship SRO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties. */
type Relationship struct {
	objects.CommonObjectProperties
	RelationshipType string `json:"relationship_type,omitempty"`
	properties.DescriptionProperty
	SourceRef string `json:"source_ref,omitempty"`
	TargetRef string `json:"target_ref,omitempty"`
	StartTime string `json:"start_time,omitempty"`
	StopTime  string `json:"stop_time,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/* New - This function will create a new STIX Relationship object and return
it as a pointer. It will also initialize the object by setting all of the basic
properties. */
func New() *Relationship {
	var obj Relationship
	obj.InitSRO("relationship")
	return &obj
}
