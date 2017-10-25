// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

// LabelsPropertyType - A property used by one or more STIX objects that
// captures a list of labels or tags for a STIX object. On some objects the
// labels property is defined as coming from an open-vocab.
type LabelsPropertyType struct {
	Labels []string `json:"labels,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - LabelsPropertyType
// ----------------------------------------------------------------------

// AddLabel - This method takes in a string value that represents a label for a
// STIX object and adds it to the list of labels in the lables property.
func (p *LabelsPropertyType) AddLabel(s string) {
	p.Labels = append(p.Labels, s)
}
