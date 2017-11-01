// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package resources

import (
	"github.com/freetaxii/libstix2/resources/apiRoot"
	"github.com/freetaxii/libstix2/resources/collection"
	"github.com/freetaxii/libstix2/resources/collections"
	"github.com/freetaxii/libstix2/resources/discovery"
	// "github.com/freetaxii/libstix2/resources/manifest"
	// "github.com/freetaxii/libstix2/resources/status"
)

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// NewAPIRoot - This function will create a new TAXII API Root object
//
// return: api_root.ApiRootType
func NewAPIRoot() apiRoot.APIRootType {
	return apiRoot.New()
}

// NewCollection - This function will create a new TAXII Collection object
//
// return: collection.CollectionType
func NewCollection() collection.CollectionType {
	return collection.New()
}

// NewCollections - This function will create a new TAXII Collections object
//
// return: collections.CollectionsType
func NewCollections() collections.CollectionsType {
	return collections.New()
}

// NewDiscovery - This function will create a new TAXII Discovery object
//
// return: discovery.DiscoeryType
func NewDiscovery() discovery.DiscoveryType {
	return discovery.New()
}

// NewManifest - This function will create a new TAXII Manifest object
//
// return: manifest.ManifestType
// func NewManifest() manifest.ManifestType {
// 	return manifest.New()
// }

// NewStatus - This function will create a new TAXII Status object
//
// return: status.StatusType
// func NewStatus() status.StatusType {
// 	return status.New()
// }
