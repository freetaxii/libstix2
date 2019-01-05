// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package baseobject

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
ExternalReferencesProperty - A property used by one or more STIX objects
that captures a list of external references as defined by STIX.
*/
type ExternalReferencesProperty struct {
	ExternalReferences []ExternalReference `json:"external_references,omitempty"`
}

/*
ExternalReference - This type defines all of the properties associated with
the STIX External Reference type. All of the methods not defined local to this
type are inherited from the individual properties.
*/
type ExternalReference struct {
	SourceName  string            `json:"source_name,omitempty"`
	Description string            `json:"description,omitempty"`
	URL         string            `json:"url,omitempty"`
	Hashes      map[string]string `json:"hashes,omitempty"`
	ExternalID  string            `json:"external_id,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - ExternalReferencesProperty
// ----------------------------------------------------------------------

/*
NewExternalReference - This method creates a new external reference and
returns a reference to a slice location. This will enable the code to update an
object located at that slice location.
*/
func (o *ExternalReferencesProperty) NewExternalReference() (*ExternalReference, error) {
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
// Public Methods - ExternalReference
// ----------------------------------------------------------------------

/*
SetSourceName - This method takes in a string value representing the name of
a source for an external reference and udpates the source name property.
*/
func (o *ExternalReference) SetSourceName(s string) error {
	o.SourceName = s
	return nil
}

/*
GetSourceName - This method will return the source name.
*/
func (o *ExternalReference) GetSourceName() string {
	return o.SourceName
}

/*
SetDescription - This method takes in a string value representing a text
description and updates the description property.
*/
func (o *ExternalReference) SetDescription(s string) error {
	o.Description = s
	return nil
}

/*
GetDescription - This method returns the description for an object as a string.
*/
func (o *ExternalReference) GetDescription() string {
	return o.Description
}

/*
SetURL - This method takes in a string value representing a URL location of a
source for an external reference and updates the url property.
*/
func (o *ExternalReference) SetURL(s string) error {
	o.URL = s
	return nil
}

/*
GetURL - This method returns the url for this external reference.
*/
func (o *ExternalReference) GetURL() string {
	return o.URL
}

/*
AddHash - This method takes in two parameters and adds the hash to the map.
The first is a string value representing a hash type from the STIX hashes
vocabulary. The second is a string value representing the actual hash of the
content from the remote external reference.
*/
func (o *ExternalReference) AddHash(k, v string) error {
	if o.Hashes == nil {
		m := make(map[string]string, 0)
		o.Hashes = m
	}
	o.Hashes[k] = v
	return nil
}

/*
SetExternalID - This method takes in a string value representing an general
purpose id in a remote system for the source of this external reference and
updates the external id property.
*/
func (o *ExternalReference) SetExternalID(s string) error {
	o.ExternalID = s
	return nil
}

/*
GetExternalID - This method returns the external id for this reference.
*/
func (o *ExternalReference) GetExternalID() string {
	return o.ExternalID
}
