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
AttackPatternType - This type implements the STIX 2 Attack Pattern SDO and defines
all of the properties methods needed to create and work with the STIX Attack Pattern
SDO. All of the methods not defined local to this type are inherited from
the individual properties.

The following information comes directly from the STIX 2 specification documents.

Attack Patterns are a type of TTP that describe ways that adversaries attempt to
compromise targets. Attack Patterns are used to help categorize attacks,
generalize specific attacks to the patterns that they follow, and provide
detailed information about how attacks are performed. An example of an attack
pattern is "spear phishing": a common type of attack where an attacker sends a
carefully crafted e-mail message to a party with the intent of getting them to
click a link or open an attachment to deliver malware. Attack Patterns can also
be more specific; spear phishing as practiced by a particular threat actor
(e.g., they might generally say that the target won a contest) can also be an
Attack Pattern.

The Attack Pattern SDO contains textual descriptions of the pattern along with
references to externally-defined taxonomies of attacks such as CAPEC [CAPEC].
Relationships from Attack Pattern can be used to relate it to what it targets
(Vulnerabilities and Identities) and which tools and malware use it (Tool and
Malware).
*/
type AttackPatternType struct {
	properties.CommonObjectPropertiesType
	properties.NamePropertyType
	properties.DescriptionPropertyType
	properties.KillChainPhasesPropertyType
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
InitAttackPattern - This function will create a new STIX Attack Pattern object
and return it as a pointer.
*/
func InitAttackPattern(ver string) *AttackPatternType {
	var obj AttackPatternType
	obj.InitObjectProperties("attack-pattern", ver)
	return &obj
}

// ----------------------------------------------------------------------
// Public Methods - AttackPatternType
// ----------------------------------------------------------------------
