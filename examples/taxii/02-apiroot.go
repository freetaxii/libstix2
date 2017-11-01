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
	d := resources.NewAPIRoot()

	d.SetTitle("FreeTAXII API Root 1")
	d.SetDescription("This API Root contains OSINT.")
	d.SetMaxContentLength(10485760)
	d.AddVersion("taxii-2.0")

	var data []byte
	data, _ = json.MarshalIndent(d, "", "    ")

	fmt.Println(string(data))
}
