// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package baseobject

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
ObjectIDProperty - A property used by all STIX objects that captures the unique
object ID used for storage. This is not included in the JSON serialization, but
is used for writing to the database.
*/
type ObjectIDProperty struct {
	ObjectID int `json:"-"`
}

// ----------------------------------------------------------------------
// Public Methods - ObjectIDProperty
// ----------------------------------------------------------------------

/*
SetObjectID - This method takes in a int64 representing an object ID and
updates the Version property.
*/
func (o *ObjectIDProperty) SetObjectID(i int) error {
	o.ObjectID = i
	return nil
}

/*
GetObjectID - This method returns the object ID value as a int64.
*/
func (o *ObjectIDProperty) GetObjectID() int {
	return o.ObjectID
}
