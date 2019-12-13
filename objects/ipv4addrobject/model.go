// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package ipv4addrobject

import (
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
)

/*
IPv4AddrObject - This type implements the STIX 2.1 IPv4AddrObject SCO.
The Domain Name object represents the properties of a network domain name.
*/
type IPv4AddrObject struct {
	objects.CommonObjectProperties
	properties.ResolvesToRefsProperty
	properties.ValueProperty
	properties.BelongsToRefsProperty
}
