# FreeTAXII/libstix2 #

[![Go Report Card](https://goreportcard.com/badge/github.com/freetaxii/libstix2)](https://goreportcard.com/report/github.com/freetaxii/libstix2) [![GoDoc](https://godoc.org/github.com/freetaxii/libstix2?status.png)](https://godoc.org/github.com/freetaxii/libstix2)

libstix2 an API for generating JSON based STIX objects and TAXII messages with 
the Go (Golang) programming language. Please see the examples directory and the 
README files in each of the sub packages for more information. This API is built 
to support STIX 2.x and TAXII 2.x.

## Installation ##

This package can be installed with the go get command:

```
go get github.com/freetaxii/libstix2
```

## Dependencies ##

This software uses the following external libraries:
```
uuid
	go get github.com/pborman/uuid
	Copyright (c) 2014 Google Inc. All rights reserved.
```

This software uses the following builtin libraries:
```
fmt, time
	Copyright 2009 The Go Authors
```


## Features ##

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

TAXII Features
- [x] Multiple Discovery Services
- [ ] Authentication
- [ ] Persistent storage
- [ ] Version checking
- [ ] URL parameters
- [ ] Object by ID

TAXII Resources
- [x] Discovery
- [x] API Root
- [x] Collections
- [x] Collection
- [ ] Objects
- [ ] Manifest
- [ ] Status
- [ ] Error


## License ##

This is free software, licensed under the Apache License, Version 2.0.


## Copyright ##

Copyright 2017 Bret Jordan, All rights reserved.

