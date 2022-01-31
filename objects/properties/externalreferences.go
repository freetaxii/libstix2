// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

import "fmt"

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/*
ExternalReferencesProperty - A property used by one or more STIX objects that
captures a list of external references as defined by STIX.
*/
type ExternalReferencesProperty struct {
	ExternalReferences []ExternalReference `json:"external_references,omitempty" bson:"external_references,omitempty"`
}

/*
ExternalReference - This type defines all of the properties associated with
the STIX External Reference type. All of the methods not defined local to this
type are inherited from the individual properties.
*/
type ExternalReference struct {
	SourceName  string            `json:"source_name,omitempty" bson:"source_name,omitempty"`
	Description string            `json:"description,omitempty" bson:"description,omitempty"`
	URL         string            `json:"url,omitempty" bson:"url,omitempty"`
	Hashes      map[string]string `json:"hashes,omitempty" bson:"hashes,omitempty"`
	ExternalID  string            `json:"external_id,omitempty" bson:"external_id,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - ExternalReferencesProperty - Setters
// ----------------------------------------------------------------------

/*
NewExternalReference - This method creates a new external reference and
returns a reference to a slice location. This will enable the code to update an
object located at that slice location.
*/
func (o *ExternalReferencesProperty) NewExternalReference() (*ExternalReference, error) {
	var s ExternalReference

	// if o.ExternalReferences == nil {
	// 	a := make([]ExternalReference, 0)
	// 	o.ExternalReferences = a
	// }

	positionThatAppendWillUse := len(o.ExternalReferences)
	o.ExternalReferences = append(o.ExternalReferences, s)
	return &o.ExternalReferences[positionThatAppendWillUse], nil
}

// ----------------------------------------------------------------------
// Public Methods - ExternalReference - Setters
// ----------------------------------------------------------------------

/*
SetSourceName - This method takes in a string value representing the name of
a source for an external reference and updates the source name property.
*/
func (o *ExternalReference) SetSourceName(s string) error {
	o.SourceName = s
	return nil
}

/*
GetSourceName - This method will return the source name.
*/
func (o *ExternalReference) GetSourceName() string {
	return o.SourceName
}

/*
SetDescription - This method takes in a string value representing a text
description and updates the description property.
*/
func (o *ExternalReference) SetDescription(s string) error {
	o.Description = s
	return nil
}

/*
GetDescription - This method returns the description for an object as a
string.
*/
func (o *ExternalReference) GetDescription() string {
	return o.Description
}

/*
SetURL - This method takes in a string value representing a URL location of a
source for an external reference and updates the URL property.
*/
func (o *ExternalReference) SetURL(s string) error {
	o.URL = s
	return nil
}

/*
GetURL - This method returns the URL for this external reference.
*/
func (o *ExternalReference) GetURL() string {
	return o.URL
}

/*
AddHash - This method takes in two parameters and adds the hash to the map.
The first is a string value representing a hash type from the STIX hashes
vocabulary. The second is a string value representing the actual hash of the
content from the remote external reference.
*/
func (o *ExternalReference) AddHash(k, v string) error {
	if o.Hashes == nil {
		m := make(map[string]string, 0)
		o.Hashes = m
	}
	o.Hashes[k] = v
	return nil
}

/*
SetExternalID - This method takes in a string value representing an general
purpose id in a remote system for the source of this external reference and
updates the external id property.
*/
func (o *ExternalReference) SetExternalID(s string) error {
	o.ExternalID = s
	return nil
}

/*
GetExternalID - This method returns the external id for this reference.
*/
func (o *ExternalReference) GetExternalID() string {
	return o.ExternalID
}

// ----------------------------------------------------------------------
// Public Methods - ExternalReferencesProperty - Checks
// ----------------------------------------------------------------------

/*
Compare - This function will compare two properties to make sure they are the
same and will return a boolean, an integer that tracks the number of problems
found, and a slice of strings that contain the detailed results, whether good or
bad.
*/
func (o *ExternalReferencesProperty) Compare(obj2 *ExternalReferencesProperty) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check External References
	if len(o.ExternalReferences) != len(obj2.ExternalReferences) {
		problemsFound++
		str := fmt.Sprintf("-- The number of entries in external references do not match: %d | %d", len(o.ExternalReferences), len(obj2.ExternalReferences))
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The number of entries in external references match: %d | %d", len(o.ExternalReferences), len(obj2.ExternalReferences))
		resultDetails = append(resultDetails, str)
		for index := range o.ExternalReferences {

			// Check External Reference Source Name
			if o.ExternalReferences[index].SourceName != obj2.ExternalReferences[index].SourceName {
				problemsFound++
				str := fmt.Sprintf("-- The source name values do not match: %s | %s", o.ExternalReferences[index].SourceName, obj2.ExternalReferences[index].SourceName)
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ The source name values match: %s | %s", o.ExternalReferences[index].SourceName, obj2.ExternalReferences[index].SourceName)
				resultDetails = append(resultDetails, str)
			}

			// Check External Reference Descriptions
			if o.ExternalReferences[index].Description != obj2.ExternalReferences[index].Description {
				problemsFound++
				str := fmt.Sprintf("-- The description values do not match: %s | %s", o.ExternalReferences[index].Description, obj2.ExternalReferences[index].Description)
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ The description values match: %s | %s", o.ExternalReferences[index].Description, obj2.ExternalReferences[index].Description)
				resultDetails = append(resultDetails, str)
			}

			// Check External Reference URLs
			if o.ExternalReferences[index].URL != obj2.ExternalReferences[index].URL {
				problemsFound++
				str := fmt.Sprintf("-- The url values do not match: %s | %s", o.ExternalReferences[index].URL, obj2.ExternalReferences[index].URL)
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ The url values match: %s | %s", o.ExternalReferences[index].URL, obj2.ExternalReferences[index].URL)
				resultDetails = append(resultDetails, str)
			}

			// Check External Reference Hashes
			if len(o.ExternalReferences[index].Hashes) != len(obj2.ExternalReferences[index].Hashes) {
				problemsFound++
				str := fmt.Sprintf("-- The number of entries in hashes do not match: %d | %d", len(o.ExternalReferences[index].Hashes), len(obj2.ExternalReferences[index].Hashes))
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ The number of entries in hashes match: %d | %d", len(o.ExternalReferences[index].Hashes), len(obj2.ExternalReferences[index].Hashes))
				resultDetails = append(resultDetails, str)

				// If lengths are the same, then check each value
				for key := range o.ExternalReferences[index].Hashes {
					if o.ExternalReferences[index].Hashes[key] != obj2.ExternalReferences[index].Hashes[key] {
						problemsFound++
						str := fmt.Sprintf("-- The hash values do not match: %s | %s", o.ExternalReferences[index].Hashes[key], obj2.ExternalReferences[index].Hashes[key])
						resultDetails = append(resultDetails, str)
					} else {
						str := fmt.Sprintf("++ The hash values match: %s | %s", o.ExternalReferences[index].Hashes[key], obj2.ExternalReferences[index].Hashes[key])
						resultDetails = append(resultDetails, str)
					}
				}
			}

			// Check External Reference External IDs
			if o.ExternalReferences[index].ExternalID != obj2.ExternalReferences[index].ExternalID {
				problemsFound++
				str := fmt.Sprintf("-- The external id values do not match: %s | %s", o.ExternalReferences[index].ExternalID, obj2.ExternalReferences[index].ExternalID)
				resultDetails = append(resultDetails, str)
			} else {
				str := fmt.Sprintf("++ The external id values match: %s | %s", o.ExternalReferences[index].ExternalID, obj2.ExternalReferences[index].ExternalID)
				resultDetails = append(resultDetails, str)
			}
		}
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
