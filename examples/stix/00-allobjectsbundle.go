// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"encoding/json"
	"fmt"

	"github.com/freetaxii/libstix2/objects/attackpattern"
	"github.com/freetaxii/libstix2/objects/bundle"
	"github.com/freetaxii/libstix2/objects/campaign"
	"github.com/freetaxii/libstix2/objects/courseofaction"
	"github.com/freetaxii/libstix2/objects/grouping"
	"github.com/freetaxii/libstix2/objects/identity"
	"github.com/freetaxii/libstix2/objects/indicator"
	"github.com/freetaxii/libstix2/objects/infrastructure"
	"github.com/freetaxii/libstix2/objects/intrusionset"
	"github.com/freetaxii/libstix2/objects/location"
	"github.com/freetaxii/libstix2/objects/malware"
	"github.com/freetaxii/libstix2/objects/malwareanalysis"
	"github.com/freetaxii/libstix2/objects/note"
	"github.com/freetaxii/libstix2/objects/observeddata"
	"github.com/freetaxii/libstix2/objects/opinion"
	"github.com/freetaxii/libstix2/objects/relationship"
	"github.com/freetaxii/libstix2/objects/report"
	"github.com/freetaxii/libstix2/objects/sighting"
	"github.com/freetaxii/libstix2/objects/threatactor"
	"github.com/freetaxii/libstix2/objects/tool"
	"github.com/freetaxii/libstix2/objects/vulnerability"
)

func main() {
	sm := bundle.New()

	sm.AddObject(attackpattern.New())
	sm.AddObject(campaign.New())
	sm.AddObject(courseofaction.New())
	sm.AddObject(grouping.New())
	sm.AddObject(identity.New())
	sm.AddObject(indicator.New())
	sm.AddObject(infrastructure.New())
	sm.AddObject(intrusionset.New())
	sm.AddObject(location.New())
	sm.AddObject(malware.New())
	sm.AddObject(malwareanalysis.New())
	sm.AddObject(note.New())
	sm.AddObject(observeddata.New())
	sm.AddObject(opinion.New())
	sm.AddObject(relationship.New())
	sm.AddObject(report.New())
	sm.AddObject(sighting.New())
	sm.AddObject(threatactor.New())
	sm.AddObject(tool.New())
	sm.AddObject(vulnerability.New())

	var data []byte
	data, _ = json.MarshalIndent(sm, "", "    ")

	fmt.Println(string(data))
}
