// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

import (
	"fmt"

	"github.com/freetaxii/libstix2/defs"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// STIXObject - This interface defines what methods an object must have to be
// considered a STIX Object. So any new object that is created that inherits
// the CommonObjectProperties is considered a STIX Object by this code. This
// interface is currently used by the Bundle object to add objects to the
// Bundle.
type STIXObject interface {
	GetCommonProperties() *CommonObjectProperties
}

// CommonObjectProperties - This type defines the properties that are common to
// most STIX objects. If an object does not use all of these properties, then
// the Encode() function for that object will clean up and remove the
// properties that might get populated by mistake. Also, there will be Init
// () functions for each type of STIX object to help with populating the right
// properties for that type of object. This was done so that we would only need
// one type that could be used by all objects, to simplify the code.
type CommonObjectProperties struct {
	DatastoreID        int                    `json:"-" bson:"-"`
	ObjectType         string                 `json:"type,omitempty" bson:"type,omitempty"`
	SpecVersion        string                 `json:"spec_version,omitempty" bson:"spec_version,omitempty"`
	ID                 string                 `json:"id,omitempty" bson:"id,omitempty"`
	CreatedByRef       string                 `json:"created_by_ref,omitempty" bson:"created_by_ref,omitempty"`
	Created            string                 `json:"created,omitempty" bson:"created,omitempty"`
	Modified           string                 `json:"modified,omitempty" bson:"modified,omitempty"`
	Revoked            bool                   `json:"revoked,omitempty" bson:"revoked,omitempty"`
	Labels             []string               `json:"labels,omitempty" bson:"labels,omitempty"`
	Confidence         int                    `json:"confidence,omitempty" bson:"confidence,omitempty"`
	Lang               string                 `json:"lang,omitempty" bson:"lang,omitempty"`
	ExternalReferences []ExternalReference    `json:"external_references,omitempty" bson:"external_references,omitempty"`
	ObjectMarkingRefs  []string               `json:"object_marking_refs,omitempty" bson:"object_marking_refs,omitempty"`
	GranularMarkings   []GranularMarking      `json:"granular_markings,omitempty" bson:"granular_markings,omitempty"`
	Extensions         map[string]interface{} `json:"extensions,omitempty" bson:"extensions,omitempty"`
	Custom             map[string][]byte      `json:"custom,omitempty" bson:"custom,omitempty"`
	Raw                []byte                 `json:"-" bson:"-"`
}

// ExternalReference - This type defines all of the properties associated with
// the STIX External Reference type. All of the methods not defined local to
// this type are inherited from the individual properties.
type ExternalReference struct {
	SourceName  string            `json:"source_name,omitempty" bson:"source_name,omitempty"`
	Description string            `json:"description,omitempty" bson:"description,omitempty"`
	URL         string            `json:"url,omitempty" bson:"url,omitempty"`
	Hashes      map[string]string `json:"hashes,omitempty" bson:"hashes,omitempty"`
	ExternalID  string            `json:"external_id,omitempty" bson:"external_id,omitempty"`
}

// GranularMarking - This type defines all of the properties associated with the
// STIX Granular Marking type. All of the methods not defined local to this type
// are inherited from the individual properties.
type GranularMarking struct {
	Lang       string   `json:"lang,omitempty" bson:"lang,omitempty"`
	MarkingRef string   `json:"marking_ref,omitempty" bson:"marking_ref,omitempty"`
	Selectors  []string `json:"selectors,omitempty" bson:"selectors,omitempty"`
}

// GetCommonPropertyList - This method will return a list of all of the
// properties that are common to all objects. This is used by the
// FindCustomProperties method. It is defined here in this file to make it easy
// to keep in sync as new properties are added.
func (o *CommonObjectProperties) GetCommonPropertyList() []string {
	return []string{
		"type",
		"spec_version",
		"id",
		"created_by_ref",
		"created",
		"modified",
		"revoked",
		"labels",
		"confidence",
		"lang",
		"external_references",
		"object_marking_refs",
		"granular_markings",
		"extensions",
	}
}

// This type is used to capture results from the Valid() and Compare() functions
type results struct {
	debug         bool
	problemsFound int
	resultDetails []string
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
InitSDO - This method will initialize a STIX Domain Object by setting all
of the basic properties and is called by the New() function from each object.
*/
func (o *CommonObjectProperties) InitSDO(objectType string) error {
	if defs.STRICT_TYPES {
		if valid := ValidObjectType(objectType); valid != true {
			return fmt.Errorf("invalid object type for InitSDO with strict checks enabled")
		}
	}

	o.SetSpecVersion(defs.CurrentSTIXVersion)
	o.SetObjectType(objectType)
	o.SetNewSTIXID(objectType)
	o.SetCreatedToCurrentTime()
	o.SetModified(o.GetCreated())
	return nil
}

/*
InitSRO - This method will initialize a STIX Relationship Object by setting
all of the basic properties and is called by the New() function from each
object.
*/
func (o *CommonObjectProperties) InitSRO(objectType string) error {
	if defs.STRICT_TYPES {
		if valid := ValidObjectType(objectType); valid != true {
			return fmt.Errorf("invalid object type for InitSRO with strict checks enabled")
		}
	}

	o.SetSpecVersion(defs.CurrentSTIXVersion)
	o.SetObjectType(objectType)
	o.SetNewSTIXID(objectType)
	o.SetCreatedToCurrentTime()
	o.SetModified(o.GetCreated())
	return nil
}

/*
InitSCO - This method will initialize a STIX Cyber Observable Object by
setting all of the basic properties and is called by the New() function from
each object.
*/
func (o *CommonObjectProperties) InitSCO(objectType string) error {
	if defs.STRICT_TYPES {
		if valid := ValidObjectType(objectType); valid != true {
			return fmt.Errorf("invalid object type for InitSCO with strict checks enabled")
		}
	}

	o.SetSpecVersion(defs.CurrentSTIXVersion)
	o.SetObjectType(objectType)
	o.SetNewSTIXID(objectType)
	return nil
}

/*
InitBundle - This method will initialize a STIX Bundle by setting all of the
basic properties and is called by the New() function from that object.
*/
func (o *CommonObjectProperties) InitBundle() error {
	o.SetObjectType("bundle")
	o.SetNewSTIXID("bundle")
	return nil
}
