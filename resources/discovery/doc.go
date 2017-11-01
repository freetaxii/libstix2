// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

/*
Package discovery implements the TAXII 2 Discovery Resource.
This package defines the properties and methods needed to create and work with
the TAXII Discovery Resource.

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
package discovery
