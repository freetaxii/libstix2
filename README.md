# FreeTAXII/libstix2

[![Go Report Card](https://goreportcard.com/badge/github.com/freetaxii/libstix2)](https://goreportcard.com/report/github.com/freetaxii/libstix2) [![GoDoc](https://godoc.org/github.com/freetaxii/libstix2?status.png)](https://godoc.org/github.com/freetaxii/libstix2)

libstix2 an API for generating JSON based STIX objects and TAXII messages with
the Go (Golang) programming language. Please see the examples directory and the
README files in each of the sub packages for more information. This API is built
to support STIX 2.x and TAXII 2.x.

## Version

0.7.2

## Installation

This package can be installed with the go get command:

```
go get github.com/freetaxii/libstix2
```

## Special Thanks

I would like to thank the following contributors for there support of this project:

```
Oleksii Morozov
```

## Dependencies

This software uses the following external libraries:

```
uuid
 go get github.com/google/uuid
 Copyright (c) 2014 Google Inc. All rights reserved. (Google License)
```

This software uses the following builtin libraries:

```
crypto/sha1, encoding/base64, errors, fmt, log, os, regexp, strings, testing, time
 Copyright 2009 The Go Authors
```

## Specification Compliance

This library implements **STIX 2.1** and **TAXII 2.1** as defined by OASIS.

**Specification Reference**: [STIX Version 2.1 - OASIS Committee Specification Draft 01](https://docs.oasis-open.org/cti/stix/v2.1/csprd01/stix-v2.1-csprd01.html)

All implemented objects include documentation with direct references to the relevant sections of the STIX 2.1 specification for complete traceability and compliance verification.

## Features

Below is a list of major features and which ones have been implemented:

### STIX Domain Objects (SDOs) - 18/18 ✅

All STIX 2.1 Domain Objects are fully implemented:

- [x] Attack Pattern (Section 4.1)
- [x] Campaign (Section 4.2)
- [x] Course of Action (Section 4.3)
- [x] Grouping (Section 4.4)
- [x] Identity (Section 4.5)
- [x] Indicator (Section 4.6)
- [x] Infrastructure (Section 4.7)
- [x] Intrusion Set (Section 4.8)
- [x] Location (Section 4.9)
- [x] Malware (Section 4.10)
- [x] Malware Analysis (Section 4.11)
- [x] Note (Section 4.12)
- [x] Observed Data (Section 4.13)
- [x] Opinion (Section 4.14)
- [x] Report (Section 4.15)
- [x] Threat Actor (Section 4.16)
- [x] Tool (Section 4.17)
- [x] Vulnerability (Section 4.18)

### STIX Relationship Objects (SROs) - 2/2 ✅

- [x] Relationship (Section 5.1)
- [x] Sighting (Section 5.2)

### STIX Cyber-observable Objects (SCOs) - 18/18 ✅

All STIX 2.1 Cyber-observable Objects are implemented:

- [x] Artifact (Section 6.1)
- [x] Autonomous System (Section 6.2)
- [x] Directory (Section 6.3)
- [x] Domain Name (Section 6.3)
- [x] Email Address (Section 6.4)
- [x] Email Message (Section 6.5)
- [x] File (Section 6.6)
- [x] IPv4 Address (Section 6.7)
- [x] IPv6 Address (Section 6.8)
- [x] MAC Address (Section 6.8)
- [x] Mutex (Section 6.9)
- [x] Network Traffic (Section 6.10)
- [x] Process (Section 6.11)
- [x] Software (Section 6.13)
- [x] URL (Section 6.15)
- [x] User Account (Section 6.14)
- [x] Windows Registry Key (Section 6.16)
- [x] X.509 Certificate (Section 6.17)

**Note**: Some SCO objects have complete implementations while others have framework structures in place and require property completion per the specification. See individual package documentation for details.

### STIX Meta Objects - 2/2 ✅

- [x] Language Content (Section 7.2)
- [x] Marking Definition (Section 7.3)

### STIX Bundle Object - 1/1 ✅

- [x] Bundle (Section 8)

### TAXII Resources - 9/9 ✅

- [x] Discovery
- [x] API Root
- [x] Collections
- [x] Collection
- [x] Objects
- [x] Manifest
- [x] Envelope
- [x] Status
- [x] Error


## Implementation Notes

### Complete vs. Framework Implementations

**Fully Implemented Objects** (with all properties, validation, and tests):

- All STIX Domain Objects (SDOs)
- All STIX Relationship Objects (SROs)
- STIX Meta Objects: Marking Definition, Language Content
- STIX Bundle Object
- Selected SCOs: Domain Name, IPv4 Address, IPv6 Address, URL, MAC Address, File, Artifact, Mutex

**Framework Implementations** (structure in place, properties need completion):

- Some SCOs have basic structure but require property completion per the specification
- Each framework implementation includes TODO comments indicating what needs to be completed
- See individual package documentation for specific implementation status

All objects follow the same consistent pattern:

- `model.go` - Object structure definition
- `doc.go` - Documentation with STIX 2.1 spec reference
- `json.go` - JSON marshaling/unmarshaling
- `setters.go` - Property setter methods
- `valid.go` - Validation logic

## Vocabularies

The `vocabs` package includes all STIX 2.1 defined vocabularies as specified in the standard:

- Account Type
- Attack Motivation
- Attack Resource Level
- Encryption Algorithm
- Extension Types
- Grouping Context
- Hashing Algorithm
- Identity Class
- Implementation Languages
- Indicator Types
- Industry Sector
- Infrastructure Types
- Malware AV Results
- Malware Capabilities
- Malware Types
- Network Socket Address Family
- Network Socket Type
- Opinion
- Pattern Type
- Processor Architecture
- Region
- Report Types
- Threat Actor Types
- Threat Actor Roles
- Threat Actor Sophistication
- Tool Types
- Windows Integrity Levels
- Windows PE Binary Types
- Windows Registry Datatypes
- Windows Service Start Types
- Windows Service Types
- Windows Service Status

Each vocabulary function includes complete values as defined in the STIX 2.1 specification.

## Naming Conventions

While Go does not require getters and setters, setters are used in libstix2 to enable validation and verification checks. All setters in libstix2 return an error type, even if they currently just return “nil”. This will ensure that the API will not change if/when additional validation / verification checks are added in the future.

Libstix2 uses the following naming conventions for methods on objects and resources.

- Methods that setup / create a new object have a name of "New" or "New"+ object type. These constructors return a pointer to the object.

- Methods that are setting a value have a name of “Set” + the property name. Example: “SetConfidence” is used for setting a value on the Confidence property.

- Methods that are getting a value have a name of “Get” + the property name. Example: “GetConfidence” is used for getting the value stored in the Confidence property.

- Methods that take in a value and add that value to a slice have a name of “Add” + the property name in the singular. Example: “AddLabel” is used to add a sting label to the labels property.

- Methods that take in an object and add that object to a slice have a name of “Add” + the object type in the singular. Example: “AddManifestEntry” is used to add a Manifest Entry to the Objects slice in the Manifest resource. It is important to note that these methods take in a pointer to the object instead of a copy of the object itself. Some examples with full signatures:

```
func (o *CollectionsType) AddCollection(o *CollectionType) (int, error) {}
func (o *ManifestType) AddManifestEntry(o *ManifestEntryType) (int, error) {}
```

- Methods that create a new object inside another object and return a pointer to a slice location for the new object have a name of “New” + the object type in the singular. Example: “NewManifestEntry” is used to create a new Manifest Entry in the Objects slice in the Manifest resource. Some examples with full signatures:

```
func (o *ExternalReferencesPropertyType) NewExternalReference() (*ExternalReferenceType, error) {}
func (o *KillChainPhasesPropertyType) NewKillChainPhase() (*KillChainPhaseType, error) {}
func (o *CollectionsType) NewCollection() (*CollectionType, error) {}
func (o *ManifestType) NewManifestEntry() (*ManifestEntryType, error) {}
```

- Methods that create and populate a new object in a single step have a name of “Create” + the object type in the singular. Example: “CreateManifestEntry” is used to create a new Manifest Entry in the Objects slice in the Manifest resource and populates it in one step. Some examples with full signatures:

```
func (o *KillChainPhasesPropertyType) CreateKillChainPhase(name, phase string) error {}
func (o *ManifestType) CreateManifestEntry(id, date, ver, media string) error {}
```

## License

This is free software, licensed under the Apache License, Version 2.0. [Read this](https://tldrlegal.com/license/apache-license-2.0-(apache-2.0)) for a summary.

## Copyright

Copyright 2015-2022 Bret Jordan, All rights reserved.
