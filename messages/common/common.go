// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package common

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/freetaxii/libstix2/messages/defs"
	"time"
)

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

func CreateId(t string) string {
	// TODO Add check to validate input value
	id := t + "--" + uuid.New()
	return id
}

func GetCurrentTime() string {
	t := time.Now().UTC().Format(defs.TIME_RFC_3339)
	return t
}
