// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type ConfidencePropertyType struct {
	Confidence int `json:"confidence,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - ConfidencePropertyType
// ----------------------------------------------------------------------

// SetConfidence takes in one parameter
// param: i - an integer representing a STIX confidence level 0-100
func (this *ConfidencePropertyType) SetConfidence(i int) {
	this.Confidence = i
}

func (this *ConfidencePropertyType) GetConfidence() int {
	return this.Confidence
}
