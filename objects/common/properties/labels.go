// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type LabelsPropertyType struct {
	Labels []string `json:"labels,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - LabelsPropertyType
// ----------------------------------------------------------------------

// AddAlias takes in one parameter
// param: s - a string value that represents a label for a STIX object
func (this *LabelsPropertyType) AddLabel(s string) {
	this.Labels = append(this.Labels, s)
}
