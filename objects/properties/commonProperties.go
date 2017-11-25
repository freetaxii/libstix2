// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Common Property Types - Used to populate the common object properties
// ----------------------------------------------------------------------

// CommonObjectPropertiesType - This type includes all of the common properties
// that are used by all STIX SDOs and SROs
type CommonObjectPropertiesType struct {
	ObjectIDPropertyType
	STIXVersionPropertyType
	ObjectTypePropertyType
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
	STIXVersionPropertyType
	ObjectTypePropertyType
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
	ObjectTypePropertyType
	IDPropertyType
}

// ----------------------------------------------------------------------
// Public Methods - CommonObjectPropertiesType
// ----------------------------------------------------------------------

// InitNewObject is a helper function to initialize a new object with common
// elements.
//
// params: objectType - a string value of the STIX object type
// params: version - the STIX spec version of the object, ex. "2.0". This is
// 		stored and used in TAXII.
func (ezt *CommonObjectPropertiesType) InitNewObject(objectType, version string) {
	// TODO make sure that the value coming in a a valid STIX object type
	ezt.SetSpecVersion(version)
	ezt.SetObjectType(objectType)
	ezt.NewID(objectType)
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
