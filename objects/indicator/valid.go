// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package indicator

import (
	"fmt"
	"time"

	"github.com/freetaxii/libstix2/resources"
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
func (o *Indicator) Valid() (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check common base properties first
	_, pBase, dBase := o.CommonObjectProperties.ValidSDO()
	problemsFound += pBase
	resultDetails = append(resultDetails, dBase...)

	if len(o.IndicatorTypes) == 0 {
		problemsFound++
		str := fmt.Sprintf("-- The indicator types property is required but missing")
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The indicator types property is required and is present")
		resultDetails = append(resultDetails, str)
	}

	if o.Pattern == "" {
		problemsFound++
		str := fmt.Sprintf("-- The pattern property is required but missing")
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The pattern property is required and is present")
		resultDetails = append(resultDetails, str)
	}

	// TODO, check value to see if it comes from open vocabulary
	if o.PatternType == "" {
		problemsFound++
		str := fmt.Sprintf("-- The pattern type property is required but missing")
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The pattern type property is required and is present")
		resultDetails = append(resultDetails, str)
	}

	if o.ValidFrom == "" {
		problemsFound++
		str := fmt.Sprintf("-- The valid from property is required but missing")
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ The valid from property is required and is present")
		resultDetails = append(resultDetails, str)
	}

	if valid := resources.IsTimestampValid(o.ValidFrom); valid == false {
		problemsFound++
		str := fmt.Sprintf("-- the valid from property does not contain a valid STIX timestamp")
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the valid from property does contain a valid STIX timestamp")
		resultDetails = append(resultDetails, str)
	}

	if valid := resources.IsTimestampValid(o.ValidUntil); valid == false {
		problemsFound++
		str := fmt.Sprintf("-- the valid until property does not contain a valid STIX timestamp")
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the valid until property does contain a valid STIX timestamp")
		resultDetails = append(resultDetails, str)
	}

	validFrom, _ := time.Parse(time.RFC3339, o.ValidFrom)
	validUntil, _ := time.Parse(time.RFC3339, o.ValidUntil)
	if yes := validUntil.After(validFrom); yes != true {
		problemsFound++
		str := fmt.Sprintf("-- the valid until timestamp is not later than the valid from timestamp")
		resultDetails = append(resultDetails, str)
	} else {
		str := fmt.Sprintf("++ the valid until timestamp is later than the valid from timestamp")
		resultDetails = append(resultDetails, str)
	}

	return true, 0, resultDetails
}
