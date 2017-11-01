// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package resources

import (
	"github.com/freetaxii/libstix2/resources/properties"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
CollectionType - This type implements the TAXII 2 Collction Resource and defines
all of the properties methods needed to create and work with the TAXII Collection
Resource. All of the methods not defined local to this type are inherited from
the individual properties.

The following information comes directly from the TAXII 2 specificaton documents.

This Endpoint provides general information about a Collection, which can be used
to help users and clients decide whether and how they want to interact with it.
For example, it will tell clients what it's called and what permissions they
have to it.

The collection resource contains general information about a Collection, such as
its id, a human-readable title and description, an optional list of supported
media_types (representing the media type of objects can be requested from or
added to it), and whether the TAXII Client, as authenticated, can get objects
from the Collection and/or add objects to it.
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
// Public Methods - CollectionType
// ----------------------------------------------------------------------

// SetCanRead - This method will set the can_read boolean to true.
func (ezt *CollectionType) SetCanRead() {
	ezt.CanRead = true
}

// GetCanRead - This method will return the value of Can Read.
func (ezt *CollectionType) GetCanRead() bool {
	return ezt.CanRead
}

// SetCanWrite - This method will set the can_write boolean to true.
func (ezt *CollectionType) SetCanWrite() {
	ezt.CanWrite = true
}

// GetCanWrite - This method will return the value of Can Write.
func (ezt *CollectionType) GetCanWrite() bool {
	return ezt.CanWrite
}

// AddMediaType - This method takes in a string value that represents a version
// of the TAXII api that is supported and adds it to the list in media types
// property.
func (ezt *CollectionType) AddMediaType(s string) {
	if ezt.MediaTypes == nil {
		a := make([]string, 0)
		ezt.MediaTypes = a
	}
	ezt.MediaTypes = append(ezt.MediaTypes, s)
}
