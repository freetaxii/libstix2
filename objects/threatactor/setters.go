// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package threatactor

import "github.com/freetaxii/libstix2/resources"

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/* AddTypes - This method takes in a string value, a comma separated list of
string values, or a slice of string values that represents an threat actor type
and adds it to the threat actor types property. The values SHOULD come from the
threat-actor-type-ov open vocabulary. */
func (o *ThreatActor) AddTypes(values interface{}) error {
	return resources.AddValuesToList(&o.ThreatActorTypes, values)
}

/*
SetSophistication - This method takes in a string value representing the
sophistication level of a threat actor from the threat-actor-sophistication-ov
and adds it to the sophistication property.
*/
func (o *ThreatActor) SetSophistication(s string) error {
	o.Sophistication = s
	return nil
}

/*
AddPersonalMotivation - This method takes in a string value representing the
motivation of a threat actor from the threat-actor-motivation-ov and adds it
to the personal motivations property.
*/
func (o *ThreatActor) AddPersonalMotivation(s string) error {
	o.PersonalMotivations = append(o.PersonalMotivations, s)
	return nil
}
