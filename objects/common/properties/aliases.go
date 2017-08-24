// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type AliasesPropertyType struct {
	Aliases []string `json:"aliases,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - AliasesType
// ----------------------------------------------------------------------

// AddAlias takes in one parameter
// param: s - a string value that represents an alias for something in STIX
func (this *AliasesPropertyType) AddAlias(s string) {
	if this.Aliases == nil {
		a := make([]string, 0)
		this.Aliases = a
	}
	this.Aliases = append(this.Aliases, s)
}
