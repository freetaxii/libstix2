// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package stixid

import (
	"regexp"
)

// IsValidUUID - This function will take in a string and return true if the string
// represents an actual UUIDv4 value.
func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
	return r.MatchString(uuid)
}
