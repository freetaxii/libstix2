// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/freetaxii/libstix2/defs"
)

// ValidObjectType - This function will take in a STIX object type and return
// true if the string represents an actual STIX object type. This is used for
// determining if input from an outside source is actually a defined STIX object or
// not.
func ValidObjectType(t string) bool {

	var m = map[string]int{
		// SDOs
		"attack-pattern":   1,
		"campaign":         1,
		"course-of-action": 1,
		"grouping":         1,
		"identity":         1,
		"indicator":        1,
		"infrastructure":   1,
		"intrusion-set":    1,
		"location":         1,
		"malware":          1,
		"malware-analysis": 1,
		"note":             1,
		"observed-data":    1,
		"opinion":          1,
		"report":           1,
		"threat-actor":     1,
		"tool":             1,
		"vulnerability":    1,
		// SROs
		"relationship": 1,
		"sighting":     1,
		// SCOs
		"artifact":             1,
		"autonomous-system":    1,
		"directory":            1,
		"domain-name":          1,
		"email-addr":           1,
		"email-message":        1,
		"file":                 1,
		"ipv4-addr":            1,
		"ipv6-addr":            1,
		"mac-addr":             1,
		"mutex":                1,
		"network-traffic":      1,
		"process":              1,
		"software":             1,
		"url":                  1,
		"user-account":         1,
		"windows-registry-key": 1,
		"x509-certificate":     1,
		// Meta Objects
		"language-content":   1,
		"marking-definition": 1,
		// Bundle
		"bundle": 1,
	}

	if _, ok := m[t]; ok {
		return true
	}
	return false
}

// GetCommonProperties - This method will return a pointer to the common
// properties of this object.
func (o *CommonObjectProperties) GetCommonProperties() *CommonObjectProperties {
	return o
}

// ----------------------------------------------------------------------
// Helper Functions
// ----------------------------------------------------------------------

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

// IsVocabEntryValid - This function determine will take in a vocabulary and a
// value and return true if it is found and false if it is not found.
func IsVocabEntryValid(vocab map[string]bool, s string) bool {
	if _, found := vocab[s]; found == true {
		return true
	}
	return false
}
