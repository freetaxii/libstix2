// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package observeddata

import "github.com/freetaxii/libstix2/objects"

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/*
SetFirstObservedToCurrentTime - This methods sets the first observed time to
the current time
*/
func (o *ObservedData) SetFirstObservedToCurrentTime() error {
	o.FirstObserved = objects.GetCurrentTime("micro")
	return nil
}

/*
SetFirstObserved - This method takes in a timestamp in either time.Time or
string format and updates the first observed property.
*/
func (o *ObservedData) SetFirstObserved(t interface{}) error {
	ts, _ := objects.TimeToString(t, "micro")
	o.FirstObserved = ts
	return nil
}

/*
SetLastObservedToCurrentTime - This methods sets the last observed time to
the current time
*/
func (o *ObservedData) SetLastObservedToCurrentTime() error {
	o.LastObserved = objects.GetCurrentTime("micro")
	return nil
}

/*
SetLastObserved - This method takes in a timestamp in either time.Time or
string format and updates the last observed property.
*/
func (o *ObservedData) SetLastObserved(t interface{}) error {
	ts, _ := objects.TimeToString(t, "micro")
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
