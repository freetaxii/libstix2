// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package common

import (
	"github.com/freetaxii/libstix2/objects/common/properties"
	"github.com/freetaxii/libstix2/objects/common/timestamp"
)

// ----------------------------------------------------------------------
// Common Property Types - Used to populate the common object properties
// ----------------------------------------------------------------------

// This type includes all of the common properties that are used by all objects
// The only object that uses this one by itself is the STIX Bundle
type CommonBasePropertiesType struct {
	properties.MessageTypePropertyType
	properties.IdPropertyType
}

type CommonCreatorPropertiesType struct {
	properties.CreatedByRefPropertyType
	properties.CreatedPropertyType
}

type CommonVersionPropertiesType struct {
	properties.ModifiedPropertyType
	properties.RevokedPropertyType
}

// This type includes all of the common properties that are used by
// SDOs and SROs
type CommonObjectPropertiesType struct {
	CommonBasePropertiesType
	CommonCreatorPropertiesType
	CommonVersionPropertiesType
	properties.LabelsPropertyType
	properties.ConfidencePropertyType
	properties.LangPropertyType
	properties.ExternalReferencesPropertyType
	properties.ObjectMarkingRefsPropertyType
	properties.GranularMarkingsPropertyType
}

// This type includes all of the common properties that are used by
// the Marking Definition object
type CommonMarkingDefinitionPropertiesType struct {
	CommonBasePropertiesType
	CommonCreatorPropertiesType
	properties.ExternalReferencesPropertyType
	properties.ObjectMarkingRefsPropertyType
	properties.GranularMarkingsPropertyType
}

// ----------------------------------------------------------------------
// Property Types - Helper Functions
// ----------------------------------------------------------------------
// These functions are helper functions to prevent needing to import the
// properties object locally

type AliasesPropertiesType struct {
	properties.AliasesPropertyType
}

type DescriptionPropertiesType struct {
	properties.DescriptionPropertyType
}

type DescriptivePropertiesType struct {
	properties.NamePropertyType
	DescriptionPropertiesType
}

type FirstLastSeenPropertiesType struct {
	properties.FirstSeenPropertyType
	properties.LastSeenPropertyType
}

type GoalsPropertiesType struct {
	properties.GoalsPropertyType
}

type KillChainPhasesPropertyType struct {
	properties.KillChainPhasesPropertyType
}

type PrimaryMotivationPropertyType struct {
	properties.PrimaryMotivationPropertyType
}

type ResourceLevelPropertyType struct {
	properties.ResourceLevelPropertyType
}

type SecondaryMotivationsPropertyType struct {
	properties.SecondaryMotivationsPropertyType
}

// ----------------------------------------------------------------------
// Public Methods - CommonObjectPropertiesType
// ----------------------------------------------------------------------

// InitNewObject is a helper function to init a new object with common elements
// It takes in one parameter
// param: s - a string value of the STIX object type
func (this *CommonObjectPropertiesType) InitNewObject(s string) {
	// TODO make sure that the value coming in a a valid STIX object type
	this.SetMessageType(s)
	this.CreateNewId(s)
	this.SetCreatedToCurrentTime()
	this.SetModifiedToCreated()
}

// SetModifiedToCreated sets the object modified time to be the same as the
// created time. This has to be done at this level, since at the individual
// properties type say "ModifiedPropertyType" this.Created does not exist.
// But it will exist at this level of inheritance
func (this *CommonObjectPropertiesType) SetModifiedToCreated() {
	this.Modified = this.Created
}

// VerifyTimestamp is a helper function to prevent needing to import the timestamp property object locally.
// It takes in one parameter and returns a string version of the timestamp
// param: t - a timestamp in either time.Time or string format
func (this *CommonObjectPropertiesType) VerifyTimestamp(t interface{}) string {
	return timestamp.Verify(t)
}

// func (this *CommonObjectPropertiesType) GetCurrentTime() string {
// 	return timestamp.GetCurrentTime()
// }
