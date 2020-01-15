// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"encoding/json"
	"fmt"

	"github.com/freetaxii/libstix2/objects/attackpattern"
)

type myCustomAttackPattern struct {
	*attackpattern.AttackPattern
	SomeCustomProperty string `json:"some_custom_propety,omitempty"`
}

func main() {
	o := attackpattern.New()
	customAP := myCustomAttackPattern{AttackPattern: o}
	// var customAP myCustomAttackPattern
	// customAP.AttackPattern = o

	customAP.SomeCustomProperty = "some custom string data"
	customAP.SetName("Phishing 123")

	data2, err1 := json.MarshalIndent(customAP, "", "    ")
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(string(data2))

}
