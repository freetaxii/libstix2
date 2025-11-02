// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package tests

import (
	"encoding/json"
	"testing"

	"github.com/freetaxii/libstix2/objects/identity"
)

// TestIdentityObject tests that the provided identity object is properly handled
func TestIdentityObject(t *testing.T) {
	// Test identity object from the user
	testJSON := `{
        "type": "identity",
        "spec_version": "2.1",
        "id": "identity--f5eec3f7-b411-41ff-8493-5bcf59f575d4",
        "created_by_ref": "identity--f5eec3f7-b411-41ff-8493-5bcf59f575d4",
        "created": "2021-08-05T11:36:37.679168Z",
        "modified": "2022-08-02T19:03:33.005815Z",
        "name": "Davis-Gallagher",
        "description": "Profound coherent circuit",
        "identity_class": "organization",
        "sectors": [
            "technology",
            "retail"
        ],
        "labels": [
            "ais-consent-everyone-cisa-proprietary",
            "pfte-batch-1"
        ],
        "confidence": 79,
        "lang": "en",
        "object_marking_refs": [
            "marking-definition--b11561e7-fd83-4ea5-8b46-c79afbaf3e75"
        ]
    }`

	// Decode JSON into identity struct
	var id identity.Identity
	if err := json.Unmarshal([]byte(testJSON), &id); err != nil {
		t.Fatalf("Failed to decode identity: %v", err)
	}

	// Validate the identity
	if valid, _, details := id.Valid(false); !valid {
		t.Errorf("Identity validation failed: %v", details)
	} else {
		t.Logf("Validation passed: %v", details)
	}

	// Check key fields
	if id.GetName() != "Davis-Gallagher" {
		t.Errorf("Expected name 'Davis-Gallagher', got '%s'", id.GetName())
	}

	if id.GetDescription() != "Profound coherent circuit" {
		t.Errorf("Expected description 'Profound coherent circuit', got '%s'", id.GetDescription())
	}

	if len(id.Sectors) != 2 {
		t.Errorf("Expected 2 sectors, got %d", len(id.Sectors))
	} else {
		expectedSectors := []string{"technology", "retail"}
		for i, sector := range expectedSectors {
			if id.Sectors[i] != sector {
				t.Errorf("Expected sector '%s' at index %d, got '%s'", sector, i, id.Sectors[i])
			}
		}
	}

	if id.Confidence != 79 {
		t.Errorf("Expected confidence 79, got %d", id.Confidence)
	}

	if id.Lang != "en" {
		t.Errorf("Expected lang 'en', got '%s'", id.Lang)
	}

	if len(id.Labels) != 2 {
		t.Errorf("Expected 2 labels, got %d", len(id.Labels))
	}

	if len(id.ObjectMarkingRefs) != 1 {
		t.Errorf("Expected 1 object marking ref, got %d", len(id.ObjectMarkingRefs))
	}

	// Test round-trip encoding
	reencodedJSON, err := json.MarshalIndent(id, "", "  ")
	if err != nil {
		t.Fatalf("Failed to re-encode identity: %v", err)
	}

	t.Logf("Re-encoded JSON:\n%s", string(reencodedJSON))

	// Parse both JSON structures for comparison
	var originalObj, reencodedObj map[string]interface{}
	if err := json.Unmarshal([]byte(testJSON), &originalObj); err != nil {
		t.Fatalf("Failed to parse original JSON: %v", err)
	}
	if err := json.Unmarshal(reencodedJSON, &reencodedObj); err != nil {
		t.Fatalf("Failed to parse re-encoded JSON: %v", err)
	}

	// Check key fields are preserved
	keyFields := []string{"type", "spec_version", "id", "name", "description", "confidence", "lang"}
	for _, field := range keyFields {
		if originalObj[field] != reencodedObj[field] {
			t.Errorf("Field '%s' mismatch: original=%v, re-encoded=%v", field, originalObj[field], reencodedObj[field])
		}
	}

	t.Log("✅ Identity object test completed successfully")
}

// TestInvalidIdentityObject tests that invalid identity objects are properly rejected
func TestInvalidIdentityObject(t *testing.T) {
	// Test identity with invalid identity_class
	invalidJSON := `{
        "type": "identity",
        "spec_version": "2.1",
        "id": "identity--f5eec3f7-b411-41ff-8493-5bcf59f575d4",
        "created": "2021-08-05T11:36:37.679168Z",
        "modified": "2022-08-02T19:03:33.005815Z",
        "name": "Test Identity",
        "identity_class": "invalid-class"
    }`

	var id identity.Identity
	if err := json.Unmarshal([]byte(invalidJSON), &id); err != nil {
		t.Fatalf("Failed to decode identity: %v", err)
	}

	// Validate the identity - should fail
	if valid, _, details := id.Valid(false); valid {
		t.Error("Identity with invalid identity_class should not be valid")
	} else {
		t.Logf("✅ Correctly rejected invalid identity_class: %v", details)
	}

	// Test identity with invalid sector
	invalidSectorJSON := `{
        "type": "identity",
        "spec_version": "2.1",
        "id": "identity--f5eec3f7-b411-41ff-8493-5bcf59f575d4",
        "created": "2021-08-05T11:36:37.679168Z",
        "modified": "2022-08-02T19:03:33.005815Z",
        "name": "Test Identity",
        "identity_class": "organization",
        "sectors": ["invalid-sector"]
    }`

	var id2 identity.Identity
	if err := json.Unmarshal([]byte(invalidSectorJSON), &id2); err != nil {
		t.Fatalf("Failed to decode identity: %v", err)
	}

	// Validate the identity - should fail
	if valid, _, details := id2.Valid(false); valid {
		t.Error("Identity with invalid sector should not be valid")
	} else {
		t.Logf("✅ Correctly rejected invalid sector: %v", details)
	}

	t.Log("✅ Invalid identity object test completed successfully")
}
