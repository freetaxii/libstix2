// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package baseobject

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
DatabaseIDProperty - A property used by all STIX objects that captures the
unique database ID for this object. This is not included in the JSON
serialization, but is used with some datastores.
*/
type DatabaseIDProperty struct {
	DatabaseID int `json:"-"`
}

// ----------------------------------------------------------------------
// Public Methods - DatabaseIDProperty
// ----------------------------------------------------------------------

/*
SetDatabaseID - This method takes in a int representing the database ID and
updates the DatabaseID property.
*/
func (o *DatabaseIDProperty) SetDatabaseID(i int) error {
	o.DatabaseID = i
	return nil
}

/*
GetDatabaseID - This method returns the database ID value.
*/
func (o *DatabaseIDProperty) GetDatabaseID() int {
	return o.DatabaseID
}
