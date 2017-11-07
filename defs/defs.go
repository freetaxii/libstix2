// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package defs

// These are the STIX timestamp formats.
// TIME_RFC_3339 is the general purpose timestamp.
// TIME_RFC_3339_MICRO is the timestamp for the created and modified properties.
const (
	STIX_MEDIA_TYPE  = "application/vnd.oasis.stix+json"
	STIX_VERSION     = "version=2.1"
	TAXII_MEDIA_TYPE = "application/vnd.oasis.taxii+json"
	TAXII_VERSION    = "version=2.0"

	TIME_RFC_3339       = "2006-01-02T15:04:05Z07:00"
	TIME_RFC_3339_MILLI = "2006-01-02T15:04:05.999Z07:00"
	TIME_RFC_3339_MICRO = "2006-01-02T15:04:05.999999Z07:00"
)
