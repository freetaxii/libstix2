// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
//
// Types
//
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
	SourceName string `json:"source_name,omitempty"`
	DescriptionProperty
	URL        string            `json:"url,omitempty"`
	Hashes     map[string]string `json:"hashes,omitempty"`
	ExternalID string            `json:"external_id,omitempty"`
}

// ----------------------------------------------------------------------
//
// Public Methods - ExternalReferencesProperty
//
// ----------------------------------------------------------------------

/*
NewExternalReference - This method creates a new external reference and
returns a reference to a slice location. This will enable the code to update an
object located at that slice location.
*/
func (p *ExternalReferencesProperty) NewExternalReference() (*ExternalReference, error) {
	var s ExternalReference

	// if p.ExternalReferences == nil {
	// 	a := make([]ExternalReference, 0)
	// 	p.ExternalReferences = a
	// }

	positionThatAppendWillUse := len(p.ExternalReferences)
	p.ExternalReferences = append(p.ExternalReferences, s)
	return &p.ExternalReferences[positionThatAppendWillUse], nil
}

// ----------------------------------------------------------------------
//
// Public Methods - ExternalReference
//
// ----------------------------------------------------------------------

/*
SetSourceName - This method takes in a string value representing the name of
a source for an external reference and udpates the source name property.
*/
func (p *ExternalReference) SetSourceName(s string) error {
	p.SourceName = s
	return nil
}

/*
GetSourceName - This method will return the source name.
*/
func (p *ExternalReference) GetSourceName() string {
	return p.SourceName
}

/*
SetURL - This method takes in a string value representing a URL location of a
source for an external reference and updates the url property.
*/
func (p *ExternalReference) SetURL(s string) error {
	p.URL = s
	return nil
}

/*
GetURL - This method returns the url for this external reference.
*/
func (p *ExternalReference) GetURL() string {
	return p.URL
}

/*
AddHash - This method takes in two parameters and adds the hash to the map.
The first is a string value representing a hash type from the STIX hashes
vocabulary. The second is a string value representing the actual hash of the
content from the remote external reference.
*/
func (p *ExternalReference) AddHash(k, v string) error {
	if p.Hashes == nil {
		m := make(map[string]string, 0)
		p.Hashes = m
	}
	p.Hashes[k] = v
	return nil
}

/*
SetExternalID - This method takes in a string value representing an general
purpose id in a remote system for the source of this external reference and
updates the external id property.
*/
func (p *ExternalReference) SetExternalID(s string) error {
	p.ExternalID = s
	return nil
}

/*
GetExternalID - This method returns the external id for this reference.
*/
func (p *ExternalReference) GetExternalID() string {
	return p.ExternalID
}
