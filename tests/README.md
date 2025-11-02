# STIX Test Suite

This directory contains comprehensive test suites for the libstix2 library to ensure compatibility and correctness when working with STIX 2.x data.

## Test Files

### `stix_roundtrip_test.go`

A comprehensive test suite that validates the library's ability to decode and re-encode STIX objects without data loss.

**Features:**

- Tests against 8,000+ real STIX objects from the MISP STIX test suite
- Validates all supported STIX object types (SDO, SRO, SCO, and meta objects)
- Performs round-trip encoding/decoding tests
- Validates object constraints and vocabulary compliance
- Provides detailed reporting of test results

**Test Functions:**

- `TestMISPSTIXRoundTrip` - Tests that STIX data can be decoded and re-encoded without losing information
- `TestMISPSTIXValidation` - Tests that all STIX objects pass validation when decoded using libstix2

### `stix_simple_test.go`

A simple test that demonstrates basic round-trip functionality with a single STIX indicator object.

**Features:**

- Tests basic JSON decoding and encoding
- Validates that all key fields are preserved
- Provides clear logging of before/after JSON structures

## Setup

The tests use the MISP STIX test suite as a submodule. To initialize:

```bash
git submodule update --init --recursive
```

## Running Tests

### Run all tests

```bash
go test ./tests/ -v
```

### Run specific test suites

```bash
# Run the comprehensive round-trip test
go test ./tests/ -v -run TestMISPSTIXRoundTrip

# Run the validation test
go test ./tests/ -v -run TestMISPSTIXValidation

# Run the simple test
go test ./tests/ -v -run TestSimpleSTIXRoundTrip
```

## Test Data

The tests use real-world STIX data from:

- **MISP STIX Test Suite**: Located in `tests/data/misp-stix-tests/files/`
  - Contains 8,000+ STIX objects covering all object types
  - Includes various TLP levels and consent configurations
  - Features real-world threat intelligence data

## Supported STIX Object Types

The test suite validates all STIX 2.1 object types supported by libstix2:

### STIX Domain Objects (SDOs)

- `attack-pattern`
- `campaign`
- `course-of-action`
- `grouping`
- `identity`
- `indicator`
- `infrastructure`
- `intrusion-set`
- `location`
- `malware`
- `malware-analysis`
- `note`
- `observed-data`
- `opinion`
- `report`
- `threat-actor`
- `tool`
- `vulnerability`

### STIX Relationship Objects (SROs)

- `relationship`
- `sighting`

### STIX Cyber Observable Objects (SCOs)

- `artifact`
- `autonomous-system`
- `directory`
- `domain-name`
- `email-addr`
- `email-message`
- `file`
- `ipv4-addr`
- `ipv6-addr`
- `mac-addr`
- `mutex`
- `network-traffic`
- `process`
- `software`
- `url`
- `user-account`
- `windows-registry-key`
- `x509-certificate`

### Meta Objects

- `language-content`
- `marking-definition`

### Container Objects

- `bundle`

## Test Results

The comprehensive test suite provides detailed statistics including:

- Total files processed
- Number of objects successfully processed
- Validation pass/fail counts
- Detailed error information for failures

## Notes

- The round-trip test may show JSON formatting differences (field order, whitespace) which are normal and expected
- The validation test ensures all objects meet STIX 2.1 specification requirements
- Some test objects are intentionally invalid to test the library's validation capabilities
- The test suite is designed to be comprehensive and can identify regressions in object handling
