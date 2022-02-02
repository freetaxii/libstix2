// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package apiroot

import (
	"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
APIRoot - This type implements the TAXII 2 API Root and defines all of the
properties and methods needed to create and work with this resource. All of the
methods not defined local to this type are inherited from the individual
properties.

The following information comes directly from the TAXII 2.1 specification.

The API Root resource contains general information about the API Root, such as a
human-readable title and description, the TAXII versions it supports, and the
maximum size (max_content_length) of the content body it will accept in a PUT or
POST request.
*/
type APIRoot struct {
	objects.TitleProperty
	objects.DescriptionProperty
	Versions         []string `json:"versions"`
	MaxContentLength int      `json:"max_content_length"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new TAXII API Root resource and return
it as a pointer.
*/
func New() *APIRoot {
	var obj APIRoot
	return &obj
}
