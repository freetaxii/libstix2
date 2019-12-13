// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import "fmt"

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
ConfidenceProperty - A property used by one or more STIX objects that
captures the STIX confidence score, which is a value from 0-100.
*/
type ConfidenceProperty struct {
	Confidence int `json:"confidence,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - ConfidenceProperty
// ----------------------------------------------------------------------

/*
SetConfidence - This method takes in an integer representing a STIX
confidence level 0-100 and updates the Confidence property.
*/
func (o *ConfidenceProperty) SetConfidence(i int) error {
	o.Confidence = i
	return nil
}

/*
GetConfidence - This method returns the confidence value as an integer.
*/
func (o *ConfidenceProperty) GetConfidence() int {
	return o.Confidence
}

// ----------------------------------------------------------------------
// Public Functions - ConfidenceProperty
// ----------------------------------------------------------------------

/* CompareConfidenceProperties - This function will compare two properties to
make sure they are the same and will return a boolean, an integer that tracks
the number of problems found, and a slice of strings that contain the detailed
results, whether good or bad. */
func CompareConfidenceProperties(obj1, obj2 *ConfidenceProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Confidence Value
	if obj1.Confidence != obj2.Confidence {
		problemsFound++
		str := fmt.Sprintf("-- The Confidence values do not match: %d | %d", obj1.Confidence, obj2.Confidence)
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The Confidence values match: %d | %d", obj1.Confidence, obj2.Confidence)
		resultDetails = append(resultDetails, str)
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
