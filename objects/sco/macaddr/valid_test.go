// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package macaddr

import (
	"testing"
)

// ----------------------------------------------------------------------
// Tests
// ----------------------------------------------------------------------

/*
TestValid1 - Make sure we get a value of false when MACAddr obj is blank.
*/
func TestValid1(t *testing.T) {
	m := New()
	want := false

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail MACAddr Object should be invalid when empty")
		t.Log(err)
	}
}

/*
TestValid2 - Valid with colon-separated MAC address
*/
func TestValid2(t *testing.T) {
	m := New()
	want := true

	m.SetValue("00:11:22:33:44:55")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail MACAddr Object should be valid with colon-separated MAC")
		t.Log(err)
	}
}

/*
TestValid3 - Valid with dash-separated MAC address
*/
func TestValid3(t *testing.T) {
	m := New()
	want := true

	m.SetValue("00-11-22-33-44-55")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail MACAddr Object should be valid with dash-separated MAC")
		t.Log(err)
	}
}

/*
TestValid4 - Valid with Cisco-style MAC address
*/
func TestValid4(t *testing.T) {
	m := New()
	want := true

	m.SetValue("0011.2233.4455")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail MACAddr Object should be valid with Cisco-style MAC")
		t.Log(err)
	}
}

/*
TestValid5 - Invalid with malformed MAC address
*/
func TestValid5(t *testing.T) {
	m := New()
	want := false

	m.SetValue("00:11:22:33:44")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail MACAddr Object should be invalid with incomplete MAC")
		t.Log(err)
	}
}

/*
TestValid6 - Invalid with non-hex characters
*/
func TestValid6(t *testing.T) {
	m := New()
	want := false

	m.SetValue("00:11:22:33:44:GG")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail MACAddr Object should be invalid with non-hex characters")
		t.Log(err)
	}
}

/*
TestNew - Ensure New() creates proper object
*/
func TestNew(t *testing.T) {
	m := New()
	
	if m.ObjectType != "mac-addr" {
		t.Error("Fail: Object type not set correctly")
	}
	
	if m.SpecVersion == "" {
		t.Error("Fail: Spec version not set")
	}
}
