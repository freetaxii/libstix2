// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package relationship

import (
	"errors"
	"github.com/freetaxii/libstix2/messages/defs"
	"github.com/freetaxii/libstix2/messages/stix"
	"time"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type RelationshipType struct {
	stix.CommonProperties
	Relationship_type string `json:"name,relationship_type"`
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
	obj.Id = stix.NewId("relationship")
	obj.Created = stix.GetCurrentTime().UTC().Format(defs.TIME_RFC_3339)
	obj.Modified = obj.Created
	obj.Version = 1
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - Common Properties
// ----------------------------------------------------------------------

func (this *RelationshipType) SetCreatedBy(s string) {
	this.Created_by_ref = s
}

func (this *RelationshipType) SetModified(d time.Time) {
	this.Modified = d.UTC().Format(defs.TIME_RFC_3339)
}

func (this *RelationshipType) SetVersion(i int) error {
	if i < defs.MIN_VERSION_SIZE {
		return errors.New("No change made, new version is smaller than min size")
	}

	if i > defs.MAX_VERSION_SIZE {
		return errors.New("No change made, new version is larger than max size")
	}

	if i <= this.Version {
		return errors.New("No change made, new version is not larger than original")
	}

	this.Version = i
	return nil
}

func (this *RelationshipType) SetRevoked() {
	this.Revoked = true
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
