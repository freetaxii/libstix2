// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/freetaxii/libstix2/messages/report"
	"github.com/freetaxii/libstix2/messages/stix"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

func main() {
	r := report.New()

	r.AddLabel("Attack Report")
	r.SetName("Malware Foo Report 2016")
	r.SetDescription("This report gives us details about Malware Foo")
	r.SetPublished(time.Now())

	r.AddObject(stix.NewId("malware"))
	// r.AddObject(stix.NewId("campaign"))
	// r.AddObject(stix.NewId("sighting"))
	// r.AddObject(stix.NewId("sighting"))
	// r.AddObject(stix.NewId("threat-actor"))
	// r.AddObject(stix.NewId("threat-actor"))
	// r.AddObject(stix.NewId("relationship"))
	// r.AddObject(stix.NewId("relationship"))
	// r.AddObject(stix.NewId("relationship"))
	// r.AddObject(stix.NewId("relationship"))
	// r.AddObject(stix.NewId("relationship"))
	// r.AddObject(stix.NewId("relationship"))
	// r.AddObject(stix.NewId("relationship"))
	// r.AddObject(stix.NewId("relationship"))

	// for j := 0; j <= 4; j++ {
	// 	r.AddObject(stix.NewId("indicator"))
	// }

	// Open connection to database
	filename := "/opt/go/src/github.com/freetaxii/libstix2/examples/db/freetaxii.sqlite"
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatalf("Unable to open file %s due to error %v", filename, err)
	}
	defer db.Close()

	r.AddToDatabase(db)

	var data []byte
	data, _ = json.MarshalIndent(r, "", "    ")

	fmt.Println(string(data))
}
