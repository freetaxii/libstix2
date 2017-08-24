// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type MessageTypePropertyType struct {
	MessageType string `json:"type,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - MessageTypePropertyType
// ----------------------------------------------------------------------

// SetMessageType takes in one parameter
// param: s - a string value representing a STIX object type
func (this *MessageTypePropertyType) SetMessageType(s string) {
	this.MessageType = s
}

func (this *MessageTypePropertyType) GetMessageType() string {
	return this.MessageType
}
