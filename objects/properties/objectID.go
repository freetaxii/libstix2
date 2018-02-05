// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
ObjectIDPropertyType - A property used by one or more STIX objects that captures
the unique object ID. This is not included in the JSON serialization, but is
used for writing to the database.
*/
type ObjectIDPropertyType struct {
	ObjectID int `json:"-"`
}

// ----------------------------------------------------------------------
// Public Methods - ObjectIDPropertyType
// ----------------------------------------------------------------------

/*
SetObjectID - This method takes in a int64 representing an object ID and
updates the Version property.
*/
func (ezt *ObjectIDPropertyType) SetObjectID(i int) error {
	ezt.ObjectID = i
	return nil
}

/*
GetObjectID - This method returns the object ID value as a int64.
*/
func (ezt *ObjectIDPropertyType) GetObjectID() int {
	return ezt.ObjectID
}
