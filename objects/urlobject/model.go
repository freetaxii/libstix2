// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package urlobject

import (
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
)

/*
UrlObject - This type implements the STIX 2.1 UrlObject SCO.
The URL object represents the properties of a uniform resource locator (URL).
*/
type UrlObject struct {
	objects.CommonObjectProperties
	properties.ValueProperty
}
