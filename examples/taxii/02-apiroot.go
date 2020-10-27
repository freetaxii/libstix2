// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"fmt"

	"github.com/wxj95/libstix2/objects/taxii/apiroot"
)

func main() {
	a := apiroot.New()

	a.SetTitle("FreeTAXII API Root 1")
	a.SetDescription("This API Root contains OSINT.")
	a.SetMaxContentLength(10485760)
	a.AddVersions("taxii-2.0")

	data, _ := a.EncodeToString()
	fmt.Println(data)
}
