// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package discovery

import (
	"github.com/freetaxii/libstix2/resources/common/properties"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
DiscoveryType defines all of the properties associated with the TAXII Discovery
Resource. All of the methods not defined local to this type are inherited
from the individual properties.
*/
type DiscoveryType struct {
	properties.TitlePropertyType
	properties.DescriptionPropertyType
	Contact  string   `json:"contact,omitempty"`
	Default  string   `json:"default,omitempty"`
	APIRoots []string `json:"api_roots,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// New - This function will create and return a new TAXII discovery object.
func New() DiscoveryType {
	var obj DiscoveryType
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - DiscoveryType
// ----------------------------------------------------------------------

// SetContact - This methods takes in a string value representing contact
// information and updates the contact property.
func (p *DiscoveryType) SetContact(s string) {
	p.Contact = s
}

// GetContact - This method returns the contact information from the contact
// property.
func (p *DiscoveryType) GetContact() string {
	return p.Contact
}

// SetDefault - This methods takes in a string value representing a default
// api-root and updates the default property.
func (p *DiscoveryType) SetDefault(s string) {
	p.Default = s
}

// GetDefault - This methods returns the default api-root.
func (p *DiscoveryType) GetDefault() string {
	return p.Default
}

// AddAPIRoot - This method takes in a string value that represents an api-root
// and adds it to the list in the APIRoots property.
func (p *DiscoveryType) AddAPIRoot(s string) {
	if p.APIRoots == nil {
		a := make([]string, 0)
		p.APIRoots = a
	}
	p.APIRoots = append(p.APIRoots, s)
}
