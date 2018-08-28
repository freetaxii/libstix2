// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

import (
	"github.com/freetaxii/libstix2/objects/properties"
)

func VerifyCommonProperties(o properties.CommonObjectProperties) error {

	if err := o.TypeProperty.Verify(); err != nil {
		return err
	}
	if err := o.IDProperty.Verify(); err != nil {
		return err
	}
	return nil
}
