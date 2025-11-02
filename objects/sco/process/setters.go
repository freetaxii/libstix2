// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package process

import (
	"errors"
)

// ----------------------------------------------------------------------
// Public Methods - Process - Setters
// ----------------------------------------------------------------------

/*
SetPid - This method will set the process ID for the process object.
*/
func (o *Process) SetPid(i int) error {
	if i >= 0 {
		o.Pid = i
	}
	return nil
}

/*
SetCreatedTime - This method will set the creation time for the process object.
*/
func (o *Process) SetCreatedTime(t interface{}) error {
	switch v := t.(type) {
	case string:
		o.CreatedTime = v
	case nil:
		o.CreatedTime = ""
	default:
		return errors.New("created_time must be a string or nil")
	}
	return nil
}

/*
SetCwd - This method will set the current working directory for the process object.
*/
func (o *Process) SetCwd(s string) error {
	o.Cwd = s
	return nil
}

/*
SetCommandLine - This method will set the command line for the process object.
*/
func (o *Process) SetCommandLine(s string) error {
	o.CommandLine = s
	return nil
}

/*
SetEnvironmentVariables - This method will set the environment variables map for the process object.
*/
func (o *Process) SetEnvironmentVariables(vars map[string]string) error {
	o.EnvironmentVariables = vars
	return nil
}

/*
AddEnvironmentVariable - This method will add a single environment variable to the process object.
*/
func (o *Process) AddEnvironmentVariable(key, value string) error {
	if o.EnvironmentVariables == nil {
		o.EnvironmentVariables = make(map[string]string)
	}
	if key != "" {
		o.EnvironmentVariables[key] = value
	}
	return nil
}

/*
AddOpenedConnectionRef - This method will add an object reference to the opened_connection_refs list.
*/
func (o *Process) AddOpenedConnectionRef(ref string) error {
	if ref != "" {
		o.OpenedConnectionRefs = append(o.OpenedConnectionRefs, ref)
	}
	return nil
}

/*
SetOpenedConnectionRefs - This method will set the opened_connection_refs list for the process object.
*/
func (o *Process) SetOpenedConnectionRefs(refs []string) error {
	o.OpenedConnectionRefs = refs
	return nil
}

/*
SetCreatorUserRef - This method will set the creator user reference for the process object.
*/
func (o *Process) SetCreatorUserRef(s string) error {
	o.CreatorUserRef = s
	return nil
}

/*
SetImageRef - This method will set the image reference for the process object.
*/
func (o *Process) SetImageRef(s string) error {
	o.ImageRef = s
	return nil
}

/*
SetParentRef - This method will set the parent process reference for the process object.
*/
func (o *Process) SetParentRef(s string) error {
	o.ParentRef = s
	return nil
}

/*
AddChildRef - This method will add a child process reference to the child_refs list.
*/
func (o *Process) AddChildRef(ref string) error {
	if ref != "" {
		o.ChildRefs = append(o.ChildRefs, ref)
	}
	return nil
}

/*
SetChildRefs - This method will set the child_refs list for the process object.
*/
func (o *Process) SetChildRefs(refs []string) error {
	o.ChildRefs = refs
	return nil
}

/*
SetOwnerSid - This method will set the owner SID for the process object.
*/
func (o *Process) SetOwnerSid(s string) error {
	o.OwnerSid = s
	return nil
}
