// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package note

import "github.com/freetaxii/libstix2/resources"

/* AddAuthors - This method takes in a string value, a comma separated list of
string values, or a slice of string values that represents an author and adds it
to the authors property. */
func (o *Note) AddAuthors(values interface{}) error {
	return resources.AddValuesToList(&o.Authors, values)
}
