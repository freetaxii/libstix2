// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package infrastructure

import (
	"testing"
	"time"

	"github.com/freetaxii/libstix2/defs"
)

/*
TestAddInfrastructureType -
*/
func TestAddInfrastructureType(t *testing.T) {
	m := New()
	want := "botnet"
	m.AddInfrastructureType(want)

	if got := m.InfrastructureTypes[0]; got != want {
		t.Error("Fail Infrastructure Add Infrastructure Type Check")
	}
}

/*
TestAddInfrastructureTypes -
*/
func TestAddInfrastructureTypes(t *testing.T) {
	m := New()
	m2 := New()
	want := "command-and-control, botnet"
	want2 := []string{" botnet ", "command-and-control"}
	wantRes := []string{"command-and-control", "botnet"}
	wantRes2 := []string{"botnet", "command-and-control"}

	m.AddInfrastructureTypes(want)
	m2.AddInfrastructureTypes(want2)

	if got := len(m.InfrastructureTypes); got != 2 {
		t.Error("Fail Infrastructure Add Infrastructure Types Check. List contains less items")
	}

	if got := len(m2.InfrastructureTypes); got != 2 {
		t.Error("Fail Infrastructure Add Infrastructure Types Check. Added less items")
	}

	if got := m.InfrastructureTypes[0]; got != wantRes[0] {
		t.Error("Fail Infrastructure Add Infrastructure Types Check. First value.")
	}

	if got := m.InfrastructureTypes[1]; got != wantRes[1] {
		t.Error("Fail Infrastructure Add Infrastructure Types Check. Method does not trim secont value string.")
	}

	if got := m2.InfrastructureTypes[0]; got != wantRes2[0] {
		t.Error("Fail Infrastructure Add Infrastructure Types Check. First value. Should add array of values")
	}

	if got := m2.InfrastructureTypes[1]; got != wantRes2[1] {
		t.Error("Fail Infrastructure Add Infrastructure Types Check. Method does not trim secont value string. Should add array of values")
	}
}

/*
TestAddAliases -
*/
func TestAddAliases(t *testing.T) {
	m := New()
	m2 := New()
	want := "command-and-control, botnet"
	want2 := []string{" botnet ", "command-and-control"}
	wantRes := []string{"command-and-control", "botnet"}
	wantRes2 := []string{"botnet", "command-and-control"}

	m.AddAliases(want)
	m2.AddAliases(want2)

	if got := m.Aliases[0]; got != wantRes[0] {
		t.Error("Fail Infrastructure Add Aliases Check. First value.")
	}

	if got := m.Aliases[1]; got != wantRes[1] {
		t.Error("Fail Infrastructure Add Aliases Check. Method does not trim secont value string.")
	}

	if got := m2.Aliases[0]; got != wantRes2[0] {
		t.Error("Fail Infrastructure Add Aliases Check. First value. Should add array of values")
	}

	if got := m2.Aliases[1]; got != wantRes2[1] {
		t.Error("Fail Infrastructure Add Aliases Check. Method does not trim secont value string. Should add array of values")
	}
}

/*
TestAddAliase -
*/
func TestAddAlias(t *testing.T) {
	m := New()
	want := "botnet"
	m.AddAlias(want)

	if got := m.Aliases[0]; got != want {
		t.Error("Fail Infrastructure Add Aliase Check")
	}
}

/*
TestSetFirstSeen -
*/
func TestSetFirstSeen(t *testing.T) {
	m := New()
	wantStr := "2019-10-04T00:00:00Z"
	want, _ := time.Parse(defs.TIME_RFC_3339_MILLI, wantStr)

	m.SetFirstSeen(want)

	if got := m.FirstSeen; got != wantStr {
		t.Error("Fail Infrastructure Set First Seen Check got:" + got + "  want: " + wantStr)
	}
}

/*
TestSetLastSeen -
*/
func TestSetLastSeen(t *testing.T) {
	m := New()
	wantStr := "2019-10-04T00:00:00Z"
	want, _ := time.Parse(defs.TIME_RFC_3339_MILLI, wantStr)

	m.SetLastSeen(want)

	if got := m.LastSeen; got != wantStr {
		t.Error("Fail Infrastructure Set Last Seen Check got:" + got + "  want: " + wantStr)
	}
}
