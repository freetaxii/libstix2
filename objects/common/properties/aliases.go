// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

// AliasesPropertyType - A property used by one or more STIX objects.
type AliasesPropertyType struct {
	Aliases []string `json:"aliases,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - AliasesType
// ----------------------------------------------------------------------

// AddAlias - This method takes in a takes in a string value that represents an
// alias for something in STIX and adds it to the property.
func (this *AliasesPropertyType) AddAlias(s string) {
	this.Aliases = append(this.Aliases, s)
}
