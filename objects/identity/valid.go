// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package identity

import (
	"fmt"

	"github.com/freetaxii/libstix2/vocabs"
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
func (o *Identity) Valid(debug bool) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check common base properties first
	_, pBase, dBase := o.CommonObjectProperties.ValidSDO(debug)
	problemsFound += pBase
	resultDetails = append(resultDetails, dBase...)

	// Verify object Name property is present
	// _, pName, dName := o.NameProperty.VerifyExists()
	// problemsFound += pName
	// resultDetails = append(resultDetails, dName...)

	if o.IdentityClass == "" {
		// in the STIX 2.1 definition, these are required, but many real-world objects do not contain these fields.
		// TODO: can make this into a "strict" validation mechanism
		// problemsFound++
		str := fmt.Sprintf("-- The identity_class property is required but missing")
		resultDetails = append(resultDetails, str)
	} else {
		// Validate that identity_class is from the vocabulary
		validVocab := vocabs.GetIdentityClassVocab()
		if !validVocab[o.IdentityClass] {
			// this is a SHOULD not a MUST so we won't add it as a problem
			// problemsFound++
			str := fmt.Sprintf("** The identity_class '%s' is not in the allowed vocabulary", o.IdentityClass)
			resultDetails = append(resultDetails, str)
		} else {
			str := fmt.Sprintf("++ The identity_class property is required and is present")
			resultDetails = append(resultDetails, str)
		}
	}

	// Validate sectors if present
	if len(o.Sectors) > 0 {
		validVocab := vocabs.GetIndustrySectorVocab()
		for _, sector := range o.Sectors {
			if !validVocab[sector] {
				// this is a SHOULD not a MUST so we won't add it as a problem
				// problemsFound++
				str := fmt.Sprintf("** The sector '%s' is not in the allowed vocabulary", sector)
				resultDetails = append(resultDetails, str)
			}
		}
	}

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
