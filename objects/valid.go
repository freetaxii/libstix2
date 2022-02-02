// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

import (
	"fmt"
	"strings"
	"time"
)

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

// ValidSDO - This method will verify and test all of the properties on a STIX
// Domain Object to make sure they are valid per the specification. It will
// return a boolean, an integer that tracks the number of problems found, and a
// slice of strings that contain the detailed results, whether good or bad.
func (o *CommonObjectProperties) ValidSDO(debug bool) (bool, int, []string) {
	r := new(results)
	r.debug = debug

	// Check each property in the model
	o.checkObjectType(r)
	o.checkSpecVersion(r)
	o.checkID(r)
	o.checkCreatedByRef(r)
	o.checkCreated(r)
	o.checkModified(r)

	// Return real values not pointers
	if r.problemsFound > 0 {
		return false, r.problemsFound, r.resultDetails
	}
	return true, r.problemsFound, r.resultDetails
}

// ----------------------------------------------------------------------
// Private Common Functions
// ----------------------------------------------------------------------

// These functions will handle common logging tasks for the various checks.

func requiredButMissing(r *results, propertyName string) {
	str := fmt.Sprintf("-- the %s property is required but missing", propertyName)
	logProblem(r, str)
}

func requiredAndFound(r *results, propertyName string) {
	str := fmt.Sprintf("++ the %s property is required and is found", propertyName)
	logValid(r, str)
}

func logProblem(r *results, msg string) {
	r.problemsFound++
	r.resultDetails = append(r.resultDetails, msg)
}

func logValid(r *results, msg string) {
	if r.debug {
		r.resultDetails = append(r.resultDetails, msg)
	}
}

// ----------------------------------------------------------------------
// Private Functions
// ----------------------------------------------------------------------

// isObjectTypeValid - This function will take in a string representing an
// object type and return true or false if it is an officially supported
// object.
func isObjectTypeValid(s string) bool {
	// objectTypes := map[string]bool{
	// 	"indicator":          true,
	// }

	// if _, found := objectTypes[s]; found == true {
	// 	return true
	// }
	// return false
	return true
}

// isIDValid - This function will take in an ID and check to see if it is
// a valid identifier per the specification for a STIX object.
func isIDValid(id string) bool {
	idparts := strings.Split(id, "--")

	if idparts == nil {
		return false
	}

	// First check to see if the object type is valid, if not return false.
	if valid := isObjectTypeValid(idparts[0]); valid == false {
		// Short circuit if the object type part is wrong
		return false
	}

	// If the type is valid, then check to see if the ID is a UUID, if not return
	// false.
	valid := IsUUIDValid(idparts[1])

	return valid
}

// IsCreatedByIDValid - This function will take in an ID and check to see
// if it is a valid identifier per the specification for an identity object.
func isCreatedByIDValid(id string) bool {
	idparts := strings.Split(id, "--")

	if idparts == nil {
		return false
	}

	// First check to see if the object type is valid, if not return false.
	if idparts[0] != "identity" {
		// Short circuit if the object type part is wrong
		return false
	}

	// If the type is valid, then check to see if the ID is a UUID, if not return
	// false.
	valid := IsUUIDValid(idparts[1])

	return valid
}

// ----------------------------------------------------------------------
// Private Methods
// ----------------------------------------------------------------------

// Each of these methods will check a specific property. It is done this way
// to reduce the complexity of the main valid() function. This way all of the
// checks for each property are self contained in their own function.

func (o *CommonObjectProperties) checkObjectType(r *results) {
	if o.ObjectType == "" {
		requiredButMissing(r, "type")
	} else {
		requiredAndFound(r, "type")

		if valid := isObjectTypeValid(o.ObjectType); valid == false {
			logProblem(r, "-- the type property does not contain a valid value")
		} else {
			str := fmt.Sprintf("++ the type property contains a valid type value of \"%s\"", o.ObjectType)
			logValid(r, str)
		}
	}
}

func (o *CommonObjectProperties) checkSpecVersion(r *results) {
	if o.SpecVersion == "" {
		requiredButMissing(r, "spec_version")
	} else {
		requiredAndFound(r, "spec_version")

		if o.SpecVersion != "2.1" {
			logProblem(r, "-- the spec_version property does not contain a value of 2.1")
		} else {
			str := fmt.Sprintf("++ the spec_version property contains a valid spec_version value of \"%s\"", o.SpecVersion)
			logValid(r, str)
		}
	}
}

func (o *CommonObjectProperties) checkID(r *results) {
	if o.ID == "" {
		requiredButMissing(r, "id")
	} else {
		requiredAndFound(r, "id")

		if valid := isIDValid(o.ID); valid == false {
			logProblem(r, "-- the id property does not contain a valid identifier")
		} else {
			str := fmt.Sprintf("++ the id property contains a valid identifier value of \"%s\"", o.ID)
			logValid(r, str)
		}
	}
}

func (o *CommonObjectProperties) checkCreatedByRef(r *results) {
	if valid := isCreatedByIDValid(o.CreatedByRef); valid == false {
		logProblem(r, "-- the created_by_ref property does not contain a valid identifier")
	} else {
		str := fmt.Sprintf("++ the created_by_ref property contains a valid identifier value of \"%s\"", o.CreatedByRef)
		logValid(r, str)
	}

}

func (o *CommonObjectProperties) checkCreated(r *results) {
	if o.Created == "" {
		requiredButMissing(r, "created")
	} else {
		requiredAndFound(r, "created")

		if valid := IsTimestampValid(o.Created); valid == false {
			logProblem(r, "-- the created property does not contain a valid timestamp")
		} else {
			str := fmt.Sprintf("++ the created property contains a valid timestamp value of \"%s\"", o.Created)
			logValid(r, str)
		}
	}
}

func (o *CommonObjectProperties) checkModified(r *results) {
	if o.Modified == "" {
		requiredButMissing(r, "modified")
	} else {
		requiredAndFound(r, "modified")

		if valid := IsTimestampValid(o.Modified); valid == false {
			logProblem(r, "-- the modified property does not contain a valid timestamp")
		} else {
			str := fmt.Sprintf("++ the modified property contains a valid timestamp value of \"%s\"", o.Modified)
			logValid(r, str)
		}

		// Make sure the modified timestampe is equal to or greater than created
		if o.Created != "" {
			created, _ := time.Parse(time.RFC3339, o.Created)
			modified, _ := time.Parse(time.RFC3339, o.Modified)
			if modified.After(created) || modified.Equal(created) {
				logValid(r, "++ the modified timestamp is later than or equal to the created timestamp")
			} else {
				logProblem(r, "-- the modified timestamp is not later than or eqaul to the created timestamp")
			}
		}
	}
}
