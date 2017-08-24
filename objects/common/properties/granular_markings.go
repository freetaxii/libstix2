// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type GranularMarkingsPropertyType struct {
	Granular_markings []GranularMarkingType `json:"granular_markings,omitempty"`
}

type GranularMarkingType struct {
	LangPropertyType
	Marking_ref string   `json:"marking_ref,omitempty"`
	Selectors   []string `json:"selectors,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - GranularMarkingType
// ----------------------------------------------------------------------

// SetMarkingRef takes in one parameter
// param: s - a string value representing a STIX Identifier
func (this *GranularMarkingType) SetMarkingRef(s string) {
	this.Marking_ref = s
}

func (this *GranularMarkingType) GetMarkingRef() string {
	return this.Marking_ref
}

// AddSelector takes in one parameter
// param: s - a string value representing a STIX granular marking selector
func (this *GranularMarkingType) AddSelector(s string) {
	if this.Selectors == nil {
		a := make([]string, 0)
		this.Selectors = a
	}
	this.Selectors = append(this.Selectors, s)
}
