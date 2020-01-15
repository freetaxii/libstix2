// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package resources

import (
	"errors"
	"strings"
)

/*
AddValuesToList - This function will add a single value, a comma separated
list of values, or a slice of values to an slice.
*/
func AddValuesToList(list *[]string, values interface{}) error {

	switch values.(type) {
	case string:
		sliceOfValues := strings.Split(values.(string), ",")
		// Get rid of any leading or trailing whitespace
		// example: values = "test, test1 , test2"
		for i, v := range sliceOfValues {
			sliceOfValues[i] = strings.TrimSpace(v)
		}
		*list = append(*list, sliceOfValues...)
	case []string:
		// Get rid of any leading or trailing whitespace
		for i, v := range values.([]string) {
			values.([]string)[i] = strings.TrimSpace(v)
		}
		*list = append(*list, values.([]string)...)
	default:
		return errors.New("invalid data passed in to AddValuesToList()")
	}

	return nil
}
