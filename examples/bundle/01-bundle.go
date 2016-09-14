// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/freetaxii/libstix2/messages/bundle"
)

func main() {
	sm := bundle.New()

	// Create a campagin
	c := sm.NewCampaign()
	c.SetName("Bank Attack 2016")
	c.SetObjective("Compromise SWIFT system and steal money")

	// Create an indicator
	i := sm.NewIndicator()
	i.SetName("Malware C2 Indicator 2016")
	i.SetPattern("file-object:hashes.md5 = 84714c100d2dfc88629531f6456b8276")

	// Define a family of malware
	m1 := sm.NewMalware()
	m1.SetName("Zeus")
	m1.AddLabel("trojan")
	m1.AddLabel("malware-family")

	// Define a piece of malware
	m2 := sm.NewMalware()
	m2.SetName("SpyEye")
	m2.AddLabel("trojan")
	m2.AddFilename("cleansweep.exe")
	m2.AddFilename("spyeye2_exe")
	m2.AddFilename("build_1_.exe")
	m2.AddHash("md5", "84714c100d2dfc88629531f6456b8276")
	m2.AddHash("sha256", "861aa9c5ddcb5284e1ba4e5d7ebacfa297567c353446506ee4b4e39c84454b09")

	// Define some scan data for the malware sample
	m2s1 := m2.NewScanData()
	m2s1.SetScannedText("2016-08-30T06:31:48Z")
	m2s1.SetProduct("avg")
	m2s1.SetClassification("Generic16.BFGI")

	m2s2 := m2.NewScanData()
	m2s2.SetScannedText("2016-08-30T06:31:48Z")
	m2s2.SetProduct("avast")
	m2s2.SetClassification("Win32:Downloader-NTU [PUP]")

	// Connect the malware sample to a malware family
	r1 := sm.NewRelationship()
	r1.SetRelationshipType("member-of")
	r1.SetSourceRef(m1.GetId())
	r1.SetTargetRef(m2.GetId())

	// Identify that this campaign uses this piece of malware
	r2 := sm.NewRelationship()
	r2.SetRelationshipType("uses")
	r2.SetSourceRef(c.GetId())
	r2.SetTargetRef(m2.GetId())

	// Identify that this indicator can indicate the presence of this malware
	r3 := sm.NewRelationship()
	r3.SetRelationshipType("indicates")
	r3.SetSourceRef(i.GetId())
	r3.SetTargetRef(m2.GetId())

	// Add a sighting for the malware
	s1 := sm.NewSighting()
	s1.SetFirstSeenText("2016-09-01T00:00:00Z")
	s1.SetLastSeenText("2016-09-01T10:30:00Z")
	s1.SetCount(3)
	s1.SetSightingOfRef(m2.GetId())

	var data []byte
	data, _ = json.MarshalIndent(sm, "", "    ")

	fmt.Println(string(data))
}
