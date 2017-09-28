// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

// ExternalReferencesPropertyType - A property used by one or more STIX objects
// that captures a list of external references as defined by STIX.
type ExternalReferencesPropertyType struct {
	ExternalReferences []ExternalReferenceType `json:"external_references,omitempty"`
}

// ExternalReferenceType -
// This type defines all of the properties associated with the STIX External Reference type.
// All of the methods not defined local to this type are inherited from the individual properties.
type ExternalReferenceType struct {
	SourceName string `json:"source_name,omitempty"`
	DescriptionPropertyType
	URL        string            `json:"url,omitempty"`
	Hashes     map[string]string `json:"hashes,omitempty"`
	ExternalID string            `json:"external_id,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - ExternalReferenceType
// ----------------------------------------------------------------------

// SetSourceName - This method takes in a string value representing the name of
// a source for an external reference and udpates the source name property.
func (ezt *ExternalReferenceType) SetSourceName(s string) {
	ezt.SourceName = s
}

// GetSourceName - This method will return the source name.
func (ezt *ExternalReferenceType) GetSourceName() string {
	return ezt.SourceName
}

// SetURL - This method takes in a string value representing a URL location of a
// source for an external reference and updates the url property.
func (ezt *ExternalReferenceType) SetURL(s string) {
	ezt.URL = s
}

// GetURL - This method returns the url for this external reference.
func (ezt *ExternalReferenceType) GetURL() string {
	return ezt.URL
}

// AddHash - This method takes in two parameters and adds the hash to the map.
// The first is a string value representing a hash type from the STIX hashes
// vocabulary. The second is a string value representing the actual hash of the
// content from the remote external reference.
func (ezt *ExternalReferenceType) AddHash(k, v string) {
	if ezt.Hashes == nil {
		m := make(map[string]string, 0)
		ezt.Hashes = m
	}
	ezt.Hashes[k] = v
}

// SetExternalID - This method takes in a string value representing an general
// purpose id in a remote system for the source of this external reference and
// updates the external id property.
func (ezt *ExternalReferenceType) SetExternalID(s string) {
	ezt.ExternalID = s
}

// GetExternalID - This method returns the external id for this reference.
func (ezt *ExternalReferenceType) GetExternalID() string {
	return ezt.ExternalID
}
