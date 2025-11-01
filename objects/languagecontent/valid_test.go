// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package languagecontent

import (
	"testing"
)

// ----------------------------------------------------------------------
// Tests
// ----------------------------------------------------------------------

/*
TestValid1 - Make sure we get a value of false when LanguageContent obj is blank.
*/
func TestValid1(t *testing.T) {
	m := New()
	want := false

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail LanguageContent Object should be invalid when empty")
		t.Log(err)
	}
}

/*
TestValid2 - Invalid with only object_ref (missing contents)
*/
func TestValid2(t *testing.T) {
	m := New()
	want := false

	m.SetObjectRef("indicator--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail LanguageContent Object should be invalid without contents")
		t.Log(err)
	}
}

/*
TestValid3 - Valid with required properties
*/
func TestValid3(t *testing.T) {
	m := New()
	want := true

	m.SetObjectRef("indicator--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f")
	m.AddContent("es", "name", "Indicador de Malware")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail LanguageContent Object should be valid with required properties")
		t.Log(err)
	}
}

/*
TestValid4 - Valid with multiple languages
*/
func TestValid4(t *testing.T) {
	m := New()
	want := true

	m.SetObjectRef("indicator--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f")
	m.AddContent("es", "name", "Indicador de Malware")
	m.AddContent("es", "description", "Este es un indicador de malware")
	m.AddContent("fr", "name", "Indicateur de Malware")
	m.AddContent("fr", "description", "Ceci est un indicateur de malware")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail LanguageContent Object should be valid with multiple languages")
		t.Log(err)
	}
}

/*
TestValid5 - Valid with object_modified
*/
func TestValid5(t *testing.T) {
	m := New()
	want := true

	m.SetObjectRef("indicator--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f")
	m.SetObjectModified("2016-05-12T08:17:27.000Z")
	m.AddContent("es", "name", "Indicador de Malware")

	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail LanguageContent Object should be valid with object_modified")
		t.Log(err)
	}
}

/*
TestNew - Ensure New() creates proper object
*/
func TestNew(t *testing.T) {
	m := New()

	if m.ObjectType != "language-content" {
		t.Error("Fail: Object type not set correctly")
	}

	if m.SpecVersion == "" {
		t.Error("Fail: Spec version not set")
	}
}

/*
TestContentManagement - Test content operations
*/
func TestContentManagement(t *testing.T) {
	m := New()

	m.AddContent("es", "name", "Nombre en español")
	m.AddContent("es", "description", "Descripción en español")
	m.AddContent("fr", "name", "Nom en français")

	if len(m.Contents) != 2 {
		t.Error("Fail: Should have 2 languages")
	}

	if len(m.Contents["es"]) != 2 {
		t.Error("Fail: Spanish should have 2 selectors")
	}

	if len(m.Contents["fr"]) != 1 {
		t.Error("Fail: French should have 1 selector")
	}

	if m.Contents["es"]["name"] != "Nombre en español" {
		t.Error("Fail: Spanish name not set correctly")
	}
}
