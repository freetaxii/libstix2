// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/freetaxii/libstix2/messages/report"
	"github.com/freetaxii/libstix2/messages/stix"
	"time"
)

func main() {
	r := report.New()

	r.SetName("Malware Foo Report 2016")
	r.SetDescription("This report gives us details about Malware Foo1")
	r.SetPublished(time.Now())

	r.AddObject(stix.NewId("malware"))
	r.AddObject(stix.NewId("campaign"))
	r.AddObject(stix.NewId("sighting"))
	r.AddObject(stix.NewId("sighting"))
	r.AddObject(stix.NewId("threat-actor"))
	r.AddObject(stix.NewId("threat-actor"))
	r.AddObject(stix.NewId("relationship"))
	r.AddObject(stix.NewId("relationship"))
	r.AddObject(stix.NewId("relationship"))
	r.AddObject(stix.NewId("relationship"))
	r.AddObject(stix.NewId("relationship"))
	r.AddObject(stix.NewId("relationship"))
	r.AddObject(stix.NewId("relationship"))
	r.AddObject(stix.NewId("relationship"))

	for j := 0; j <= 4; j++ {
		r.AddObject(stix.NewId("indicator"))
	}

	var data []byte
	data, _ = json.MarshalIndent(r, "", "    ")

	fmt.Println(string(data))
}
