// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package mutex

import (
	"testing"
)

// ----------------------------------------------------------------------
// Tests
// ----------------------------------------------------------------------

/*
TestValid1 - Make sure we get a value of false when Mutex obj is blank.
*/
func TestValid1(t *testing.T) {
	m := New()
	want := false

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail Mutex Object should be invalid when empty (requires name)")
		t.Log(err)
	}
}

/*
TestValid2 - Valid with name
*/
func TestValid2(t *testing.T) {
	m := New()
	want := true

	m.SetName("Global\\MyMutex")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail Mutex Object should be valid with name")
		t.Log(err)
	}
}

/*
TestValid3 - Valid with typical Windows mutex name
*/
func TestValid3(t *testing.T) {
	m := New()
	want := true

	m.SetName("__SYSTEM__")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail Mutex Object should be valid with Windows mutex name")
		t.Log(err)
	}
}

/*
TestNew - Ensure New() creates proper object
*/
func TestNew(t *testing.T) {
	m := New()

	if m.ObjectType != "mutex" {
		t.Error("Fail: Object type not set correctly")
	}

	if m.SpecVersion == "" {
		t.Error("Fail: Spec version not set")
	}
}
