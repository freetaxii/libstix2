// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"fmt"
	"time"

	"github.com/freetaxii/libstix2/objects/indicator"
)

func main() {
	i := indicator.New()

	i.SetName("Malware C2 Indicator 2016")
	i.AddLabel("BadStuff")
	i.AddType("compromised")

	// Set modified time to be one hour from now
	//modifiedTime := time.Now().Add(time.Hour)
	//i.SetModified(modifiedTime)

	i.SetValidFrom(time.Now())
	i.CreateKillChainPhase("lockheed-martin-cyber-kill-chain", "delivery")

	data, _ := i.EncodeToString()
	fmt.Println(data)
}
