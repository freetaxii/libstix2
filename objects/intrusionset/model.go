// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package intrusionset

import (
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Type
// ----------------------------------------------------------------------

/* IntrusionSet - This type implements the STIX 2 Intrusion Set SDO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties. */
type IntrusionSet struct {
	objects.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	properties.AliasesProperty
	properties.SeenProperties
	properties.GoalsProperty
	properties.ResourceLevelProperty
	properties.MotivationProperties
}

/* GetProperties - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync. */
func (o *IntrusionSet) GetPropertyList() []string {
	return []string{"name", "description", "aliases", "first_seen", "last_seen", "goals", "resource_level", "primary_motivation", "secondary_motivations"}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/* New - This function will create a new STIX Intrusion Set object and return
it as a pointer. It will also initialize the object by setting all of the basic
properties. */
func New() *IntrusionSet {
	var obj IntrusionSet
	obj.InitSDO("intrusion-set")
	return &obj
}
