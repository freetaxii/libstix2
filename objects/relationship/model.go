// Copyright 2015-2022 Bret Jordan, All rights reserved.
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

/*
Relationship - This type implements the STIX 2 Relationship SRO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties.
*/
type Relationship struct {
	objects.CommonObjectProperties
	RelationshipType string `json:"relationship_type,omitempty" bson:"relationship_type,omitempty"`
	properties.DescriptionProperty
	SourceRef string `json:"source_ref,omitempty" bson:"source_ref,omitempty"`
	TargetRef string `json:"target_ref,omitempty" bson:"target_ref,omitempty"`
	StartTime string `json:"start_time,omitempty" bson:"start_time,omitempty"`
	StopTime  string `json:"stop_time,omitempty" bson:"stop_time,omitempty"`
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *Relationship) GetPropertyList() []string {
	return []string{"relationship_type", "description", "source_ref", "target_ref", "start_time", "stop_time"}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Relationship object and return
it as a pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *Relationship {
	var obj Relationship
	obj.InitSRO("relationship")
	return &obj
}
