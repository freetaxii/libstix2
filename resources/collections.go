// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package resources

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
CollectionsType - This type implements the TAXII 2 Collections Resource and defines
all of the properties and methods needed to create and work with the TAXII Collections
Resource. All of the methods not defined local to this type are inherited from
the individual properties.

The following information comes directly from the TAXII 2 specification documents.

This Endpoint provides information about the Collections hosted under this API
Root. This is similar to the response to get a Collection (see section 5.2), but
rather than providing information about one Collection it provides information
about all of the Collections. Most importantly, it provides the Collection's id,
which is used to request objects or manifest entries from the Collection.

The collections resource is a simple wrapper around a list of collection
resources.
*/
type CollectionsType struct {
	Collections []CollectionType `json:"collections,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - CollectionsType
// ----------------------------------------------------------------------

// AddCollection - This method takes in an object that represents a collection
// and adds it to the list in the collections property and returns an integer of
// the location in the slice where the collection object was added. This method
// would be used if the collection was created separately and it just needs to be
// added in whole to the collections list.
func (ezt *CollectionsType) AddCollection(o CollectionType) int {
	ezt.initCollectionsProperty()
	positionThatAppendWillUse := len(ezt.Collections)
	ezt.Collections = append(ezt.Collections, o)
	return positionThatAppendWillUse
}

// NewCollection - This method is used to create a collection and automatically
// add it to the collections array. It returns a resources.CollectionType which
// is a pointer to the actual Collection that was created in the collections
// slice.
func (ezt *CollectionsType) NewCollection() *CollectionType {
	ezt.initCollectionsProperty()
	o := NewCollection()
	positionThatAppendWillUse := len(ezt.Collections)
	ezt.Collections = append(ezt.Collections, o)
	return &ezt.Collections[positionThatAppendWillUse]
}

// ----------------------------------------------------------------------
// Private Methods - CollectionsType
// ----------------------------------------------------------------------

// initCollectionsProperty - This method will initialize the Collections
// slice if it has not already been initialized.
func (ezt *CollectionsType) initCollectionsProperty() {
	if ezt.Collections == nil {
		a := make([]CollectionType, 0)
		ezt.Collections = a
	}
}
