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
DiscoveryType - This type implements the TAXII 2 Discovery Resource and defines
all of the properties methods needed to create and work with the TAXII Discovery
Resource. All of the methods not defined local to this type are inherited from
the individual properties.

The following information comes directly from the TAXII 2 specificaton documents.

This Endpoint provides general information about a TAXII Server, including the
advertised API Roots. It's a common entry point for TAXII Clients into the data
and services provided by a TAXII Server. For example, clients auto-discovering
TAXII Servers via the DNS SRV record defined in section 1.4.1 will be able to
automatically retrieve a discovery response for that server by requesting the
/taxii/ path on that domain.

Discovery API responses MAY advertise any TAXII API Root that they have
permission to advertise, included those hosted on other servers.

The discovery resource contains information about a TAXII Server, such as a
human-readable title, description, and contact information, as well as a list of
API Roots that it is advertising. It also has an indication of which API Root it
considers the default, or the one to use in the absence of other
information/user choice.
*/
type DiscoveryType struct {
	properties.TitlePropertyType
	properties.DescriptionPropertyType
	Contact  string   `json:"contact,omitempty"`
	Default  string   `json:"default,omitempty"`
	APIRoots []string `json:"api_roots,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - DiscoveryType
// ----------------------------------------------------------------------

// SetContact - This methods takes in a string value representing contact
// information and updates the contact property.
func (ezt *DiscoveryType) SetContact(s string) {
	ezt.Contact = s
}

// GetContact - This method returns the contact information from the contact
// property.
func (ezt *DiscoveryType) GetContact() string {
	return ezt.Contact
}

// SetDefault - This methods takes in a string value representing a default
// api-root and updates the default property.
func (ezt *DiscoveryType) SetDefault(s string) {
	ezt.Default = s
}

// GetDefault - This methods returns the default api-root.
func (ezt *DiscoveryType) GetDefault() string {
	return ezt.Default
}

// AddAPIRoot - This method takes in a string value that represents an api-root
// and adds it to the list in the APIRoots property.
func (ezt *DiscoveryType) AddAPIRoot(s string) {
	if ezt.APIRoots == nil {
		a := make([]string, 0)
		ezt.APIRoots = a
	}
	ezt.APIRoots = append(ezt.APIRoots, s)
}
