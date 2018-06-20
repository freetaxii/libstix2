// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/freetaxii/libstix2/objects"
)

func main() {
	data := getdata()

	var b objects.BundleDecodeType
	json.NewDecoder(strings.NewReader(data)).Decode(&b)

	count := 0
	for _, v := range b.Objects {

		var o objects.BundleObjectType
		err := json.Unmarshal(v, &o)
		if err != nil {
			continue
		}
		count++

		switch o.ObjectType {
		case "campaign":
			var o objects.CampaignType
			err = json.Unmarshal(v, &o)
			fmt.Println(o.ID)
		case "indicator":
			var o objects.IndicatorType
			err = json.Unmarshal(v, &o)
			fmt.Println(o.ID)
		case "infrastructure":
			var o objects.InfrastructureType
			err = json.Unmarshal(v, &o)
			fmt.Println(o.ID)
		case "malware":
			var o objects.MalwareType
			err = json.Unmarshal(v, &o)
			fmt.Println(o.ID)
		case "observed-data":
			var o objects.ObservedDataType
			err = json.Unmarshal(v, &o)
			fmt.Println(o.ID)
		case "relationship":
			var o objects.RelationshipType
			err = json.Unmarshal(v, &o)
			fmt.Println(o.ID)
		case "sighting":
			var o objects.SightingType
			err = json.Unmarshal(v, &o)
			fmt.Println(o.ID)
		}
	}
	fmt.Println("Total number of objects", count)

}

func getdata() string {
	s := `
{
    "type": "bundle",
    "id": "bundle--7e95ec95-b71a-45d9-a3c1-dcbb21f3fdcf",
    "spec_version": "2.0",
    "objects": [
        {
            "type": "campaign",
            "id": "campaign--d62e1eff-eb93-42e2-bd90-dabff3b93427",
            "created": "2018-06-05T18:25:15.917Z",
            "modified": "2018-06-05T18:25:15.917Z",
            "name": "Bank Attack 2016",
            "objective": "Compromise SWIFT system and steal money"
        },
        {
            "type": "indicator",
            "id": "indicator--2210fa4b-23bc-40b2-a3d7-2282530e8f5f",
            "created": "2018-06-05T18:25:15.917Z",
            "modified": "2018-06-05T18:25:15.917Z",
            "name": "Malware C2 Indicator 2016",
            "description": "This indicator should detect the SpyEye malware by looking for this MD5 hash",
            "pattern": "file-object:hashes.md5 = 84714c100d2dfc88629531f6456b8276"
        },
        {
            "type": "infrastructure",
            "id": "infrastructure--baab384b-e737-47c8-a58d-382f6862f9f0",
            "created": "2018-06-05T18:25:15.917Z",
            "modified": "2018-06-05T18:25:15.917Z",
            "name": "SpyEye Command and Control Servers",
            "description": "These servers are located in a datacenter in the Netherlands and the IPs change on a weekly basis"
        },
        {
            "type": "observed-data",
            "id": "observed-data--779862bb-90c0-42f0-b9ad-a7b3cb4aa2a7",
            "created": "2018-06-05T18:25:15.917Z",
            "modified": "2018-06-05T18:25:15.917Z",
            "first_observed": "2016-09-01T00:00:01Z",
            "last_observed": "2016-09-07T00:00:01Z",
            "number_observed": 3
        },
        {
            "type": "observed-data",
            "id": "observed-data--4939b7a0-44b8-46fb-a66b-f7127db985f4",
            "created": "2018-06-05T18:25:15.917Z",
            "modified": "2018-06-05T18:25:15.917Z",
            "first_observed": "2016-09-07T00:00:01Z",
            "last_observed": "2016-09-14T00:00:01Z",
            "number_observed": 3
        },
        {
            "type": "observed-data",
            "id": "observed-data--64e33489-0b86-46c2-9b13-d757c3bf4334",
            "created": "2018-06-05T18:25:15.917Z",
            "modified": "2018-06-05T18:25:15.917Z",
            "first_observed": "2016-09-07T00:00:01Z",
            "last_observed": "2016-09-14T00:00:01Z",
            "number_observed": 1
        },
        {
            "type": "malware",
            "id": "malware--c1587bf8-dda6-42a7-8636-0865b19192f5",
            "created": "2018-06-05T18:25:15.917Z",
            "modified": "2018-06-05T18:25:15.917Z",
            "labels": [
                "trojan",
                "malware-family"
            ],
            "name": "Zeus"
        },
        {
            "type": "malware",
            "id": "malware--c5460773-10d2-4d06-80c1-7b0d83e73c0b",
            "created": "2018-06-05T18:25:15.917Z",
            "modified": "2018-06-05T18:25:15.917Z",
            "labels": [
                "trojan"
            ],
            "name": "SpyEye"
        },
        {
            "type": "relationship",
            "id": "relationship--971c0aa3-1490-4931-8804-00acbd381242",
            "created": "2018-06-05T18:25:15.917Z",
            "modified": "2018-06-05T18:25:15.917Z",
            "relationship_type": "member-of",
            "source_ref": "malware--c1587bf8-dda6-42a7-8636-0865b19192f5",
            "target_ref": "malware--c5460773-10d2-4d06-80c1-7b0d83e73c0b"
        },
        {
            "type": "relationship",
            "id": "relationship--c3484590-8c09-4e76-877d-7aae13b64f8a",
            "created": "2018-06-05T18:25:15.917Z",
            "modified": "2018-06-05T18:25:15.917Z",
            "relationship_type": "uses",
            "source_ref": "campaign--d62e1eff-eb93-42e2-bd90-dabff3b93427",
            "target_ref": "malware--c5460773-10d2-4d06-80c1-7b0d83e73c0b"
        },
        {
            "type": "relationship",
            "id": "relationship--249ece43-a4a3-421b-9e0d-fc51a19eed52",
            "created": "2018-06-05T18:25:15.917Z",
            "modified": "2018-06-05T18:25:15.917Z",
            "relationship_type": "uses",
            "source_ref": "campaign--d62e1eff-eb93-42e2-bd90-dabff3b93427",
            "target_ref": "infrastructure--baab384b-e737-47c8-a58d-382f6862f9f0"
        },
        {
            "type": "relationship",
            "id": "relationship--184a85b4-4c2a-473f-8184-e47639450d2a",
            "created": "2018-06-05T18:25:15.917Z",
            "modified": "2018-06-05T18:25:15.917Z",
            "relationship_type": "uses",
            "source_ref": "malware--c5460773-10d2-4d06-80c1-7b0d83e73c0b",
            "target_ref": "infrastructure--baab384b-e737-47c8-a58d-382f6862f9f0"
        },
        {
            "type": "relationship",
            "id": "relationship--63d5bb66-244d-4dc9-9f0e-8cb97be46433",
            "created": "2018-06-05T18:25:15.917Z",
            "modified": "2018-06-05T18:25:15.917Z",
            "relationship_type": "indicates",
            "source_ref": "indicator--2210fa4b-23bc-40b2-a3d7-2282530e8f5f",
            "target_ref": "malware--c5460773-10d2-4d06-80c1-7b0d83e73c0b"
        },
        {
            "type": "relationship",
            "id": "relationship--4aaa36ea-29e6-4655-ae66-a661f56ffdcf",
            "created": "2018-06-05T18:25:15.917Z",
            "modified": "2018-06-05T18:25:15.917Z",
            "relationship_type": "part-of",
            "source_ref": "observed-data--779862bb-90c0-42f0-b9ad-a7b3cb4aa2a7",
            "target_ref": "infrastructure--baab384b-e737-47c8-a58d-382f6862f9f0"
        },
        {
            "type": "relationship",
            "id": "relationship--9e34ad0d-f7b7-46db-9e20-582f6222746f",
            "created": "2018-06-05T18:25:15.917Z",
            "modified": "2018-06-05T18:25:15.917Z",
            "relationship_type": "part-of",
            "source_ref": "observed-data--4939b7a0-44b8-46fb-a66b-f7127db985f4",
            "target_ref": "infrastructure--baab384b-e737-47c8-a58d-382f6862f9f0"
        },
        {
            "type": "sighting",
            "id": "sighting--0005022b-8ef3-4026-814b-c72aeea7e87e",
            "created": "2018-06-05T18:25:15.917Z",
            "modified": "2018-06-05T18:25:15.917Z",
            "first_seen": "2016-09-01T00:00:01Z",
            "last_seen": "2016-09-01T10:30:00Z",
            "count": 3,
            "sighting_of_ref": "malware--c5460773-10d2-4d06-80c1-7b0d83e73c0b"
        },
        {
            "type": "sighting",
            "id": "sighting--75b12ef7-5ef5-4c97-8ab8-e6b8b21d9c0b",
            "created": "2018-06-05T18:25:15.917Z",
            "modified": "2018-06-05T18:25:15.917Z",
            "first_seen": "2016-09-01T00:00:01Z",
            "last_seen": "2016-09-01T10:30:00Z",
            "count": 10,
            "sighting_of_ref": "infrastructure--baab384b-e737-47c8-a58d-382f6862f9f0",
            "observed_data_refs": [
                "observed-data--64e33489-0b86-46c2-9b13-d757c3bf4334"
            ]
        }
    ]
}
`

	return s
}
