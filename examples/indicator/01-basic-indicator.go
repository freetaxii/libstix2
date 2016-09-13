// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/freetaxii/libstix2/messages/indicator"
	//"log"
	"time"
)

func main() {
	i := indicator.New()

	i.SetName("Malware C2 Indicator 2016")

	// Set modified time to be one hour from now
	//modifiedTime := time.Now().Add(time.Hour)
	//i.SetModified(modifiedTime)

	validFrom := time.Now()
	i.SetValidFrom(validFrom)
	i.SetPrecisionYear("valid_from")
	i.AddKillChainPhase("lockheed-martin-cyber-kill-chain", "delivery")

	// err := i.SetVersion(2)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var data []byte
	data, _ = json.MarshalIndent(i, "", "    ")

	fmt.Println(string(data))
}
