// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

/*
Package discovery implements the TAXII 2.1 API Root resource.

The following information comes directly from the TAXII 2.1 specification.

This Endpoint provides general information about a TAXII Server, including the
advertised API Roots. It's a common entry point for TAXII Clients into the data
and services provided by a TAXII Server. For example, clients auto-discovering
TAXII Servers via the DNS SRV record defined in section 1.6.1 will be able to
automatically retrieve a discovery response for that server by requesting the
/taxii2/ path on that domain.

Discovery API responses MAY advertise any TAXII API Root that they have
permission to advertise, included those hosted on other servers.
*/
package discovery
