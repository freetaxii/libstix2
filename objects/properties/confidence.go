// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
//
// Types
//
// ----------------------------------------------------------------------

/*
ConfidencePropertyType - A property used by one or more STIX objects that
captures the STIX confidence score, which is a value from 0-100.
*/
type ConfidencePropertyType struct {
	Confidence int `json:"confidence,omitempty"`
}

// ----------------------------------------------------------------------
//
// Public Methods - ConfidencePropertyType
//
// ----------------------------------------------------------------------

/*
SetConfidence - This method takes in an integer representing a STIX
confidence level 0-100 and updates the Confidence property.
*/
func (p *ConfidencePropertyType) SetConfidence(i int) error {
	p.Confidence = i
	return nil
}

/*
GetConfidence - This method returns the confidence value as an integer.
*/
func (p *ConfidencePropertyType) GetConfidence() int {
	return p.Confidence
}
