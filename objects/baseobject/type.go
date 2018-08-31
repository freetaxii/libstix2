// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package baseobject

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
TypeProperty - A property used by one or more STIX objects that
captures the STIX object type in string format.
*/
type TypeProperty struct {
	ObjectType string `json:"type,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - TypeProperty
// ----------------------------------------------------------------------

/*
SetObjectType - This method takes in a string value representing a STIX object
type and updates the type property.
*/
func (o *TypeProperty) SetObjectType(s string) error {
	o.ObjectType = s
	return nil
}

/*
GetObjectType - This method returns the object type.
*/
func (o *TypeProperty) GetObjectType() string {
	return o.ObjectType
}
