// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

/*
Package apiroot implements the TAXII 2.1 API Root resource.

The following information comes directly from the TAXII 2.1 specification.

API Roots are logical groupings of TAXII Collections, Channels, and related
functionality. A TAXII server instance can support one or more API Roots. API
Roots can be thought of as instances of the TAXII API available at different
URLs, where each API Root is the "root" URL of that particular instance of the
TAXII API. Organizing the Collections and Channels into API Roots allows for a
division of content and access control by trust group or any other logical
grouping. For example, a single TAXII Server could host multiple API Roots - one
API Root for Collections and Channels used by Sharing Group A and another API
Root for Collections and Channels used by Sharing Group B.

Each API Root contains a set of Endpoints that a TAXII Client contacts in order
to interact with the TAXII Server. This interaction can take several forms:

  - Server Discovery, as described above, can be used to learn about the API Roots

hosted by a TAXII Server.

  - Each API Root might support zero or more Collections. Interactions with

Collections include discovering the type of CTI contained in that Collection,
pushing new CTI to that Collection, and/or retrieving CTI from that Collection.
Each piece of CTI content in a Collection is referred to as an Object.

  - Each API Root might host zero or more Channels.

  - Each API Root also allows TAXII Clients to check on the Status of certain types

of requests to the TAXII Server. For example, if a TAXII Client submitted new
CTI, a Status request can allow the Client to check on whether the new CTI was
accepted.

This Endpoint provides general information about an API Root, which can be used
to help users and clients decide whether and how they want to interact with it.
Multiple API Roots MAY be hosted on a single TAXII Server. Often, an API Root
represents a single trust group.

  - Each API Root MUST have a unique URL.

  - Each API Root MAY have different authentication and authorization schemes.
*/
package apiroot
