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

// ----------------------------------------------------------------------
// Test the public Valid method - Make sure we hit each level and make sure
// required property checks are working when they are left blank.
// ----------------------------------------------------------------------

/*
TestValid1 - Make sure we get a value of false when Indicator Type is blank.
*/
func TestValid1(t *testing.T) {
	i := New()
	want := false

	if got, err := i.Valid(); got != want {
		t.Error("Fail Indicator Type Check 0")
		t.Log(err)
	}
}

/*
TestValid2 - Make sure we get a value of false when Pattern is blank.
*/
func TestValid2(t *testing.T) {
	i := New()
	want := false
	// Set the Indicator Type value so we can move to next test.
	i.IndicatorTypes = append(i.IndicatorTypes, "TestValue")

	if got, err := i.Valid(); got != want {
		t.Error("Fail Pattern Check 0")
		t.Log(err)
	}
}

/*
TestValid3 - Make sure we get a value of false when Pattern Type is blank.
*/
func TestValid3(t *testing.T) {
	i := New()
	want := false
	// Set the Indicator Type and Pattern value so we can move to the next test.
	i.IndicatorTypes = append(i.IndicatorTypes, "TestValue")
	i.Pattern = "TestPattern"

	if got, err := i.Valid(); got != want {
		t.Error("Fail Pattern Type Check 0")
		t.Log(err)
	}
}

/*
TestValid4 - Make sure we get a value of false when Valid From is blank.
*/
func TestValid4(t *testing.T) {
	i := New()
	want := false
	// Set the Indicator Type, Pattern, and Pattern Type value so we can move to
	// the next test.
	i.IndicatorTypes = append(i.IndicatorTypes, "TestValue")
	i.Pattern = "TestPattern"
	i.PatternType = "stix"

	if got, err := i.Valid(); got != want {
		t.Error("Fail Valid From Check 0")
		t.Log(err)
	}
}

/*
TestValid5 - Make sure we get a value of false when Valid Until is invalid.
*/
func TestValid5(t *testing.T) {
	i := New()
	want := false

	// Set the Indicator Type, Pattern, and Pattern Type value so we can move to.
	// the next test
	i.IndicatorTypes = append(i.IndicatorTypes, "TestValue")
	i.Pattern = "TestPattern"
	i.PatternType = "stix"

	// We need the next test for Valid Until to fail so lets set a bad time.
	i.ValidUntil = "2019-0924T20:49:12.123456Z"

	if got, err := i.Valid(); got != want {
		t.Error("Fail Valid Until Check 0")
		t.Log(err)
	}
}

/*
TestValid6 - Make sure we get a value of false when Valid Until is before Valid
From.
*/
func TestValid6(t *testing.T) {
	i := New()
	want := false

	// Set the Indicator Type, Pattern, and Pattern Type value so we can move to
	// the next test.
	i.IndicatorTypes = append(i.IndicatorTypes, "TestValue")
	i.Pattern = "TestPattern"
	i.PatternType = "stix"

	// We need the next test for Valid Until to fail so lets set a bad time.
	i.ValidFrom = "2019-09-24T20:49:13.123456Z"
	i.ValidUntil = "2019-09-24T20:49:12.123456Z"

	if got, err := i.Valid(); got != want {
		t.Error("Fail Valid Until Check 0")
		t.Log(err)
	}
}

/*
TestValid7 - Make sure we get a value of true when everything is filled out
correctly.
*/
func TestValid7(t *testing.T) {
	i := New()
	want := true

	// Set the Indicator Type, Pattern, and Pattern Type value so we can move to
	// the next test.
	i.IndicatorTypes = append(i.IndicatorTypes, "TestValue")
	i.Pattern = "TestPattern"
	i.PatternType = "stix"

	// Set the timestamps correctly
	i.ValidFrom = "2019-09-24T20:49:12.123456Z"
	i.ValidUntil = "2019-09-24T20:49:13.123456Z"

	if got, err := i.Valid(); got != want {
		t.Error("Fail Valid Until Check 0")
		t.Log(err)
	}
}

// ----------------------------------------------------------------------
// Test individual private validation methods
// ----------------------------------------------------------------------

/*
TestValidIndicatorType1 - Make sure we get a value of false when the required
Indicator Type property is left blank.
*/
func TestValidIndicatorType1(t *testing.T) {
	i := New()
	want := false

	if got, err := i.validIndicatorType(); got != want {
		t.Error("Fail Indicator Type Check 1")
		t.Log(err)
	}
}

/*
TestValidIndicatorType2 - Make sure we get a value of true when the Indicator
Type is populated.
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
TestValidPattern1 - Make sure we get a value of false when the required Pattern
property is left blank.
*/
func TestValidPattern1(t *testing.T) {
	i := New()
	want := false

	if got, err := i.validPattern(); got != want {
		t.Error("Fail Pattern Check 1")
		t.Log(err)
	}
}

/*
TestValidPattern2 - Make sure we get a value of true when the Pattern is
populated.
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
TestValidPatternType1 - Make sure we get a value of false when the required
Pattern Type property is left blank.
*/
func TestValidPatternType1(t *testing.T) {
	i := New()
	want := false

	if got, err := i.validPatternType(); got != want {
		t.Error("Fail Pattern Type Check 1")
		t.Log(err)
	}
}

/*
TestValidPatternType2 - Make sure we get a value of true when the Pattern Type
is populated correctly.
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
is populated incorrectly.
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
TestValidValidFrom1 - Make sure we get a value of false when the required Valid
From property is left blank.
*/
func TestValidValidFrom1(t *testing.T) {
	i := New()
	want := false

	if got, err := i.validValidFrom(); got != want {
		t.Error("Fail Valid From Check 1")
		t.Log(err)
	}
}

/*
TestValidValidFrom2 - Make sure we get a value of true when the Valid From
timestamp is populated correctly.
*/
func TestValidValidFrom2(t *testing.T) {
	i := New()
	want := true
	i.ValidFrom = "2019-09-24T20:49:12.123456Z"
	if got, err := i.validValidFrom(); got != want {
		t.Error("Fail Valid From Check 2")
		t.Log(err)
	}
}

/*
TestValidValidFrom3 - Make sure we get a value of false when the Valid From
timestamp is populated incorrectly.
*/
func TestValidValidFrom3(t *testing.T) {
	i := New()
	want := false
	i.ValidFrom = "20190924T20:49:12.123456Z"
	if got, err := i.validValidFrom(); got != want {
		t.Error("Fail Valid From Check 3")
		t.Log(err)
	}
}

/*
TestValidValidUntil2 - Make sure we get a value of true when the Valid From and
Valid Until timestamps are populated correctly.
*/
func TestValidValidUntil2(t *testing.T) {
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
TestValidValidUntil3 - Make sure we get a value of false when Valid Until
timestamp is populated correctly but there is no Valid From timestamp.
*/
func TestValidValidUntil3(t *testing.T) {
	i := New()
	want := false
	i.ValidUntil = "2019-29-24T20:49:13.123456Z"
	if got, err := i.validValidUntil(); got != want {
		t.Error("Fail Valid Until Check 4")
		t.Log(err)
	}
}

/*
TestValidValidUntil4 - Make sure we get a value of false when timestamps are
populated incorrectly.
*/
func TestValidValidUntil4(t *testing.T) {
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
TestValidValidUntil5 - Make sure we get a value of false when timestamps are
populated correct and are valid but the Valid Until timestamp is before the
Valid From timestamp.
*/
func TestValidValidUntil5(t *testing.T) {
	i := New()
	want := false
	i.ValidFrom = "2019-09-24T20:49:13.123456Z"
	i.ValidUntil = "2019-29-24T20:49:12.123456Z"
	if got, err := i.validValidUntil(); got != want {
		t.Error("Fail Valid Until Check 5")
		t.Log(err)
	}
}
