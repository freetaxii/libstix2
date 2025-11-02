// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package file

import (
	"testing"
)

// ----------------------------------------------------------------------
// Tests
// ----------------------------------------------------------------------

/*
TestValid1 - Make sure we get a value of false when File obj is blank (missing required fields).
*/
func TestValid1(t *testing.T) {
	m := New()
	want := false

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail File Object should be invalid when empty (requires name or hashes)")
		t.Log(err)
	}
}

/*
TestValid2 - Valid with name only
*/
func TestValid2(t *testing.T) {
	m := New()
	want := true

	m.SetName("malware.exe")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail File Object should be valid with name")
		t.Log(err)
	}
}

/*
TestValid3 - Valid with hashes only
*/
func TestValid3(t *testing.T) {
	m := New()
	want := true

	m.AddHash("MD5", "d41d8cd98f00b204e9800998ecf8427e")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail File Object should be valid with hashes")
		t.Log(err)
	}
}

/*
TestValid4 - Valid with both name and hashes
*/
func TestValid4(t *testing.T) {
	m := New()
	want := true

	m.SetName("malware.exe")
	m.AddHash("MD5", "d41d8cd98f00b204e9800998ecf8427e")
	m.AddHash("SHA-256", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail File Object should be valid with name and hashes")
		t.Log(err)
	}
}

/*
TestValid5 - Valid with all properties
*/
func TestValid5(t *testing.T) {
	m := New()
	want := true

	m.SetName("document.pdf")
	m.AddHash("SHA-256", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
	m.SetSize(1024)
	m.SetMimeType("application/pdf")
	m.SetMagicNumberHex("25504446")
	m.SetParentDirectoryRef("directory--93c0a9b0-520d-545d-9094-1a08ddf46b05")
	m.AddContainsRefs("file--fb0419a8-f09c-57f8-be64-71a80417591c")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail File Object should be valid with all properties")
		t.Log(err)
	}
}

/*
TestValid6 - Invalid with negative size
*/
func TestValid6(t *testing.T) {
	m := New()
	want := false

	m.SetName("test.txt")
	m.SetSize(-100)

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail File Object should be invalid with negative size")
		t.Log(err)
	}
}

/*
TestNew - Ensure New() creates proper object
*/
func TestNew(t *testing.T) {
	m := New()

	if m.ObjectType != "file" {
		t.Error("Fail: Object type not set correctly")
	}

	if m.SpecVersion == "" {
		t.Error("Fail: Spec version not set")
	}
}

/*
TestHashManagement - Test hash map operations
*/
func TestHashManagement(t *testing.T) {
	m := New()

	m.AddHash("MD5", "test-md5")
	m.AddHash("SHA-256", "test-sha256")

	if len(m.Hashes) != 2 {
		t.Error("Fail: Should have 2 hashes")
	}

	if m.Hashes["MD5"] != "test-md5" {
		t.Error("Fail: MD5 hash not set correctly")
	}

	if m.Hashes["SHA-256"] != "test-sha256" {
		t.Error("Fail: SHA-256 hash not set correctly")
	}
}
