// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package urlobject

import (
	"testing"
)

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
TestValid1 - Make sure we get a value of false when UrlObject obj is blank.
*/
func TestValid1(t *testing.T) {
	m := New()
	want := false

	if got, err := m.Valid(false); got != want {
		t.Error("Fail UrlObject Object should be invalid when empty")
		t.Log(err)
	}
}

//TestValid2 -
func TestValid2(t *testing.T) {
	m := New()
	want := true

	m.SetValue("https://example.com/research/index.html")

	if got, err := m.Valid(false); got != want {
		t.Error("Fail UrlObject Object should be valid when required fields are not empty")
		t.Log(err)
	}
}
