// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/freetaxii/libstix2/objects"
)

func main() {
	r := objects.NewReport()

	r.SetName("Malware Foo Report 2016")
	r.SetDescription("This report gives us details about Malware Foo1")
	r.SetPublished(time.Now())

	m, _ := r.CreateSTIXUUID("malware")
	r.AddObject(m)

	c, _ := r.CreateSTIXUUID("campaign")
	r.AddObject(c)

	s1, _ := r.CreateSTIXUUID("sighting")
	r.AddObject(s1)

	s2, _ := r.CreateSTIXUUID("sighting")
	r.AddObject(s2)

	t1, _ := r.CreateSTIXUUID("threat-actor")
	r.AddObject(t1)

	t2, _ := r.CreateSTIXUUID("threat-actor")
	r.AddObject(t2)

	for i := 0; i <= 8; i++ {
		r1, _ := r.CreateSTIXUUID("relationship")
		r.AddObject(r1)

	}

	for j := 0; j <= 4; j++ {
		i, _ := r.CreateSTIXUUID("indicator")
		r.AddObject(i)
	}

	var data []byte
	data, _ = json.MarshalIndent(r, "", "    ")

	fmt.Println(string(data))
}
