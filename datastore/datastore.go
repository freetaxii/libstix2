// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package datastore

import (
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/resources"
)

// Datastorer - This type enables access to the TAXII datastore.
// GetObject returns a STIX object
type Datastorer interface {
	Close() error
	GetEnabledCollections() resources.CollectionsType
	GetObject(stixid string) (interface{}, error)
	GetObjectsInCollection(collectionid string) objects.BundleType
	GetListOfObjectsInCollection(collectionid string) ([]string, error)
}

// QueryType - This struct will hold all of the variables that a user can
// query a collection on.
type QueryType struct {
	CollectionID string
	STIXID       string
	STIXType     string
	STIXVersion  string
	AddedAfter   string
	RangeBegin   int
	RangeEnd     int
	RangeMax     int
}
