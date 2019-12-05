// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package ipv4addrobject

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
TestValid1 - Make sure we get a value of false when IPv4AddrObject obj is blank.
*/
func TestValid1(t *testing.T) {
	m := New()
	want := false

	if got, err := m.Valid(); got != want {
		t.Error("Fail IPv4AddrObject Object should be invalid when empty")
		t.Log(err)
	}
}

//TestValid2 -
func TestValid2(t *testing.T) {
	m := New()
	want := true

	m.SetValue("127.0.0.1")

	if got, err := m.Valid(); got != want {
		t.Error("Fail IPv4AddrObject Object should be valid when required fields are not empty")
		t.Log(err)
	}
}

//TestValid9 - should be valid
func TestValid9(t *testing.T) {
	m := New()
	want := true

	m.SetValue("127.0.0.1")
	m.AddResolvesToRefs([]string{"mac-addr--efcd5e80-570d-4131-b213-62cb18eaa6a8", "mac-addr--efcd5e80-570d-4131-b213-62cb18eaa6a9"})
	m.AddResolvesToRef("mac-addr--efcd5e80-570d-4131-b213-62cb18eaa6a7")
	m.AddBelongsToRefs([]string{"3", "4"})
	m.AddBelongsToRef("5")

	if got, err := m.Valid(); got != want {
		t.Error("Fail IPv4AddrObject Object shoulf be valid")
		t.Log(err)
	}
}
