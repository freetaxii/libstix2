// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
TitleProperty - A property used by one or more TAXII resources.
*/
type TitleProperty struct {
	Title string `json:"title"`
}

// ----------------------------------------------------------------------
// Public Methods - TitleProperty
// ----------------------------------------------------------------------

/*
SetTitle - This method takes in a string value representing a title or name
and updates the title property.
*/
func (p *TitleProperty) SetTitle(s string) error {
	p.Title = s
	return nil
}

/*
GetTitle - This method returns the title.
*/
func (p *TitleProperty) GetTitle() string {
	return p.Title
}
