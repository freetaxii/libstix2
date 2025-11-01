// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

//go:build example
// +build example

package main

import (
	"fmt"

	"github.com/freetaxii/libstix2/objects/taxii/discovery"
)

func main() {

	var d discovery.Discovery

	d.SetTitle("FreeTAXII Discovery Service")
	d.SetDescription("This service will display API roots that this TAXII knows about.")
	d.SetContact("FreeTAXII")
	d.SetDefault("https://www.freetaxii.com/api1")
	d.AddAPIRoots("https://www.freetaxii.com/api1")
	d.AddAPIRoots("https://www.freetaxii.com/api2")
	data, _ := d.EncodeToString()

	fmt.Println(data)
}
