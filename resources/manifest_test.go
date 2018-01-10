// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package resources

import (
	"testing"
)

// ----------------------------------------------------------------------
//
// func (ezt *ManifestType) AddManifestEntry(o ManifestEntryType) (int, error)
//
// ----------------------------------------------------------------------
func Test_sqlObjectList(t *testing.T) {
	m := InitManifestObject()
	me := InitManifestEntryObject()

	t.Log("Test 1: get an index integer of 0 for adding a manifest")
	if int1, _ := m.AddManifestEntry(me); int1 != 0 {
		t.Error("incorrect index value")
	}

	t.Log("Test 2: get an index integer of 1 for adding a manifest")
	if int2, _ := m.AddManifestEntry(me); int2 != 1 {
		t.Error("incorrect index value")
	}

}
