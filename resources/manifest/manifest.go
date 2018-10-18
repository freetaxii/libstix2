// Copyright 2015-2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package manifest

import (
	"github.com/freetaxii/libstix2/resources/properties"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
Manifest - This type implements the TAXII 2 Manifest Resource and defines
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
type Manifest struct {
	Objects []ManifestRecord `json:"objects,omitempty"`
}

/*
ManifestRecord - This type implements the TAXII 2 Manifest Record Type and
defines all of the properties and methods needed to create and work with the TAXII
Manifest Record.

The following information comes directly from the TAXII 2 specification documents.

The manifest-record type captures metadata about a single versions of an object,
indicated by the id property. The metadata includes information such as when that
versions of the object was added to the Collection, the version of the object
itself, and the media type that this specific version of the object is available
in.
*/
type ManifestRecord struct {
	properties.IDProperty
	DateAdded string `json:"date_added,omitempty"`
	Version   string `json:"version,omitempty"`
	MediaType string `json:"media_type,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new TAXII Manifest object and return it as a
pointer.
*/
func New() *Manifest {
	var obj Manifest
	return &obj
}

/*
NewRecord - This function will create a new TAXII Manifest Record object and
return it as a pointer.
*/
func NewRecord() *ManifestRecord {
	var obj ManifestRecord
	return &obj
}

// ----------------------------------------------------------------------
// Public Methods - Manifest
// ----------------------------------------------------------------------

/*
AddRecord - This method takes in an object that represents a manifest
record and adds it to the list in the objects property and returns an integer of
the location in the slice where the manifest entry object was added. This method
would be used if the manifest entry was created separately and it just needs to
be added in whole to the manifest list.
*/
func (r *Manifest) AddRecord(o *ManifestRecord) (int, error) {
	positionThatAppendWillUse := len(r.Objects)
	r.Objects = append(r.Objects, *o)
	return positionThatAppendWillUse, nil
}

/*
NewRecord - This method is used to create a manifest entry and automatically
add it to the objects array. It returns a resources.ManifestRecord which is a
pointer to the actual manifest entry that was created in the manifest slice.
*/
func (r *Manifest) NewRecord() (*ManifestRecord, error) {
	o := NewRecord()
	positionThatAppendWillUse := len(r.Objects)
	r.Objects = append(r.Objects, *o)
	return &r.Objects[positionThatAppendWillUse], nil
}

/*
CreateRecord - This method is used to create and add a manifest entry in
a single step, by taking in all of the values as parameters.
*/
func (r *Manifest) CreateRecord(id, date, ver, media string) error {
	m, _ := r.NewRecord()
	m.SetID(id)
	m.SetDateAdded(date)
	m.SetVersion(ver)
	m.SetMediaType(media)
	return nil
}

// ----------------------------------------------------------------------
// Public Methods - ManifestRecord
// ----------------------------------------------------------------------

/*
SetDateAdded - This method will add the date added to the manifest entry
*/
func (r *ManifestRecord) SetDateAdded(s string) error {
	r.DateAdded = s
	return nil
}

/*
SetVersion - This method will add the version to the manifest entry
*/
func (r *ManifestRecord) SetVersion(s string) error {
	r.Version = s
	return nil
}

/*
SetMediaType - This method will add the media type to the manifest entry
*/
func (r *ManifestRecord) SetMediaType(s string) error {
	r.MediaType = s
	return nil
}
