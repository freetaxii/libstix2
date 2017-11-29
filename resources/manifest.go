// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package resources

import (
	"github.com/freetaxii/libstix2/resources/properties"
	"strings"
)

// ----------------------------------------------------------------------
// Define Message Type
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

// ----------------------------------------------------------------------
// Public Methods - ManifestType
// ----------------------------------------------------------------------

// AddManifestEntry - This method takes in an object that represents a manifest entry
// and adds it to the list in the objects property and returns an integer of
// the location in the slice where the manifest entry object was added. This method
// would be used if the manifest entry was created separately and it just needs to be
// added in whole to the manifest list.
func (ezt *ManifestType) AddManifestEntry(o ManifestEntryType) int {
	ezt.initManifestProperty()
	positionThatAppendWillUse := len(ezt.Objects)
	ezt.Objects = append(ezt.Objects, o)
	return positionThatAppendWillUse
}

// NewManifestEntry - This method is used to create a manifest entry and automatically
// add it to the manifest array. It returns a resources.ManifestEntryType which
// is a pointer to the actual manifest entry that was created in the manifest
// slice.
func (ezt *ManifestType) NewManifestEntry() *ManifestEntryType {
	ezt.initManifestProperty()
	o := NewManifestEntry()
	positionThatAppendWillUse := len(ezt.Objects)
	ezt.Objects = append(ezt.Objects, o)
	return &ezt.Objects[positionThatAppendWillUse]
}

func (ezt *ManifestType) CreateManifestEntry(id, date, ver, media string) {
	m := ezt.NewManifestEntry()
	m.SetID(id)
	m.SetDateAdded(date)

	versions := strings.Split(ver, ",")
	for _, v := range versions {
		m.AddVersion(v)
	}

	mediatypes := strings.Split(media, ",")
	for i, mt := range mediatypes {

		if i > 0 && mt == mediatypes[i-1] {
			continue
		}
		m.AddMediaType(mt)
	}
}

// ----------------------------------------------------------------------
// Private Methods - ManifestType
// ----------------------------------------------------------------------

// initManifestProperty - This method will initialize the Manifest
// slice if it has not already been initialized.
func (ezt *ManifestType) initManifestProperty() {
	if ezt.Objects == nil {
		a := make([]ManifestEntryType, 0)
		ezt.Objects = a
	}
}

/*
ManifestEntryType - This type implements the TAXII 2 Manifest Entry Type and
defines all of the properties and methods needed to create and work with the TAXII
Manifest Entry.

The following information comes directly from the TAXII 2 specification documents.

The manifest-entry type captures metadata about a single object, indicated by
the id property. The metadata includes information such as when the object was
added to the Collection, what versions of the object are available, and what
media types the object is available in.
*/
type ManifestEntryType struct {
	properties.IDPropertyType
	DateAdded  string   `json:"date_added,omitempty"`
	Versions   []string `json:"versions,omitempty"`
	MediaTypes []string `json:"media_types,omitempty"`
}

// SetDateAdded - This method will add the date added to the manifest entry
func (ezt *ManifestEntryType) SetDateAdded(s string) {
	ezt.DateAdded = s
}

// AddVersion - This method takes in a string value that represents an object
// version and adds it to the list of versions that are available for this object.
func (ezt *ManifestEntryType) AddVersion(s string) {
	if ezt.Versions == nil {
		a := make([]string, 0)
		ezt.Versions = a
	}
	ezt.Versions = append(ezt.Versions, s)
}

// AddMediaType - This method takes in a string value that represents a version
// of the STIX specification that is supported and adds it to the list in media types
// that this object is available in.
func (ezt *ManifestEntryType) AddMediaType(s string) {
	if ezt.MediaTypes == nil {
		a := make([]string, 0)
		ezt.MediaTypes = a
	}
	ezt.MediaTypes = append(ezt.MediaTypes, s)
}
