// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package common

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type FirstLastSeenType struct {
	First_seen           string `json:"first_seen,omitempty"`
	First_seen_precision string `json:"first_seen_precision,omitempty"`
	Last_seen            string `json:"last_seen,omitempty"`
	Last_seen_precision  string `json:"last_seen_precision,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - FirstLastSeenType
// ----------------------------------------------------------------------

// SetFirstSeen takes in two parameters and returns and error if there is one
// param: t a timestamp in either time.Time or string format
// param: s a timestamp precision in string format
func (this *FirstLastSeenType) SetFirstSeen(t interface{}, s string) error {

	ts, err := VerifyTimestamp(t)
	if err != nil {
		return err
	}
	this.First_seen = ts

	p, err := VerifyPrecision(s)
	if err != nil {
		return err
	}
	this.First_seen_precision = p

	return nil
}

// SetLastSeen takes in two parameters and returns and error if there is one
// param: t a timestamp in either time.Time or string format
// param: s a timestamp precision in string format
func (this *FirstLastSeenType) SetLastSeen(t interface{}, s string) error {

	ts, err := VerifyTimestamp(t)
	if err != nil {
		return err
	}
	this.Last_seen = ts

	p, err := VerifyPrecision(s)
	if err != nil {
		return err
	}
	this.Last_seen_precision = p

	return nil
}
