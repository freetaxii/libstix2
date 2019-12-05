# FreeTAXII/libstix2

[![Go Report Card](https://goreportcard.com/badge/github.com/freetaxii/libstix2)](https://goreportcard.com/report/github.com/freetaxii/libstix2) [![GoDoc](https://godoc.org/github.com/freetaxii/libstix2?status.png)](https://godoc.org/github.com/freetaxii/libstix2)

libstix2 an API for generating JSON based STIX objects and TAXII messages with 
the Go (Golang) programming language. Please see the examples directory and the 
README files in each of the sub packages for more information. This API is built 
to support STIX 2.x and TAXII 2.x.

## Version
0.6.1

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
	go get github.com/pborman/uuid
	Copyright (c) 2014 Google Inc. All rights reserved. (Google License)

qo-sqlite3
	go get github.com/mattn/go-sqlite3
	go install github.com/mattn/go-sqlite3
	Copyright (c) 2014 Yasuhiro Matsumoto (MIT License)
```

This software uses the following builtin libraries:
```
crypto/sha1, database/sql, encoding/base64, errors, fmt, log, os, regexp, strings, testing, time
	Copyright 2009 The Go Authors
```


## Features

Below is a list of major features and which ones have been implemented:

STIX Domain Objects
- [x] Attack Pattern
- [x] Campaign
- [x] Course of Action
- [x] Identity
- [x] Indicator
- [x] Intrusion Set
- [ ] Location
- [x] Malware
- [ ] Note
- [x] Observed Data
- [ ] Opinion
- [x] Report
- [x] Threat Actor
- [x] Tool
- [x] Vulnerability

STIX Relationship Objects
- [x] Relationship
- [x] Sighting

Other STIX Objects
- [x] Bundle
- [ ] Language Content
- [ ] Marking Definition

TAXII Resources
- [x] Discovery
- [x] API Root
- [x] Collections
- [x] Collection
- [x] Objects
- [x] Manifest
- [x] Envelope
- [x] Status
- [x] Error

Datastore
- [ ] SQLite 3


## Naming Conventions

While Go does not require getters and setters, setters are used in libstix2 to enable validation and verification checks. All setters in libstix2 return an error type, even if they currently just return “nil”. This will ensure that the API will not change if/when additional validation / verification checks are added in the future. 

Libstix2 uses the following naming conventions for methods on objects and resources.

* Methods that setup / create a new object have a name of "New" or "New"+ object type. These constructors return a pointer to the object. 

* Methods that are setting a value have a name of “Set” + the property name. Example: “SetConfidence” is used for setting a value on the Confidence property.

* Methods that are getting a value have a name of “Get” + the property name. Example: “GetConfidence” is used for getting the value stored in the Confidence property.

* Methods that take in a value and add that value to a slice have a name of “Add” + the property name in the singular. Example: “AddLabel” is used to add a sting label to the labels property. 

* Methods that take in an object and add that object to a slice have a name of “Add” + the object type in the singular. Example: “AddManifestEntry” is used to add a Manifest Entry to the Objects slice in the Manifest resource. It is important to note that these methods take in a pointer to the object instead of a copy of the object itself. Some examples with full signatures:

```
func (o *CollectionsType) AddCollection(o *CollectionType) (int, error) {}
func (o *ManifestType) AddManifestEntry(o *ManifestEntryType) (int, error) {}
```

* Methods that create a new a new object inside another object and return a pointer to a slice location for the new object have a name of “New” + the object type in the singular. Example: “NewManifestEntry” is used to create a new Manifest Entry in the Objects slice in the Manifest resource. Some examples with full signatures:

```
func (o *ExternalReferencesPropertyType) NewExternalReference() (*ExternalReferenceType, error) {}
func (o *KillChainPhasesPropertyType) NewKillChainPhase() (*KillChainPhaseType, error) {}
func (o *CollectionsType) NewCollection() (*CollectionType, error) {}
func (o *ManifestType) NewManifestEntry() (*ManifestEntryType, error) {}
```

* Methods that create and populate a new object in a single step have a name of “Create” + the object type in the singular. Example: “CreateManifestEntry” is used to create a new Manifest Entry in the Objects slice in the Manifest resource and populates it in one step. Some examples with full signatures:

```
func (o *KillChainPhasesPropertyType) CreateKillChainPhase(name, phase string) error {}
func (o *ManifestType) CreateManifestEntry(id, date, ver, media string) error {}
```


## License

This is free software, licensed under the Apache License, Version 2.0. [Read this](https://tldrlegal.com/license/apache-license-2.0-(apache-2.0)) for a summary.


## Copyright

Copyright 2015-2019 Bret Jordan, All rights reserved.

