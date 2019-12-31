// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import (
	"fmt"
	"strings"
)

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/* LabelsProperty - A property used by one or more STIX objects that captures a
list of labels or tags for a STIX object. On some objects the labels property is
defined as coming from an open-vocab. */
type LabelsProperty struct {
	Labels []string `json:"labels,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - LabelsProperty - Setters
// ----------------------------------------------------------------------

/* AddLabel - This method takes in a string value that represents one or more
labels separated by a command for a STIX object and adds it to the list of
labels in the labels property. */
func (o *LabelsProperty) AddLabel(s string) error {

	labels := strings.Split(s, ",")
	for _, label := range labels {
		o.Labels = append(o.Labels, label)
	}

	return nil
}

// ----------------------------------------------------------------------
// Public Methods - LabelsProperty - Checks
// ----------------------------------------------------------------------

/* Compare - This method will compare two properties to make sure they are the
same and will return a boolean, an integer that tracks the number of problems
found, and a slice of strings that contain the detailed results, whether good or
bad. */
func (o *LabelsProperty) Compare(obj2 *LabelsProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Labels Values
	if len(o.Labels) != len(obj2.Labels) {
		problemsFound++
		str := fmt.Sprintf("-- The number of entries in labels do not match: %d | %d", len(o.Labels), len(obj2.Labels))
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The number of entries in labels match: %d | %d", len(o.Labels), len(obj2.Labels))
		resultDetails = append(resultDetails, str)

		// If lengths are the same, then check each value
		for index := range o.Labels {
			if o.Labels[index] != obj2.Labels[index] {
				problemsFound++
				str := fmt.Sprintf("-- The label values do not match: %s | %s", o.Labels[index], obj2.Labels[index])
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ The label values match: %s | %s", o.Labels[index], obj2.Labels[index])
				resultDetails = append(resultDetails, str)
			}
		}
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
