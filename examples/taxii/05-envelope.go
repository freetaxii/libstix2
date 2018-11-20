// Copyright 2015-2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"encoding/json"
	"fmt"

	"github.com/freetaxii/libstix2/objects/campaign"
	"github.com/freetaxii/libstix2/resources/envelope"
)

func main() {

	var e envelope.Envelope

	e.SetMore()
	b, _ := e.NewBundle()

	// Create a campaign
	c := campaign.New()
	c.SetName("Bank Attack 2016")
	c.SetObjective("Compromise SWIFT system and steal money")
	b.AddObject(c)

	var data []byte
	data, _ = json.MarshalIndent(e, "", "    ")

	fmt.Println(string(data))
}
