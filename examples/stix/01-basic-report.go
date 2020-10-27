// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/wxj95/libstix2/objects/report"
)

func main() {
	r := report.New()

	r.SetName("Malware Foo Report 2016")
	r.SetDescription("This report gives us details about Malware Foo1")
	r.SetPublished(time.Now())

	m, _ := r.CreateSTIXUUID("malware")
	r.AddObjectRefs(m)

	c, _ := r.CreateSTIXUUID("campaign")
	r.AddObjectRefs(c)

	s1, _ := r.CreateSTIXUUID("sighting")
	r.AddObjectRefs(s1)

	s2, _ := r.CreateSTIXUUID("sighting")
	r.AddObjectRefs(s2)

	t1, _ := r.CreateSTIXUUID("threat-actor")
	r.AddObjectRefs(t1)

	t2, _ := r.CreateSTIXUUID("threat-actor")
	r.AddObjectRefs(t2)

	for i := 0; i <= 8; i++ {
		r1, _ := r.CreateSTIXUUID("relationship")
		r.AddObjectRefs(r1)

	}

	for j := 0; j <= 4; j++ {
		i, _ := r.CreateSTIXUUID("indicator")
		r.AddObjectRefs(i)
	}

	var data []byte
	data, _ = json.MarshalIndent(r, "", "    ")

	fmt.Println(string(data))
}
