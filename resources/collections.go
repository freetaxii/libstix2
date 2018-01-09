// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package resources

import (
	"github.com/freetaxii/libstix2/resources/properties"
)

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

/*
CollectionType - This type implements the TAXII 2 Collection Resource and defines
all of the properties and methods needed to create and work with the TAXII Collection
Resource. All of the methods not defined local to this type are inherited from
the individual properties.

The following information comes directly from the TAXII 2 specification documents.

This Endpoint provides general information about a Collection, which can be used
to help users and clients decide whether and how they want to interact with it.
For example, it will tell clients what it's called and what permissions they
have to it.

The collection resource contains general information about a Collection, such as
its id, a human-readable title and description, an optional list of supported
media_types (representing the media type of objects can be requested from or
added to it), and whether the TAXII Client, as authenticated, can get objects
from the Collection and/or add objects to it.
*/
type CollectionType struct {
	DateAdded string `json:"-"`
	Enabled   bool   `json:"-"`
	Hidden    bool   `json:"-"`
	properties.IDPropertyType
	properties.TitlePropertyType
	properties.DescriptionPropertyType
	CanRead    bool     `json:"can_read"`
	CanWrite   bool     `json:"can_write"`
	MediaTypes []string `json:"media_types,omitempty"`
}

/*
CollectionRecordType - This type will hold the data for adding an object to
a collection and is stored in the t_collection_data database table.
*/
type CollectionRecordType struct {
	CollectionID string
	STIXID       string
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
InitCollections - This function will create a new TAXII Collections object and return
it as a pointer.
*/
func InitCollections() *CollectionsType {
	var obj CollectionsType
	return &obj
}

/*
InitCollection - This function will create a new TAXII Collection object and return
it as a pointer.
*/
func InitCollection() *CollectionType {
	var obj CollectionType
	return &obj
}

/*
InitCollectionRecord - This function will create a new TAXII Collection object and return
it as a pointer.
*/
func InitCollectionRecord() *CollectionRecordType {
	var obj CollectionRecordType
	return &obj
}

// ----------------------------------------------------------------------
// Public Methods - CollectionsType
// ----------------------------------------------------------------------

/*
AddCollection - This method takes in an object that represents a collection
and adds it to the list in the collections property and returns an integer of
the location in the slice where the collection object was added. This method
would be used if the collection was created separately and it just needs to be
added in whole to the collections list.
*/
func (ezt *CollectionsType) AddCollection(o CollectionType) int {
	ezt.initCollectionsProperty()
	positionThatAppendWillUse := len(ezt.Collections)
	ezt.Collections = append(ezt.Collections, o)
	return positionThatAppendWillUse
}

/*
NewCollection - This method is used to create a collection and automatically
add it to the collections array. It returns a resources.CollectionType which
is a pointer to the actual Collection that was created in the collections
slice.
*/
func (ezt *CollectionsType) NewCollection() *CollectionType {
	ezt.initCollectionsProperty()
	o := NewCollection()
	positionThatAppendWillUse := len(ezt.Collections)
	ezt.Collections = append(ezt.Collections, *o)
	return &ezt.Collections[positionThatAppendWillUse]
}

// ----------------------------------------------------------------------
// Private Methods - CollectionsType
// ----------------------------------------------------------------------

/*
initCollectionsProperty - This method will initialize the Collections
slice if it has not already been initialized.
*/
func (ezt *CollectionsType) initCollectionsProperty() {
	if ezt.Collections == nil {
		a := make([]CollectionType, 0)
		ezt.Collections = a
	}
}

// ----------------------------------------------------------------------
// Public Methods - CollectionType
// ----------------------------------------------------------------------

/*
SetEnabled - This method will set the collection to be enabled.
*/
func (ezt *CollectionType) SetEnabled() error {
	ezt.Enabled = true
	return nil
}

/*
SetDisabled - This method will set the collection to be disabled.
*/
func (ezt *CollectionType) SetDisabled() error {
	ezt.Enabled = false
	return nil
}

/*
SetHidden - This method will set the collection to be hidden.
*/
func (ezt *CollectionType) SetHidden() error {
	ezt.Hidden = true
	return nil
}

/*
SetVisible - This method will set the collection to be visible.
*/
func (ezt *CollectionType) SetVisible() error {
	ezt.Hidden = false
	return nil
}

/*
SetCanRead - This method will set the can_read boolean to true.
*/
func (ezt *CollectionType) SetCanRead() {
	ezt.CanRead = true
}

/*
GetCanRead - This method will return the value of Can Read.
*/
func (ezt *CollectionType) GetCanRead() (bool, error) {
	return ezt.CanRead, nil
}

/*
SetCanWrite - This method will set the can_write boolean to true.
*/
func (ezt *CollectionType) SetCanWrite() error {
	ezt.CanWrite = true
	return nil
}

/*
GetCanWrite - This method will return the value of Can Write.
*/
func (ezt *CollectionType) GetCanWrite() (bool, error) {
	return ezt.CanWrite, nil
}

/*
AddMediaType - This method takes in a string value that represents a version
of the TAXII api that is supported and adds it to the list in media types
property.
*/
func (ezt *CollectionType) AddMediaType(s string) error {
	if ezt.MediaTypes == nil {
		a := make([]string, 0)
		ezt.MediaTypes = a
	}
	ezt.MediaTypes = append(ezt.MediaTypes, s)
	return nil
}

// ----------------------------------------------------------------------
// Public Methods - CollectionRecordType
// ----------------------------------------------------------------------

/*
NewCollectionRecord - This function will take in a collection ID and a STIX ID
and return a collection record type which is used for storying a record in
the database in the t_collection_data table.
*/
func NewCollectionRecord(cid, sid string) *CollectionRecordType {
	obj := InitCollectionRecord()
	obj.CollectionID = cid
	obj.STIXID = sid
	return &obj
}
