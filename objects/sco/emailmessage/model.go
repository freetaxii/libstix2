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
IPv4Addr - This type implements the STIX 2 IPv4 Address SCO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties.
*/
type EmailMessage struct {
	objects.CommonObjectProperties
	IsMultipart            bool           `json:"is_multipart,omitempty" bson:"is_multipart,omitempty"`
	Date                   string         `json:"date,omitempty" bson:"date,omitempty"`
	ContentType            string         `json:"content_type,omitempty" bson:"content_type,omitempty"`
	FromRef                string         `json:"from_ref,omitempty" bson:"from_ref,omitempty"`
	SenderRef              string         `json:"sender_ref,omitempty" bson:"sender_ref,omitempty"`
	ToRefs                 []string       `json:"to_refs,omitempty" bson:"to_refs,omitempty"`
	CcRefs                 []string       `json:"cc_refs,omitempty" bson:"cc_refs,omitempty"`
	BccRefs                []string       `json:"bcc_refs,omitempty" bson:"bcc_refs,omitempty"`
	MessageId              string         `json:"message_id,omitempty" bson:"message_id,omitempty"`
	Subject                string         `json:"subject,omitempty" bson:"subject,omitempty"`
	ReceivedLines          []string       `json:"received_lines,omitempty" bson:"received_lines,omitempty"`
	AdditionalHeaderFields map[string]any `json:"additional_header_fields,omitempty" bson:"additional_header_fields,omitempty"`
	Body                   string         `json:"body,omitempty" bson:"body,omitempty"`
	BodyMultipart          []string       `json:"body_multipart,omitempty" bson:"body_multipart,omitempty"`
	RawEmailRef            string         `json:"raw_email_ref,omitempty" bson:"raw_email_ref,omitempty"`
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *EmailMessage) GetPropertyList() []string {
	return []string{
		"is_multipart",
		"date",
		"content_type",
		"from_ref",
		"sender_ref",
		"to_refs",
		"cc_refs",
		"bcc_refs",
		"message_id",
		"subject",
		"received_lines",
		"additional_header_fields",
		"body",
		"body_multipart",
		"raw_email_ref",
	}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX IPv4 Address SCO and return it as
a pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *EmailMessage {
	var obj EmailMessage
	obj.InitSCO("email-message")
	return &obj
}
