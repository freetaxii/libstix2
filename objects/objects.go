// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

import (
	"fmt"

	"github.com/freetaxii/libstix2/defs"
)

/*
STIXObject - This interface defines what methods an object must have to be
considered a STIX Object.
*/
type STIXObject interface {
	GetObjectType() string
	GetID() string
	Valid() (bool, error)
	SetRawData([]byte) error
	GetCommonProperties() *CommonObjectProperties
}

/*
ValidType - This function will take in a STIX Object Type and return true if
the string represents an actual STIX object type. This is used for determining
if input from an outside source is actually a defined STIX object or not.
*/
func ValidType(t string) bool {

	var m = map[string]int{
		"attack-pattern":     1,
		"campaign":           1,
		"course-of-action":   1,
		"identity":           1,
		"indicator":          1,
		"intrusion-set":      1,
		"location":           1,
		"malware":            1,
		"marking-definition": 1,
		"note":               1,
		"observed-data":      1,
		"opinion":            1,
		"relationship":       1,
		"report":             1,
		"sighting":           1,
		"threat-actor":       1,
		"tool":               1,
		"vulnerability":      1,
	}

	if _, ok := m[t]; ok {
		return true
	}
	return false
}

/* InitObject - This method will initialize the object by setting all of the
basic properties and is called by the New() function on each object. */
func (o *CommonObjectProperties) InitObject(stixType string) error {
	if defs.STRICT_TYPES {
		if valid := ValidType(stixType); valid != true {
			return fmt.Errorf("invalid object type for InitObject with strict checks enabled")
		}
	}
	// TODO make sure that the value coming in is a valid STIX object type
	o.SetSpecVersion(defs.STIX_VERSION)
	o.SetObjectType(stixType)
	o.SetNewID(stixType)
	o.SetCreatedToCurrentTime()
	o.SetModifiedToCreated()
	return nil
}

/*
GetCommonProperties - This method will return a pointer to the common properties
of this object.
*/
func (o *CommonObjectProperties) GetCommonProperties() *CommonObjectProperties {
	return o
}

/*
Valid - This method will ensure that all of the required properties are
populated and try to ensure all of values are valid.
*/
func (o *CommonObjectProperties) Valid() (bool, error) {

	if valid, err := o.TypeProperty.Valid(); valid != true {
		return valid, err
	}

	if valid, err := o.SpecVersionProperty.Valid(); valid != true {
		return valid, err
	}

	if valid, err := o.IDProperty.Valid(); valid != true {
		return valid, err
	}

	if valid, err := o.CreatedModifiedProperty.Valid(); valid != true {
		return valid, err
	}

	return true, nil
}
