// Copyright 2015-2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package datastore

import "github.com/freetaxii/libstix2/resources/collections"

/*
Datastorer - This interface enables access to the STIX/TAXII datastore.

Close           - This will close the connection to the datastore
GetObject       - This takes in a STIX ID and modified timestamp and returns a specific STIX object
AddObject       - This takes in a STIX object and writes it to the datastore
AddTAXIIObject  - This takes in a TAXII object and writes it to the datastore
GetBundle       - This will return a STIX Bundle object based on the query parameters provided
GetManifestData - This will return a TAXII Manifest resource based on the query parameters provide
*/
type Datastorer interface {
	Close() error
	GetObject(id, version string) (interface{}, error)
	AddObject(obj interface{}) error
	AddTAXIIObject(obj interface{}) error
	AddToCollection(collectionid, stixid string) error
	GetBundle(query collections.CollectionQuery) (*collections.CollectionQueryResult, error)
	GetManifestData(query collections.CollectionQuery) (*collections.CollectionQueryResult, error)
}
