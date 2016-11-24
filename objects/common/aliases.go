// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package common

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type AliasesType struct {
	Aliases []string `json:"aliases,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - AliasesType
// ----------------------------------------------------------------------

func (this *AliasesType) AddAlias(value string) {
	if this.Aliases == nil {
		a := make([]string, 0)
		this.Aliases = a
	}
	this.Aliases = append(this.Aliases, value)
}
