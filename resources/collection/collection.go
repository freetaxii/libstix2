// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package collection

import (
	"github.com/freetaxii/libstix2/resources/common/properties"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
CollectionType defines all of the properties associated with the TAXII
Collection Resource. All of the methods not defined local to this type are
inherited from the individual properties.
*/
type CollectionType struct {
	properties.IDPropertyType
	properties.TitlePropertyType
	properties.DescriptionPropertyType
	CanRead    bool     `json:"can_read"`
	CanWrite   bool     `json:"can_write"`
	MediaTypes []string `json:"media_types,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// New - This function will create and return a new TAXII collection object.
func New() CollectionType {
	var obj CollectionType
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - CollectionType
// ----------------------------------------------------------------------

// SetCanRead - This method will set the can_read boolean to true.
func (p *CollectionType) SetCanRead() {
	p.CanRead = true
}

// GetCanRead - This method will return the value of Can Read.
func (p *CollectionType) GetCanRead() bool {
	return p.CanRead
}

// SetCanWrite - This method will set the can_write boolean to true.
func (p *CollectionType) SetCanWrite() {
	p.CanWrite = true
}

// GetCanWrite - This method will return the value of Can Write.
func (p *CollectionType) GetCanWrite() bool {
	return p.CanWrite
}

// AddMediaType - This method takes in a string value that represents a version
// of the TAXII api that is supported and adds it to the list in media types
// property.
func (p *CollectionType) AddMediaType(s string) {
	if p.MediaTypes == nil {
		a := make([]string, 0)
		p.MediaTypes = a
	}
	p.MediaTypes = append(p.MediaTypes, s)
}
