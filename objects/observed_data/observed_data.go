// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package observed_data

import (
	"github.com/freetaxii/libstix2/objects/common"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type ObservedDataType struct {
	common.CommonObjectPropertiesType
	First_observed  string `json:"first_observed,omitempty"`
	Last_observed   string `json:"last_observed,omitempty"`
	Number_observed int    `json:"number_observed,omitempty"`
	Objects         string `json:"objects,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() ObservedDataType {
	var obj ObservedDataType
	obj.InitNewObject("observed-data")
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - ObservedDataType
// ----------------------------------------------------------------------

// SetFirstObserved takes in one parameter
// param: t - a timestamp in either time.Time or string format
func (this *ObservedDataType) SetFirstObserved(t interface{}) {
	ts := this.VerifyTimestamp(t)
	this.First_observed = ts
}

// SetLastObserved takes in one parameter
// param: t - a timestamp in either time.Time or string format
func (this *ObservedDataType) SetLastObserved(t interface{}) {
	ts := this.VerifyTimestamp(t)
	this.Last_observed = ts
}

// SetNumberObserved takes in one parameter
// param: i - an integer that represents the number of objects that were observed
func (this *ObservedDataType) SetNumberObserved(i int) {
	this.Number_observed = i
}

// SetObjects takes in one parameter and returns
// param: s - an string value that represents represents a cyber observable JSON object
func (this *ObservedDataType) SetObjects(s string) {
	this.Objects = s
}
