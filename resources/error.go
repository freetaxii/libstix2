// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package resources

import (
	"github.com/freetaxii/libstix2/resources/properties"
)

// ----------------------------------------------------------------------
//
// Define Message Type
//
// ----------------------------------------------------------------------

/*
ErrorType - This type implements the TAXII 2 Error Message and defines
all of the properties and methods needed to create and work with the TAXII Error
Message. All of the methods not defined local to this type are inherited from
the individual properties.

The following information comes directly from the TAXII 2 specification documents.

TAXII primarily relies on the standard HTTP error semantics (400-series and
500-series status codes, defined by sections 6.5 and 6.6 of [RFC7231]) to allow
TAXII Servers to indicate when an error has occurred. For example, an HTTP 404
(Not Found) status code in response to a request to get information about a
Collection means that the Collection could not be found. The tables defining the
Endpoints in sections 4 and 5 identify common errors and which response should
be used, but are not exhaustive and do not describe all possible errors.

In addition to this, TAXII defines an error message structure that is provided
in the response body when an error status is being returned. It does not,
however, define any error codes or error conditions beyond those defined by HTTP.

The error message is provided by TAXII Servers in the response body when
returning an HTTP error status and contains more information describing the
error, including a human-readable title and description, an error_code and
error_id, and a details structure to capture further structured information
about the error. All of the fields are application-specific and clients
shouldn't assume consistent meaning across TAXII Servers even if the codes, IDs,
or titles are the same.
*/
type ErrorType struct {
	properties.TitlePropertyType
	properties.DescriptionPropertyType
	ErrorID         string                 `json:"error_id,omitempty"`
	ErrorCode       string                 `json:"error_code,omitempty"`
	HTTPStatus      string                 `json:"http_status,omitempty"`
	ExternalDetails string                 `json:"external_details,omitempty"`
	Details         map[string]interface{} `json:"details,omitempty"`
}

// ----------------------------------------------------------------------
//
// Initialization Functions
//
// ----------------------------------------------------------------------

/*
NewError - This functions will create a new TAXII Error Message object and return
it as a pointer.
*/
func NewError() *ErrorType {
	var obj ErrorType
	return &obj
}

// ----------------------------------------------------------------------
//
// Public Methods - ErrorType
//
// ----------------------------------------------------------------------

/*
SetErrorID - This methods takes in a string value representing an identifier
for this particular error instnace and updates the Error ID property. A TAXII
Server might choose to assign each error occurrence it's own identifier in
order to facilitate debugging.
*/
func (r *ErrorType) SetErrorID(s string) error {
	r.ErrorID = s
	return nil
}

/*
SetErrorCode - This method takes in a string value representing the error
code for this error type and updates the Error Code property. A TAXII Server
might choose to assign a common error code to all errors of the same type.
Error codes are application-specific and not intended to be meaningful across
different TAXII Servers.
*/
func (r *ErrorType) SetErrorCode(s string) error {
	r.ErrorCode = s
	return nil
}

/*
SetHTTPStatus - This method takes in a string value representing the HTTP
status code applicable to this error and updates the HTTP Status property.
*/
func (r *ErrorType) SetHTTPStatus(s string) error {
	r.HTTPStatus = s
	return nil
}

/*
SetExternalDetails - This method takes in a string value representing a URL
that points to additional details and updates the External Details property.
For example, this could be a URL pointing to a knowledge base article
describing the error code. Absence of this field indicates that there are no
additional details.
*/
func (r *ErrorType) SetExternalDetails(s string) error {
	r.ExternalDetails = s
	return nil
}
