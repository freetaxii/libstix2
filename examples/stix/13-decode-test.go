// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/bundle"
)

func main() {
	data := getdata()

	b, err := bundle.Decode(strings.NewReader(data))
	if err != nil {
		for _, v := range err {
			slog.Info(v.Error())
		}
		os.Exit(1)
	}

	count := 0
	for _, v := range b.Objects {

		o, err := objects.Decode(v)
		if err != nil {
			fmt.Println("ERROR:", err)
			continue
		}
		if o == nil {
			continue
		}

		fmt.Printf("Type: %s\t\tID: %s\tVersion: %s\n", o.GetObjectType(), o.GetID(), o.GetModified())
		fmt.Println(o, "\n")

		count++
	}
	fmt.Println("===========================\nTotal number of objects", count)

}

func getdata() string {
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
        }
    ]
}
`

	return s
}
