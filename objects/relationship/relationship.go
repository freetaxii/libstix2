// Copyright 2017 Bret Jordan, All rights reserved.
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
	common.CommonObjectPropertiesType
	Relationship_type string `json:"relationship_type,omitempty"`
	common.DescriptionPropertiesType
	Source_ref string `json:"source_ref,omitempty"`
	Target_ref string `json:"target_ref,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() RelationshipType {
	var obj RelationshipType
	obj.InitNewObject("relationship")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - RelationshipType
// ----------------------------------------------------------------------

// SetRelationshipType takes in one parameter
// param: s - a string value that represents the type name of the releationship
func (this *RelationshipType) SetRelationshipType(s string) {
	this.Relationship_type = s
}

// SetSourceRef takes in one parameter
// param: s - a string value that represents a STIX identifier to the source STIX object
func (this *RelationshipType) SetSourceRef(s string) {
	this.Source_ref = s
}

// SetTargetRef takes in one parameter
// param: s - a string value that represents a STIX identifier to the target STIX object
func (this *RelationshipType) SetTargetRef(s string) {
	this.Target_ref = s
}

// SetSourceTarget takes in two parameters
// This function is a convenience function for setting both ends of the relationship at the same time.
// param: s - a string value that represents a STIX identifier to the source STIX object
// param: t - a string value that represents a STIX identifier to the target STIX object
func (this *RelationshipType) SetSourceTarget(s, t string) {
	this.Source_ref = s
	this.Target_ref = t
}
