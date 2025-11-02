// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package directory

import (
	"errors"
)

// ----------------------------------------------------------------------
// Public Methods - Directory - Setters
// ----------------------------------------------------------------------

/*
SetPath - This method will set the path for the directory object.
The path field is required for this object.
*/
func (o *Directory) SetPath(s string) error {
	if s == "" {
		return errors.New("the directory path cannot be empty")
	}
	o.Path = s
	return nil
}

/*
SetPathEnc - This method will set the encoded path for the directory object.
*/
func (o *Directory) SetPathEnc(s string) error {
	o.PathEnc = s
	return nil
}

/*
SetCtime - This method will set the creation time for the directory object.
*/
func (o *Directory) SetCtime(t interface{}) error {
	switch v := t.(type) {
	case string:
		o.Ctime = v
	case nil:
		o.Ctime = ""
	default:
		return errors.New("ctime must be a string or nil")
	}
	return nil
}

/*
SetMtime - This method will set the modification time for the directory object.
*/
func (o *Directory) SetMtime(t interface{}) error {
	switch v := t.(type) {
	case string:
		o.Mtime = v
	case nil:
		o.Mtime = ""
	default:
		return errors.New("mtime must be a string or nil")
	}
	return nil
}

/*
SetAtime - This method will set the access time for the directory object.
*/
func (o *Directory) SetAtime(t interface{}) error {
	switch v := t.(type) {
	case string:
		o.Atime = v
	case nil:
		o.Atime = ""
	default:
		return errors.New("atime must be a string or nil")
	}
	return nil
}

/*
AddContainsRef - This method will add an object reference to the contains_refs list.
*/
func (o *Directory) AddContainsRef(ref string) error {
	if ref != "" {
		o.ContainsRefs = append(o.ContainsRefs, ref)
	}
	return nil
}

/*
SetContainsRefs - This method will set the contains_refs list for the directory object.
*/
func (o *Directory) SetContainsRefs(refs []string) error {
	o.ContainsRefs = refs
	return nil
}
