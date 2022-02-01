// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package resources

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/freetaxii/libstix2/defs"
)

/*
AddValuesToList - This function will add a single value, a comma separated
list of values, or a slice of values to an slice.
*/
// func AddValuesToList(list *[]string, values interface{}) error {

// 	switch values.(type) {
// 	case string:
// 		sliceOfValues := strings.Split(values.(string), ",")
// 		// Get rid of any leading or trailing whitespace
// 		// example: values = "test, test1 , test2"
// 		for i, v := range sliceOfValues {
// 			sliceOfValues[i] = strings.TrimSpace(v)
// 		}
// 		*list = append(*list, sliceOfValues...)
// 	case []string:
// 		// Get rid of any leading or trailing whitespace
// 		for i, v := range values.([]string) {
// 			values.([]string)[i] = strings.TrimSpace(v)
// 		}
// 		*list = append(*list, values.([]string)...)
// 	default:
// 		return errors.New("invalid data passed in to AddValuesToList()")
// 	}

// 	return nil
// }

// IsUUIDValid - This function will take in a string and return true if the
// string represents an actual UUID v4 or v5 value.
func IsUUIDValid(uuid string) bool {
	r := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4-5][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
	return r.MatchString(uuid)
}

// GetCurrentSpecVersion - This function returns the current specification version
// that this library is using.
func GetCurrentSpecVersion() string {
	return defs.CurrentSTIXVersion
}

// GetCurrentTime - This function takes in a value of either milli or micro and
// returns the current time in RFC 3339 format
func GetCurrentTime(precision string) string {
	if precision == "milli" {
		return time.Now().UTC().Format(defs.TimeRFC3339Milli)
	} else if precision == "micro" {
		return time.Now().UTC().Format(defs.TimeRFC3339Micro)
	}
	return time.Now().UTC().Format(defs.TimeRFC3339)
}

// TimeToString - This function takes in a timestamp in either time.Time or string
// format and returns a string version of the timestamp.
func TimeToString(t interface{}, precision string) (string, error) {
	// TODO: One potential problem is if the time is created with the time package
	// at a precision less than micro and we set it to micro for some object.

	var format string
	if precision == "milli" {
		format = defs.TimeRFC3339Milli
	} else if precision == "micro" {
		format = defs.TimeRFC3339Micro
	} else {
		format = defs.TimeRFC3339
	}

	switch ts := t.(type) {
	case time.Time:
		return ts.UTC().Format(format), nil
	case string:
		//TODO verify format of timestamp when in string format
		return ts, nil
	default:
		return "", fmt.Errorf("the timestamp format of \"%s\" is not a valid format", ts)
	}
}

// IsTimestampValid - This function will take in a timestamp and check to see if
// it is valid per the specification.
func IsTimestampValid(t string) bool {
	re1 := regexp.MustCompile(`^[12]\d{3}-[01]\d{1}-[0-3]\d{1}T[0-2]\d{1}:[0-5]\d{1}:[0-5]\d{1}(\.\d{0,6})?Z$`)
	//re2 := regexp.MustCompile(`^[12]\d{3}-[01]\d{1}-[0-3]\d{1}$`)
	//re3 := regexp.MustCompile(`^[12]\d{3}$`)
	if re1.MatchString(t) {
		return true
	}
	//else if re2.MatchString(t) {
	//	return true
	//} else if re3.MatchString(t) {
	//	return true
	//}
	return false
}

// AddValuesToList - This function will add a single value, a comma separated
// list of values, or a slice of values to an slice.
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
