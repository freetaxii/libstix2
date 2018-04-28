// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package defs

// These are the STIX timestamp formats.
// TIME_RFC_3339 is the general purpose timestamp.
// TIME_RFC_3339_MICRO is the timestamp for the created and modified properties.
const (
	STIX_MEDIA_TYPE    = "application/stix+json"
	STIX_VERSION       = "2.1"
	TAXII_MEDIA_TYPE   = "application/taxii+json"
	TAXII_VERSION      = "2.1"
	CONTENT_TYPE_STIX  = STIX_MEDIA_TYPE + "; version=" + STIX_VERSION + "; charset=utf-8"
	CONTENT_TYPE_TAXII = TAXII_MEDIA_TYPE + "; version=" + TAXII_VERSION + "; charset=utf-8"
	CONTENT_TYPE_JSON  = "application/json; charset=utf-8"
	CONTENT_TYPE_HTML  = "text/html; charset=utf-8"

	TIME_RFC_3339       = "2006-01-02T15:04:05Z07:00"
	TIME_RFC_3339_MILLI = "2006-01-02T15:04:05.999Z07:00"
	TIME_RFC_3339_MICRO = "2006-01-02T15:04:05.999999Z07:00"
)
