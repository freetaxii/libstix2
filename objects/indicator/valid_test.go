// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package indicator

import "testing"

// ----------------------------------------------------------------------
// Tests
// These tests will not use the setters as some setters will have their
// own logic and verification steps in them.
// ----------------------------------------------------------------------

/*
TestValid - Make sure we hit each level and make sure required property checks
are working when they are left blank
*/
func TestValid(t *testing.T) {
	i := New()
	want := false

	// Make sure we get a value of false when blank
	if got, err := i.Valid(); got != want {
		t.Error("Fail Indicator Type Check 0")
		t.Log(err)
	}

	// Set the Indicator Type value so we can move to next test
	i.IndicatorTypes = append(i.IndicatorTypes, "TestValue")

	if got, err := i.Valid(); got != want {
		t.Error("Fail Pattern Check 0")
		t.Log(err)
	}

	// Set the Pattern value so we can move to the next test
	i.Pattern = "TestPattern"

	if got, err := i.Valid(); got != want {
		t.Error("Fail Pattern Type Check 0")
		t.Log(err)
	}

	// Set the Pattern Type value so we can move to the next test
	i.PatternType = "stix"

	if got, err := i.Valid(); got != want {
		t.Error("Fail Valid From Check 0")
		t.Log(err)
	}

	// We need the next test for Valid Until to fail so lets set a bad time
	i.ValidUntil = "2019-0924T20:49:12.123456Z"

	if got, err := i.Valid(); got != want {
		t.Error("Fail Valid Until Check 0")
		t.Log(err)
	}

	// We need the next test for Valid Until to fail so lets set a bad time
	i.ValidFrom = "2019-09-24T20:49:13.123456Z"
	i.ValidUntil = "2019-09-24T20:49:12.123456Z"

	if got, err := i.Valid(); got != want {
		t.Error("Fail Valid Until Check 0")
		t.Log(err)
	}

	// Everything is now set correctly and we should pass this test
	i.ValidFrom = "2019-09-24T20:49:12.123456Z"
	i.ValidUntil = "2019-09-24T20:49:13.123456Z"
	want = true
	if got, err := i.Valid(); got != want {
		t.Error("Fail Valid Until Check 0")
		t.Log(err)
	}

}

/*
TestValidRequiredProperties - Make sure required property checks are working
when properties are left blank we should get a valid response of "false" for
each of the tests.
*/
func TestValidRequiredProperties(t *testing.T) {
	i := New()
	want := false

	// Make sure we get a value of false when blank
	if got, err := i.validIndicatorType(); got != want {
		t.Error("Fail Indicator Type Check 1")
		t.Log(err)
	}

	if got, err := i.validPattern(); got != want {
		t.Error("Fail Pattern Check 1")
		t.Log(err)
	}

	if got, err := i.validPatternType(); got != want {
		t.Error("Fail Pattern Type Check 1")
		t.Log(err)
	}

	if got, err := i.validValidFrom(); got != want {
		t.Error("Fail Valid From Check 1")
		t.Log(err)
	}

	if got, err := i.validValidFrom(); got != want {
		t.Error("Fail Valid From Check 1")
		t.Log(err)
	}

}

/*
TestValidIndicatorType2 - Make sure we get a value of true when the Indicator
Type is populated
*/
func TestValidIndicatorType2(t *testing.T) {
	i := New()
	want := true
	i.IndicatorTypes = append(i.IndicatorTypes, "TestValue")
	if got, err := i.validIndicatorType(); got != want {
		t.Error("Fail Indicator Type Check 2")
		t.Log(err)
	}
}

/*
TestValidPattern2 - Make sure we get a value of true when the Pattern is
populated
*/
func TestValidPattern2(t *testing.T) {
	i := New()
	want := true
	i.Pattern = "TestPattern"
	if got, err := i.validPattern(); got != want {
		t.Error("Fail Pattern Check 2")
		t.Log(err)
	}
}

/*
TestValidPatternType2 - Make sure we get a value of true when the Pattern Type
is populated correctly
*/
func TestValidPatternType2(t *testing.T) {
	i := New()
	want := true
	values := []string{"stix", "snort", "yara"}
	for index := range values {
		i.PatternType = values[index]
		if got, err := i.validPatternType(); got != want {
			t.Error("Fail PatternType Check 2")
			t.Log(err)
		}
	}
}

/*
TestValidPatternType3 - Make sure we get a value of false when the Pattern Type
is populated incorrectly
*/
func TestValidPatternType3(t *testing.T) {
	i := New()
	want := false
	i.PatternType = "TestPatternType"
	if got, err := i.validPatternType(); got != want {
		t.Error("Fail PatternType Check 3")
		t.Log(err)
	}
}

/*
TestValidFrom2 - Make sure we get a value of true when the Valid From timestamp
is populated correctly
*/
func TestValidFrom2(t *testing.T) {
	i := New()
	want := true
	i.ValidFrom = "2019-09-24T20:49:12.123456Z"
	if got, err := i.validValidFrom(); got != want {
		t.Error("Fail Valid From Check 2")
		t.Log(err)
	}
}

/*
TestValidFrom3 - Make sure we get a value of false when the Valid From timestamp
is populated incorrectly
*/
func TestValidFrom3(t *testing.T) {
	i := New()
	want := false
	i.ValidFrom = "20190924T20:49:12.123456Z"
	if got, err := i.validValidFrom(); got != want {
		t.Error("Fail Valid From Check 3")
		t.Log(err)
	}
}

/*
TestValidUntil2 - Make sure we get a value of true when the Valid From and Valid
Until timestamps are populated correctly
*/
func TestValidUntil2(t *testing.T) {
	i := New()
	want := true
	i.ValidFrom = "2019-09-24T20:49:12.123456Z"
	i.ValidUntil = "2019-09-24T20:49:13.123456Z"
	if got, err := i.validValidUntil(); got != want {
		t.Error("Fail Valid Until Check 2")
		t.Log(err)
	}
}

/*
TestValidUntil3 - Make sure we get a value of false when Valid Until timestamp
is populated correctly but there is no Valid From timestamp
*/
func TestValidUntil3(t *testing.T) {
	i := New()
	want := false
	i.ValidUntil = "2019-29-24T20:49:13.123456Z"
	if got, err := i.validValidUntil(); got != want {
		t.Error("Fail Valid Until Check 4")
		t.Log(err)
	}
}

/*
TestValidUntil4 - Make sure we get a value of false when timestamps are
populated incorrectly
*/
func TestValidUntil4(t *testing.T) {
	i := New()
	want := false
	i.ValidFrom = "2019-09-24T20:49:12.123456Z"
	i.ValidUntil = "2019-2924T20:49:13.123456Z"
	if got, err := i.validValidUntil(); got != want {
		t.Error("Fail Valid Until Check 4")
		t.Log(err)
	}
}

/*
TestValidUntil5 - Make sure we get a value of false when timestamps are
populated correct and are valid but the Valid Until timestamp is before the
Valid From timestamp
*/
func TestValidUntil5(t *testing.T) {
	i := New()
	want := false
	i.ValidFrom = "2019-09-24T20:49:13.123456Z"
	i.ValidUntil = "2019-29-24T20:49:12.123456Z"
	if got, err := i.validValidUntil(); got != want {
		t.Error("Fail Valid Until Check 5")
		t.Log(err)
	}
}
