// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

import "github.com/google/uuid"

// ----------------------------------------------------------------------
// Public Methods - DatastoreIDProperty - Setters
// ----------------------------------------------------------------------

// SetDatastoreID - This method takes in a int representing the database ID and
// updates the DatastoreID property.
func (o *CommonObjectProperties) SetDatastoreID(i int) error {
	o.DatastoreID = i
	return nil
}

// GetDatastoreID - This method returns the database ID value.
func (o *CommonObjectProperties) GetDatastoreID() int {
	return o.DatastoreID
}

// ----------------------------------------------------------------------
// Public Methods - TypeProperty - Setters
// ----------------------------------------------------------------------

// SetObjectType - This method takes in a string value representing a STIX
// object type and updates the type property.
func (o *CommonObjectProperties) SetObjectType(s string) error {
	o.ObjectType = s
	return nil
}

// GetObjectType - This method returns the object type.
func (o *CommonObjectProperties) GetObjectType() string {
	return o.ObjectType
}

// ----------------------------------------------------------------------
// Public Methods - SpecVersionProperty - Setters
// ----------------------------------------------------------------------

// SetSpecVersion20 - This method will set the specification version to 2.0.
func (o *CommonObjectProperties) SetSpecVersion20() error {
	o.SpecVersion = "2.0"
	return nil
}

// SetSpecVersion21 - This method will set the specification version to 2.1.
func (o *CommonObjectProperties) SetSpecVersion21() error {
	o.SpecVersion = "2.1"
	return nil
}

// SetSpecVersion - This method takes in a string representing a STIX
// specification version and updates the Version property.
func (o *CommonObjectProperties) SetSpecVersion(s string) error {
	o.SpecVersion = s
	return nil
}

// GetSpecVersion - This method returns the version value as a string.
func (o *CommonObjectProperties) GetSpecVersion() string {
	return o.SpecVersion
}

// ----------------------------------------------------------------------
// Public Methods - IDProperty - Setters
// ----------------------------------------------------------------------

// CreateSTIXUUID - This method takes in a string value representing a STIX
// object type and creates and returns a new ID based on the approved STIX UUIDv4
// format.
func (o *CommonObjectProperties) CreateSTIXUUID(s string) (string, error) {
	// TODO add check to validate that s is a valid type
	id := s + "--" + uuid.New().String()
	return id, nil
}

// CreateTAXIIUUID - This method does not take in any parameters. It is used to
// create a new ID based on the approved TAXII UUIDv4 format.
func (o *CommonObjectProperties) CreateTAXIIUUID() (string, error) {
	id := uuid.New().String()
	return id, nil
}

// SetNewTAXIIID - This method does not take in any parameters. It is used to
// create a new ID based on the approved TAXII UUIDv4 format and assigns it to the
// ID property.
func (o *CommonObjectProperties) SetNewTAXIIID() error {
	o.ID, _ = o.CreateTAXIIUUID()
	return nil
}

// SetNewSTIXID - This method takes in a string value representing a STIX object
// type and creates a new ID based on the approved STIX UUIDv4 format and update
// the id property for the object.
func (o *CommonObjectProperties) SetNewSTIXID(s string) error {
	// TODO Add check to validate input value
	o.ID, _ = o.CreateSTIXUUID(s)
	return nil
}

// SetID - This method takes in a string value representing an existing STIX id
// and updates the id property for the object.
func (o *CommonObjectProperties) SetID(s string) error {
	o.ID = s
	return nil
}

// GetID - This method will return the id for a given STIX object.
func (o *CommonObjectProperties) GetID() string {
	return o.ID
}

// ----------------------------------------------------------------------
// Public Methods - CreatedByRefProperty
// ----------------------------------------------------------------------

// SetCreatedByRef - This method takes in a string value representing a STIX
// identifier and updates the Created By Ref property.
func (o *CommonObjectProperties) SetCreatedByRef(s string) error {
	o.CreatedByRef = s
	return nil
}

// GetCreatedByRef - This method returns the STIX identifier for the identity
// that created this object.
func (o *CommonObjectProperties) GetCreatedByRef() string {
	return o.CreatedByRef
}

// ----------------------------------------------------------------------
// Public Methods - CreatedProperty - Setters
// ----------------------------------------------------------------------

// SetCreatedToCurrentTime - This methods sets the object created time to the
// current time
func (o *CommonObjectProperties) SetCreatedToCurrentTime() error {
	o.Created = GetCurrentTime("milli")
	return nil
}

// SetCreated - This method takes in a timestamp in either time.Time or string
// format and updates the created property with it. The value is stored as a
// string, so if the value is in time.Time format, it will be converted to the
// correct STIX timestamp format.
func (o *CommonObjectProperties) SetCreated(t interface{}) error {
	ts, _ := TimeToString(t, "milli")
	o.Created = ts
	return nil
}

// GetCreated - This method will return the created timestamp as a string.
func (o *CommonObjectProperties) GetCreated() string {
	return o.Created
}

// ----------------------------------------------------------------------
// Public Methods - ModifiedProperty - Setters
// ----------------------------------------------------------------------

// SetModifiedToCurrentTime - This methods sets the object created time to the
// current time
func (o *CommonObjectProperties) SetModifiedToCurrentTime() error {
	o.Modified = GetCurrentTime("milli")
	return nil
}

// SetModified - This method takes in a timestamp in either time.Time or string
// format and updates the modified property with it. The value is stored as a
// string, so if the value is in time.Time format, it will be converted to the
// correct STIX timestamp format.
func (o *CommonObjectProperties) SetModified(t interface{}) error {
	ts, _ := TimeToString(t, "milli")
	o.Modified = ts
	return nil
}

// GetModified - This method will return the modified timestamp as a string. If
// the value is the same as the created timestamp, then this object is the first
// version of the object.
func (o *CommonObjectProperties) GetModified() string {
	return o.Modified
}

// ----------------------------------------------------------------------
// Public Methods - RevokedProperty - Setters
// ----------------------------------------------------------------------

// SetRevoked - This method sets the revoked boolean to true
func (o *CommonObjectProperties) SetRevoked() error {
	o.Revoked = true
	return nil
}

// GetRevoked - This method returns the current value of the revoked property.
func (o *CommonObjectProperties) GetRevoked() bool {
	return o.Revoked
}

// ----------------------------------------------------------------------
// Public Methods - LabelsProperty - Setters
// ----------------------------------------------------------------------

// AddLabels - This method takes in a string value, a comma separated list of
// string values, or a slice of string values that all representing a label and
// adds it to the labels property.
func (o *CommonObjectProperties) AddLabels(values interface{}) error {
	return AddValuesToList(&o.Labels, values)
}

// ----------------------------------------------------------------------
// Public Methods - ConfidenceProperty - Setters
// ----------------------------------------------------------------------

// SetConfidence - This method takes in an integer representing a STIX
// confidence level 0-100 and updates the Confidence property.
func (o *CommonObjectProperties) SetConfidence(i int) error {
	o.Confidence = i
	return nil
}

// GetConfidence - This method returns the confidence value as an integer.
func (o *CommonObjectProperties) GetConfidence() int {
	return o.Confidence
}

// ----------------------------------------------------------------------
// Public Methods - LangProperty - Setters
// ----------------------------------------------------------------------

// SetLang - This method takes in a string value representing an ISO 639-2
// encoded language code as defined in RFC 5646 and updates the lang property.
func (o *CommonObjectProperties) SetLang(s string) error {
	o.Lang = s
	return nil
}

// GetLang - This method returns the current language code for a given object.
func (o *CommonObjectProperties) GetLang() string {
	return o.Lang
}

// ----------------------------------------------------------------------
// Public Methods - ExternalReferencesProperty - Setters
// ----------------------------------------------------------------------

// NewExternalReference - This method creates a new external reference and
// returns a reference to a slice location. This will enable the code to update an
// object located at that slice location.
func (o *CommonObjectProperties) NewExternalReference() (*ExternalReference, error) {
	var s ExternalReference

	// if o.ExternalReferences == nil {
	// 	a := make([]ExternalReference, 0)
	// 	o.ExternalReferences = a
	// }

	positionThatAppendWillUse := len(o.ExternalReferences)
	o.ExternalReferences = append(o.ExternalReferences, s)
	return &o.ExternalReferences[positionThatAppendWillUse], nil
}

// ----------------------------------------------------------------------
// Public Methods - ExternalReference - Setters
// ----------------------------------------------------------------------

// SetSourceName - This method takes in a string value representing the name of
// a source for an external reference and updates the source name property.
func (o *ExternalReference) SetSourceName(s string) error {
	o.SourceName = s
	return nil
}

// GetSourceName - This method will return the source name.
func (o *ExternalReference) GetSourceName() string {
	return o.SourceName
}

// SetDescription - This method takes in a string value representing a text
// description and updates the description property.
func (o *ExternalReference) SetDescription(s string) error {
	o.Description = s
	return nil
}

// GetDescription - This method returns the description for an object as a
// string.
func (o *ExternalReference) GetDescription() string {
	return o.Description
}

// SetURL - This method takes in a string value representing a URL location of a
// source for an external reference and updates the URL property.
func (o *ExternalReference) SetURL(s string) error {
	o.URL = s
	return nil
}

// GetURL - This method returns the URL for this external reference.
func (o *ExternalReference) GetURL() string {
	return o.URL
}

// AddHash - This method takes in two parameters and adds the hash to the map.
// The first is a string value representing a hash type from the STIX hashes
// vocabulary. The second is a string value representing the actual hash of the
// content from the remote external reference.
func (o *ExternalReference) AddHash(k, v string) error {
	if o.Hashes == nil {
		m := make(map[string]string, 0)
		o.Hashes = m
	}
	o.Hashes[k] = v
	return nil
}

// SetExternalID - This method takes in a string value representing an general
// purpose id in a remote system for the source of this external reference and
// updates the external id property.
func (o *ExternalReference) SetExternalID(s string) error {
	o.ExternalID = s
	return nil
}

// GetExternalID - This method returns the external id for this reference.
func (o *ExternalReference) GetExternalID() string {
	return o.ExternalID
}

// ----------------------------------------------------------------------
// Public Methods - MarkingProperty - Setters
// ----------------------------------------------------------------------

// AddObjectMarkingRef - This method takes in a string value that represents a
// STIX identifier for a marking definition object and adds it to the list of object
// marking refs.
func (o *CommonObjectProperties) AddObjectMarkingRef(s string) error {
	o.ObjectMarkingRefs = append(o.ObjectMarkingRefs, s)
	return nil
}

// ----------------------------------------------------------------------
// Public Methods - GranularMarking - Setters
// ----------------------------------------------------------------------

// SetMarkingRef - This method takes in a string value representing a STIX
// identifier of a marking definition object and sets the marking ref property to
// that value.
func (o *GranularMarking) SetMarkingRef(s string) error {
	o.MarkingRef = s
	return nil
}

// GetMarkingRef - This method returns the STIX identifier of the marking
// definition object that was recorded in this granular marking type.
func (o *GranularMarking) GetMarkingRef() string {
	return o.MarkingRef
}

// AddSelector - This method takes in a string value representing a STIX
// granular marking selector and adds it to the list of selectors.
func (o *GranularMarking) AddSelector(s string) error {
	o.Selectors = append(o.Selectors, s)
	return nil
}

// ----------------------------------------------------------------------
// Public Methods - RawProperty - Setters
// ----------------------------------------------------------------------

// SetRawData - This method takes in a slice of bytes representing a full JSON
// object and updates the raw property for the object.
func (o *CommonObjectProperties) SetRawData(data []byte) error {
	o.Raw = data
	return nil
}

// GetRawData - This method will return the raw bytes for a given STIX object.
func (o *CommonObjectProperties) GetRawData() []byte {
	return o.Raw
}
