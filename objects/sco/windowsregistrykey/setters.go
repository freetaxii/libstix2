// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package windowsregistrykey

import (
	"errors"
)

// ----------------------------------------------------------------------
// Public Methods - WindowsRegistryKey - Setters
// ----------------------------------------------------------------------

/*
SetKey - This method will set the key for the Windows registry key object.
The key field is required for this object.
*/
func (o *WindowsRegistryKey) SetKey(s string) error {
	if s == "" {
		return errors.New("the registry key cannot be empty")
	}
	o.Key = s
	return nil
}

/*
AddValue - This method will add a Windows registry value to the values list.
*/
func (o *WindowsRegistryKey) AddValue(value WindowsRegistryValue) error {
	o.Values = append(o.Values, value)
	return nil
}

/*
SetModifiedTime - This method will set the last modified time for this registry key.
*/
func (o *WindowsRegistryKey) SetModifiedTime(t interface{}) error {
	switch v := t.(type) {
	case string:
		o.ModifiedTime = v
	case nil:
		o.ModifiedTime = ""
	default:
		return errors.New("modified time must be a string or nil")
	}
	return nil
}

/*
SetCreatorUserRef - This method will set the user account reference that created this registry key.
*/
func (o *WindowsRegistryKey) SetCreatorUserRef(s string) error {
	if s != "" {
		o.CreatorUserRef = s
	}
	return nil
}

/*
SetNumberOfSubkeys - This method will set the number of subkeys contained by this registry key.
*/
func (o *WindowsRegistryKey) SetNumberOfSubkeys(i int) error {
	if i >= 0 {
		o.NumberOfSubkeys = i
	}
	return nil
}

/*
SetValueName - This method will set the name for a registry value at the given index.
*/
func (o *WindowsRegistryKey) SetValueName(index int, name string) error {
	if index < 0 || index >= len(o.Values) {
		return errors.New("invalid index for registry value")
	}
	o.Values[index].Name = name
	return nil
}

/*
SetValueData - This method will set the data for a registry value at the given index.
*/
func (o *WindowsRegistryKey) SetValueData(index int, data string) error {
	if index < 0 || index >= len(o.Values) {
		return errors.New("invalid index for registry value")
	}
	o.Values[index].Data = data
	return nil
}

/*
SetValueDataType - This method will set the data type for a registry value at the given index.
*/
func (o *WindowsRegistryKey) SetValueDataType(index int, dataType string) error {
	if index < 0 || index >= len(o.Values) {
		return errors.New("invalid index for registry value")
	}
	o.Values[index].DataType = dataType
	return nil
}

/*
SetValueHive - This method will set the hive for a registry value at the given index.
*/
func (o *WindowsRegistryKey) SetValueHive(index int, hive string) error {
	if index < 0 || index >= len(o.Values) {
		return errors.New("invalid index for registry value")
	}
	o.Values[index].Hive = hive
	return nil
}
