// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package apiroot

import (
	"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
// Public Methods - APIRoot
// ----------------------------------------------------------------------

/*
AddVersions - This method takes in a string value, a comma separated list of
string values, or a slice of string values that represents a version of the
TaXII API that is supported and adds it to the versions property.
*/
func (o *APIRoot) AddVersions(values interface{}) error {
	// if o.Versions == nil {
	// 	a := make([]string, 0)
	// 	o.Versions = a
	// }
	//o.Versions = append(o.Versions, s)
	return objects.AddValuesToList(&o.Versions, values)
}

/*
SetMaxContentLength - This method takes in an integer value representing the
max content length that the server can support and updates the max content
length property.
*/
func (o *APIRoot) SetMaxContentLength(i int) error {
	o.MaxContentLength = i
	return nil
}

/*
GetMaxContentLength - This method returns the max content length as an
integer.
*/
func (o *APIRoot) GetMaxContentLength() int {
	return o.MaxContentLength
}
