// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package common

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type ExteralReferenceType struct {
	Source_name string `json:"source_name,omitempty"`
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
	External_id string `json:"external_id,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - ExteralReferenceType
// ----------------------------------------------------------------------

func (this *ExteralReferenceType) SetSourceName(s string) {
	this.Source_name = s
}

func (this *ExteralReferenceType) SetDescription(s string) {
	this.Description = s
}

func (this *ExteralReferenceType) SetURL(s string) {
	this.Url = s
}

func (this *ExteralReferenceType) SetExternalID(s string) {
	this.External_id = s
}
