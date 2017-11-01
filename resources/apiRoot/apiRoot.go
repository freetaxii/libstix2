// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package apiRoot

import (
	"github.com/freetaxii/libstix2/resources/common/properties"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
APIRootType defines all of the properties associated with the TAXII API Root
Resource. All of the methods not defined local to this type are inherited
from the individual properties.
*/
type APIRootType struct {
	properties.TitlePropertyType
	properties.DescriptionPropertyType
	Versions         []string `json:"versions"`
	MaxContentLength int      `json:"max_content_length"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// New - This function will create and return a new TAXII api root object.
func New() APIRootType {
	var obj APIRootType
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - APIRootType
// ----------------------------------------------------------------------

// AddVersion - This method takes in a string value that represents a version of
// the TAXII api that is supported and adds it to the versions property.
func (p *APIRootType) AddVersion(s string) {
	if p.Versions == nil {
		a := make([]string, 0)
		p.Versions = a
	}
	p.Versions = append(p.Versions, s)
}

// SetMaxContentLength - This method takes in an integer value representing the
// max content length the server can support and updates the max content length
// property.
func (p *APIRootType) SetMaxContentLength(i int) {
	p.MaxContentLength = i
}

// GetMaxContentLength - This method returns the max content length as an
// integer.
func (p *APIRootType) GetMaxContentLength() int {
	return p.MaxContentLength
}
