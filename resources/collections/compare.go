// Copyright 2015-2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package collections

import "fmt"

/*
Compare - This method will compare two collections to make sure they
are the same. The collection receiver is the master and represent the correct
data, the indicator passed in as i represents the one we need to test.
*/
func (r *Collection) Compare(c *Collection) (bool, int, []string) {
	problemsFound := 0
	details := make([]string, 0)

	// Check ID Value
	if c.ID != r.ID {
		problemsFound++
		str := fmt.Sprintf("-- IDs Do Not Match: %s | %s", r.ID, c.ID)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ IDs Match: %s | %s", r.ID, c.ID)
		details = append(details, str)
	}

	// Check Title Value
	if c.Title != r.Title {
		problemsFound++
		str := fmt.Sprintf("-- Titles Do Not Match: %s | %s", r.Title, c.Title)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Titles Match: %s | %s", r.Title, c.Title)
		details = append(details, str)
	}

	// Check Description Value
	if c.Description != r.Description {
		problemsFound++
		str := fmt.Sprintf("-- Descriptions Do Not Match: %s | %s", r.Description, c.Description)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Descriptions Match: %s | %s", r.Description, c.Description)
		details = append(details, str)
	}

	// Check Can Read Value
	if c.CanRead != r.CanRead {
		problemsFound++
		str := fmt.Sprintf("-- Can Read Values Do Not Match: %t | %t", r.CanRead, c.CanRead)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Can Read Values Match: %t | %t", r.CanRead, c.CanRead)
		details = append(details, str)
	}

	// Check Can Write Value
	if c.CanWrite != r.CanWrite {
		problemsFound++
		str := fmt.Sprintf("-- Can Write Values Do Not Match: %t | %t", r.CanWrite, c.CanWrite)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Can Write Values Match: %t | %t", r.CanWrite, c.CanWrite)
		details = append(details, str)
	}

	// Check Media Type Property Length
	if len(c.MediaTypes) != len(r.MediaTypes) {
		problemsFound++
		str := fmt.Sprintf("-- Media Type Lengths Do Not Match: %d | %d", r.MediaTypes, c.MediaTypes)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Media Type Lengths Match: %d | %d", r.MediaTypes, c.MediaTypes)
		details = append(details, str)

		// If lengths are the same, then check each value
		for index, _ := range r.MediaTypes {
			if c.MediaTypes[index] != r.MediaTypes[index] {
				problemsFound++
				str := fmt.Sprintf("-- Media Types Do Not Match: %s | %s", r.MediaTypes[index], c.MediaTypes[index])
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Media Types Match: %s | %s", r.MediaTypes[index], c.MediaTypes[index])
				details = append(details, str)
			}
		}
	}

	if problemsFound > 0 {
		return false, problemsFound, details
	}

	return true, 0, details
}
