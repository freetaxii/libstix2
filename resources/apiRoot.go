// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package resources

import (
	"github.com/freetaxii/libstix2/resources/properties"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
APIRootType - This type implements the TAXII 2 API Root Resource and defines
all of the properties and methods needed to create and work with the TAXII API Root
Resource. All of the methods not defined local to this type are inherited from
the individual properties.

The following information comes directly from the TAXII 2 specification documents.

This Endpoint provides general information about an API Root, which can be used
to help users and clients decide whether and how they want to interact with it.
Multiple API Roots MAY be hosted on a single TAXII Server. Often, an API Root
represents a single trust group.

The api-root resource contains general information about the API Root, such as a
human-readable title and description, the TAXII versions it supports, and the
maximum size of the content body it will accept in a PUT or POST
(max_content_length).
*/
type APIRootType struct {
	properties.TitlePropertyType
	properties.DescriptionPropertyType
	Versions         []string `json:"versions"`
	MaxContentLength int      `json:"max_content_length"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
InitAPIRoot - This function will create a new TAXII API Root object and return
it as a pointer.
*/
func InitAPIRoot() *APIRootType {
	var obj APIRootType
	return &obj
}

// ----------------------------------------------------------------------
// Public Methods - APIRootType
// ----------------------------------------------------------------------

/*
AddVersion - This method takes in a string value that represents a version of
the TAXII api that is supported and adds it to the versions property.
*/
func (ezt *APIRootType) AddVersion(s string) error {
	if ezt.Versions == nil {
		a := make([]string, 0)
		ezt.Versions = a
	}
	ezt.Versions = append(ezt.Versions, s)
	return nil
}

/*
SetMaxContentLength - This method takes in an integer value representing the
max content length that the server can support and updates the max content
length property.
*/
func (ezt *APIRootType) SetMaxContentLength(i int) error {
	ezt.MaxContentLength = i
	return nil
}

/*
GetMaxContentLength - This method returns the max content length as an
integer.
*/
func (ezt *APIRootType) GetMaxContentLength() (int, error) {
	return ezt.MaxContentLength, nil
}
