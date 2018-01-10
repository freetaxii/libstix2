// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package objects

import (
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
IntrusionSetType - This type implements the STIX 2 Intrusion Set SDO and defines
all of the properties methods needed to create and work with the STIX Intrusion Set
SDO. All of the methods not defined local to this type are inherited from
the individual properties.

The following information comes directly from the STIX 2 specification documents.

An Intrusion Set is a grouped set of adversarial behaviors and resources with
common properties that is believed to be orchestrated by a single organization.
An Intrusion Set may capture multiple Campaigns or other activities that are all
tied together by shared attributes indicating a common known or unknown Threat
Actor. New activity can be attributed to an Intrusion Set even if the Threat
Actors behind the attack are not known. Threat Actors can move from supporting
one Intrusion Set to supporting another, or they may support multiple Intrusion
Sets.

Where a Campaign is a set of attacks over a period of time against a specific
set of targets to achieve some objective, an Intrusion Set is the entire attack
package and may be used over a very long period of time in multiple Campaigns to
achieve potentially multiple purposes.

While sometimes an Intrusion Set is not active, or changes focus, it is usually
difficult to know if it has truly disappeared or ended. Analysts may have
varying level of fidelity on attributing an Intrusion Set back to Threat Actors
and may be able to only attribute it back to a nation state or perhaps back to
an organization within that nation state.
*/
type IntrusionSetType struct {
	properties.CommonObjectPropertiesType
	properties.NamePropertyType
	properties.DescriptionPropertyType
	properties.AliasesPropertyType
	properties.FirstSeenPropertyType
	properties.LastSeenPropertyType
	properties.GoalsPropertyType
	properties.ResourceLevelPropertyType
	properties.PrimaryMotivationPropertyType
	properties.SecondaryMotivationsPropertyType
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
InitIntrusionSet - This function will create a new STIX Intrusion Set object
and return it as a pointer.
*/
func InitIntrusionSet(ver string) *IntrusionSetType {
	var obj IntrusionSetType
	obj.InitObjectProperties("intrusion-set", ver)
	return &obj
}

// ----------------------------------------------------------------------
// Public Methods - IntrusionSetType
// ----------------------------------------------------------------------
