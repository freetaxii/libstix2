// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

/*
ValidObjectType - This function will take in a STIX object type and return
true if the string represents an actual STIX object type. This is used for
determining if input from an outside source is actually a defined STIX object or
not.
*/
func ValidObjectType(t string) bool {

	var m = map[string]int{
		"attack-pattern":     1,
		"campaign":           1,
		"course-of-action":   1,
		"grouping":           1,
		"identity":           1,
		"indicator":          1,
		"infrastructure":     1,
		"intrusion-set":      1,
		"language-content":   1,
		"location":           1,
		"malware":            1,
		"malware-analysis":   1,
		"marking-definition": 1,
		"note":               1,
		"observed-data":      1,
		"opinion":            1,
		"relationship":       1,
		"report":             1,
		"sighting":           1,
		"threat-actor":       1,
		"tool":               1,
		"vulnerability":      1,
	}

	if _, ok := m[t]; ok {
		return true
	}
	return false
}

/*
GetCommonProperties - This method will return a pointer to the common
properties of this object.
*/
func (o *CommonObjectProperties) GetCommonProperties() *CommonObjectProperties {
	return o
}
