// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package datastore

import (
	"github.com/freetaxii/libstix2/resources"
)

/*
Datastorer - This interface enables access to the STIX/TAXII datastore.

Close           - This will close the connection to the datastore
GetSTIXObject   - This takes in a STIX ID and modified timestamp and returns a specific STIX object
AddSTIXObject   - This takes in a STIX object and writes it to the datastore
AddTAXIIObject  - This takes in a TAXII object and writes it to the datastore
GetBundle       - This will return a STIX Bundle object based on the query parameters provided
GetManifestData - This will return a TAXII Manifest resource based on the query parameters provide
*/
type Datastorer interface {
	Close() error
	GetSTIXObject(stixid, version string) (interface{}, error)
	AddSTIXObject(obj interface{}) error
	AddTAXIIObject(obj interface{}) error
	GetBundle(query resources.CollectionQueryType) (*resources.CollectionQueryResultType, error)
	GetManifestData(query resources.CollectionQueryType) (*resources.CollectionQueryResultType, error)
}

/*
CacheType - This struct will hold a cache of various database elements
that will be loaded during initialization and updated along the way.
*/
type CacheType struct {
	BaseObjectIDIndex int
	Collections       map[string]*resources.CollectionType
}
