// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package discovery

import (
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
Discovery - This type implements the TAXII 2 Discovery and defines all of the
properties and methods needed to create and work with this resource. All of the
methods not defined local to this type are inherited from the individual
properties.

The following information comes directly from the TAXII 2.1 specification.

The discovery resource contains information about a TAXII Server, such as a
human-readable title, description, and contact information, as well as a list of
API Roots that it is advertising. It also has an indication of which API Root it
considers the default, or the one to use in the absence of another
information/user choice.
*/
type Discovery struct {
	properties.TitleProperty
	properties.DescriptionProperty
	Contact  string   `json:"contact,omitempty"`
	Default  string   `json:"default,omitempty"`
	APIRoots []string `json:"api_roots,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new TAXII Discovery resource and return
it as a pointer.
*/
func New() *Discovery {
	var obj Discovery
	return &obj
}
