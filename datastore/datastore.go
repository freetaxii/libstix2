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
	GetObject(stixid, version string) (interface{}, error)
	GetObjectsFromCollection(query QueryType) (*objects.BundleType, *QueryReturnDataType, error)
	GetListOfObjectsFromCollection(query QueryType) (*[]CollectionRawDataType, *QueryReturnDataType, error)
	GetManifestFromCollection(query QueryType) (*resources.ManifestType, *QueryReturnDataType, error)
}

// QueryType - This struct will hold all of the variables that a user can
// query a collection on.
type QueryType struct {
	CollectionID string
	STIXID       []string
	STIXType     []string
	STIXVersion  []string
	AddedAfter   []string
	RangeBegin   int
	RangeEnd     int
	RangeMax     int
}

// QueryReturnDataType - This struct contains the various bits of meta data that
// are returned from a query against a TAXII server. This is done so that the
// method signatures do not need to change as time goes on and we add more meta
// data that needs to be returned.
type QueryReturnDataType struct {
	Size           int
	RangeBegin     int
	RangeEnd       int
	DateAddedFirst string
	DateAddedLast  string
}

// CollectionRawDataType -
type CollectionRawDataType struct {
	STIXID      string
	DateAdded   string
	STIXVersion string
	SpecVersion string
}
