// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/freetaxii/libstix2/objects"
	"time"
)

func main() {
	container := make([]interface{}, 0)

	// Create a report
	r := objects.NewReport()
	r.SetName("Malware Foo Report 2016")
	r.SetDescription("This report gives us details about Malware Foo1")
	r.SetPublished(time.Now())

	// Create a campagin
	c := objects.NewCampaign()
	c.SetName("Bank Attack 2016")
	c.SetObjective("Compromise SWIFT system and steal money")
	r.AddObject(c.GetId())
	container = append(container, c)

	// Define a family of malware
	m1 := objects.NewMalware()
	m1.SetName("Zeus")
	m1.AddLabel("trojan")
	m1.AddLabel("malware-family")
	r.AddObject(m1.GetId())
	container = append(container, m1)

	// Define a piece of malware
	m2 := objects.NewMalware()
	m2.SetName("SpyEye")
	m2.AddLabel("trojan")
	m2.AddFilename("cleansweep.exe")
	m2.AddFilename("spyeye2_exe")
	m2.AddFilename("build_1_.exe")
	m2.AddHash("md5", "84714c100d2dfc88629531f6456b8276")
	m2.AddHash("sha256", "861aa9c5ddcb5284e1ba4e5d7ebacfa297567c353446506ee4b4e39c84454b09")
	m2.AddKillChainPhase("lockheed-martin-cyber-kill-chain", "command-and-control")
	r.AddObject(m2.GetId())
	container = append(container, m2)

	// Create an indicator
	i := objects.NewIndicator()
	i.SetName("Malware C2 Indicator 2016")
	i.SetDescription("This indicator should detect the SpyEye malware by looking for this MD5 hash")
	i.SetPattern("file-object:hashes.md5 = 84714c100d2dfc88629531f6456b8276")
	container = append(container, c)

	r.AddObject(r.NewId("sighting"))
	r.AddObject(r.NewId("sighting"))
	r.AddObject(r.NewId("threat-actor"))
	r.AddObject(r.NewId("threat-actor"))
	r.AddObject(r.NewId("relationship"))
	r.AddObject(r.NewId("relationship"))
	r.AddObject(r.NewId("relationship"))
	r.AddObject(r.NewId("relationship"))
	r.AddObject(r.NewId("relationship"))
	r.AddObject(r.NewId("relationship"))
	r.AddObject(r.NewId("relationship"))
	r.AddObject(r.NewId("relationship"))

	for j := 0; j <= 4; j++ {
		r.AddObject(r.NewId("indicator"))
	}

	container = append(container, r)
	var data []byte
	data, _ = json.MarshalIndent(container, "", "    ")

	fmt.Println(string(data))
}
