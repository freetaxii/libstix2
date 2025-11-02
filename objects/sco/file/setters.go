// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package file

import (
	"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
// Public Methods - File - Setters
// ----------------------------------------------------------------------

/*
SetSize - This method takes in an integer representing the size of the file in
bytes and updates the size property.
*/
func (o *File) SetSize(i int64) error {
	o.Size = i
	return nil
}

/*
SetNameEnc - This method takes in a string value representing the name encoding
and updates the name_enc property.
*/
func (o *File) SetNameEnc(s string) error {
	o.NameEnc = s
	return nil
}

/*
SetMagicNumberHex - This method takes in a string value representing the magic
number in hexadecimal and updates the magic_number_hex property.
*/
func (o *File) SetMagicNumberHex(s string) error {
	o.MagicNumberHex = s
	return nil
}

/*
SetMimeType - This method takes in a string value representing the MIME type and
updates the mime_type property.
*/
func (o *File) SetMimeType(s string) error {
	o.MimeType = s
	return nil
}

/*
SetCtime - This method takes in a timestamp in either time.Time or string format
and updates the ctime (creation time) property.
*/
func (o *File) SetCtime(t interface{}) error {
	ts, _ := objects.TimeToString(t, "micro")
	o.Ctime = ts
	return nil
}

/*
SetMtime - This method takes in a timestamp in either time.Time or string format
and updates the mtime (modification time) property.
*/
func (o *File) SetMtime(t interface{}) error {
	ts, _ := objects.TimeToString(t, "micro")
	o.Mtime = ts
	return nil
}

/*
SetAtime - This method takes in a timestamp in either time.Time or string format
and updates the atime (access time) property.
*/
func (o *File) SetAtime(t interface{}) error {
	ts, _ := objects.TimeToString(t, "micro")
	o.Atime = ts
	return nil
}

/*
SetParentDirectoryRef - This method takes in a string value representing a
reference to a directory object and updates the parent_directory_ref property.
*/
func (o *File) SetParentDirectoryRef(s string) error {
	o.ParentDirectoryRef = s
	return nil
}

/*
AddContainsRefs - This method takes in a string value, a comma separated list of
string values, or a slice of string values that represents a reference to other
objects contained in the file and adds it to the contains_refs property.
*/
func (o *File) AddContainsRefs(values interface{}) error {
	return objects.AddValuesToList(&o.ContainsRefs, values)
}

/*
SetContentRef - This method takes in a string value representing a reference to
an artifact object and updates the content_ref property.
*/
func (o *File) SetContentRef(s string) error {
	o.ContentRef = s
	return nil
}

/*
AddHash - This method takes in two parameters and adds the hash to the map. The
first is a string value representing a hash type from the STIX hashing-algorithm-ov
vocabulary. The second is a string value representing the actual hash of the file.
*/
func (o *File) AddHash(k, v string) error {
	if o.Hashes == nil {
		o.Hashes = make(map[string]string, 0)
	}
	o.Hashes[k] = v
	return nil
}
