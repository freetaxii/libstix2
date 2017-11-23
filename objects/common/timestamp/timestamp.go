// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package timestamp

import (
	"fmt"
	"github.com/freetaxii/libstix2/defs"
	"time"
)

// ----------------------------------------------------------------------
// Public Functions
// ----------------------------------------------------------------------

// GetCurrentTime - This function will return the current time in STIX timestamp
// format, which is in RFC 3339 format.
func GetCurrentTime() string {
	return time.Now().UTC().Format(defs.TIME_RFC_3339)
}

func GetCurrentTimeMilli() string {
	return time.Now().UTC().Format(defs.TIME_RFC_3339_MILLI)
}

// Verify - This function takes in a timestamp in either time.Time or string
// format and returns a string version of the timestamp.
func Verify(t interface{}) string {
	switch ts := t.(type) {
	case time.Time:
		return ts.UTC().Format(defs.TIME_RFC_3339)
	case string:
		//TODO verify format of timestamp when in string format
		return ts
	default:
		return fmt.Sprintf("The timestamp format of \"%s\" is not a valid format", ts)
	}
}

func VerifyMilli(t interface{}) string {
	switch ts := t.(type) {
	case time.Time:
		return ts.UTC().Format(defs.TIME_RFC_3339_MILLI)
	case string:
		//TODO verify format of timestamp when in string format
		return ts
	default:
		return fmt.Sprintf("The timestamp format of \"%s\" is not a valid format", ts)
	}
}

// // VerifyPrecision will verify the supplied precision string to make sure it
// // is valid per the STIX specification.
// func VerifyPrecision(s string) (string, error) {

// 	if s == "" {
// 		return "", nil
// 	}

// 	s = strings.ToLower(s)
// 	switch s {
// 	case "year":
// 		return s, nil
// 	case "month":
// 		return s, nil
// 	case "day":
// 		return s, nil
// 	case "hour":
// 		return s, nil
// 	case "minute":
// 		return s, nil
// 	case "full":
// 		return s, nil
// 	default:
// 		return "", fmt.Errorf("invalid precision \"%s\", setting requested precision to \"\"", s)
// 	}
// }
