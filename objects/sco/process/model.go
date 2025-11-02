// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package process

import (
	"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
Process - This type implements the STIX 2.1 Process SCO and defines all of the
properties and methods needed to create and work with this object. All of the
methods not defined local to this type are inherited from the individual
properties.

Reference: STIX 2.1 specification section 4.11
*/
type Process struct {
	objects.CommonObjectProperties
	Pid                  int                `json:"pid,omitempty" bson:"pid,omitempty"`
	CreatedTime          string             `json:"created_time,omitempty" bson:"created_time,omitempty"`
	Cwd                  string             `json:"cwd,omitempty" bson:"cwd,omitempty"`
	CommandLine          string             `json:"command_line,omitempty" bson:"command_line,omitempty"`
	EnvironmentVariables map[string]string  `json:"environment_variables,omitempty" bson:"environment_variables,omitempty"`
	OpenedConnectionRefs []string           `json:"opened_connection_refs,omitempty" bson:"opened_connection_refs,omitempty"`
	CreatorUserRef       string             `json:"creator_user_ref,omitempty" bson:"creator_user_ref,omitempty"`
	ImageRef             string             `json:"image_ref,omitempty" bson:"image_ref,omitempty"`
	ParentRef            string             `json:"parent_ref,omitempty" bson:"parent_ref,omitempty"`
	ChildRefs            []string           `json:"child_refs,omitempty" bson:"child_refs,omitempty"`
	OwnerSid             string             `json:"owner_sid,omitempty" bson:"owner_sid,omitempty"`
	WindowsProcessExt    *WindowsProcessExt `json:"extensions,omitempty" bson:"extensions,omitempty"`
	WindowsServiceExt    *WindowsServiceExt `json:"-" bson:"-"` // Not part of JSON, internal
	UnixAccountExt       *UnixProcessExt    `json:"-" bson:"-"` // Not part of JSON, internal
}

// WindowsProcessExt - Windows Process Extension
type WindowsProcessExt struct {
	AslrEnabled    bool   `json:"aslr_enabled,omitempty" bson:"aslr_enabled,omitempty"`
	DepEnabled     bool   `json:"dep_enabled,omitempty" bson:"dep_enabled,omitempty"`
	Priority       string `json:"priority,omitempty" bson:"priority,omitempty"`
	OwnerSid       string `json:"owner_sid,omitempty" bson:"owner_sid,omitempty"`
	WindowTitle    string `json:"window_title,omitempty" bson:"window_title,omitempty"`
	StartupInfo    string `json:"startup_info,omitempty" bson:"startup_info,omitempty"`
	IntegrityLevel string `json:"integrity_level,omitempty" bson:"integrity_level,omitempty"`
}

// WindowsServiceExt - Windows Service Extension
type WindowsServiceExt struct {
	ServiceName   string      `json:"service_name,omitempty" bson:"service_name,omitempty"`
	DisplayName   string      `json:"display_name,omitempty" bson:"display_name,omitempty"`
	ServiceType   string      `json:"service_type,omitempty" bson:"service_type,omitempty"`
	StartType     string      `json:"start_type,omitempty" bson:"start_type,omitempty"`
	ServiceStatus string      `json:"service_status,omitempty" bson:"service_status,omitempty"`
	StartedAs     []string    `json:"-" bson:"-"` // Internal, not in STIX spec
	Actions       []string    `json:"-" bson:"-"` // Internal, not in STIX spec
	Info          interface{} `json:"-" bson:"-"` // Internal, not in STIX spec
}

// UnixProcessExt - UNIX Process Extension
type UnixProcessExt struct {
	StartedAs []string    `json:"-" bson:"-"` // Internal, not in STIX spec
	Actions   []string    `json:"-" bson:"-"` // Internal, not in STIX spec
	Info      interface{} `json:"-" bson:"-"` // Internal, not in STIX spec
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *Process) GetPropertyList() []string {
	return []string{"pid", "created_time", "cwd", "command_line", "environment_variables",
		"opened_connection_refs", "creator_user_ref", "image_ref", "parent_ref", "child_refs",
		"owner_sid"}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Process SCO and return it as a
pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *Process {
	var obj Process
	obj.InitSCO("process")
	return &obj
}
