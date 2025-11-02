// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package tests

import (
	"encoding/json"
	"testing"

	"github.com/freetaxii/libstix2/objects/markingdefinition"
)

// TestMarkingDefinitionWithExtensions tests that marking-definitions with extensions are properly handled
func TestMarkingDefinitionWithExtensions(t *testing.T) {
	// Test marking definition with extensions from the user
	testJSON := `{
        "type": "marking-definition",
        "spec_version": "2.1",
        "id": "marking-definition--75ad3822-3fbe-4de8-ac94-f02f01bf778b",
        "created": "2021-08-14T01:08:09.163015Z",
        "modified": "2021-08-14T01:08:09.163015Z",
        "extensions": {
            "extension-definition--3a65884d-005a-4290-8335-cb2d778a83ce": {
                "extension_type": "property-extension",
                "identifier": "isa:guide.19001.ACS3-e0c28f68-5570-434a-9764-3e8498f0a59a",
                "name": "Serve much idea usually",
                "create_date_time": "2021-08-14T01:08:08.432065Z",
                "responsible_entity_custodian": "USA.DHS.US-CERT",
                "responsible_entity_originator": "USA.DHS.US-CERT",
                "authority_reference": [
                    "urn:isa:authority:misa"
                ],
                "policy_reference": "urn:isa:policy:acs:ns:v3.0?privdefault=deny&sharedefault=deny",
                "access_privilege": [
                    {
                        "privilege_action": "CISAUSES",
                        "privilege_scope": {
                            "permitted_nationalities": [
                                "ALL"
                            ],
                            "permitted_organizations": [
                                "ALL"
                            ],
                            "shareability": [
                                "ALL"
                            ],
                            "entity": [
                                "ALL"
                            ]
                        },
                        "rule_effect": "permit"
                    }
                ],
                "further_sharing": [
                    {
                        "sharing_scope": [
                            "USA.USG"
                        ],
                        "rule_effect": "permit"
                    }
                ],
                "control_set": {
                    "classification": "U",
                    "formal_determination": [
                        "FOUO",
                        "AIS"
                    ]
                }
            }
        }
    }`

	// Decode JSON into marking definition struct
	var md markingdefinition.MarkingDefinition
	if err := json.Unmarshal([]byte(testJSON), &md); err != nil {
		t.Fatalf("Failed to decode marking definition: %v", err)
	}

	// Validate the marking definition
	if valid, _, details := md.Valid(false); !valid {
		t.Errorf("Marking definition validation failed: %v", details)
	} else {
		t.Logf("Validation passed: %v", details)
	}

	// Check that extensions are properly captured
	if len(md.Extensions) == 0 {
		t.Error("Extensions were not captured in the marking definition")
	} else {
		t.Logf("Found %d extensions", len(md.Extensions))

		// Check for the specific extension
		extID := "extension-definition--3a65884d-005a-4290-8335-cb2d778a83ce"
		if ext, ok := md.Extensions[extID]; ok {
			t.Logf("Found extension: %s", extID)
			if extMap, ok := ext.(map[string]interface{}); ok {
				if name, ok := extMap["name"].(string); ok {
					t.Logf("Extension name: %s", name)
					if name != "Serve much idea usually" {
						t.Errorf("Expected extension name 'Serve much idea usually', got '%s'", name)
					}
				} else {
					t.Error("Extension name not found or not a string")
				}
			} else {
				t.Error("Extension is not a map[string]interface{}")
			}
		} else {
			t.Errorf("Expected extension %s not found", extID)
		}
	}

	// Test round-trip encoding
	reencodedJSON, err := json.MarshalIndent(md, "", "  ")
	if err != nil {
		t.Fatalf("Failed to re-encode marking definition: %v", err)
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
	keyFields := []string{"type", "spec_version", "id", "created"}
	for _, field := range keyFields {
		if originalObj[field] != reencodedObj[field] {
			t.Errorf("Field '%s' mismatch: original=%v, re-encoded=%v", field, originalObj[field], reencodedObj[field])
		}
	}

	// Check that extensions are preserved
	if originalObj["extensions"] == nil && reencodedObj["extensions"] != nil {
		t.Error("Extensions appeared in re-encoded JSON but not in original")
	} else if originalObj["extensions"] != nil && reencodedObj["extensions"] == nil {
		t.Error("Extensions disappeared in re-encoded JSON")
	}

	t.Log("Marking definition with extensions test completed successfully")
}

// TestTraditionalMarkingDefinition tests that traditional marking definitions still work
func TestTraditionalMarkingDefinition(t *testing.T) {
	// Traditional TLP marking definition
	testJSON := `{
        "type": "marking-definition",
        "spec_version": "2.1",
        "id": "marking-definition--613f2e26-407d-48c7-9eca-b8e91df99dc9",
        "created": "2017-01-20T00:00:00.000Z",
        "modified": "2017-01-20T00:00:00.000Z",
        "definition_type": "tlp",
        "name": "TLP:WHITE",
        "definition": {
            "tlp": "white"
        }
    }`

	// Decode JSON into marking definition struct
	var md markingdefinition.MarkingDefinition
	if err := json.Unmarshal([]byte(testJSON), &md); err != nil {
		t.Fatalf("Failed to decode traditional marking definition: %v", err)
	}

	// Validate the marking definition
	if valid, _, details := md.Valid(false); !valid {
		t.Errorf("Traditional marking definition validation failed: %v", details)
	} else {
		t.Logf("Traditional validation passed: %v", details)
	}

	// Check that no extensions are present
	if len(md.Extensions) > 0 {
		t.Errorf("Traditional marking definition should not have extensions, found %d", len(md.Extensions))
	}

	t.Log("Traditional marking definition test completed successfully")
}
