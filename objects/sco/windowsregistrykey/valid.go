// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package windowsregistrykey

import (
	"fmt"
	"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/*
Valid - This method will verify and test all of the properties on an object
to make sure they are valid per the specification. It will return a boolean, an
integer that tracks the number of problems found, and a slice of strings that
contain the detailed results, whether good or bad.
*/
func (o *WindowsRegistryKey) Valid(debug bool) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check common SCO properties (type, spec_version, id) - these are required for SCOs
	if o.ObjectType == "" {
		problemsFound++
		resultDetails = append(resultDetails, "-- the type property is required but missing")
	} else {
		resultDetails = append(resultDetails, "++ the type property is present")
	}

	if o.SpecVersion == "" {
		problemsFound++
		resultDetails = append(resultDetails, "-- the spec_version property is required but missing")
	} else {
		resultDetails = append(resultDetails, "++ the spec_version property is present")
	}

	if o.ID == "" {
		problemsFound++
		resultDetails = append(resultDetails, "-- the id property is required but missing")
	} else {
		resultDetails = append(resultDetails, "++ the id property is present")
	}

	// Windows Registry Key specific validations
	if o.Key == "" {
		problemsFound++
		resultDetails = append(resultDetails, "-- the key property is required but missing")
	} else {
		resultDetails = append(resultDetails, "++ the key property is present")
	}

	// Validate values if present
	if o.Values != nil {
		for i, value := range o.Values {
			if value.Name != "" {
				resultDetails = append(resultDetails, "++ registry value has a name")
			}
			if value.DataType != "" {
				// Check if the data type is valid according to the specification
				if !isValidRegistryDataType(value.DataType) {
					problemsFound++
					resultDetails = append(resultDetails, "-- registry value data type is not valid")
				} else {
					resultDetails = append(resultDetails, "++ registry value data type is valid")
				}
			}
			if value.Data != "" {
				resultDetails = append(resultDetails, "++ registry value has data")
			}
			if debug {
				resultDetails = append(resultDetails, "++ processing registry value at index "+fmt.Sprintf("%d", i))
			}
		}
	}

	// Validate modified_time if present (optional field)
	if o.ModifiedTime != "" {
		if valid := objects.IsTimestampValid(o.ModifiedTime); !valid {
			problemsFound++
			resultDetails = append(resultDetails, "-- the modified_time property does not contain a valid timestamp")
		} else {
			resultDetails = append(resultDetails, "++ the modified_time property contains a valid timestamp")
		}
	}

	// Validate creator_user_ref if present (optional field)
	if o.CreatorUserRef != "" {
		if valid := objects.IsIDValid(o.CreatorUserRef); !valid {
			problemsFound++
			resultDetails = append(resultDetails, "-- the creator_user_ref property does not contain a valid identifier")
		} else {
			resultDetails = append(resultDetails, "++ the creator_user_ref property contains a valid identifier")
		}
	}

	// Validate number_of_subkeys if present (should be non-negative)
	if o.NumberOfSubkeys < 0 {
		problemsFound++
		resultDetails = append(resultDetails, "-- the number_of_subkeys property cannot be negative")
	} else {
		resultDetails = append(resultDetails, "++ the number_of_subkeys property is non-negative")
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}

// isValidRegistryDataType - Check if the registry data type is valid according to STIX spec
func isValidRegistryDataType(dataType string) bool {
	validTypes := map[string]bool{
		"REG_NONE":                       true,
		"REG_SZ":                         true,
		"REG_EXPAND_SZ":                  true,
		"REG_BINARY":                     true,
		"REG_DWORD":                      true,
		"REG_DWORD_BIG_ENDIAN":           true,
		"REG_LINK":                       true,
		"REG_MULTI_SZ":                   true,
		"REG_RESOURCE_LIST":              true,
		"REG_FULL_RESOURCE_DESCRIPTOR":   true,
		"REG_RESOURCE_REQUIREMENTS_LIST": true,
		"REG_QWORD":                      true,
	}

	_, exists := validTypes[dataType]
	return exists
}
