// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/wxj95/libstix2/objects/envelope"
	"github.com/wxj95/libstix2/objects/indicator"
)

func main() {

	e := envelope.New()

	// Create an indicator
	i := indicator.New()

	i.SetName("Malware C2 Indicator 2016")
	i.AddLabel("BadStuff")
	i.AddType("compromised")
	i.SetPattern("[ ipv4-addr:value = '192.168.100.100' ]")
	i.SetValidFrom(time.Now())
	i.CreateKillChainPhase("lockheed-martin-cyber-kill-chain", "delivery")

	e.AddObject(i)

	var data []byte
	data, _ = json.MarshalIndent(e, "", "    ")

	fmt.Println(string(data))
}
