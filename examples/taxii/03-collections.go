// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"encoding/json"
	"fmt"

	"github.com/freetaxii/libstix2/objects/taxii/collections"
	"github.com/google/uuid"
)

func main() {
	o := collections.New()

	c, _ := o.NewCollection()
	// create a new UUID for the collection
	c.ID = uuid.New().String()
	c.SetCanRead()

	// You can do this manually by creating your own Collection object
	// and adding data to and then manually adding it to the collections array
	// c := objects.NewCollection()
	// c.NewId()
	// c.SetCanRead()
	// i := o.AddCollection(c)
	// You can even add stuff to it after the fact, even if this is bad form
	// from an OO perspective
	// o.Collections[i].SetCanWrite()

	var data []byte
	data, _ = json.MarshalIndent(o, "", "    ")

	fmt.Println(string(data))
}
