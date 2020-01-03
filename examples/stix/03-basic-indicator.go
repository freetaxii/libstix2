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
	i.AddLabels("BadStuff")
	i.AddTypes("compromised")
	i.AddTypes("test1,test2 , test3")
	t := []string{"test4 ", " test5", "test6"}
	i.AddTypes(t)

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

	data, _ := i.EncodeToString()
	fmt.Println(data)
}
