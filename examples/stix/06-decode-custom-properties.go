// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

//go:build ignore
// +build ignore

package main

import (
	"encoding/json"
	"fmt"

	"github.com/freetaxii/libstix2/objects/attackpattern"
)

func main() {

	// Decode the data defined down below with the custom property data being
	// stored in a map that is called "custom"
	o, err := attackpattern.Decode([]byte(getdata()))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(o)

	fmt.Println("Just the custom properties:")
	fmt.Println(o.Custom)

	// Since we know the data is a string, lets create a variable to unmarshal the data to
	var foo string
	json.Unmarshal(o.Custom["some_custom_property"], &foo)
	fmt.Println(foo)

	data2, _ := o.EncodeToString()
	fmt.Println(data2)

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
    "aliases": ["Banking1", "ATM2"],
    "some_custom_property": "some_custom_value",
    "some_custom_property1": "some_custom_value1"
}
`
	return s
}
