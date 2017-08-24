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

type CreatedPropertyType struct {
	Created string `json:"created,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - CreatedPropertyType
// ----------------------------------------------------------------------

// SetCreatedToCurrentTime sets the object created time to the current time
func (this *CreatedPropertyType) SetCreatedToCurrentTime() {
	this.Created = timestamp.GetCurrentTime()
}

// SetCreated takes in one parameter
// param: t - a timestamp in either time.Time or string format
func (this *CreatedPropertyType) SetCreated(t interface{}) {
	ts := timestamp.Verify(t)
	this.Created = ts
}

func (this *CreatedPropertyType) GetCreated() string {
	return this.Created
}
