// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"fmt"

	"github.com/wxj95/libstix2/objects/attackpattern"
)

func main() {
	o := attackpattern.New()

	o.SetName("Phishing")
	o.AddAliases("BadStuff")

	o.CreateKillChainPhase("lockheed-martin-cyber-kill-chain", "delivery")

	data, _ := o.EncodeToString()
	fmt.Println("Step 1: Print a basic attack pattern created in this script")
	fmt.Println(data)

	o1, err := attackpattern.Decode([]byte(getdata()))
	if err != nil {
		fmt.Println(err)
	}
	data1, _ := o1.EncodeToString()
	fmt.Println("Step 2: Print a basic attack pattern from data found in the script")
	fmt.Println("This data has been decoded and then re-encoded")
	fmt.Println(data1)

}

func getdata() string {
	s := `
{
	"type": "attack-pattern",
    "spec_version": "2.1",
    "id": "attack-pattern--d62e1eff-eb93-42e2-bd90-dabff3b93427",
    "created": "2018-06-05T18:25:15.917Z",
    "modified": "2018-06-05T18:25:15.917Z",
    "name": "Phishing",
    "aliases": ["Banking1", "ATM2"]
}
`
	return s
}
