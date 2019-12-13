// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package bundle

import "github.com/freetaxii/libstix2/objects"

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/* AddObject - This method will take in an object as an interface and add it to
the list of objects in the bundle. */
func (o *Bundle) AddObject(i objects.STIXObject) error {
	o.Objects = append(o.Objects, i)
	return nil
}
