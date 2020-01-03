// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package observeddata

import (
	"encoding/json"

	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/timestamp"
)

// ----------------------------------------------------------------------
// Define Object Type
// ----------------------------------------------------------------------

/* ObservedData - This type implements the STIX 2 Observed Data SDO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties. */
type ObservedData struct {
	objects.CommonObjectProperties
	FirstObserved  string `json:"first_observed,omitempty"`
	LastObserved   string `json:"last_observed,omitempty"`
	NumberObserved int    `json:"number_observed,omitempty"`
	Objects        string `json:"objects,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/* New - This function will create a new STIX Observed Data object and return
it as a pointer. It will also initialize the object by setting all of the basic
properties. */
func New() *ObservedData {
	var obj ObservedData
	obj.InitSDO("observed-data")
	return &obj
}

// ----------------------------------------------------------------------
// Public Methods - Observed Data - Core Functionality
// ----------------------------------------------------------------------

/*
Decode - This function will decode some JSON data encoded as a slice of bytes
into an actual struct. It will return the object as a pointer and any errors found.
*/
func Decode(data []byte) (*ObservedData, error) {
	var o ObservedData
	err := json.Unmarshal(data, &o)
	if err != nil {
		return nil, err
	}

	if valid, err := o.Valid(); valid != true {
		return nil, err
	}

	o.SetRawData(data)
	return &o, nil
}

/*
Encode - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *ObservedData) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
EncodeToString - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *ObservedData) EncodeToString() (string, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

/*
Valid - This method will verify all of the properties on the object.
*/
func (o *ObservedData) Valid() (bool, error) {

	// Check common base properties first
	if valid, err := o.CommonObjectProperties.Valid(); valid != true {
		return false, err
	}

	return true, nil
}

// ----------------------------------------------------------------------
// Public Methods - ObservedData
// ----------------------------------------------------------------------

/*
SetFirstObservedToCurrentTime - This methods sets the first observed time to the
current time
*/
func (o *ObservedData) SetFirstObservedToCurrentTime() error {
	o.FirstObserved = timestamp.CurrentTime("micro")
	return nil
}

/*
SetFirstObserved - This method takes in a timestamp in either time.Time or
string format and updates the first observed property.
*/
func (o *ObservedData) SetFirstObserved(t interface{}) error {
	ts, _ := timestamp.ToString(t, "micro")
	o.FirstObserved = ts
	return nil
}

/*
SetLastObservedToCurrentTime - This methods sets the last observed time to the
current time
*/
func (o *ObservedData) SetLastObservedToCurrentTime() error {
	o.LastObserved = timestamp.CurrentTime("micro")
	return nil
}

/*
SetLastObserved - This method takes in a timestamp in either time.Time or
string format and updates the last observed property.
*/
func (o *ObservedData) SetLastObserved(t interface{}) error {
	ts, _ := timestamp.ToString(t, "micro")
	o.LastObserved = ts
	return nil
}

/*
SetNumberObserved - This method takes in an integer that represents the
number of objects that were observed and updates the number observed property.
*/
func (o *ObservedData) SetNumberObserved(i int) error {
	o.NumberObserved = i
	return nil
}

/*
SetObjects - This takes in a string value that represents represents a cyber
observable JSON object and updates the objects property.
*/
func (o *ObservedData) SetObjects(s string) error {
	o.Objects = s
	return nil
}
