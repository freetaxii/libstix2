// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package relationship

import (
	"github.com/freetaxii/libstix2/objects/baseobject"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
//
// Define Object Type
//
// ----------------------------------------------------------------------

/*
Relationship - This type implements the STIX 2 Relationship SRO and defines
all of the properties methods needed to create and work with the STIX Relationship
SRO. All of the methods not defined local to this type are inherited from
the individual properties.

The following information comes directly from the STIX 2 specification documents.

The Relationship object is used to link together two SDOs in order to describe
how they are related to each other. If SDOs are considered "nodes" or "vertices"
in the graph, the Relationship Objects (SROs) represent "edges".

STIX defines many relationship types to link together SDOs. These relationships
are contained in the "Relationships" table under each SDO definition.
Relationship types defined in the specification SHOULD be used to ensure
consistency. An example of a specification-defined relationship is that an
indicator indicates a campaign. That relationship type is listed in the
Relationships section of the Indicator SDO definition.

STIX also allows relationships from any SDO to any SDO that have not been
defined in this specification. These relationships MAY use the related-to
relationship type or MAY use a custom relationship type. As an example, a user
might want to link malware directly to a tool. They can do so using related-to
to say that the Malware is related to the Tool but not describe how, or they
could use delivered-by (a custom name they determined) to indicate more detail.

Note that some relationships in STIX may seem like "shortcuts". For example, an
Indicator doesn't really detect a Campaign: it detects activity (Attack
Patterns, Malware, etc.) that are often used by that campaign. While some
analysts might want all of the source data and think that shortcuts are
misleading, in many cases it's helpful to provide just the key points
(shortcuts) and leave out the low-level details. In other cases, the low-level
analysis may not be known or sharable, while the high-level analysis is. For
these reasons, relationships that might appear to be "shortcuts" are not
excluded from STIX.
*/
type Relationship struct {
	baseobject.CommonObjectProperties
	RelationshipType string `json:"relationship_type,omitempty"`
	properties.DescriptionProperty
	SourceRef string `json:"source_ref,omitempty"`
	TargetRef string `json:"target_ref,omitempty"`
}

// ----------------------------------------------------------------------
//
// Initialization Functions
//
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Relationship object and return it as
a pointer.
*/
func New() *Relationship {
	var obj Relationship
	obj.InitObject("relationship")
	return &obj
}

// ----------------------------------------------------------------------
//
// Public Methods - Relationship
//
// ----------------------------------------------------------------------

/*
SetRelationship - This method takes in a string value that represents the
type name of the releationship and updates the relationship type property.
*/
func (o *Relationship) SetRelationshipType(s string) error {
	o.RelationshipType = s
	return nil
}

/*
SetSourceRef - This method takes in a string value that represents a STIX
identifier of the source STIX object in the relationship and updates the
source ref property.
*/
func (o *Relationship) SetSourceRef(s string) error {
	o.SourceRef = s
	return nil
}

/*
SetTargetRef - This method takes in a string value that represents a STIX
identifier of the target STIX object in the relationship and updates the
target ref property.
*/
func (o *Relationship) SetTargetRef(s string) error {
	o.TargetRef = s
	return nil
}

/*
SetSourceTarget - This methods takes in two string values where both
represent a STIX identifier. This is a convenience function for setting both
ends of the relationship at the same time. The first identifier is for the
source and the second is for the target.
*/
func (o *Relationship) SetSourceTarget(s, t string) error {
	o.SourceRef = s
	o.TargetRef = t
	return nil
}
