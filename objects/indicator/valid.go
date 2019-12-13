// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package indicator

import (
	"errors"
	"time"

	"github.com/freetaxii/libstix2/timestamp"
)

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/* Valid - This method will verify and test all of the properties on the object
to make sure they are valid per the specification. */
func (o *Indicator) Valid() (bool, error) {

	// Check common base properties first
	if valid, err := o.CommonObjectProperties.Valid(); valid != true {
		return false, err
	}

	// Check Indicator Specific Properties
	if valid, err := o.validIndicatorType(); valid == false {
		return valid, err
	}

	if valid, err := o.validPattern(); valid == false {
		return valid, err
	}

	if valid, err := o.validPatternType(); valid == false {
		return valid, err
	}

	if valid, err := o.validValidFrom(); valid == false {
		return valid, err
	}

	if valid, err := o.validValidUntil(); valid == false {
		return valid, err
	}

	return true, nil
}

// ----------------------------------------------------------------------
// Private Methods
// ----------------------------------------------------------------------

func (o *Indicator) validIndicatorType() (bool, error) {
	if len(o.IndicatorTypes) == 0 {
		return false, errors.New("the indicator types property is required, but missing")
	}
	return true, nil
}

// First check to see if present and then see if it is valid
func (o *Indicator) validPattern() (bool, error) {
	if o.Pattern == "" {
		return false, errors.New("the pattern property is required, but missing")
	} //else {
	// TODO verify the pattern is correct
	//}
	return true, nil
}

// First check to see if present and then see if it is valid
func (o *Indicator) validPatternType() (bool, error) {
	if o.PatternType == "" {
		return false, errors.New("the pattern_type property is required, but missing")
	}
	if o.PatternType != "stix" && o.PatternType != "snort" && o.PatternType != "yara" {
		return false, errors.New("pattern_type contains a value other than stix, snort, or yara")
	}
	return true, nil
}

// First check to see if present and then see if it is valid
func (o *Indicator) validValidFrom() (bool, error) {
	if o.ValidFrom == "" {
		return false, errors.New("the valid_from property is required, but missing")
	}
	if valid := timestamp.Valid(o.ValidFrom); valid == false {
		return false, errors.New("the valid_from property does not contain a valid STIX timestamp")

	}
	return true, nil
}

// If the value_until property is populated, we need to check it. First lets
// see if it is in a valid format and then see if it is newer than the
// timestamp in the valid_from property.
func (o *Indicator) validValidUntil() (bool, error) {
	if o.ValidUntil != "" {
		// Valid From must be present if Valid Until is present
		if valid := timestamp.Valid(o.ValidFrom); valid != true {
			return false, errors.New("the valid_from property does not contain a valid STIX timestamp")
		}
		validFrom, _ := time.Parse(time.RFC3339, o.ValidFrom)
		validUntil, _ := time.Parse(time.RFC3339, o.ValidUntil)
		if yes := validUntil.After(validFrom); yes != true {
			return false, errors.New("the valid_until timestamp is not later than the valid_from timestamp")
		}
	}
	return true, nil
}
