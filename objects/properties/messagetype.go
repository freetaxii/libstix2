// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

// MessageTypePropertyType - A property used by one or more STIX objects that
// captures the STIX object type in string format.
type MessageTypePropertyType struct {
	MessageType string `json:"type,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - MessageTypePropertyType
// ----------------------------------------------------------------------

// SetMessageType - This method takes in a string value representing a STIX
// object type and updates the type property.
func (ezt *MessageTypePropertyType) SetMessageType(s string) {
	ezt.MessageType = s
}

// GetMessageType - This method returns the object type.
func (ezt *MessageTypePropertyType) GetMessageType() string {
	return ezt.MessageType
}
