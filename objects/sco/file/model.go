// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package file

import (
	"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
File - This type implements the STIX 2 File SCO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties.
*/
type File struct {
	objects.CommonObjectProperties
	Extensions map[string]string `json:"extensions,omitempty" bson:"extensions,omitempty"`
	Hashes     map[string]string `json:"hashes,omitempty" bson:"hashes,omitempty"`
	Size       int               `json:"size,omitempty" bson:"size,omitempty"`
	objects.NameProperty
	NameEnc            string   `json:"name_enc,omitempty" bson:"name_enc,omitempty"`
	MagicNumberHex     string   `json:"magic_number_hex,omitempty" bson:"magic_number_hex,omitempty"`
	MimeType           string   `json:"mime_type,omitempty" bson:"mime_type,omitempty"`
	Ctime              string   `json:"ctime,omitempty" bson:"ctime,omitempty"`
	Mtime              string   `json:"mtime,omitempty" bson:"mtime,omitempty"`
	Atime              string   `json:"atime,omitempty" bson:"atime,omitempty"`
	ParentDirectoryRef string   `json:"parent_directory_ref,omitempty" bson:"parent_directory_ref,omitempty"`
	ContainsRef        []string `json:"contains_ref,omitempty" bson:"contains_ref,omitempty"`
	ContentRef         string   `json:"content_ref,omitempty" bson:"content_ref,omitempty"`
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *File) GetPropertyList() []string {
	return []string{
		"extensions",
		"hashes",
		"size",
		"name",
		"name_enc",
		"magic_number_hex",
		"mime_type",
		"ctime",
		"mtime",
		"atime",
		"parent_directory_ref",
		"contains_ref",
		"content_ref",
	}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX File SCO and return it as a
pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *File {
	var obj File
	obj.InitSCO("file")
	return &obj
}
