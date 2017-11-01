// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/freetaxii/libstix2/resources"
)

func main() {
	//d := resources.NewDiscovery()

	var d resources.DiscoveryType

	d.SetTitle("FreeTAXII Discovery Service")
	d.SetDescription("This service will display API roots that this TAXII knows about.")
	d.SetContact("FreeTAXII")
	d.SetDefault("https://www.freetaxii.com/api1")
	d.AddAPIRoot("https://www.freetaxii.com/api1")
	d.AddAPIRoot("https://www.freetaxii.com/api2")

	var data []byte
	data, _ = json.MarshalIndent(d, "", "    ")

	fmt.Println(string(data))
}
