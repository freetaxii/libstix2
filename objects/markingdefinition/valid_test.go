package markingdefinition

import (
	"testing"
)

func TestValid1(t *testing.T) {
	m := New()
	want := false
	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail Malware Object should be invalid when empty")
		t.Log(err)
	}
}

func TestValid2(t *testing.T) {
	m := New()
	want := true
	m.SetName("test")
	m.SetDefinitionType("tlp")
	m.SetDefinition(m.GetDefinitionType(), "red")
	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail Malware Object should be invalid when empty")
		t.Log(err)
	}
}

func TestValid3(t *testing.T) {
	m := New()
	want := false
	m.SetName("test")
	m.SetDefinition("tlp", "red")
	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail Malware Object should be invalid when empty")
		t.Log(err)
	}
}

func TestValid4(t *testing.T) {
	m := New()
	want := false
	m.SetDefinitionType("tlp")
	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail Malware Object should be invalid when empty")
		t.Log(err)
	}
}

func TestValid5(t *testing.T) {
	m := New()
	want := false
	m.SetName("test")
	m.SetDefinitionType("tlp2")
	m.SetDefinition("tlp", "red")
	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail Malware Object should be invalid when empty")
		t.Log(err)
	}
}

func TestValid6(t *testing.T) {
	m := New()
	want := false
	m.SetDefinitionType("tlp")
	m.SetDefinition("tlp", "red")
	if got, _, err := m.Valid(false); got != want {
		t.Error("Fail Malware Object should be invalid when empty")
		t.Log(err)
	}
}
