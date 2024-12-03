// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"encoding/json"
	"fmt"

	"github.com/freetaxii/libstix2/objects/taxii/manifest"
)

func main() {
	o := manifest.New()

	o.CreateRecord("indicator--623d9b92-28cd-49f4-9e53-557fd648fd8c", "2017-01-01T01:01:01.123456Z", "2017-01-01T01:01:01.123456Z", "stix")
	o.CreateRecord("indicator--9d463cd7-3be6-4cd1-a720-75c9886de896", "2018-01-01T01:01:01.123456Z", "2016-01-01T01:01:01.123456Z,2017-01-01T01:01:01.123456Z", "stix 2.0")

	var data []byte
	data, _ = json.MarshalIndent(o, "", "    ")

	fmt.Println(string(data))
}
