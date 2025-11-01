// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package emailmessage

import (
	"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
EmailMessage - This type implements the STIX 2.1 EmailMessage SCO and defines all of the
properties and methods needed to create and work with this object. All of the
methods not defined local to this type are inherited from the individual
properties.

Reference: STIX 2.1 specification section 6.5

TODO: Complete implementation of all properties per specification
*/
type EmailMessage struct {
	objects.CommonObjectProperties
	// TODO: Add specific properties for EmailMessage based on STIX 2.1 spec section 6.5
	// TODO: Add is_multipart, date, content_type, from_ref, sender_ref, to_refs, cc_refs, bcc_refs, message_id, subject, received_lines, additional_header_fields, body, body_multipart, raw_email_ref
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *EmailMessage) GetPropertyList() []string {
	// TODO: Update with actual property names
	return []string{}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX EmailMessage SCO and return it as a
pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *EmailMessage {
	var obj EmailMessage
	obj.InitSCO("emailmessage")
	return &obj
}
