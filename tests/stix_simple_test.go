// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package tests

import (
	"encoding/json"
	"testing"

	"github.com/freetaxii/libstix2/objects/indicator"
)

// TestSimpleSTIXRoundTrip tests a simple STIX object for round-trip compatibility
func TestSimpleSTIXRoundTrip(t *testing.T) {
	// Simple indicator JSON
	originalJSON := `{
		"type": "indicator",
		"spec_version": "2.1",
		"id": "indicator--44af4c39-c09b-49d5-b9a7-3770316149f7",
		"created": "2017-01-01T00:00:00.000Z",
		"modified": "2017-01-01T00:00:00.000Z",
		"indicator_types": ["malicious-activity"],
		"pattern": "[file:hashes.MD5 = 'd41d8cd98f00b204e9800998ecf8427e']",
		"pattern_type": "stix",
		"valid_from": "2017-01-01T00:00:00Z"
	}`

	// Decode JSON into indicator struct
	var ind indicator.Indicator
	if err := json.Unmarshal([]byte(originalJSON), &ind); err != nil {
		t.Fatalf("Failed to decode indicator: %v", err)
	}

	// Validate the indicator
	if valid, _, details := ind.Valid(false); !valid {
		t.Errorf("Indicator validation failed: %v", details)
	}

	// Re-encode back to JSON
	reencodedJSON, err := json.MarshalIndent(ind, "", "  ")
	if err != nil {
		t.Fatalf("Failed to re-encode indicator: %v", err)
	}

	t.Logf("Original JSON:\n%s", originalJSON)
	t.Logf("Re-encoded JSON:\n%s", string(reencodedJSON))

	// Decode both JSON structures for comparison
	var originalObj, reencodedObj map[string]interface{}
	if err := json.Unmarshal([]byte(originalJSON), &originalObj); err != nil {
		t.Fatalf("Failed to parse original JSON: %v", err)
	}
	if err := json.Unmarshal(reencodedJSON, &reencodedObj); err != nil {
		t.Fatalf("Failed to parse re-encoded JSON: %v", err)
	}

	// Check key fields are preserved
	keyFields := []string{"type", "spec_version", "pattern", "pattern_type", "valid_from"}
	for _, field := range keyFields {
		if originalObj[field] != reencodedObj[field] {
			t.Errorf("Field '%s' mismatch: original=%v, re-encoded=%v", field, originalObj[field], reencodedObj[field])
		}
	}

	// Special handling for indicator_types (slice)
	if originalIndTypes, ok := originalObj["indicator_types"].([]interface{}); ok {
		if reencodedIndTypes, ok := reencodedObj["indicator_types"].([]interface{}); ok {
			if len(originalIndTypes) != len(reencodedIndTypes) {
				t.Errorf("indicator_types length mismatch: original=%d, re-encoded=%d", len(originalIndTypes), len(reencodedIndTypes))
			} else if len(originalIndTypes) > 0 && originalIndTypes[0] != reencodedIndTypes[0] {
				t.Errorf("indicator_types value mismatch: original=%v, re-encoded=%v", originalIndTypes[0], reencodedIndTypes[0])
			}
		} else {
			t.Error("indicator_types not found or not a slice in re-encoded object")
		}
	} else {
		t.Error("indicator_types not found or not a slice in original object")
	}

	t.Log("Simple round-trip test completed successfully")
}