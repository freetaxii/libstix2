// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
DatastoreIDProperty - A property used by all STIX objects that captures the
unique database ID for this object. This is not included in the JSON
serialization, but is used with some datastores.
*/
type DatastoreIDProperty struct {
	DatastoreID int `json:"-"`
}

// ----------------------------------------------------------------------
// Public Methods - DatastoreIDProperty
// ----------------------------------------------------------------------

/*
SetDatastoreID - This method takes in a int representing the database ID and
updates the DatastoreID property.
*/
func (o *DatastoreIDProperty) SetDatastoreID(i int) error {
	o.DatastoreID = i
	return nil
}

/*
GetDatastoreID - This method returns the database ID value.
*/
func (o *DatastoreIDProperty) GetDatastoreID() int {
	return o.DatastoreID
}
