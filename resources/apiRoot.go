// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package resources

import (
	"github.com/freetaxii/libstix2/resources/properties"
)

// ----------------------------------------------------------------------
//
// Define Message Type
//
// ----------------------------------------------------------------------

/*
APIRoot - This type implements the TAXII 2 API Root Resource and defines
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
type APIRoot struct {
	properties.TitleProperty
	properties.DescriptionProperty
	Versions         []string `json:"versions"`
	MaxContentLength int      `json:"max_content_length"`
}

// ----------------------------------------------------------------------
//
// Initialization Functions
//
// ----------------------------------------------------------------------

/*
NewAPIRoot - This function will create a new TAXII API Root object and return
it as a pointer.
*/
func NewAPIRoot() *APIRoot {
	var obj APIRoot
	return &obj
}

// ----------------------------------------------------------------------
//
// Public Methods - APIRoot
//
// ----------------------------------------------------------------------

/*
AddVersion - This method takes in a string value that represents a version of
the TAXII api that is supported and adds it to the versions property.
*/
func (r *APIRoot) AddVersion(s string) error {
	if r.Versions == nil {
		a := make([]string, 0)
		r.Versions = a
	}
	r.Versions = append(r.Versions, s)
	return nil
}

/*
SetMaxContentLength - This method takes in an integer value representing the
max content length that the server can support and updates the max content
length property.
*/
func (r *APIRoot) SetMaxContentLength(i int) error {
	r.MaxContentLength = i
	return nil
}

/*
GetMaxContentLength - This method returns the max content length as an
integer.
*/
func (r *APIRoot) GetMaxContentLength() int {
	return r.MaxContentLength
}
