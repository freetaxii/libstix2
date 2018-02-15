// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package resources

import (
	"github.com/freetaxii/libstix2/resources/properties"
)

// ----------------------------------------------------------------------
//
// Define Message Type
//
// ----------------------------------------------------------------------

/*
ManifestType - This type implements the TAXII 2 Manifest Resource and defines
all of the properties and methods needed to create and work with the TAXII Manifest
Resource.

The following information comes directly from the TAXII 2 specification documents.

This Endpoint retrieves a manifest about objects from a Collection. It supports
filtering and pagination identical to the get objects Endpoint (see section 5.3)
but rather than returning the object itself it returns metadata about the object.
It can be used to retrieve metadata to decide whether it's worth retrieving the
actual objects.

If the Collection specifies can_read as false, this Endpoint SHOULD return a
HTTP 403 error.

This Endpoint supports filtering, which is applied against the source object
rather than the manifest entry for an object. Thus, searching the manifest for a
type of indicator will return the manifest entries for objects with a type of
indicator, even though the manifest doesn't have a type field.
*/
type ManifestType struct {
	Objects []ManifestEntryType `json:"objects,omitempty"`
}

/*
ManifestEntryType - This type implements the TAXII 2 Manifest Entry Type and
defines all of the properties and methods needed to create and work with the TAXII
Manifest Entry.

The following information comes directly from the TAXII 2 specification documents.

The manifest-entry type captures metadata about a single versions of an object,
indicated by the id property. The metadata includes information such as when that
versions of the object was added to the Collection, the version of the object
itself, and the media type that this specific version of the object is available
in.
*/
type ManifestEntryType struct {
	properties.IDPropertyType
	DateAdded string `json:"date_added,omitempty"`
	Version   string `json:"version,omitempty"`
	MediaType string `json:"media_type,omitempty"`
}

// ----------------------------------------------------------------------
//
// Initialization Functions
//
// ----------------------------------------------------------------------

/*
NewManifest - This function will create a new TAXII Manifest object and return
it as a pointer.
*/
func NewManifest() *ManifestType {
	var obj ManifestType
	return &obj
}

/*
NewManifestEntry - This function will create a new TAXII Manifest Entry object
and return it as a pointer.
*/
func NewManifestEntry() *ManifestEntryType {
	var obj ManifestEntryType
	return &obj
}

// ----------------------------------------------------------------------
//
// Public Methods - ManifestType
//
// ----------------------------------------------------------------------

/*
AddManifestEntry - This method takes in an object that represents a manifest
entry and adds it to the list in the objects property and returns an integer of
the location in the slice where the manifest entry object was added. This method
would be used if the manifest entry was created separately and it just needs to
be added in whole to the manifest list.
*/
func (r *ManifestType) AddManifestEntry(o *ManifestEntryType) (int, error) {
	positionThatAppendWillUse := len(r.Objects)
	r.Objects = append(r.Objects, *o)
	return positionThatAppendWillUse, nil
}

/*
NewManifestEntry - This method is used to create a manifest entry and automatically
add it to the objects array. It returns a resources.ManifestEntryType which is a
pointer to the actual manifest entry that was created in the manifest slice.
*/
func (r *ManifestType) NewManifestEntry() (*ManifestEntryType, error) {
	o := NewManifestEntry()
	positionThatAppendWillUse := len(r.Objects)
	r.Objects = append(r.Objects, *o)
	return &r.Objects[positionThatAppendWillUse], nil
}

/*
CreateManifestEntry - This method is used to create and add a manifest entry in
a single step, by taking in all of the values as parameters.
*/
func (r *ManifestType) CreateManifestEntry(id, date, ver, media string) error {
	m, _ := r.NewManifestEntry()
	m.SetID(id)
	m.SetDateAdded(date)
	m.SetVersion(ver)
	m.SetMediaType(media)
	return nil
}

// ----------------------------------------------------------------------
// Public Methods - ManifestEntryType
// ----------------------------------------------------------------------

/*
SetDateAdded - This method will add the date added to the manifest entry
*/
func (r *ManifestEntryType) SetDateAdded(s string) error {
	r.DateAdded = s
	return nil
}

/*
SetVersion - This method will add the version to the manifest entry
*/
func (r *ManifestEntryType) SetVersion(s string) error {
	r.Version = s
	return nil
}

/*
SetMediaType - This method will add the media type to the manifest entry
*/
func (r *ManifestEntryType) SetMediaType(s string) error {
	r.MediaType = s
	return nil
}
