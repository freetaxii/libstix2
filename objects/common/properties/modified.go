// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

import (
	"github.com/freetaxii/libstix2/objects/common/timestamp"
)

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

// ModifiedPropertyType - A property used by one or more STIX objects that
// captures the time that this object was modified in STIX timestamp format,
// which is an RFC3339 format.
type ModifiedPropertyType struct {
	Modified string `json:"modified,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - ModifiedPropertyType
// ----------------------------------------------------------------------

// SetModified - This method takes in a timestamp in either time.Time or string
// format and updates the modifed property with it. The value is stored as a
// string, so if the value is in time.Time format, it will be converted to the
// correct STIX timestamp format.
func (p *ModifiedPropertyType) SetModified(t interface{}) {
	ts := timestamp.Verify(t)
	p.Modified = ts
}

// GetModified - This method will return the modified timestamp as a string. If
// the value is the same as the created timestamp, then this object is the
// first version of the object.
func (p *ModifiedPropertyType) GetModified() string {
	return p.Modified
}
