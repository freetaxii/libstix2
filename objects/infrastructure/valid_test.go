// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package infrastructure

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
TestValid1 - Make sure we get a value of false when Infrastructure obj is blank.
*/
func TestValid1(t *testing.T) {
	m := New()
	want := false

	if got, err := m.Valid(); got != want {
		t.Error("Fail Infrastructure Object should be invalid when empty")
		t.Log(err)
	}
}

//TestValid2 -
func TestValid2(t *testing.T) {
	m := New()
	want := true

	m.AddInfrastructureTypes("botnet")

	if got, err := m.Valid(); got != want {
		t.Error("Fail Infrastructure Object should be valid when require fields not empty")
		t.Log(err)
	}
}

//TestValid5 - InfrastructureTypes should add only of vocab value
func TestValid5(t *testing.T) {
	m := New()
	want := false
	wantMessage := "the InfrastructureTypes property should be one of list: amplification, anonymization, botnet, command-and-control, exfiltration, hosting-malware, hosting-target-lists, phishing, reconnaissance, staging, undefined"

	m.AddInfrastructureTypes("asdasdasd")

	if got, err := m.Valid(); got != want {
		t.Error("Fail Infrastructure Object InfrastructureTypes added value not from vocab")
		t.Log(err)
	}

	if _, err := m.Valid(); err.Error() != wantMessage {
		t.Error("Fail Infrastructure Object InfrastructureTypes added value not from vocab. Wrong error message")
		t.Log(err)
	}
}

//TestValid9 - should be valid
func TestValid9(t *testing.T) {
	m := New()
	want := true

	m.AddInfrastructureTypes("botnet")

	m.AddAliases("botnet2")

	if got, err := m.Valid(); got != want {
		t.Error("Fail Infrastructure Object shoulf be valid")
		t.Log(err)
	}
}
