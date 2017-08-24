// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type ExternalReferencesPropertyType struct {
	External_references []ExteralReferenceType `json:"external_references,omitempty"`
}

type ExteralReferenceType struct {
	Source_name string            `json:"source_name,omitempty"`
	Description string            `json:"description,omitempty"`
	Url         string            `json:"url,omitempty"`
	Hashes      map[string]string `json:"hashes,omitempty"`
	External_id string            `json:"external_id,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - ExteralReferenceType
// ----------------------------------------------------------------------

// SetSourceName takes in one parameter
// param: s - a string value representing the name of a source for an external refernce
func (this *ExteralReferenceType) SetSourceName(s string) {
	this.Source_name = s
}

func (this *ExteralReferenceType) GetSourceName() string {
	return this.Source_name
}

// SetDescription takes in one parameter
// param: s - a string value representing a description of a source for an external refernce
func (this *ExteralReferenceType) SetDescription(s string) {
	this.Description = s
}

func (this *ExteralReferenceType) GetDescription() string {
	return this.Description
}

// SetUrl takes in one parameter
// param: s - a string value representing a URL location of a source for an external refernce
func (this *ExteralReferenceType) SetURL(s string) {
	this.Url = s
}

func (this *ExteralReferenceType) GetUrl() string {
	return this.Url
}

// AddHash takes in two parameters
// param: k - a string value representing a hash type from the STIX hashes vocabulary
// param: v - a string value representing the actual hash of the content from the remote external reference
func (this *ExteralReferenceType) AddHash(k, v string) {
	if this.Hashes == nil {
		m := make(map[string]string, 0)
		this.Hashes = m
	}
	this.Hashes[k] = v
}

// SetExternalId takes in one parameter
// param: s - a string value representing an id in a remote system for the source of this external refernce
func (this *ExteralReferenceType) SetExternalId(s string) {
	this.External_id = s
}

func (this *ExteralReferenceType) GetExternalId() string {
	return this.External_id
}
