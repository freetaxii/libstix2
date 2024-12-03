// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package indicator

import "testing"

// TestAddType -
func TestAddType(t *testing.T) {
	i := New()
	want := "testData"
	i.AddTypes(want)

	if got := i.IndicatorTypes[0]; got != want {
		t.Error("Fail Indicator Add Type Check")
	}
}

// TestSetPattern -
func TestSetPattern(t *testing.T) {
	i := New()
	want := "testData"
	i.SetPattern(want)

	if got := i.Pattern; got != want {
		t.Error("Fail Indicator Set Pattern Check")
	}
}

// TestSetPatternType1 -
func TestSetPatternType1(t *testing.T) {
	i := New()
	want := ""
	i.SetPatternType("testData")

	if got := i.PatternType; got != want {
		t.Error("Fail Indicator Set Pattern Type Check 1")
	}
}

// TestSetPatternType2 -
func TestSetPatternType2(t *testing.T) {
	i := New()
	want := "stix"
	i.SetPatternType(want)

	if got := i.PatternType; got != want {
		t.Error("Fail Indicator Set Pattern Type Check 2")
	}
}

// TestSetPatternVersion -
func TestSetPatternVersion(t *testing.T) {
	i := New()
	want := "testData"
	i.SetPatternVersion(want)

	if got := i.PatternVersion; got != want {
		t.Error("Fail Indicator Set Pattern Version Check")
	}
}

// TODO Finish fleshing this out. We need the valid from and valid until tests
// both the positive and negative tests.
