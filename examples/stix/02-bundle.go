// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"encoding/json"
	"fmt"

	"github.com/freetaxii/libstix2/objects/bundle"
	"github.com/freetaxii/libstix2/objects/campaign"
	"github.com/freetaxii/libstix2/objects/indicator"
	"github.com/freetaxii/libstix2/objects/infrastructure"
	"github.com/freetaxii/libstix2/objects/malware"
	"github.com/freetaxii/libstix2/objects/observeddata"
	"github.com/freetaxii/libstix2/objects/relationship"
	"github.com/freetaxii/libstix2/objects/sighting"
)

func main() {
	sm := bundle.New()

	// Create a campaign
	c := campaign.New()
	c.SetName("Bank Attack 2016")
	c.SetObjective("Compromise SWIFT system and steal money")
	sm.AddObject(c)

	// Create an indicator
	i := indicator.New()
	i.SetName("Malware C2 Indicator 2016")
	i.SetDescription("This indicator should detect the SpyEye malware by looking for this MD5 hash")
	i.SetPattern("file-object:hashes.md5 = 84714c100d2dfc88629531f6456b8276")
	sm.AddObject(i)

	// Define some infrastructure as used by this campaign and malware.
	infra := infrastructure.New()
	infra.SetName("SpyEye Command and Control Servers")
	infra.SetDescription("These servers are located in a datacenter in the Netherlands and the IPs change on a weekly basis")
	// infra.AddKillChainPhase("lockheed-martin-cyber-kill-chain", "command-and-control")
	// infra.SetFirstSeenText("2016-09-01T00:00:01Z")
	// infra.SetRegion("Europe")
	// infra.SetCountry("NL")
	sm.AddObject(infra)

	// Define some Observed Data for the Infrastructure
	od1 := observeddata.New()
	od1.SetFirstObserved("2016-09-01T00:00:01Z")
	od1.SetLastObserved("2016-09-07T00:00:01Z")
	od1.SetNumberObserved(3)
	//od1.SetCybox("This will be a CybOX container object using the ipv4-addr object pointing to 5.79.68.0/24")
	sm.AddObject(od1)

	od2 := observeddata.New()
	od2.SetFirstObserved("2016-09-07T00:00:01Z")
	od2.SetLastObserved("2016-09-14T00:00:01Z")
	od2.SetNumberObserved(3)
	//od2.SetCybox("This will be a CybOX container object using the ipv4-addr object pointing to 5.79.52.0/24")
	sm.AddObject(od2)

	// Define some Observed Data for the sighting of the Infrastructure
	od3 := observeddata.New()
	od3.SetFirstObserved("2016-09-07T00:00:01Z")
	od3.SetLastObserved("2016-09-14T00:00:01Z")
	od3.SetNumberObserved(1)
	//od3.SetCybox("This will be a CybOX container object using the ipv4-addr object pointing to 5.79.52.100")
	sm.AddObject(od3)

	// Define a family of malware
	m1 := malware.New()
	m1.SetName("Zeus")
	m1.AddLabels("trojan")
	m1.AddLabels("malware-family")
	sm.AddObject(m1)

	// Define a piece of malware
	m2 := malware.New()
	m2.SetName("SpyEye")
	m2.AddLabels("trojan")
	// m2.AddFilename("cleansweep.exe")
	// m2.AddFilename("spyeye2_exe")
	// m2.AddFilename("build_1_.exe")
	// m2.AddHash("md5", "84714c100d2dfc88629531f6456b8276")
	// m2.AddHash("sha256", "861aa9c5ddcb5284e1ba4e5d7ebacfa297567c353446506ee4b4e39c84454b09")
	// m2.AddKillChainPhase("lockheed-martin-cyber-kill-chain", "command-and-control")
	sm.AddObject(m2)

	// Define some scan data for the malware sample
	// m2s1 := m2.NewScanData()
	// m2s1.SetScannedText("2016-08-30T06:31:48Z")
	// m2s1.SetProduct("avg")
	// m2s1.SetClassification("Generic16.BFGI")

	// m2s2 := m2.NewScanData()
	// m2s2.SetScannedText("2016-08-30T06:31:48Z")
	// m2s2.SetProduct("avast")
	// m2s2.SetClassification("Win32:Downloader-NTU [PUP]")

	// Connect the malware sample to a malware family
	r1 := relationship.New()
	r1.SetType("member-of")
	r1.SetSourceTarget(m1.GetID(), m2.GetID())
	sm.AddObject(r1)

	// Identify that this campaign uses this piece of malware
	r2 := relationship.New()
	r2.SetType("uses")
	r2.SetSourceTarget(c.GetID(), m2.GetID())
	sm.AddObject(r2)

	// Identify that this campaign uses this infrastructure
	r3 := relationship.New()
	r3.SetType("uses")
	r3.SetSourceTarget(c.GetID(), infra.GetID())
	sm.AddObject(r3)

	// Identify that this malware uses this infrastructure
	r4 := relationship.New()
	r4.SetType("uses")
	r4.SetSourceTarget(m2.GetID(), infra.GetID())
	sm.AddObject(r4)

	// Identify that this indicator can indicate the presence of this malware
	r5 := relationship.New()
	r5.SetType("indicates")
	r5.SetSourceTarget(i.GetID(), m2.GetID())
	sm.AddObject(r5)

	// Attach some Observed Data to an Infrastructure Object
	r6 := relationship.New()
	r6.SetType("part-of")
	r6.SetSourceTarget(od1.GetID(), infra.GetID())
	sm.AddObject(r6)

	// Attach some Observed Data to an Infrastructure Object
	r7 := relationship.New()
	r7.SetType("part-of")
	r7.SetSourceTarget(od2.GetID(), infra.GetID())
	sm.AddObject(r7)

	// Add a sighting for the malware
	s1 := sighting.New()
	s1.SetFirstSeen("2016-09-01T00:00:01Z")
	s1.SetLastSeen("2016-09-01T10:30:00Z")
	s1.SetCount(3)
	s1.SetSightingOfRef(m2.GetID())
	sm.AddObject(s1)

	// Add a sighting for the infrastructure
	s2 := sighting.New()
	s2.SetFirstSeen("2016-09-01T00:00:01Z")
	s2.SetLastSeen("2016-09-01T10:30:00Z")
	s2.SetCount(10)
	s2.SetSightingOfRef(infra.GetID())
	s2.AddObservedDataRefs(od3.GetID())
	sm.AddObject(s2)

	var data []byte
	data, _ = json.MarshalIndent(sm, "", "    ")

	fmt.Println(string(data))
}
