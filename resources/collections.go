// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package resources

import (
	"fmt"

	"github.com/freetaxii/libstix2/objects/bundle"
	"github.com/freetaxii/libstix2/resources/properties"
)

// ----------------------------------------------------------------------
//
// Define Message Type
//
// ----------------------------------------------------------------------

/*
Collections - This type implements the TAXII 2 Collections Resource and defines
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
type Collections struct {
	Collections []Collection `json:"collections,omitempty"`
}

/*
Collection - This type implements the TAXII 2 Collection Resource and defines
all of the properties and methods needed to create and work with the TAXII
Collection Resource. All of the methods not defined local to this type are
inherited from the individual properties.

DatastireID = A unique integer that represents this collection
DateAdded   = The date that this collection was added to the system
Enabled     = Is this collection currently enabled
Hidden      = Is this collection currently hidden for the directory listing
Size        = The current size of the collection
ID 		    = The collection ID, a UUIDv4 value
Title 	    = The title of this collection
Description = A long description about this collection
CanRead     = A boolean flag that indicates if one can read from this collection
CanWrite    = A boolean flag that indicates if one can write to this collection
MediaTypes  = A slice of strings of the media types that are found in this collection

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
type Collection struct {
	DatastoreID int    `json:"-"`
	DateAdded   string `json:"-"`
	Enabled     bool   `json:"-"`
	Hidden      bool   `json:"-"`
	Size        int    `json:"-"`
	properties.IDProperty
	properties.TitleProperty
	properties.DescriptionProperty
	CanRead    bool     `json:"can_read"`
	CanWrite   bool     `json:"can_write"`
	MediaTypes []string `json:"media_types,omitempty"`
}

/*
CollectionRecord - This type will hold the data for adding an object to
a collection and is stored in the t_collection_data database table.

CollectionID - The collection ID, a UUIDv4 value.  The database will store a DatastoreID integer.
STIXID       - The STIX object ID. This is done so all versions of an object show up in the collection.
*/
type CollectionRecord struct {
	CollectionID string
	STIXID       string
}

/*
CollectionQuery - This struct will hold all of the variables that a user can
use to query a collection.
*/
type CollectionQuery struct {
	CollectionID          string
	CollectionDatastoreID int
	STIXID                []string // Passed in from the URL
	STIXType              []string // Passed in from the URL
	STIXVersion           []string // Passed in from the URL
	AddedAfter            []string // Passed in from the URL
	AddedBefore           []string // Passed in from the URL
	Limit                 []string // Passed in from the URL
	ServerRecordLimit     int      // Server defined value in the configuration file
	// RangeBegin            int      // Passed in from Range Headers
	// RangeEnd              int      // Passed in from Range Headers
}

/*
CollectionQueryResult - This struct contains the various bits of meta data
that are returned from a query against a collection on a TAXII server. This is
done so that the method signatures do not need to change as time goes on and we
add more meta data that needs to be returned. It is important to note that a
collection may have more entries than the server or client wants to transmit. So
it is important to keep track of which records are actually being delivered to
the client.

Size           = The total size of the dataset returned from the database query.
DateAddedFirst = The added date of the first record being sent to the client.
DateAddedLast  = The added date of the last record being sent to the client.
BundleData     = The STIX bundle that contains the requested data from the collection.
ManifestData   = The TAXII manifest resource that contains the requested data from the collection.
RangeBegin     = The range value of the first record being sent to the client.
RangeEnd       = The range value of the last record being sent to the client.
*/
type CollectionQueryResult struct {
	Size           int
	DateAddedFirst string
	DateAddedLast  string
	BundleData     bundle.Bundle
	ManifestData   Manifest
	// RangeBegin     int
	// RangeEnd       int
}

// ----------------------------------------------------------------------
//
// Initialization Functions
//
// ----------------------------------------------------------------------

/*
NewCollections - This function will create a new TAXII Collections object and return
it as a pointer.
*/
func NewCollections() *Collections {
	var obj Collections
	return &obj
}

/*
NewCollection - This function will create a new TAXII Collection object and return
it as a pointer.
*/
func NewCollection() *Collection {
	var obj Collection
	return &obj
}

/*
CreateCollectionRecord - This function will take in a collection ID and a STIX ID
and create a new TAXII Collection Record object and return it as a pointer. This
is used for storying a record in the database in the t_collection_data table.
*/
func CreateCollectionRecord(cid, sid string) *CollectionRecord {
	var obj CollectionRecord
	obj.CollectionID = cid
	obj.STIXID = sid
	return &obj
}

/*
NewCollectionQuery - This function will take in a collection ID as a string
and the Server Record Limit and return a CollectionQueryType object.
*/
func NewCollectionQuery(id string, limit int) *CollectionQuery {
	var obj CollectionQuery
	obj.CollectionID = id
	obj.ServerRecordLimit = limit
	return &obj
}

// ----------------------------------------------------------------------
//
// Public Methods - Collections
//
// ----------------------------------------------------------------------

/*
AddCollection - This method takes in an object that represents a collection
and adds it to the list in the collections property and returns an integer of
the location in the slice where the collection object was added. This method
would be used if the collection was created separately and it just needs to be
added in whole to the collections list.
*/
func (r *Collections) AddCollection(o *Collection) (int, error) {
	//r.initCollectionsProperty()
	positionThatAppendWillUse := len(r.Collections)
	r.Collections = append(r.Collections, *o)
	return positionThatAppendWillUse, nil
}

/*
NewCollection - This method is used to create a collection and automatically
add it to the collections array. It returns a resources.Collection which
is a pointer to the actual Collection that was created in the collections
slice.
*/
func (r *Collections) NewCollection() (*Collection, error) {
	//r.initCollectionsProperty()
	o := NewCollection()
	positionThatAppendWillUse := len(r.Collections)
	r.Collections = append(r.Collections, *o)
	return &r.Collections[positionThatAppendWillUse], nil
}

// ----------------------------------------------------------------------
// Private Methods - Collections
// ----------------------------------------------------------------------

/*
initCollectionsProperty - This method will initialize the Collections
slice if it has not already been initialized.
*/
// func (r *Collections) initCollectionsProperty() error {
// 	if r.Collections == nil {
// 		a := make([]Collection, 0)
// 		r.Collections = a
// 	}
// 	return nil
// }

// ----------------------------------------------------------------------
// Public Methods - Collection
// ----------------------------------------------------------------------

/*
SetEnabled - This method will set the collection to be enabled.
*/
func (r *Collection) SetEnabled() error {
	r.Enabled = true
	return nil
}

/*
SetDisabled - This method will set the collection to be disabled.
*/
func (r *Collection) SetDisabled() error {
	r.Enabled = false
	return nil
}

/*
SetHidden - This method will set the collection to be hidden.
*/
func (r *Collection) SetHidden() error {
	r.Hidden = true
	return nil
}

/*
SetVisible - This method will set the collection to be visible.
*/
func (r *Collection) SetVisible() error {
	r.Hidden = false
	return nil
}

/*
SetCanRead - This method will set the can_read boolean to true.
*/
func (r *Collection) SetCanRead() error {
	r.CanRead = true
	return nil
}

/*
GetCanRead - This method will return the value of Can Read.
*/
func (r *Collection) GetCanRead() bool {
	return r.CanRead
}

/*
SetCanWrite - This method will set the can_write boolean to true.
*/
func (r *Collection) SetCanWrite() error {
	r.CanWrite = true
	return nil
}

/*
GetCanWrite - This method will return the value of Can Write.
*/
func (r *Collection) GetCanWrite() bool {
	return r.CanWrite
}

/*
AddMediaType - This method takes in a string value that represents a version
of the TAXII api that is supported and adds it to the list in media types
property.
*/
func (r *Collection) AddMediaType(s string) error {
	if r.MediaTypes == nil {
		a := make([]string, 0)
		r.MediaTypes = a
	}
	r.MediaTypes = append(r.MediaTypes, s)
	return nil
}

/*
Compare - This method will compare two collections to make sure they
are the same. The collection receiver is the master and represent the correct
data, the indicator passed in as i represents the one we need to test.
*/
func (r *Collection) Compare(c *Collection) (bool, int, []string) {
	problemsFound := 0
	details := make([]string, 0)

	// Check ID Value
	if c.ID != r.ID {
		problemsFound++
		str := fmt.Sprintf("-- IDs Do Not Match: %s | %s", r.ID, c.ID)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ IDs Match: %s | %s", r.ID, c.ID)
		details = append(details, str)
	}

	// Check Title Value
	if c.Title != r.Title {
		problemsFound++
		str := fmt.Sprintf("-- Titles Do Not Match: %s | %s", r.Title, c.Title)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Titles Match: %s | %s", r.Title, c.Title)
		details = append(details, str)
	}

	// Check Description Value
	if c.Description != r.Description {
		problemsFound++
		str := fmt.Sprintf("-- Descriptions Do Not Match: %s | %s", r.Description, c.Description)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Descriptions Match: %s | %s", r.Description, c.Description)
		details = append(details, str)
	}

	// Check Can Read Value
	if c.CanRead != r.CanRead {
		problemsFound++
		str := fmt.Sprintf("-- Can Read Values Do Not Match: %t | %t", r.CanRead, c.CanRead)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Can Read Values Match: %t | %t", r.CanRead, c.CanRead)
		details = append(details, str)
	}

	// Check Can Write Value
	if c.CanWrite != r.CanWrite {
		problemsFound++
		str := fmt.Sprintf("-- Can Write Values Do Not Match: %t | %t", r.CanWrite, c.CanWrite)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Can Write Values Match: %t | %t", r.CanWrite, c.CanWrite)
		details = append(details, str)
	}

	// Check Media Type Property Length
	if len(c.MediaTypes) != len(r.MediaTypes) {
		problemsFound++
		str := fmt.Sprintf("-- Media Type Lengths Do Not Match: %d | %d", r.MediaTypes, c.MediaTypes)
		details = append(details, str)
	} else {
		str := fmt.Sprintf("++ Media Type Lengths Match: %d | %d", r.MediaTypes, c.MediaTypes)
		details = append(details, str)

		// If lengths are the same, then check each value
		for index, _ := range r.MediaTypes {
			if c.MediaTypes[index] != r.MediaTypes[index] {
				problemsFound++
				str := fmt.Sprintf("-- Media Types Do Not Match: %s | %s", r.MediaTypes[index], c.MediaTypes[index])
				details = append(details, str)
			} else {
				str := fmt.Sprintf("++ Media Types Match: %s | %s", r.MediaTypes[index], c.MediaTypes[index])
				details = append(details, str)
			}
		}
	}

	if problemsFound > 0 {
		return false, problemsFound, details
	}

	return true, 0, details
}
