// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package defs

// These are the STIX timestamp formats.
// TIME_RFC_3339 is the general purpose timestamp.
// TIME_RFC_3339_MICRO is the timestamp for the created and modified properties.
const (
	STIX_VERSION       = "2.1"
	TAXII_VERSION      = "2.1"
	MEDIA_TYPE_STIX    = "application/stix+json"
	MEDIA_TYPE_TAXII   = "application/taxii+json"
	MEDIA_TYPE_STIX20  = "application/vnd.oasis.stix+json;version=2.0"
	MEDIA_TYPE_TAXII20 = "application/vnd.oasis.taxii+json;version=2.0"
	MEDIA_TYPE_STIX21  = "application/stix+json;version=2.1"
	MEDIA_TYPE_STIX22  = "application/stix+json;version=2.2"
	MEDIA_TYPE_STIX23  = "application/stix+json;version=2.3"
	MEDIA_TYPE_STIX24  = "application/stix+json;version=2.4"
	MEDIA_TYPE_TAXII21 = "application/taxii+json;version=2.1"
	MEDIA_TYPE_JSON    = "application/json"
	MEDIA_TYPE_HTML    = "text/html; charset=utf-8"

	TIME_RFC_3339       = "2006-01-02T15:04:05Z07:00"
	TIME_RFC_3339_MILLI = "2006-01-02T15:04:05.999Z07:00"
	TIME_RFC_3339_MICRO = "2006-01-02T15:04:05.999999Z07:00"
)
