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
	i.AddTypes("compromised")

	// Set modified time to be one hour from now
	//modifiedTime := time.Now().Add(time.Hour)
	//i.SetModified(modifiedTime)

	i.SetValidFrom(time.Now())
	i.CreateKillChainPhase("lockheed-martin-cyber-kill-chain", "delivery")
	i.SetPattern("somedata")
	i.SetPatternType("stix")

	valid, problems, details := i.Valid()
	fmt.Println("Is valid:", valid)
	fmt.Println("Problems:", problems)
	fmt.Println("Error Msg:", details)

	data1, _ := i.EncodeToString()
	fmt.Println(data1)

	i2 := indicator.New()

	i2.SetName("Malware C2 Indicator 2016")
	i2.AddLabel("BadStuff1")
	i2.AddTypes("compromised1")

	i2.SetValidFrom(time.Now())
	i2.CreateKillChainPhase("lockheed-martin-cyber-kill-chain", "delivery")
	i2.SetPattern("somedata")
	i2.SetPatternType("stix")

	data2, _ := i.EncodeToString()
	fmt.Println(data2)

	okay, number, details := i2.Compare(i)
	fmt.Println("Are they equal: ", okay)
	fmt.Println("How many issues: ", number)
	for _, value := range details {
		fmt.Println(value)
	}
}
