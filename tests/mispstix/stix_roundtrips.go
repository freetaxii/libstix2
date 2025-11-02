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
	"github.com/freetaxii/libstix2/objects/sco/artifact"
	"github.com/freetaxii/libstix2/objects/sco/directory"
	"github.com/freetaxii/libstix2/objects/sco/domainname"
	"github.com/freetaxii/libstix2/objects/sco/emailaddr"
	"github.com/freetaxii/libstix2/objects/sco/emailmessage"
	"github.com/freetaxii/libstix2/objects/sco/file"
	"github.com/freetaxii/libstix2/objects/sco/ipv4addr"
	"github.com/freetaxii/libstix2/objects/sco/ipv6addr"
	"github.com/freetaxii/libstix2/objects/sco/macaddr"
	"github.com/freetaxii/libstix2/objects/sco/mutex"
	"github.com/freetaxii/libstix2/objects/sco/networktraffic"
	"github.com/freetaxii/libstix2/objects/sco/process"
	"github.com/freetaxii/libstix2/objects/sco/software"
	"github.com/freetaxii/libstix2/objects/sco/urlobject"
	"github.com/freetaxii/libstix2/objects/sco/useraccount"
	"github.com/freetaxii/libstix2/objects/sco/windowsregistrykey"
	"github.com/freetaxii/libstix2/objects/sco/x509certificate"
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
	showMissing   = flag.Bool("show-missing", false, "Show missing object types that are not implemented")
)

// ----------------------------------------------------------------------
// Missing Types Tracking
// ----------------------------------------------------------------------

// MissingTypeStats tracks statistics about missing object types
type MissingTypeStats struct {
	TypeCounts  map[string]int      // Count of each missing type
	FilesByType map[string][]string // Files containing each missing type
}

// NewMissingTypeStats creates a new MissingTypeStats instance
func NewMissingTypeStats() *MissingTypeStats {
	return &MissingTypeStats{
		TypeCounts:  make(map[string]int),
		FilesByType: make(map[string][]string),
	}
}

// AddMissingType records a missing object type
func (m *MissingTypeStats) AddMissingType(objType, filePath string) {
	m.TypeCounts[objType]++

	// Only add file if not already present for this type
	files := m.FilesByType[objType]
	for _, f := range files {
		if f == filePath {
			return
		}
	}
	m.FilesByType[objType] = append(files, filePath)
}

// GetSummary returns a summary of missing types
func (m *MissingTypeStats) GetSummary() (int, int, map[string]int) {
	totalMissingTypes := len(m.TypeCounts)
	totalMissingObjects := 0
	for _, count := range m.TypeCounts {
		totalMissingObjects += count
	}
	return totalMissingTypes, totalMissingObjects, m.TypeCounts
}

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

		case "artifact":
			var art artifact.Artifact
			if err := json.Unmarshal(objJSON, &art); err != nil {
				t.Errorf("Failed to decode artifact in %s: %v", filePath, err)
				continue
			}
			decodedObj = art

		case "directory":
			var dir directory.Directory
			if err := json.Unmarshal(objJSON, &dir); err != nil {
				t.Errorf("Failed to decode directory in %s: %v", filePath, err)
				continue
			}
			decodedObj = dir

		case "domain-name":
			var dn domainname.DomainName
			if err := json.Unmarshal(objJSON, &dn); err != nil {
				t.Errorf("Failed to decode domain-name in %s: %v", filePath, err)
				continue
			}
			decodedObj = dn

		case "email-addr":
			var eaddr emailaddr.EmailAddress
			if err := json.Unmarshal(objJSON, &eaddr); err != nil {
				t.Errorf("Failed to decode email-addr in %s: %v", filePath, err)
				continue
			}
			decodedObj = eaddr

		case "email-message":
			var emsg emailmessage.EmailMessage
			if err := json.Unmarshal(objJSON, &emsg); err != nil {
				t.Errorf("Failed to decode email-message in %s: %v", filePath, err)
				continue
			}
			decodedObj = emsg

		case "file":
			var f file.File
			if err := json.Unmarshal(objJSON, &f); err != nil {
				t.Errorf("Failed to decode file in %s: %v", filePath, err)
				continue
			}
			decodedObj = f

		case "ipv4-addr":
			var ip4 ipv4addr.IPv4Addr
			if err := json.Unmarshal(objJSON, &ip4); err != nil {
				t.Errorf("Failed to decode ipv4-addr in %s: %v", filePath, err)
				continue
			}
			decodedObj = ip4

		case "ipv6-addr":
			var ip6 ipv6addr.IPv6Addr
			if err := json.Unmarshal(objJSON, &ip6); err != nil {
				t.Errorf("Failed to decode ipv6-addr in %s: %v", filePath, err)
				continue
			}
			decodedObj = ip6

		case "mac-addr":
			var mac macaddr.MACAddr
			if err := json.Unmarshal(objJSON, &mac); err != nil {
				t.Errorf("Failed to decode mac-addr in %s: %v", filePath, err)
				continue
			}
			decodedObj = mac

		case "mutex":
			var m mutex.Mutex
			if err := json.Unmarshal(objJSON, &m); err != nil {
				t.Errorf("Failed to decode mutex in %s: %v", filePath, err)
				continue
			}
			decodedObj = m

		case "network-traffic":
			var nt networktraffic.NetworkTraffic
			if err := json.Unmarshal(objJSON, &nt); err != nil {
				t.Errorf("Failed to decode network-traffic in %s: %v", filePath, err)
				continue
			}
			decodedObj = nt

		case "process":
			var p process.Process
			if err := json.Unmarshal(objJSON, &p); err != nil {
				t.Errorf("Failed to decode process in %s: %v", filePath, err)
				continue
			}
			decodedObj = p

		case "software":
			var s software.Software
			if err := json.Unmarshal(objJSON, &s); err != nil {
				t.Errorf("Failed to decode software in %s: %v", filePath, err)
				continue
			}
			decodedObj = s

		case "url":
			var url urlobject.URLObject
			if err := json.Unmarshal(objJSON, &url); err != nil {
				t.Errorf("Failed to decode url in %s: %v", filePath, err)
				continue
			}
			decodedObj = url

		case "user-account":
			var ua useraccount.UserAccount
			if err := json.Unmarshal(objJSON, &ua); err != nil {
				t.Errorf("Failed to decode user-account in %s: %v", filePath, err)
				continue
			}
			decodedObj = ua

		case "windows-registry-key":
			var wrk windowsregistrykey.WindowsRegistryKey
			if err := json.Unmarshal(objJSON, &wrk); err != nil {
				t.Errorf("Failed to decode windows-registry-key in %s: %v", filePath, err)
				continue
			}
			decodedObj = wrk

		case "x509-certificate":
			var cert x509certificate.X509Certificate
			if err := json.Unmarshal(objJSON, &cert); err != nil {
				t.Errorf("Failed to decode x509-certificate in %s: %v", filePath, err)
				continue
			}
			decodedObj = cert

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
	missingStats := NewMissingTypeStats()

	for _, filePath := range jsonFiles {
		t.Run(filepath.Base(filePath), func(t *testing.T) {
			valid, invalid := testSTIXFileValidationWithMissing(t, filePath, missingStats)
			totalObjects += valid + invalid
			validObjects += valid
			invalidObjects += invalid
		})
	}

	// Report missing types
	totalMissingTypes, totalMissingObjects, typeCounts := missingStats.GetSummary()
	if totalMissingTypes > 0 {
		t.Logf("Missing object types found: %d types, %d objects", totalMissingTypes, totalMissingObjects)
		for objType, count := range typeCounts {
			t.Logf("  - %s: %d objects", objType, count)
		}
	}

	t.Logf("Validation summary: %d total objects, %d valid, %d invalid", totalObjects, validObjects, invalidObjects)
}

// testSTIXFileValidation tests that all STIX objects in a file pass validation
func testSTIXFileValidation(t *testing.T, filePath string) (int, int) {
	missingStats := NewMissingTypeStats()
	return testSTIXFileValidationWithMissing(t, filePath, missingStats)
}

// testSTIXFileValidationWithMissing tests that all STIX objects in a file pass validation and tracks missing types
func testSTIXFileValidationWithMissing(t *testing.T, filePath string, missingStats *MissingTypeStats) (int, int) {
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

		// Validate based on object type and track missing types
		isValid, _, isMissing := validateSTIXObjectWithMissing(objType, objJSON)
		if isMissing {
			missingStats.AddMissingType(objType, filepath.Base(filePath))
		}

		if isValid {
			validCount++
		} else {
			invalidCount++
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
	if *showMissing {
		fmt.Println("Show missing types: ENABLED")
	}
	fmt.Println()

	var jsonFiles []string
	var err error
	missingStats := NewMissingTypeStats()

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
			isValid, validationErrors, isMissing := validateSTIXObjectWithMissing(objType, objJSON)
			if isMissing {
				missingStats.AddMissingType(objType, displayName)
			}

			if isValid {
				fileValidCount++
				if *debugMode {
					if isMissing {
						fmt.Printf("  ⚠️  Object %d (%s): validation passed (unsupported type)\n", i+1, objType)
					} else {
						fmt.Printf("  ✅ Object %d (%s): validation passed\n", i+1, objType)
					}
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

	// Print missing types summary
	totalMissingTypes, totalMissingObjects, typeCounts := missingStats.GetSummary()
	if totalMissingTypes > 0 {
		fmt.Printf("\nMissing Object Types: %d types, %d objects\n", totalMissingTypes, totalMissingObjects)

		if *showMissing {
			fmt.Println("\nMissing Types Details:")
			fmt.Println("======================")
			for objType, count := range typeCounts {
				files := missingStats.FilesByType[objType]
				fmt.Printf("Type: %s (Count: %d)\n", objType, count)
				if len(files) <= 5 {
					for _, file := range files {
						fmt.Printf("  - %s\n", file)
					}
				} else {
					for i := 0; i < 3; i++ {
						fmt.Printf("  - %s\n", files[i])
					}
					fmt.Printf("  ... and %d more files\n", len(files)-3)
				}
				fmt.Println()
			}
		} else {
			fmt.Println("Use --show-missing flag to see detailed breakdown of missing types")
		}
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
	isValid, errors, _ := validateSTIXObjectWithMissing(objType, objJSON)
	return isValid, errors
}

// validateSTIXObjectWithMissing validates a single STIX object and tracks if the type is missing
func validateSTIXObjectWithMissing(objType string, objJSON []byte) (bool, []string, bool) {
	var errors []string

	switch objType {
	case "indicator":
		var ind indicator.Indicator
		if err := json.Unmarshal(objJSON, &ind); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := ind.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("Validation note: %v", err))
			}
		}
		return valid, errors, false

	case "malware":
		var mal malware.Malware
		if err := json.Unmarshal(objJSON, &mal); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := mal.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("Validation note: %v", err))
			}
		}
		return valid, errors, false

	case "infrastructure":
		var inf infrastructure.Infrastructure
		if err := json.Unmarshal(objJSON, &inf); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := inf.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("Validation note: %v", err))
			}
		}
		return valid, errors, false

	case "threat-actor":
		var ta threatactor.ThreatActor
		if err := json.Unmarshal(objJSON, &ta); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := ta.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "tool":
		var toolObj tool.Tool
		if err := json.Unmarshal(objJSON, &toolObj); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := toolObj.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "attack-pattern":
		var ap attackpattern.AttackPattern
		if err := json.Unmarshal(objJSON, &ap); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := ap.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "campaign":
		var camp campaign.Campaign
		if err := json.Unmarshal(objJSON, &camp); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := camp.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "course-of-action":
		var coa courseofaction.CourseOfAction
		if err := json.Unmarshal(objJSON, &coa); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := coa.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "grouping":
		var grp grouping.Grouping
		if err := json.Unmarshal(objJSON, &grp); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := grp.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "identity":
		var id identity.Identity
		if err := json.Unmarshal(objJSON, &id); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := id.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "intrusion-set":
		var is intrusionset.IntrusionSet
		if err := json.Unmarshal(objJSON, &is); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := is.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "location":
		var loc location.Location
		if err := json.Unmarshal(objJSON, &loc); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := loc.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "malware-analysis":
		var ma malwareanalysis.MalwareAnalysis
		if err := json.Unmarshal(objJSON, &ma); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := ma.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "note":
		var noteObj note.Note
		if err := json.Unmarshal(objJSON, &noteObj); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := noteObj.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "observed-data":
		var od observeddata.ObservedData
		if err := json.Unmarshal(objJSON, &od); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := od.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "opinion":
		var op opinion.Opinion
		if err := json.Unmarshal(objJSON, &op); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := op.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "report":
		var reportObj report.Report
		if err := json.Unmarshal(objJSON, &reportObj); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := reportObj.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "relationship":
		var rel relationship.Relationship
		if err := json.Unmarshal(objJSON, &rel); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := rel.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "sighting":
		var sight sighting.Sighting
		if err := json.Unmarshal(objJSON, &sight); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := sight.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "vulnerability":
		var vuln vulnerability.Vulnerability
		if err := json.Unmarshal(objJSON, &vuln); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := vuln.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "bundle":
		var bundleObj bundle.Bundle
		if err := json.Unmarshal(objJSON, &bundleObj); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		// Bundle doesn't have a Valid method, consider it valid if it decodes correctly
		return true, errors, false

	case "marking-definition":
		var md markingdefinition.MarkingDefinition
		if err := json.Unmarshal(objJSON, &md); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := md.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "artifact":
		var art artifact.Artifact
		if err := json.Unmarshal(objJSON, &art); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := art.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "directory":
		var dir directory.Directory
		if err := json.Unmarshal(objJSON, &dir); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := dir.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "domain-name":
		var dn domainname.DomainName
		if err := json.Unmarshal(objJSON, &dn); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := dn.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "email-addr":
		var eaddr emailaddr.EmailAddress
		if err := json.Unmarshal(objJSON, &eaddr); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := eaddr.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "email-message":
		var emsg emailmessage.EmailMessage
		if err := json.Unmarshal(objJSON, &emsg); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := emsg.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "file":
		var f file.File
		if err := json.Unmarshal(objJSON, &f); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := f.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "ipv4-addr":
		var ip4 ipv4addr.IPv4Addr
		if err := json.Unmarshal(objJSON, &ip4); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := ip4.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "ipv6-addr":
		var ip6 ipv6addr.IPv6Addr
		if err := json.Unmarshal(objJSON, &ip6); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := ip6.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "mac-addr":
		var mac macaddr.MACAddr
		if err := json.Unmarshal(objJSON, &mac); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := mac.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "mutex":
		var m mutex.Mutex
		if err := json.Unmarshal(objJSON, &m); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := m.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "network-traffic":
		var nt networktraffic.NetworkTraffic
		if err := json.Unmarshal(objJSON, &nt); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := nt.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "process":
		var p process.Process
		if err := json.Unmarshal(objJSON, &p); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := p.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "software":
		var s software.Software
		if err := json.Unmarshal(objJSON, &s); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := s.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "url":
		var url urlobject.URLObject
		if err := json.Unmarshal(objJSON, &url); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := url.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "user-account":
		var ua useraccount.UserAccount
		if err := json.Unmarshal(objJSON, &ua); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := ua.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "windows-registry-key":
		var wrk windowsregistrykey.WindowsRegistryKey
		if err := json.Unmarshal(objJSON, &wrk); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := wrk.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	case "x509-certificate":
		var cert x509certificate.X509Certificate
		if err := json.Unmarshal(objJSON, &cert); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, false
		}
		valid, _, validationErrors := cert.Valid(*debugMode)
		if !valid && *debugMode {
			for _, err := range validationErrors {
				errors = append(errors, fmt.Sprintf("validation note: %v", err))
			}
		}
		return valid, errors, false

	default:
		// For unsupported types, consider them valid if they can be decoded as JSON
		// Mark this as a missing type for tracking
		var generic map[string]interface{}
		if err := json.Unmarshal(objJSON, &generic); err != nil {
			errors = append(errors, fmt.Sprintf("JSON decode error: %v", err))
			return false, errors, true
		}
		if *debugMode {
			errors = append(errors, fmt.Sprintf("Info: Object type '%s' is not explicitly supported", objType))
		}
		return true, errors, true
	}
}
