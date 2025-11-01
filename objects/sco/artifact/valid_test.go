// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package artifact

import (
	"testing"
)

// ----------------------------------------------------------------------
// Tests
// ----------------------------------------------------------------------

/*
TestValid1 - Make sure we get a value of false when Artifact obj is blank (missing required fields).
*/
func TestValid1(t *testing.T) {
	m := New()
	want := false

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail Artifact Object should be invalid when empty (requires payload_bin or url)")
		t.Log(err)
	}
}

/*
TestValid2 - Valid with payload_bin
*/
func TestValid2(t *testing.T) {
	m := New()
	want := true

	m.SetPayloadBin("VGhpcyBpcyBhIHRlc3Q=")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail Artifact Object should be valid with payload_bin")
		t.Log(err)
	}
}

/*
TestValid3 - Valid with url
*/
func TestValid3(t *testing.T) {
	m := New()
	want := true

	m.SetURL("https://example.com/malware.exe")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail Artifact Object should be valid with url")
		t.Log(err)
	}
}

/*
TestValid4 - Valid with both payload_bin and url
*/
func TestValid4(t *testing.T) {
	m := New()
	want := true

	m.SetPayloadBin("VGhpcyBpcyBhIHRlc3Q=")
	m.SetURL("https://example.com/malware.exe")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail Artifact Object should be valid with both payload_bin and url")
		t.Log(err)
	}
}

/*
TestValid5 - Valid with all properties
*/
func TestValid5(t *testing.T) {
	m := New()
	want := true

	m.SetPayloadBin("VGhpcyBpcyBhIHRlc3Q=")
	m.SetMimeType("application/zip")
	m.AddHash("SHA-256", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
	m.SetEncryptionAlgorithm("AES-256-GCM")
	m.SetDecryptionKey("MySecretKey123")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail Artifact Object should be valid with all properties")
		t.Log(err)
	}
}

/*
TestNew - Ensure New() creates proper object
*/
func TestNew(t *testing.T) {
	m := New()

	if m.ObjectType != "artifact" {
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
}
