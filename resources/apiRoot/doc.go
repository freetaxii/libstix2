// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

/*
Package apiRoot implements the TAXII 2 API Root Resource.
This package defines the properties and methods needed to create and work with
the TAXII API Root Resource.

The following information comes directly from the TAXII 2 specificaton documents.

This Endpoint provides general information about an API Root, which can be used
to help users and clients decide whether and how they want to interact with it.
Multiple API Roots MAY be hosted on a single TAXII Server. Often, an API Root
represents a single trust group.

The api-root resource contains general information about the API Root, such as a
human-readable title and description, the TAXII versions it supports, and the
maximum size of the content body it will accept in a PUT or POST
(max_content_length).
*/
package apiRoot
