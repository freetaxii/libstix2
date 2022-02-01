// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package timestamp

/*
CurrentTime - This function will return the current time in the STIX
timestamp format, which is in RFC 3339 format. The options are "milli",
"micro", or "" which will give you to the second.
*/
// func CurrentTime(precision string) string {
// 	if precision == "milli" {
// 		return time.Now().UTC().Format(defs.TIME_RFC_3339_MILLI)
// 	} else if precision == "micro" {
// 		return time.Now().UTC().Format(defs.TIME_RFC_3339_MICRO)
// 	}
// 	return time.Now().UTC().Format(defs.TIME_RFC_3339)
// }

/*
Valid - This function will return true if the timestamp is a valid STIX
timestamp, just the date in proper RFC3339 format, or just the year. If it
is anything else it will return false.
*/
// func Valid(t string) bool {
// 	re1 := regexp.MustCompile(`^[12]\d{3}-[01]\d{1}-[0-3]\d{1}T[0-2]\d{1}:[0-5]\d{1}:[0-5]\d{1}(\.\d{0,6})?Z$`)
// 	//re2 := regexp.MustCompile(`^[12]\d{3}-[01]\d{1}-[0-3]\d{1}$`)
// 	//re3 := regexp.MustCompile(`^[12]\d{3}$`)
// 	if re1.MatchString(t) {
// 		return true
// 	}
// 	//else if re2.MatchString(t) {
// 	//	return true
// 	//} else if re3.MatchString(t) {
// 	//	return true
// 	//}
// 	return false
// }

/*
ToString - This function takes in a timestamp in either time.Time or string
format and returns a string version of the timestamp.
*/
// func ToString(t interface{}, p string) (string, error) {
// 	// TODO: One potential problem is if the time is created with the time package
// 	// at a precision less than micro and we set it to micro in things like
// 	// indicator, observed_data, first_seen, and last_seen for example

// 	var format string
// 	if p == "milli" {
// 		format = defs.TIME_RFC_3339_MILLI
// 	} else if p == "micro" {
// 		format = defs.TIME_RFC_3339_MICRO
// 	} else {
// 		format = defs.TIME_RFC_3339
// 	}

// 	switch ts := t.(type) {
// 	case time.Time:
// 		return ts.UTC().Format(format), nil
// 	case string:
// 		//TODO verify format of timestamp when in string format
// 		return ts, nil
// 	default:
// 		return "", fmt.Errorf("the timestamp format of \"%s\" is not a valid format", ts)
// 	}
// }

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
