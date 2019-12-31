// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/* TitleProperty - A property used by one or more TAXII resources. */
type TitleProperty struct {
	Title string `json:"title"`
}

// ----------------------------------------------------------------------
// Public Methods - TitleProperty - Setters
// ----------------------------------------------------------------------

/* SetTitle - This method takes in a string value representing a title or name
and updates the title property. */
func (o *TitleProperty) SetTitle(s string) error {
	o.Title = s
	return nil
}

/* GetTitle - This method returns the title. */
func (o *TitleProperty) GetTitle() string {
	return o.Title
}
