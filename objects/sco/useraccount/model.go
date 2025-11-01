// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package useraccount

import (
"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
UserAccount - This type implements the STIX 2.1 UserAccount SCO and defines all of the
properties and methods needed to create and work with this object. All of the
methods not defined local to this type are inherited from the individual
properties.

Reference: STIX 2.1 specification section 6.14

TODO: Complete implementation of all properties per specification
*/
type UserAccount struct {
objects.CommonObjectProperties
// TODO: Add specific properties for UserAccount based on STIX 2.1 spec section 6.14
UserID           string   `json:"user_id" bson:"user_id"`
Credential       string   `json:"credential,omitempty" bson:"credential,omitempty"`
AccountLogin     string   `json:"account_login,omitempty" bson:"account_login,omitempty"`
AccountType      string   `json:"account_type,omitempty" bson:"account_type,omitempty"`
DisplayName      string   `json:"display_name,omitempty" bson:"display_name,omitempty"`
IsServiceAccount bool     `json:"is_service_account,omitempty" bson:"is_service_account,omitempty"`
IsPrivileged     bool     `json:"is_privileged,omitempty" bson:"is_privileged,omitempty"`
CanEscalatePrivs bool     `json:"can_escalate_privs,omitempty" bson:"can_escalate_privs,omitempty"`
IsDisabled       bool     `json:"is_disabled,omitempty" bson:"is_disabled,omitempty"`
AccountCreated   string   `json:"account_created,omitempty" bson:"account_created,omitempty"`
AccountExpires   string   `json:"account_expires,omitempty" bson:"account_expires,omitempty"`
CredentialLastChanged string `json:"credential_last_changed,omitempty" bson:"credential_last_changed,omitempty"`
AccountFirstLogin string   `json:"account_first_login,omitempty" bson:"account_first_login,omitempty"`
AccountLastLogin string   `json:"account_last_login,omitempty" bson:"account_last_login,omitempty"`
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *UserAccount) GetPropertyList() []string {
// TODO: Update with actual property names
return []string{}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX UserAccount SCO and return it as a
pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *UserAccount {
var obj UserAccount
obj.InitSCO("useraccount")
return &obj
}
