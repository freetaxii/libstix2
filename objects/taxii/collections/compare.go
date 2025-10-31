// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package collections

import "fmt"

/*
Compare - This method will compare two collections to make sure they
are the same. The collection receiver is the master and represent the correct
data, the collection passed in as toTest represents the one we need to test.
*/
func (r *Collection) Compare(toTest *Collection) (bool, int, []string) {
	return Compare(r, toTest)
}

/*
Compare - This function will compare two collections to make sure they
are the same. Collection correct is the master and represent the correct
data, collection toTest represents the one we need to test.
*/
func Compare(correct, toTest *Collection) (bool, int, []string) {
	problemsFound := 0
	details := make([]string, 0)

	// Check ID Value
	if toTest.ID != correct.ID {
		problemsFound++
		str := fmt.Sprintf("-- IDs Do Not Match: %s | %s", correct.ID, toTest.ID)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ IDs Match: %s | %s", correct.ID, toTest.ID)
		details = append(details, str)
	}

	// Check Title Value
	if toTest.Title != correct.Title {
		problemsFound++
		str := fmt.Sprintf("-- Titles Do Not Match: %s | %s", correct.Title, toTest.Title)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Titles Match: %s | %s", correct.Title, toTest.Title)
		details = append(details, str)
	}

	// Check Description Value
	if toTest.Description != correct.Description {
		problemsFound++
		str := fmt.Sprintf("-- Descriptions Do Not Match: %s | %s", correct.Description, toTest.Description)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Descriptions Match: %s | %s", correct.Description, toTest.Description)
		details = append(details, str)
	}

	// Check Can Read Value
	if toTest.CanRead != correct.CanRead {
		problemsFound++
		str := fmt.Sprintf("-- Can Read Values Do Not Match: %t | %t", correct.CanRead, toTest.CanRead)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Can Read Values Match: %t | %t", correct.CanRead, toTest.CanRead)
		details = append(details, str)
	}

	// Check Can Write Value
	if toTest.CanWrite != correct.CanWrite {
		problemsFound++
		str := fmt.Sprintf("-- Can Write Values Do Not Match: %t | %t", correct.CanWrite, toTest.CanWrite)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Can Write Values Match: %t | %t", correct.CanWrite, toTest.CanWrite)
		details = append(details, str)
	}

	// Check Media Type Property Length
	if len(toTest.MediaTypes) != len(correct.MediaTypes) {
		problemsFound++
		str := fmt.Sprintf("-- Media Type Lengths Do Not Match: %v | %v", correct.MediaTypes, toTest.MediaTypes)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Media Type Lengths Match: %v | %v", correct.MediaTypes, toTest.MediaTypes)
		details = append(details, str)

		// If lengths are the same, then check each value
		for index := range correct.MediaTypes {
			if toTest.MediaTypes[index] != correct.MediaTypes[index] {
				problemsFound++
				str := fmt.Sprintf("-- Media Types Do Not Match: %s | %s", correct.MediaTypes[index], toTest.MediaTypes[index])
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Media Types Match: %s | %s", correct.MediaTypes[index], toTest.MediaTypes[index])
				details = append(details, str)
			}
		}
	}

	if problemsFound > 0 {
		return false, problemsFound, details
	}

	return true, 0, details
}
