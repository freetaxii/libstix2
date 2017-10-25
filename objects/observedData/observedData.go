// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package observedData

import (
	"github.com/freetaxii/libstix2/objects/common/properties"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
ObservedDataType defines all of the properties associated with the STIX Observed
Data SDO. All of the methods not defined local to this type are inherited from
the individual properties.
*/
type ObservedDataType struct {
	properties.CommonObjectPropertiesType
	FirstObserved  string `json:"first_observed,omitempty"`
	LastObserved   string `json:"last_observed,omitempty"`
	NumberObserved int    `json:"number_observed,omitempty"`
	Objects        string `json:"objects,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// New - This function will create a new observed data object.
func New(ver string) ObservedDataType {
	var obj ObservedDataType
	obj.InitNewObject("observed-data", ver)
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - ObservedDataType
// ----------------------------------------------------------------------

// SetFirstObserved - This method takes in a timestamp in either time.Time or
// string format and updates the first observed property.
func (p *ObservedDataType) SetFirstObserved(t interface{}) {
	ts := p.VerifyTimestamp(t)
	p.FirstObserved = ts
}

// SetLastObserved - This method takes in a timestamp in either time.Time or
// string format and updates the last observed property.
func (p *ObservedDataType) SetLastObserved(t interface{}) {
	ts := p.VerifyTimestamp(t)
	p.LastObserved = ts
}

// SetNumberObserved - This method takes in an integer that represents the
// number of objects that were observed and updates the number observed property.
func (p *ObservedDataType) SetNumberObserved(i int) {
	p.NumberObserved = i
}

// SetObjects - This takes in a string value that represents represents a cyber
// observable JSON object and updates the objects property.
func (p *ObservedDataType) SetObjects(s string) {
	p.Objects = s
}
