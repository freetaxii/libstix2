// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package ipv6addr

import (
	"testing"
)

// ----------------------------------------------------------------------
// Tests
// ----------------------------------------------------------------------

/*
TestValid1 - Make sure we get a value of false when IPv6Addr obj is blank.
*/
func TestValid1(t *testing.T) {
	m := New()
	want := false

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail IPv6Addr Object should be invalid when empty")
		t.Log(err)
	}
}

/*
TestValid2 - Valid with standard IPv6 address
*/
func TestValid2(t *testing.T) {
	m := New()
	want := true

	m.SetValue("2001:0db8:85a3:0000:0000:8a2e:0370:7334")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail IPv6Addr Object should be valid with standard IPv6 address")
		t.Log(err)
	}
}

/*
TestValid3 - Valid with shortened IPv6 address
*/
func TestValid3(t *testing.T) {
	m := New()
	want := true

	m.SetValue("2001:db8::1")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail IPv6Addr Object should be valid with shortened IPv6 address")
		t.Log(err)
	}
}

/*
TestValid4 - Valid with IPv6 CIDR notation
*/
func TestValid4(t *testing.T) {
	m := New()
	want := true

	m.SetValue("2001:db8::/32")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail IPv6Addr Object should be valid with IPv6 CIDR notation")
		t.Log(err)
	}
}

/*
TestValid5 - Invalid with IPv4 address
*/
func TestValid5(t *testing.T) {
	m := New()
	want := false

	m.SetValue("192.168.1.1")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail IPv6Addr Object should be invalid with IPv4 address")
		t.Log(err)
	}
}

/*
TestValid6 - Valid with all optional properties
*/
func TestValid6(t *testing.T) {
	m := New()
	want := true

	m.SetValue("2001:db8::1")
	m.AddResolvesToRefs([]string{"mac-addr--efcd5e80-570d-4131-b213-62cb18eaa6a8"})
	m.AddBelongsToRefs("autonomous-system--f720c34b-98ae-597f-ade5-27dc241e8c74")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail IPv6Addr Object should be valid with all properties")
		t.Log(err)
	}
}

/*
TestValid7 - Invalid with malformed IPv6
*/
func TestValid7(t *testing.T) {
	m := New()
	want := false

	m.SetValue("gggg::1")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail IPv6Addr Object should be invalid with malformed IPv6")
		t.Log(err)
	}
}

/*
TestNew - Ensure New() creates proper object
*/
func TestNew(t *testing.T) {
	m := New()
	
	if m.ObjectType != "ipv6-addr" {
		t.Error("Fail: Object type not set correctly")
	}
	
	if m.SpecVersion == "" {
		t.Error("Fail: Spec version not set")
	}
}
