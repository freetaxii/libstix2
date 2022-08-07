// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"fmt"
	decodeExample "github.com/freetaxii/libstix2/examples/stix/decode-example"
	"github.com/freetaxii/libstix2/objects"
	"strings"

	"github.com/freetaxii/libstix2/objects/bundle"
	"github.com/gologme/log"
)

func main() {
	data := getBundle()

	b, errors := bundle.DecodeWithCustomObjects(strings.NewReader(data), map[string]bundle.DecodeFunc{
		"example": func(bytes []byte) (objects.STIXObject, error) {
			return decodeExample.Decode(bytes)
		},
	})
	for _, err := range errors {
		log.Fatalln(err)
	}

	count := 0
	for _, v := range b.Objects {

		o := v
		if o == nil {
			continue
		}

		fmt.Printf("Type: %s\t\tID: %s\tVersion: %s\n",
			o.GetCommonProperties().GetObjectType(),
			o.GetCommonProperties().GetID(),
			o.GetCommonProperties().GetModified(),
		)
		fmt.Println(o, "\n")

		count++
	}
	fmt.Println("===========================\nTotal number of objects", count)

}

func getBundle() string {
	s := `
{
    "type": "bundle",
    "id": "bundle--7e95ec95-b71a-45d9-a3c1-dcbb21f3fdcf",
    "objects": [
        {
            "type": "campaign",
            "spec_version": "2.1",
            "id": "campaign--d62e1eff-eb93-42e2-bd90-dabff3b93427",
            "created": "2018-06-05T18:25:15.917Z",
            "modified": "2018-06-05T18:25:15.917Z",
            "name": "Bank Attack 2016",
            "objective": "Compromise SWIFT system and steal money"
        },
        {
            "type": "indicator",
            "spec_version": "2.1",
            "id": "indicator--2210fa4b-23bc-40b2-a3d7-2282530e8f5f",
            "created": "2018-06-05T18:25:15.917Z",
            "modified": "2018-06-05T18:25:15.917Z",
            "name": "Malware C2 Indicator 2016",
            "description": "This indicator should detect the SpyEye malware by looking for this MD5 hash",
            "indicator_types": [ "compromised" ],
            "valid_from": "2018-05-05T12:12:13.142Z",
            "pattern": "file-object:hashes.md5 = 84714c100d2dfc88629531f6456b8276"
        },
        {
            "type": "example",
            "spec_version": "2.1",
            "id": "example--9faa1d36-b550-4202-9cbc-aa8994c9c5af",
            "created": "2022-08-07T14:23:36.917Z",
            "modified": "2022-08-07T14:23:36.917Z",
            "name": "Example Type",
            "example": true
        }
    ]
}
`

	return s
}
