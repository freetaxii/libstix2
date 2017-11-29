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
	o := resources.NewManifest()

	o.CreateManifestEntry("1234", "2017", "2017-10-12", "stix")
	o.CreateManifestEntry("9999", "2018", "2016-10-12,2017-01-01", "stix 2.0")

	var data []byte
	data, _ = json.MarshalIndent(o, "", "    ")

	fmt.Println(string(data))
}
