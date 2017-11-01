// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package collections

import (
	"github.com/freetaxii/libstix2/resources/collection"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
CollectionsType defines all of the properties associated with the TAXII
Collections Resource. All of the methods not defined local to this type are inherited
from the individual properties.
*/
type CollectionsType struct {
	Collections []collection.CollectionType `json:"collections,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// New - This function will create and return a new TAXII collections object.
func New() CollectionsType {
	var obj CollectionsType
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - CollectionsType
// ----------------------------------------------------------------------

// AddCollection - This method takes in an object that represents a collection
// and adds it to the list in the collections property and returns an integer of
// the location in the slice where the collection object was added. This method
// would be used if the collection was created separtely and it just needs to be
// added in whole to the collections list.
func (p *CollectionsType) AddCollection(o collection.CollectionType) int {
	p.initCollectionsProperty()
	positionThatAppendWillUse := len(p.Collections)
	p.Collections = append(p.Collections, o)
	return positionThatAppendWillUse
}

// NewCollection - This method is used to create a collection and automatically
// add it to the collections array. It returns a collection.CollectionType which
// is a pointer to the actual Collection that was created in the collections
// slice.
func (p *CollectionsType) NewCollection() *collection.CollectionType {
	p.initCollectionsProperty()
	o := collection.New()
	positionThatAppendWillUse := len(p.Collections)
	p.Collections = append(p.Collections, o)
	return &p.Collections[positionThatAppendWillUse]
}

// ----------------------------------------------------------------------
// Private Methods - CollectionsType
// ----------------------------------------------------------------------

// initCollectionsProperty - This method will initialize the Collections
// slice if it has not already been initialized.
func (p *CollectionsType) initCollectionsProperty() {
	if p.Collections == nil {
		a := make([]collection.CollectionType, 0)
		p.Collections = a
	}
}
