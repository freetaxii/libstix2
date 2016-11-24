// Copyright 2016 Bret Jordan, All rights reserved.
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
	common.CommonPropertiesType
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
	obj.MessageType = "observed-data"
	obj.Id = obj.NewId("observed-data")
	obj.Created = obj.GetCurrentTime()
	obj.Modified = obj.Created
	obj.Version = 1
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - ObservedDataType
// ----------------------------------------------------------------------

// SetFirstObserved takes in two parameters and returns and error if there is one
// param: t a timestamp in either time.Time or string format
func (this *ObservedDataType) SetFirstObserved(t interface{}) error {

	ts, err := this.VerifyTimestamp(t)
	if err != nil {
		return err
	}
	this.First_observed = ts

	return nil
}

// SetLastObserved takes in two parameters and returns and error if there is one
// param: t a timestamp in either time.Time or string format
func (this *ObservedDataType) SetLastObserved(t interface{}) error {

	ts, err := this.VerifyTimestamp(t)
	if err != nil {
		return err
	}
	this.Last_observed = ts

	return nil
}

func (this *ObservedDataType) SetNumberObserved(i int) {
	this.Number_observed = i
}

//TODO This will probably need to be changed either to an array or something
func (this *ObservedDataType) SetObjects(s string) {
	this.Objects = s
}
