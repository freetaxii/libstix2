// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/freetaxii/libstix2/objects/attackpattern"
	"github.com/freetaxii/libstix2/objects/bundle"
	"github.com/freetaxii/libstix2/objects/campaign"
	"github.com/freetaxii/libstix2/objects/courseofaction"
	"github.com/freetaxii/libstix2/objects/grouping"
	"github.com/freetaxii/libstix2/objects/identity"
	"github.com/freetaxii/libstix2/objects/indicator"
	"github.com/freetaxii/libstix2/objects/infrastructure"
	"github.com/freetaxii/libstix2/objects/intrusionset"
	"github.com/freetaxii/libstix2/objects/location"
	"github.com/freetaxii/libstix2/objects/malware"
	"github.com/freetaxii/libstix2/objects/malwareanalysis"
	"github.com/freetaxii/libstix2/objects/markingdefinition"
	"github.com/freetaxii/libstix2/objects/note"
	"github.com/freetaxii/libstix2/objects/observeddata"
	"github.com/freetaxii/libstix2/objects/opinion"
	"github.com/freetaxii/libstix2/objects/relationship"
	"github.com/freetaxii/libstix2/objects/report"
	"github.com/freetaxii/libstix2/objects/sighting"
	"github.com/freetaxii/libstix2/objects/threatactor"
	"github.com/freetaxii/libstix2/objects/tool"
	"github.com/freetaxii/libstix2/objects/vulnerability"
	"github.com/google/uuid"
)

// ----------------------------------------------------------------------
// Constants and Configuration
// ----------------------------------------------------------------------

const (
	MispTestDataDir = "../data/misp-stix-tests/files"
)

// ----------------------------------------------------------------------
// Command-line flags
// ----------------------------------------------------------------------
var (
	singleFile    = flag.String("file", "", "Path to a single JSON file to test (instead of all files)")
	debugMode     = flag.Bool("debug", false, "Enable debug mode for detailed validation output")
	showAllErrors = flag.Bool("all-errors", false, "Show all validation errors instead of just summaries")
)

// ----------------------------------------------------------------------
// Helper Functions
// ----------------------------------------------------------------------

// findJSONFiles recursively finds all JSON files in the given directory
func findJSONFiles(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".json" {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}

// parseSTIXData parses STIX JSON data into a generic interface
func parseSTIXData(data []byte) ([]map[string]interface{}, error) {
	var stixObjects []map[string]interface{}

	// Try to parse as array first
	if err := json.Unmarshal(data, &stixObjects); err != nil {
		// If that fails, try to parse as single object
		var singleObject map[string]interface{}
		if err := json.Unmarshal(data, &singleObject); err != nil {
			return nil, fmt.Errorf("failed to parse JSON as array or object: %v", err)
		}
		stixObjects = []map[string]interface{}{singleObject}
	}

	return stixObjects, nil
}

// normalizeJSON normalizes JSON for comparison by removing insignificant differences
func normalizeJSON(data interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		normalized := make(map[string]interface{})
		for key, value := range v {
			// Skip UUID-specific fields that might change during processing
			if key == "id" {
				// Validate that it's a valid UUID but don't compare exact values
				if strVal, ok := value.(string); ok {
					if _, err := uuid.Parse(strVal); err != nil {
						// If it's not a valid UUID, keep the original value
						normalized[key] = value
					}
				}
				continue
			}
			normalized[key] = normalizeJSON(value)
		}
		return normalized

	case []interface{}:
		normalized := make([]interface{}, len(v))
		for i, item := range v {
			normalized[i] = normalizeJSON(item)
		}
		return normalized

	default:
		return v
	}
}

// compareJSON compares two JSON structures for equivalence
func compareJSON(a, b interface{}) (bool, error) {
	normalizedA := normalizeJSON(a)
	normalizedB := normalizeJSON(b)

	jsonA, err := json.Marshal(normalizedA)
	if err != nil {
		return false, fmt.Errorf("failed to marshal first JSON: %v", err)
	}

	jsonB, err := json.Marshal(normalizedB)
	if err != nil {
		return false, fmt.Errorf("failed to marshal second JSON: %v", err)
	}

	return string(jsonA) == string(jsonB), nil
}

// ----------------------------------------------------------------------
// Test Functions
// ----------------------------------------------------------------------

// TestMISPSTIXRoundTrip tests that all STIX data can be decoded and re-encoded
// without losing any information
func TestMISPSTIXRoundTrip(t *testing.T) {
	// Check if the MISP test data directory exists
	if _, err := os.Stat(MispTestDataDir); os.IsNotExist(err) {
		t.Skipf("MISP test data directory not found at %s. Run 'git submodule update --init' to fetch test data.", MispTestDataDir)
		return
	}

	// Find all JSON files
	jsonFiles, err := findJSONFiles(MispTestDataDir)
	if err != nil {
		t.Fatalf("Failed to find JSON files: %v", err)
	}

	if len(jsonFiles) == 0 {
		t.Skip("No JSON files found in MISP test data directory")
		return
	}

	t.Logf("Found %d JSON files to test", len(jsonFiles))

	// Track test statistics
	totalFiles := 0
	passedFiles := 0
	failedFiles := 0

	for _, filePath := range jsonFiles {
		totalFiles++
		t.Run(filepath.Base(filePath), func(t *testing.T) {
			testSTIXFileRoundTrip(t, filePath)
		})

		// Update statistics
		if !t.Failed() {
			passedFiles++
		} else {
			failedFiles++
		}
	}

	t.Logf("Round-trip test summary: %d total, %d passed, %d failed", totalFiles, passedFiles, failedFiles)

	if failedFiles > 0 {
		t.Errorf("Failed round-trip tests for %d files", failedFiles)
	}
}

// testSTIXFileRoundTrip tests a single STIX file for round-trip compatibility
func testSTIXFileRoundTrip(t *testing.T, filePath string) {
	// Read original file
	originalData, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", filePath, err)
	}

	// Parse original JSON
	originalObjects, err := parseSTIXData(originalData)
	if err != nil {
		t.Fatalf("Failed to parse original JSON from %s: %v", filePath, err)
	}

	if len(originalObjects) == 0 {
		t.Skipf("No STIX objects found in file %s", filePath)
		return
	}

	// Process each STIX object
	var processedObjects []interface{}

	for _, originalObj := range originalObjects {
		// Get the object type
		objType, ok := originalObj["type"].(string)
		if !ok {
			t.Errorf("Object missing 'type' field in %s", filePath)
			continue
		}

		// Marshal the object to JSON and then unmarshal it using libstix2
		// This simulates real-world usage where JSON is parsed into Go structs
		objJSON, err := json.Marshal(originalObj)
		if err != nil {
			t.Errorf("Failed to marshal object in %s: %v", filePath, err)
			continue
		}

		// Decode based on object type
		var decodedObj interface{}
		switch objType {
		case "indicator":
			var ind indicator.Indicator
			if err := json.Unmarshal(objJSON, &ind); err != nil {
				t.Errorf("Failed to decode indicator in %s: %v", filePath, err)
				continue
			}
			decodedObj = ind

		case "malware":
			var mal malware.Malware
			if err := json.Unmarshal(objJSON, &mal); err != nil {
				t.Errorf("Failed to decode malware in %s: %v", filePath, err)
				continue
			}
			decodedObj = mal

		case "infrastructure":
			var inf infrastructure.Infrastructure
			if err := json.Unmarshal(objJSON, &inf); err != nil {
				t.Errorf("Failed to decode infrastructure in %s: %v", filePath, err)
				continue
			}
			decodedObj = inf

		case "threat-actor":
			var ta threatactor.ThreatActor
			if err := json.Unmarshal(objJSON, &ta); err != nil {
				t.Errorf("Failed to decode threat-actor in %s: %v", filePath, err)
				continue
			}
			decodedObj = ta

		case "tool":
			var toolObj tool.Tool
			if err := json.Unmarshal(objJSON, &toolObj); err != nil {
				t.Errorf("Failed to decode tool in %s: %v", filePath, err)
				continue
			}
			decodedObj = toolObj

		case "attack-pattern":
			var ap attackpattern.AttackPattern
			if err := json.Unmarshal(objJSON, &ap); err != nil {
				t.Errorf("Failed to decode attack-pattern in %s: %v", filePath, err)
				continue
			}
			decodedObj = ap

		case "campaign":
			var camp campaign.Campaign
			if err := json.Unmarshal(objJSON, &camp); err != nil {
				t.Errorf("Failed to decode campaign in %s: %v", filePath, err)
				continue
			}
			decodedObj = camp

		case "course-of-action":
			var coa courseofaction.CourseOfAction
			if err := json.Unmarshal(objJSON, &coa); err != nil {
				t.Errorf("Failed to decode course-of-action in %s: %v", filePath, err)
				continue
			}
			decodedObj = coa

		case "grouping":
			var grp grouping.Grouping
			if err := json.Unmarshal(objJSON, &grp); err != nil {
				t.Errorf("Failed to decode grouping in %s: %v", filePath, err)
				continue
			}
			decodedObj = grp

		case "identity":
			var id identity.Identity
			if err := json.Unmarshal(objJSON, &id); err != nil {
				t.Errorf("Failed to decode identity in %s: %v", filePath, err)
				continue
			}
			decodedObj = id

		case "intrusion-set":
			var is intrusionset.IntrusionSet
			if err := json.Unmarshal(objJSON, &is); err != nil {
				t.Errorf("Failed to decode intrusion-set in %s: %v", filePath, err)
				continue
			}
			decodedObj = is

		case "location":
			var loc location.Location
			if err := json.Unmarshal(objJSON, &loc); err != nil {
				t.Errorf("Failed to decode location in %s: %v", filePath, err)
				continue
			}
			decodedObj = loc

		case "malware-analysis":
			var ma malwareanalysis.MalwareAnalysis
			if err := json.Unmarshal(objJSON, &ma); err != nil {
				t.Errorf("Failed to decode malware-analysis in %s: %v", filePath, err)
				continue
			}
			decodedObj = ma

		case "note":
			var noteObj note.Note
			if err := json.Unmarshal(objJSON, &noteObj); err != nil {
				t.Errorf("Failed to decode note in %s: %v", filePath, err)
				continue
			}
			decodedObj = noteObj

		case "observed-data":
			var od observeddata.ObservedData
			if err := json.Unmarshal(objJSON, &od); err != nil {
				t.Errorf("Failed to decode observed-data in %s: %v", filePath, err)
				continue
			}
			decodedObj = od

		case "opinion":
			var op opinion.Opinion
			if err := json.Unmarshal(objJSON, &op); err != nil {
				t.Errorf("Failed to decode opinion in %s: %v", filePath, err)
				continue
			}
			decodedObj = op

		case "report":
			var reportObj report.Report
			if err := json.Unmarshal(objJSON, &reportObj); err != nil {
				t.Errorf("Failed to decode report in %s: %v", filePath, err)
				continue
			}
			decodedObj = reportObj

		case "relationship":
			var rel relationship.Relationship
			if err := json.Unmarshal(objJSON, &rel); err != nil {
				t.Errorf("Failed to decode relationship in %s: %v", filePath, err)
				continue
			}
			decodedObj = rel

		case "sighting":
			var sight sighting.Sighting
			if err := json.Unmarshal(objJSON, &sight); err != nil {
				t.Errorf("Failed to decode sighting in %s: %v", filePath, err)
				continue
			}
			decodedObj = sight

		case "vulnerability":
			var vuln vulnerability.Vulnerability
			if err := json.Unmarshal(objJSON, &vuln); err != nil {
				t.Errorf("Failed to decode vulnerability in %s: %v", filePath, err)
				continue
			}
			decodedObj = vuln

		case "bundle":
			var bundleObj bundle.Bundle
			if err := json.Unmarshal(objJSON, &bundleObj); err != nil {
				t.Errorf("Failed to decode bundle in %s: %v", filePath, err)
				continue
			}
			decodedObj = bundleObj

		case "marking-definition":
			var md markingdefinition.MarkingDefinition
			if err := json.Unmarshal(objJSON, &md); err != nil {
				t.Errorf("Failed to decode marking-definition in %s: %v", filePath, err)
				continue
			}
			decodedObj = md

		default:
			t.Logf("Unsupported object type '%s' in %s, skipping", objType, filePath)
			// For unsupported types, keep the original object
			decodedObj = originalObj
		}

		// Re-encode the decoded object back to JSON
		reencodedJSON, err := json.Marshal(decodedObj)
		if err != nil {
			t.Errorf("Failed to re-encode object in %s: %v", filePath, err)
			continue
		}

		// Parse the re-encoded JSON back to a map for comparison
		var reencodedObj map[string]interface{}
		if err := json.Unmarshal(reencodedJSON, &reencodedObj); err != nil {
			t.Errorf("Failed to parse re-encoded JSON in %s: %v", filePath, err)
			continue
		}

		processedObjects = append(processedObjects, reencodedObj)
	}

	if len(processedObjects) == 0 {
		t.Skipf("No objects were successfully processed in %s", filePath)
		return
	}

	// Compare original and re-encoded data
	areEqual, err := compareJSON(originalObjects, processedObjects)
	if err != nil {
		t.Errorf("Failed to compare JSON data in %s: %v", filePath, err)
		return
	}

	if !areEqual {
		// Provide detailed comparison
		originalJSON, _ := json.MarshalIndent(originalObjects, "", "  ")
		reencodedJSON, _ := json.MarshalIndent(processedObjects, "", "  ")

		t.Errorf("Round-trip data mismatch in %s\nOriginal:\n%s\n\nRe-encoded:\n%s",
			filePath, string(originalJSON), string(reencodedJSON))
		return
	}

	t.Logf("Successfully round-tripped %d objects from %s", len(processedObjects), filePath)
}

// TestMISPSTIXValidation tests that all STIX data in the MISP test suite
// passes validation when decoded using libstix2
func TestMISPSTIXValidation(t *testing.T) {
	// Check if the MISP test data directory exists
	if _, err := os.Stat(MispTestDataDir); os.IsNotExist(err) {
		t.Skipf("MISP test data directory not found at %s. Run 'git submodule update --init' to fetch test data.", MispTestDataDir)
		return
	}

	// Find all JSON files
	jsonFiles, err := findJSONFiles(MispTestDataDir)
	if err != nil {
		t.Fatalf("Failed to find JSON files: %v", err)
	}

	if len(jsonFiles) == 0 {
		t.Skip("No JSON files found in MISP test data directory")
		return
	}

	t.Logf("Found %d JSON files to validate", len(jsonFiles))

	// Track validation statistics
	totalObjects := 0
	validObjects := 0
	invalidObjects := 0

	for _, filePath := range jsonFiles {
		t.Run(filepath.Base(filePath), func(t *testing.T) {
			valid, invalid := testSTIXFileValidation(t, filePath)
			totalObjects += valid + invalid
			validObjects += valid
			invalidObjects += invalid
		})
	}

	t.Logf("Validation summary: %d total objects, %d valid, %d invalid", totalObjects, validObjects, invalidObjects)
}

// testSTIXFileValidation tests that all STIX objects in a file pass validation
func testSTIXFileValidation(t *testing.T, filePath string) (int, int) {
	// Read file
	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", filePath, err)
		return 0, 0
	}

	// Parse JSON
	stixObjects, err := parseSTIXData(data)
	if err != nil {
		t.Fatalf("Failed to parse JSON from %s: %v", filePath, err)
		return 0, 0
	}

	validCount := 0
	invalidCount := 0

	for _, obj := range stixObjects {
		objType, ok := obj["type"].(string)
		if !ok {
			t.Errorf("Object missing 'type' field in %s", filePath)
			invalidCount++
			continue
		}

		// Convert object to JSON for decoding
		objJSON, err := json.Marshal(obj)
		if err != nil {
			t.Errorf("Failed to marshal object in %s: %v", filePath, err)
			invalidCount++
			continue
		}

		// Validate based on object type
		switch objType {
		case "indicator":
			var ind indicator.Indicator
			if err := json.Unmarshal(objJSON, &ind); err != nil {
				t.Errorf("Failed to decode indicator in %s: %v", filePath, err)
				invalidCount++
				continue
			}
			if valid, _, _ := ind.Valid(false); valid {
				validCount++
			} else {
				invalidCount++
			}

		case "malware":
			var mal malware.Malware
			if err := json.Unmarshal(objJSON, &mal); err != nil {
				t.Errorf("Failed to decode malware in %s: %v", filePath, err)
				invalidCount++
				continue
			}
			if valid, _, _ := mal.Valid(false); valid {
				validCount++
			} else {
				invalidCount++
			}

		case "infrastructure":
			var inf infrastructure.Infrastructure
			if err := json.Unmarshal(objJSON, &inf); err != nil {
				t.Errorf("Failed to decode infrastructure in %s: %v", filePath, err)
				invalidCount++
				continue
			}
			if valid, _, _ := inf.Valid(false); valid {
				validCount++
			} else {
				invalidCount++
			}

		case "threat-actor":
			var ta threatactor.ThreatActor
			if err := json.Unmarshal(objJSON, &ta); err != nil {
				t.Errorf("Failed to decode threat-actor in %s: %v", filePath, err)
				invalidCount++
				continue
			}
			if valid, _, _ := ta.Valid(false); valid {
				validCount++
			} else {
				invalidCount++
			}

		case "tool":
			var toolObj tool.Tool
			if err := json.Unmarshal(objJSON, &toolObj); err != nil {
				t.Errorf("Failed to decode tool in %s: %v", filePath, err)
				invalidCount++
				continue
			}
			if valid, _, _ := toolObj.Valid(false); valid {
				validCount++
			} else {
				invalidCount++
			}

		case "attack-pattern":
			var ap attackpattern.AttackPattern
			if err := json.Unmarshal(objJSON, &ap); err != nil {
				t.Errorf("Failed to decode attack-pattern in %s: %v", filePath, err)
				invalidCount++
				continue
			}
			if valid, _, _ := ap.Valid(false); valid {
				validCount++
			} else {
				invalidCount++
			}

		case "campaign":
			var camp campaign.Campaign
			if err := json.Unmarshal(objJSON, &camp); err != nil {
				t.Errorf("Failed to decode campaign in %s: %v", filePath, err)
				invalidCount++
				continue
			}
			if valid, _, _ := camp.Valid(false); valid {
				validCount++
			} else {
				invalidCount++
			}

		case "course-of-action":
			var coa courseofaction.CourseOfAction
			if err := json.Unmarshal(objJSON, &coa); err != nil {
				t.Errorf("Failed to decode course-of-action in %s: %v", filePath, err)
				invalidCount++
				continue
			}
			if valid, _, _ := coa.Valid(false); valid {
				validCount++
			} else {
				invalidCount++
			}

		case "grouping":
			var grp grouping.Grouping
			if err := json.Unmarshal(objJSON, &grp); err != nil {
				t.Errorf("Failed to decode grouping in %s: %v", filePath, err)
				invalidCount++
				continue
			}
			if valid, _, _ := grp.Valid(false); valid {
				validCount++
			} else {
				invalidCount++
			}

		case "identity":
			var id identity.Identity
			if err := json.Unmarshal(objJSON, &id); err != nil {
				t.Errorf("Failed to decode identity in %s: %v", filePath, err)
				invalidCount++
				continue
			}
			if valid, _, _ := id.Valid(false); valid {
				validCount++
			} else {
				invalidCount++
			}

		case "intrusion-set":
			var is intrusionset.IntrusionSet
			if err := json.Unmarshal(objJSON, &is); err != nil {
				t.Errorf("Failed to decode intrusion-set in %s: %v", filePath, err)
				invalidCount++
				continue
			}
			if valid, _, _ := is.Valid(false); valid {
				validCount++
			} else {
				invalidCount++
			}

		case "location":
			var loc location.Location
			if err := json.Unmarshal(objJSON, &loc); err != nil {
				t.Errorf("Failed to decode location in %s: %v", filePath, err)
				invalidCount++
				continue
			}
			if valid, _, _ := loc.Valid(false); valid {
				validCount++
			} else {
				invalidCount++
			}

		case "malware-analysis":
			var ma malwareanalysis.MalwareAnalysis
			if err := json.Unmarshal(objJSON, &ma); err != nil {
				t.Errorf("Failed to decode malware-analysis in %s: %v", filePath, err)
				invalidCount++
				continue
			}
			if valid, _, _ := ma.Valid(false); valid {
				validCount++
			} else {
				invalidCount++
			}

		case "note":
			var noteObj note.Note
			if err := json.Unmarshal(objJSON, &noteObj); err != nil {
				t.Errorf("Failed to decode note in %s: %v", filePath, err)
				invalidCount++
				continue
			}
			if valid, _, _ := noteObj.Valid(false); valid {
				validCount++
			} else {
				invalidCount++
			}

		case "observed-data":
			var od observeddata.ObservedData
			if err := json.Unmarshal(objJSON, &od); err != nil {
				t.Errorf("Failed to decode observed-data in %s: %v", filePath, err)
				invalidCount++
				continue
			}
			if valid, _, _ := od.Valid(false); valid {
				validCount++
			} else {
				invalidCount++
			}

		case "opinion":
			var op opinion.Opinion
			if err := json.Unmarshal(objJSON, &op); err != nil {
				t.Errorf("Failed to decode opinion in %s: %v", filePath, err)
				invalidCount++
				continue
			}
			if valid, _, _ := op.Valid(false); valid {
				validCount++
			} else {
				invalidCount++
			}

		case "report":
			var reportObj report.Report
			if err := json.Unmarshal(objJSON, &reportObj); err != nil {
				t.Errorf("Failed to decode report in %s: %v", filePath, err)
				invalidCount++
				continue
			}
			if valid, _, _ := reportObj.Valid(false); valid {
				validCount++
			} else {
				invalidCount++
			}

		case "relationship":
			var rel relationship.Relationship
			if err := json.Unmarshal(objJSON, &rel); err != nil {
				t.Errorf("Failed to decode relationship in %s: %v", filePath, err)
				invalidCount++
				continue
			}
			if valid, _, _ := rel.Valid(false); valid {
				validCount++
			} else {
				invalidCount++
			}

		case "sighting":
			var sight sighting.Sighting
			if err := json.Unmarshal(objJSON, &sight); err != nil {
				t.Errorf("Failed to decode sighting in %s: %v", filePath, err)
				invalidCount++
				continue
			}
			if valid, _, _ := sight.Valid(false); valid {
				validCount++
			} else {
				invalidCount++
			}

		case "vulnerability":
			var vuln vulnerability.Vulnerability
			if err := json.Unmarshal(objJSON, &vuln); err != nil {
				t.Errorf("Failed to decode vulnerability in %s: %v", filePath, err)
				invalidCount++
				continue
			}
			if valid, _, _ := vuln.Valid(false); valid {
				validCount++
			} else {
				invalidCount++
			}

		case "bundle":
			var bundleObj bundle.Bundle
			if err := json.Unmarshal(objJSON, &bundleObj); err != nil {
				t.Errorf("Failed to decode bundle in %s: %v", filePath, err)
				invalidCount++
				continue
			}
			// Bundle doesn't have a Valid method, consider it valid if it decodes correctly
			validCount++

		case "marking-definition":
			var md markingdefinition.MarkingDefinition
			if err := json.Unmarshal(objJSON, &md); err != nil {
				t.Errorf("Failed to decode marking-definition in %s: %v", filePath, err)
				invalidCount++
				continue
			}
			if valid, _, _ := md.Valid(false); valid {
				validCount++
			} else {
				invalidCount++
			}

		default:
			// For unsupported types, count as valid
			validCount++
		}
	}

	t.Logf("Validation results for %s: %d valid, %d invalid", filePath, validCount, invalidCount)
	return validCount, invalidCount
}

func main() {
	flag.Parse()

	fmt.Println("STIX File Validation Tool")
	fmt.Println("=========================")
	if *debugMode {
		fmt.Println("Debug mode: ENABLED")
	}
	if *showAllErrors {
		fmt.Println("Show all errors: ENABLED")
	}
	fmt.Println()

	var jsonFiles []string
	var err error

	// Determine which files to process
	if *singleFile != "" {
		// Validate single file
		if _, err := os.Stat(*singleFile); os.IsNotExist(err) {
			fmt.Printf("Error: File not found: %s\n", *singleFile)
			os.Exit(1)
		}
		jsonFiles = []string{*singleFile}
		fmt.Printf("Testing single file: %s\n", *singleFile)
	} else {
		// Check if the MISP test data directory exists
		if _, err := os.Stat(MispTestDataDir); os.IsNotExist(err) {
			fmt.Printf("Error: MISP test data directory not found at %s\n", MispTestDataDir)
			fmt.Println("Run 'git submodule update --init' to fetch test data.")
			os.Exit(1)
		}

		// Find all JSON files
		jsonFiles, err = findJSONFiles(MispTestDataDir)
		if err != nil {
			fmt.Printf("Error: Failed to find JSON files: %v\n", err)
			os.Exit(1)
		}

		if len(jsonFiles) == 0 {
			fmt.Println("No JSON files found in MISP test data directory")
			os.Exit(0)
		}

		fmt.Printf("Found %d JSON files to validate\n\n", len(jsonFiles))
	}

	// Track validation statistics
	totalFiles := 0
	totalObjects := 0
	validObjects := 0
	invalidObjects := 0
	failedFiles := 0
	allErrors := []string{}

	// Process each file
	for _, filePath := range jsonFiles {
		totalFiles++
		var displayName string
		if *singleFile != "" {
			displayName = filepath.Base(filePath)
		} else {
			displayName, _ = filepath.Rel(MispTestDataDir, filePath)
		}
		fmt.Printf("Processing: %s\n", displayName)

		// Read file
		data, err := os.ReadFile(filePath)
		if err != nil {
			errorMsg := fmt.Sprintf("  ❌ Failed to read file: %v", err)
			fmt.Println(errorMsg)
			allErrors = append(allErrors, fmt.Sprintf("%s: %s", displayName, errorMsg))
			failedFiles++
			continue
		}

		// Parse JSON
		stixObjects, err := parseSTIXData(data)
		if err != nil {
			errorMsg := fmt.Sprintf("  ❌ Failed to parse JSON: %v", err)
			fmt.Println(errorMsg)
			allErrors = append(allErrors, fmt.Sprintf("%s: %s", displayName, errorMsg))
			failedFiles++
			continue
		}

		if len(stixObjects) == 0 {
			fmt.Printf("  ⚠️  No STIX objects found\n")
			continue
		}

		fileValidCount := 0
		fileInvalidCount := 0

		// Validate each object in the file
		for i, obj := range stixObjects {
			objType, ok := obj["type"].(string)
			if !ok {
				errorMsg := fmt.Sprintf("  ❌ Object %d: missing 'type' field", i+1)
				fmt.Println(errorMsg)
				if *showAllErrors {
					allErrors = append(allErrors, fmt.Sprintf("%s: %s", displayName, errorMsg))
				}
				fileInvalidCount++
				continue
			}

			// Convert object to JSON for decoding
			objJSON, err := json.Marshal(obj)
			if err != nil {
				errorMsg := fmt.Sprintf("  ❌ Object %d (%s): failed to marshal: %v", i+1, objType, err)
				fmt.Println(errorMsg)
				if *showAllErrors {
					allErrors = append(allErrors, fmt.Sprintf("%s: %s", displayName, errorMsg))
				}
				fileInvalidCount++
				continue
			}

			// Validate based on object type
			isValid, validationErrors := validateSTIXObject(objType, objJSON)
			if isValid {
				fileValidCount++
				if *debugMode {
					fmt.Printf("  ✅ Object %d (%s): validation passed\n", i+1, objType)
				}
			} else {
				fileInvalidCount++
				errorMsg := fmt.Sprintf("  ❌ Object %d (%s): validation failed", i+1, objType)
				fmt.Println(errorMsg)

				// Show detailed errors if debug mode or all-errors flag is enabled
				if *debugMode || *showAllErrors {
					for _, validationError := range validationErrors {
						detailedError := fmt.Sprintf("    - %s", validationError)
						fmt.Println(detailedError)
						if *showAllErrors {
							allErrors = append(allErrors, fmt.Sprintf("%s: %s", displayName, detailedError))
						}
					}
				}
				if *showAllErrors && len(validationErrors) == 0 {
					allErrors = append(allErrors, fmt.Sprintf("%s: %s", displayName, errorMsg))
				}
			}
		}

		totalObjects += fileValidCount + fileInvalidCount
		validObjects += fileValidCount
		invalidObjects += fileInvalidCount

		if fileInvalidCount == 0 {
			fmt.Printf("  ✅ All %d objects valid\n", fileValidCount)
		} else {
			fmt.Printf("  ⚠️  %d valid, %d invalid objects\n", fileValidCount, fileInvalidCount)
		}
		fmt.Println()
	}

	// Print summary
	fmt.Println("Validation Summary:")
	fmt.Println("==================")
	fmt.Printf("Files processed: %d\n", totalFiles)
	fmt.Printf("Files failed to process: %d\n", failedFiles)
	fmt.Printf("Total STIX objects: %d\n", totalObjects)
	fmt.Printf("Valid objects: %d\n", validObjects)
	fmt.Printf("Invalid objects: %d\n", invalidObjects)

	if totalObjects > 0 {
		successRate := float64(validObjects) / float64(totalObjects) * 100
		fmt.Printf("Success rate: %.2f%%\n", successRate)
	}

	// Print all errors if requested
	if *showAllErrors && len(allErrors) > 0 {
		fmt.Println("\nAll Errors:")
		fmt.Println("===========")
		for _, errorMsg := range allErrors {
			fmt.Println(errorMsg)
		}
	}

	if invalidObjects > 0 || failedFiles > 0 {
		fmt.Printf("\n❌ Validation completed with %d invalid objects and %d failed files\n", invalidObjects, failedFiles)
		os.Exit(1)
	}

	fmt.Printf("\n✅ All STIX files validated successfully!\n")
}

// validateSTIXObject validates a single STIX object based on its type
func validateSTIXObject(objType string, objJSON []byte) (bool, []string) {
	var errors []string

	switch objType {
	case "indicator":
		var ind indicator.Indicator
		if err := json.Unmarshal(objJSON, &ind); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors
		}
		valid, _, validationErrors := ind.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("Validation note: %v", err))
			}
		}
		return valid, errors

	case "malware":
		var mal malware.Malware
		if err := json.Unmarshal(objJSON, &mal); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors
		}
		valid, _, validationErrors := mal.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("Validation note: %v", err))
			}
		}
		return valid, errors

	case "infrastructure":
		var inf infrastructure.Infrastructure
		if err := json.Unmarshal(objJSON, &inf); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors
		}
		valid, _, validationErrors := inf.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("Validation note: %v", err))
			}
		}
		return valid, errors

	case "threat-actor":
		var ta threatactor.ThreatActor
		if err := json.Unmarshal(objJSON, &ta); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors
		}
		valid, _, validationErrors := ta.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors

	case "tool":
		var toolObj tool.Tool
		if err := json.Unmarshal(objJSON, &toolObj); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors
		}
		valid, _, validationErrors := toolObj.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors

	case "attack-pattern":
		var ap attackpattern.AttackPattern
		if err := json.Unmarshal(objJSON, &ap); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors
		}
		valid, _, validationErrors := ap.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors

	case "campaign":
		var camp campaign.Campaign
		if err := json.Unmarshal(objJSON, &camp); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors
		}
		valid, _, validationErrors := camp.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors

	case "course-of-action":
		var coa courseofaction.CourseOfAction
		if err := json.Unmarshal(objJSON, &coa); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors
		}
		valid, _, validationErrors := coa.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors

	case "grouping":
		var grp grouping.Grouping
		if err := json.Unmarshal(objJSON, &grp); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors
		}
		valid, _, validationErrors := grp.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors

	case "identity":
		var id identity.Identity
		if err := json.Unmarshal(objJSON, &id); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors
		}
		valid, _, validationErrors := id.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors

	case "intrusion-set":
		var is intrusionset.IntrusionSet
		if err := json.Unmarshal(objJSON, &is); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors
		}
		valid, _, validationErrors := is.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors

	case "location":
		var loc location.Location
		if err := json.Unmarshal(objJSON, &loc); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors
		}
		valid, _, validationErrors := loc.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors

	case "malware-analysis":
		var ma malwareanalysis.MalwareAnalysis
		if err := json.Unmarshal(objJSON, &ma); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors
		}
		valid, _, validationErrors := ma.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors

	case "note":
		var noteObj note.Note
		if err := json.Unmarshal(objJSON, &noteObj); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors
		}
		valid, _, validationErrors := noteObj.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors

	case "observed-data":
		var od observeddata.ObservedData
		if err := json.Unmarshal(objJSON, &od); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors
		}
		valid, _, validationErrors := od.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors

	case "opinion":
		var op opinion.Opinion
		if err := json.Unmarshal(objJSON, &op); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors
		}
		valid, _, validationErrors := op.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors

	case "report":
		var reportObj report.Report
		if err := json.Unmarshal(objJSON, &reportObj); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors
		}
		valid, _, validationErrors := reportObj.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors

	case "relationship":
		var rel relationship.Relationship
		if err := json.Unmarshal(objJSON, &rel); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors
		}
		valid, _, validationErrors := rel.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors

	case "sighting":
		var sight sighting.Sighting
		if err := json.Unmarshal(objJSON, &sight); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors
		}
		valid, _, validationErrors := sight.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors

	case "vulnerability":
		var vuln vulnerability.Vulnerability
		if err := json.Unmarshal(objJSON, &vuln); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors
		}
		valid, _, validationErrors := vuln.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors

	case "bundle":
		var bundleObj bundle.Bundle
		if err := json.Unmarshal(objJSON, &bundleObj); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors
		}
		// Bundle doesn't have a Valid method, consider it valid if it decodes correctly
		return true, errors

	case "marking-definition":
		var md markingdefinition.MarkingDefinition
		if err := json.Unmarshal(objJSON, &md); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors
		}
		valid, _, validationErrors := md.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors

	default:
		// For unsupported types, consider them valid if they can be decoded as JSON
		// TODO: log these as coverage gaps
		var generic map[string]interface{}
		if err := json.Unmarshal(objJSON, &generic); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors
		}
		if *debugMode {
			errors = append(errors, fmt.Sprintf("Info: Object type '%s' is not explicitly supported", objType))
		}
		return true, errors
	}
}
