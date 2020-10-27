// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package stixid

import (
	"regexp"
	"strings"

	"github.com/wxj95/libstix2/objects"
)

/*
ValidSTIXID - This function will take in a STIX ID and return true if the
string represents an actual STIX ID in the correct format.
*/
func ValidSTIXID(id string) bool {
	idparts := strings.Split(id, "--")

	if idparts == nil {
		return false
	}

	// First check to see if the object type is valid, if not return false.
	if valid := objects.ValidType(idparts[0]); valid == false {
		// Short circuit if the STIX type part is wrong
		return false
	}

	// If the type is valid, then check to see if the ID is a UUID, if not return
	// false.
	valid := ValidUUID(idparts[1])

	return valid
}

/*
ValidUUID - This function will take in a string and return true if the string
represents an actual UUIDv4 value.
*/
func ValidUUID(uuid string) bool {
	r := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
	return r.MatchString(uuid)
}
