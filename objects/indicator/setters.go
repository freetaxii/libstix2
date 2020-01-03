// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package indicator

import (
	"errors"

	"github.com/freetaxii/libstix2/resources"
	"github.com/freetaxii/libstix2/timestamp"
)

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/* AddTypes - This method takes in a string value, a comma separated list of
string values, or a slice of string values that represents an indicator type and
adds it to the indicator types property. The values SHOULD come from the
indicator-type-ov open vocabulary. */
func (o *Indicator) AddTypes(values interface{}) error {
	return resources.AddValuesToList(&o.IndicatorTypes, values)
}

/* SetPattern - This method takes in a string value representing a complete and
valid STIX pattern and will set the pattern property to that value. */
func (o *Indicator) SetPattern(s string) error {
	o.Pattern = s
	return nil
}

/* SetPatternType - This method takes in a string representing the type of
pattern used in this indicator and will set the pattern_type property to that
value. The value should be one of "stix", "snort", or "yara". */
func (o *Indicator) SetPatternType(s string) error {
	if s != "stix" && s != "snort" && s != "yara" {
		return errors.New("the supplied pattern type is not one of stix, snort, or yara")
	}
	o.PatternType = s
	return nil
}

/* SetPatternVersion - This method takes in a string representing the version of
the pattern used in this indicator and will set the pattern_version property to
that value.

For patterns that do not have a formal specification, the build or code version
that the pattern is known to work with SHOULD be used.
*/
func (o *Indicator) SetPatternVersion(s string) error {
	o.PatternVersion = s
	return nil
}

/* SetValidFromToCurrentTime - This method will set the valid_from timestamp to
the current time. */
func (o *Indicator) SetValidFromToCurrentTime() error {
	o.ValidFrom = timestamp.CurrentTime("micro")
	return nil
}

/* SetValidFrom - This method will take in a timestamp in either time.Time or
string format and will set the valid_from property to that value. */
func (o *Indicator) SetValidFrom(t interface{}) error {
	ts, _ := timestamp.ToString(t, "micro")
	o.ValidFrom = ts
	return nil
}

/* SetValidUntilToCurrentTime - This method will set the valid_until time to the
current time. */
func (o *Indicator) SetValidUntilToCurrentTime() error {
	o.ValidUntil = timestamp.CurrentTime("micro")
	return nil
}

/* SetValidUntil - This method will take in a timestamp in either time.Time or
string format and will set the valid_until property to that value. */
func (o *Indicator) SetValidUntil(t interface{}) error {
	ts, _ := timestamp.ToString(t, "micro")

	// TODO check to make sure this is later than the vaild_from
	o.ValidUntil = ts
	return nil
}
