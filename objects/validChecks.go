// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package objects

import (
	"github.com/freetaxii/libstix2/common/stixid"
	"strings"
)

/*
IsValidSTIXID - This function will take in a STIX ID and return true if the
string represents an actual STIX ID in the correct format.
*/
func IsValidSTIXID(id string) (bool, error) {
	valid := false

	idparts := strings.Split(id, "--")

	valid = IsValidSTIXObject(idparts[0])

	// Short circuit if the STIX type part is wrong
	if valid == false {
		return false
	}

	valid = stixid.IsValidUUID(idparts[1])

	return valid, nil
}

/*
IsValidSTIXObject - This function will take in a string and return true if the
string represents an actual STIX object type. This is used for determining if
input from an outside source is actually a defined STIX object or not.
*/
func IsValidSTIXObject(obj string) (bool, error) {
	valid := false

	switch obj {
	case "attack-pattern":
		valid = true
	case "campaign":
		valid = true
	case "course-of-action":
		valid = true
	case "identity":
		valid = true
	case "indicator":
		valid = true
	case "intrusion-set":
		valid = true
	case "location":
		valid = true
	case "malware":
		valid = true
	case "marking-definition":
		valid = true
	case "note":
		valid = true
	case "observed-data":
		valid = true
	case "opinion":
		valid = true
	case "relationship":
		valid = true
	case "report":
		valid = true
	case "sighting":
		valid = true
	case "threat-actor":
		valid = true
	case "tool":
		valid = true
	case "vulnerability":
		valid = true
	}
	return valid, nil
}
