// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package resources

import (
	"github.com/freetaxii/libstix2/resources/properties"
	"sort"
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

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
InitManifest - This function will create a new TAXII Manifest object and return
it as a pointer.
*/
func InitManifest() *ManifestType {
	var obj ManifestType
	return &obj
}

/*
InitManifestEntry - This function will create a new TAXII Manifest Entry object
and return it as a pointer.
*/
func InitManifestEntry() *ManifestEntryType {
	var obj ManifestEntryType
	return &obj
}

// ----------------------------------------------------------------------
// Public Methods - ManifestType
// ----------------------------------------------------------------------

/*
AddManifestEntry - This method takes in an object that represents a manifest
entry and adds it to the list in the objects property and returns an integer of
the location in the slice where the manifest entry object was added. This method
would be used if the manifest entry was created separately and it just needs to
be added in whole to the manifest list.
*/
func (ezt *ManifestType) AddManifestEntry(o *ManifestEntryType) (int, error) {
	positionThatAppendWillUse := len(ezt.Objects)
	ezt.Objects = append(ezt.Objects, *o)
	return positionThatAppendWillUse, nil
}

/*
GetNewManifestEntry - This method is used to create a manifest entry and automatically
add it to the objects array. It returns a resources.ManifestEntryType which is a
pointer to the actual manifest entry that was created in the manifest slice.
*/
func (ezt *ManifestType) GetNewManifestEntry() (*ManifestEntryType, error) {
	o := InitManifestEntry()
	positionThatAppendWillUse := len(ezt.Objects)
	ezt.Objects = append(ezt.Objects, *o)
	return &ezt.Objects[positionThatAppendWillUse], nil
}

/*
CreateManifestEntry - This method is used to create and add a manifest entry in
a single step, by taking in all of the values as parameters. Multiple values for
the version and media type properties can be provided as a comma separated list
with no spaces in between the values.
*/
func (ezt *ManifestType) CreateManifestEntry(id, date, ver, media string) error {
	m, _ := ezt.GetNewManifestEntry()
	m.SetID(id)
	m.SetDateAdded(date)

	versions := strings.Split(ver, ",")

	// The specification says that the newest objects should start at index 0 so
	// lets sorts them in reverse order.
	if len(versions) > 1 {
		sort.Sort(sort.Reverse(sort.StringSlice(versions)))
	}

	for _, v := range versions {
		m.AddVersion(v)
	}

	mediatypes := strings.Split(media, ",")
	for i, mt := range mediatypes {

		// If the media types are all the same, due to the way the SQL query
		// returns results, then only record one entry.
		if i > 0 && mt == mediatypes[i-1] {
			continue
		}
		m.AddMediaType(mt)
	}
	return nil
}

// ----------------------------------------------------------------------
// Public Methods - ManifestEntryType
// ----------------------------------------------------------------------

/*
SetDateAdded - This method will add the date added to the manifest entry
*/
func (ezt *ManifestEntryType) SetDateAdded(s string) error {
	ezt.DateAdded = s
	return nil
}

/*
AddVersion - This method takes in a string value that represents an object
version and adds it to the list of versions that are available for this object.
*/
func (ezt *ManifestEntryType) AddVersion(s string) error {
	if ezt.Versions == nil {
		a := make([]string, 0)
		ezt.Versions = a
	}
	ezt.Versions = append(ezt.Versions, s)
	return nil
}

/*
AddMediaType - This method takes in a string value that represents a version of
the STIX specification that is supported and adds it to the list in media types
that this object is available in.
*/
func (ezt *ManifestEntryType) AddMediaType(s string) error {
	if ezt.MediaTypes == nil {
		a := make([]string, 0)
		ezt.MediaTypes = a
	}
	ezt.MediaTypes = append(ezt.MediaTypes, s)
	return nil
}
