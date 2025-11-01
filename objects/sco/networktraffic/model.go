// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package networktraffic

import (
"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
NetworkTraffic - This type implements the STIX 2.1 NetworkTraffic SCO and defines all of the
properties and methods needed to create and work with this object. All of the
methods not defined local to this type are inherited from the individual
properties.

Reference: STIX 2.1 specification section 6.10

TODO: Complete implementation of all properties per specification
*/
type NetworkTraffic struct {
objects.CommonObjectProperties
// TODO: Add specific properties for NetworkTraffic based on STIX 2.1 spec section 6.10
// TODO: Add start, end, is_active, src_ref, dst_ref, src_port, dst_port, protocols, src_byte_count, dst_byte_count, src_packets, dst_packets, ipfix, src_payload_ref, dst_payload_ref, encapsulates_refs, encapsulated_by_ref
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *NetworkTraffic) GetPropertyList() []string {
// TODO: Update with actual property names
return []string{}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX NetworkTraffic SCO and return it as a
pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *NetworkTraffic {
var obj NetworkTraffic
obj.InitSCO("networktraffic")
return &obj
}
