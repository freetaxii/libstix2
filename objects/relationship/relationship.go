// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package relationship

import (
	"github.com/freetaxii/libstix2/objects/common"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type RelationshipType struct {
	common.CommonPropertiesType
	Relationship_type string `json:"relationship_type,omitempty"`
	Description       string `json:"description,omitempty"`
	Source_ref        string `json:"source_ref,omitempty"`
	Target_ref        string `json:"target_ref,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() RelationshipType {
	var obj RelationshipType
	obj.MessageType = "relationship"
	obj.Id = obj.NewId("relationship")
	obj.Created = obj.GetCurrentTime()
	obj.Modified = obj.Created
	obj.Version = 1
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - RelationshipType
// ----------------------------------------------------------------------

func (this *RelationshipType) SetRelationshipType(s string) {
	this.Relationship_type = s
}

func (this *RelationshipType) SetDescription(s string) {
	this.Description = s
}

func (this *RelationshipType) SetSourceRef(s string) {
	this.Source_ref = s
}

func (this *RelationshipType) SetTargetRef(s string) {
	this.Target_ref = s
}

// This function is a convenience function for setting both ends of the
// relationship at the same time.
func (this *RelationshipType) SetSourceTarget(s, t string) {
	this.Source_ref = s
	this.Target_ref = t
}
