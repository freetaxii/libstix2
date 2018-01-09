// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package resources

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// NewDiscovery - This function will create a new TAXII Discovery object.
func NewDiscovery() *DiscoveryType {
	var obj DiscoveryType
	return &obj
}

// NewStatus - This function will create a new TAXII Status object.
// func NewStatus() StatusType {
// 	return status.New()
// }

// NewError - This functions will create a new TAXII Error Message object.
func NewError() *ErrorType {
	var obj ErrorType
	return &obj
}
