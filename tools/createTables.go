// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package main

import (
	"github.com/freetaxii/libstix2/datastore"
	"log"
)

func main() {
	var err error
	filename := "/opt/go/src/github.com/freetaxii/libstix2/examples/db/freestix.sqlite"
	ds := datastore.NewSqlite3(filename)

	err = ds.Connect()
	handleError(err)

	//ds.CreateTables()
	//ds.CreateVocabTables()
	//ds.PopulateVocabTables()

	ds.Close()

}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
