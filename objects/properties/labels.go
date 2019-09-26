// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import (
	"fmt"
	"strings"
)

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
LabelsProperty - A property used by one or more STIX objects that
captures a list of labels or tags for a STIX object. On some objects the
labels property is defined as coming from an open-vocab.
*/
type LabelsProperty struct {
	Labels []string `json:"labels,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - LabelsProperty
// ----------------------------------------------------------------------

/*
AddLabel - This method takes in a string value that represents one or more
labels separated by a command for a STIX object and adds it to the list of
labels in the labels property.
*/
func (o *LabelsProperty) AddLabel(s string) error {

	labels := strings.Split(s, ",")
	for _, label := range labels {
		o.Labels = append(o.Labels, label)
	}

	return nil
}

/*
CompareLabelsProperties - This function will compare two labels properties
(object 1 and object 2) to make sure they are the same. This function will
return an integer that tracks the number of problems and a slice of strings that
contain the detailed results, whether good or bad.
*/
func CompareLabelsProperties(obj1, obj2 *LabelsProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Labels Values
	if len(obj1.Labels) != len(obj2.Labels) {
		problemsFound++
		str := fmt.Sprintf("-- Labels Length Do Not Match: %d | %d", len(obj1.Labels), len(obj2.Labels))
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ Labels Length Match: %d | %d", len(obj1.Labels), len(obj2.Labels))
		resultDetails = append(resultDetails, str)

		// If lengths are the same, then check each value
		for index := range obj1.Labels {
			if obj1.Labels[index] != obj2.Labels[index] {
				problemsFound++
				str := fmt.Sprintf("-- Labels Do Not Match: %s | %s", obj1.Labels[index], obj2.Labels[index])
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ Labels Match: %s | %s", obj1.Labels[index], obj2.Labels[index])
				resultDetails = append(resultDetails, str)
			}
		}
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
