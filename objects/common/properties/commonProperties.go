// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

import (
	"github.com/freetaxii/libstix2/objects/common/timestamp"
)

// ----------------------------------------------------------------------
// Common Property Types - Used to populate the common object properties
// ----------------------------------------------------------------------

// CommonObjectPropertiesType - This type includes all of the common properties
// that are used by all STIX SDOs and SROs
type CommonObjectPropertiesType struct {
	MessageTypePropertyType
	IDPropertyType
	CreatedByRefPropertyType
	CreatedPropertyType
	ModifiedPropertyType
	RevokedPropertyType
	LabelsPropertyType
	ConfidencePropertyType
	LangPropertyType
	ExternalReferencesPropertyType
	ObjectMarkingRefsPropertyType
	GranularMarkingsPropertyType
}

// CommonMarkingDefinitionPropertiesType - This type includes all of the common
// properties that are used by the STIX Marking Definition object
type CommonMarkingDefinitionPropertiesType struct {
	MessageTypePropertyType
	IDPropertyType
	CreatedByRefPropertyType
	CreatedPropertyType
	ExternalReferencesPropertyType
	ObjectMarkingRefsPropertyType
	GranularMarkingsPropertyType
}

// CommonBundlePropertiesType - This type includes all of the common properties
// that are used by the STIX Bundle object
type CommonBundlePropertiesType struct {
	MessageTypePropertyType
	IDPropertyType
}

// ----------------------------------------------------------------------
// Public Methods - CommonObjectPropertiesType
// ----------------------------------------------------------------------

// InitNewObject is a helper function to init a new object with common elements
// It takes in one parameter
// param: s - a string value of the STIX object type
func (ezt *CommonObjectPropertiesType) InitNewObject(s string) {
	// TODO make sure that the value coming in a a valid STIX object type
	ezt.SetMessageType(s)
	ezt.CreateID(s)
	ezt.SetCreatedToCurrentTime()
	ezt.SetModifiedToCreated()
}

// SetModifiedToCreated sets the object modified time to be the same as the
// created time. This has to be done at this level, since at the individual
// properties type say "ModifiedPropertyType" this.Created does not exist.
// But it will exist at this level of inheritance
func (ezt *CommonObjectPropertiesType) SetModifiedToCreated() {
	ezt.Modified = ezt.Created
}

// VerifyTimestamp is a helper function to prevent needing to import the timestamp property object locally.
// It takes in one parameter and returns a string version of the timestamp
// param: t - a timestamp in either time.Time or string format
func (ezt *CommonObjectPropertiesType) VerifyTimestamp(t interface{}) string {
	return timestamp.Verify(t)
}

// func (this *CommonObjectPropertiesType) GetCurrentTime() string {
// 	return timestamp.GetCurrentTime()
// }
