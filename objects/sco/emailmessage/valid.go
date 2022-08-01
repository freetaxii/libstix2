// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package emailmessage

import "fmt"

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/*
Valid - This method will verify and test all of the properties on an object
to make sure they are valid per the specification. It will return a boolean, an
integer that tracks the number of problems found, and a slice of strings that
contain the detailed results, whether good or bad.
*/
func (o *EmailMessage) Valid(debug bool) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check common base properties first
	_, pBase, dBase := o.CommonObjectProperties.ValidSDO(debug)
	problemsFound += pBase
	resultDetails = append(resultDetails, dBase...)

	if o.IsMultipart && len(o.BodyMultipart) == 0 {
		problemsFound++
		str := fmt.Sprintf("-- is_multipart is set but body_multipart is empty")
		resultDetails = append(resultDetails, str)
	} else if !o.IsMultipart && len(o.BodyMultipart) > 0 {
		problemsFound++
		str := fmt.Sprintf("-- is_multipart is not set but body_multipart is not empty")
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("-- is_multipart is ok")
		resultDetails = append(resultDetails, str)
	}

	// Verify object refs property is present
	// _, pObjectRefs, dObjectRefs := o.ObjectRefsProperty.VerifyExists()
	// problemsFound += pObjectRefs
	// resultDetails = append(resultDetails, dObjectRefs...)

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
