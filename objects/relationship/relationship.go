// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package relationship

import (
	"github.com/freetaxii/libstix2/objects/common/properties"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
RelationshipType defines all of the properties associated with the STIX
Relationship SRO. All of the methods not defined local to this type are
inherited from the individual properties.
*/
type RelationshipType struct {
	properties.CommonObjectPropertiesType
	RelationshipType string `json:"relationship_type,omitempty"`
	properties.DescriptionPropertyType
	SourceRef string `json:"source_ref,omitempty"`
	TargetRef string `json:"target_ref,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// New - This function will create a new relationship object.
func New() RelationshipType {
	var obj RelationshipType
	obj.InitNewObject("relationship")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - RelationshipType
// ----------------------------------------------------------------------

// SetRelationshipType - This method takes in a string value that represents the
// type name of the releationship and updates the relationship type property.
func (ezt *RelationshipType) SetRelationshipType(s string) {
	ezt.RelationshipType = s
}

// SetSourceRef - This method takes in a string value that represents a STIX
// identifier of the source STIX object in the relationship and updates the
// source ref property.
func (ezt *RelationshipType) SetSourceRef(s string) {
	ezt.SourceRef = s
}

// SetTargetRef - This method takes in a string value that represents a STIX
// identifier of the target STIX object in the relationship and updates the
// target ref property.
func (ezt *RelationshipType) SetTargetRef(s string) {
	ezt.TargetRef = s
}

// SetSourceTarget - This methods takes in two string values where both
// represent a STIX identifier. This is a convenience function for setting both
// ends of the relationship at the same time. The first identifier is for the
// source and the second is for the target.
func (ezt *RelationshipType) SetSourceTarget(s, t string) {
	ezt.SourceRef = s
	ezt.TargetRef = t
}
