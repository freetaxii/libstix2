// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package discovery

import "github.com/freetaxii/libstix2/objects"

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

/*
SetContact - This methods takes in a string value representing contact
information and updates the contact property.
*/
func (o *Discovery) SetContact(s string) error {
	o.Contact = s
	return nil
}

/*
GetContact - This method returns the contact information from the contact
property.
*/
func (o *Discovery) GetContact() string {
	return o.Contact
}

/*
SetDefault - This methods takes in a string value representing a default
api-root and updates the default property.
*/
func (o *Discovery) SetDefault(s string) error {
	o.Default = s
	return nil
}

/*
GetDefault - This methods returns the default api-root.
*/
func (o *Discovery) GetDefault() string {
	return o.Default
}

/*
AddAPIRoots - This method takes in a string value, a comma separated list of
string values, or a slice of string values that represents an API Root and adds
it to the apiroots property.
*/
func (o *Discovery) AddAPIRoots(values interface{}) error {
	// if o.APIRoots == nil {
	// 	a := make([]string, 0)
	// 	o.APIRoots = a
	// }
	// o.APIRoots = append(o.APIRoots, s)
	return objects.AddValuesToList(&o.APIRoots, values)
}
