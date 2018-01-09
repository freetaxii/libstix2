// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package properties

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

/*
TitlePropertyType - A property used by one or more TAXII resources.
*/
type TitlePropertyType struct {
	Title string `json:"title"`
}

// ----------------------------------------------------------------------
// Public Methods - TitlePropertyType
// ----------------------------------------------------------------------

/*
SetTitle - This method takes in a string value representing a title or name
and updates the title property.
*/
func (p *TitlePropertyType) SetTitle(s string) error {
	p.Title = s
	return nil
}

/*
GetTitle - This method returns the title.
*/
func (p *TitlePropertyType) GetTitle() (string, error) {
	return p.Title, nil
}
