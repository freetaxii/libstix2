// Copyright 2015-2019 Bret Jordan, All rights reserved.
// Copyright 2019 Oleksii Morozov, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package urlobject

import (
	"github.com/freetaxii/libstix2/objects/baseobject"
	"github.com/synsec/libstix2/objects/properties"
)

/*
UrlObject - This type implements the STIX 2.1 UrlObject SCO.
The URL object represents the properties of a uniform resource locator (URL).
*/
type UrlObject struct {
	baseobject.CommonObjectProperties
	properties.ValueProperty
}
