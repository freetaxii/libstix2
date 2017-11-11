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
	ListObjectsInCollection(collectionid string) []string
	GetObjectsInCollection(collectionid string) objects.BundleType
	GetObject(stixid string) (interface{}, error)
	Close() error
	GetEnabledCollections() resources.CollectionsType
}
