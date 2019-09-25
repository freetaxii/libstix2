// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package status

import (
	"github.com/freetaxii/libstix2/resources/properties"
	"github.com/freetaxii/libstix2/timestamp"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

/*
Status - This type implements the TAXII 2 Status Resource and defines
all of the properties and methods needed to create and work with the TAXII
Status Resource.

The following information comes directly from the TAXII 2 specification documents.

The status resource represents information about a request to add objects to a
Collection. It contains information about the status of the request, such as
whether or not it's completed (status) and MAY contain the status of individual
objects within the request (i.e. whether they are still pending, completed and
failed, or completed and succeeded).
*/
type Status struct {
	properties.IDProperty
	Status           string          `json:"status,omitempty"`
	RequestTimestamp string          `json:"request_timestamp,omitempty"`
	TotalCount       int             `json:"total_count,omitempty"`
	SuccessCount     int             `json:"success_count,omitempty"`
	Successes        []StatusDetails `json:"successes,omitempty"`
	FailureCount     int             `json:"failure_count,omitempty"`
	Failures         []StatusDetails `json:"failures,omitempty"`
	PendingCount     int             `json:"pending_count,omitempty"`
	Pendings         []StatusDetails `json:"pendings,omitempty"`
}

/*
StatusDetails - This type defines the details for various status elements
*/
type StatusDetails struct {
	ID      string `json:"id,omitempty"`
	Version string `json:"version,omitempty"`
	Message string `json:"message,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new TAXII Status object and return it as a
pointer.
*/
func New() *Status {
	var obj Status
	return &obj
}

/*
NewStatusDetails - This function will create a new TAXII Status Detail object and
return it as a pointer.
*/
func NewStatusDetails() *StatusDetails {
	var obj StatusDetails
	return &obj
}

// ----------------------------------------------------------------------
// Public Methods - Status
// ----------------------------------------------------------------------

/*
SetStatusCompleted - This method will set the status to completed.
*/
func (r *Status) SetStatusCompleted() error {
	r.Status = "Completed"
	return nil
}

/*
SetStatusPending - This method will set the status to pending.
*/
func (r *Status) SetStatusPending() error {
	r.Status = "Pending"
	return nil
}

/*
SetRequestTimestamp - This method will set the Request time stamp to the provided
value.
*/
func (r *Status) SetRequestTimestamp(s string) error {
	r.RequestTimestamp = s
	return nil
}

/*
SetRequestTimestampToCurrentTime - This method will set the Request time stamp
to the current time with micro second precision.
*/
func (r *Status) SetRequestTimestampToCurrentTime() error {
	r.RequestTimestamp = timestamp.CurrentTime("micro")
	return nil
}

/*
SetTotalCount - Set Total Count to value provided.
*/
func (r *Status) SetTotalCount(i int) error {
	r.TotalCount = i
	return nil
}

/*
IncreaseTotalCount - Increase Total count number.
*/
func (r *Status) IncreaseTotalCount() error {
	r.TotalCount++
	return nil
}

/*
SetSuccessCount - Set Success Count to value provided.
*/
func (r *Status) SetSuccessCount(i int) error {
	r.SuccessCount = i
	return nil
}

/*
IncreaseSuccessCount - Increase Success count number.
*/
func (r *Status) IncreaseSuccessCount() error {
	r.SuccessCount++
	return nil
}

/*
SetFailureCount - Set Failure Count to value provided.
*/
func (r *Status) SetFailureCount(i int) error {
	r.FailureCount = i
	return nil
}

/*
IncreaseFailureCount - Increase Failure count number.
*/
func (r *Status) IncreaseFailureCount() error {
	r.FailureCount++
	return nil
}

/*
SetPendingCount - Set Pending Count to value provided.
*/
func (r *Status) SetPendingCount(i int) error {
	r.PendingCount = i
	return nil
}

/*
IncreasePendingCount - Increase Pending count number.
*/
func (r *Status) IncreasePendingCount() error {
	r.PendingCount++
	return nil
}

/*
NewSuccessDetails - This method will create an empty SuccessDetails type, add it
to the Successes property, and return a pointer to that instance.
*/
func (r *Status) NewSuccessDetails() (*StatusDetails, error) {
	o := NewStatusDetails()
	positionThatAppendWillUse := len(r.Successes)
	r.Successes = append(r.Successes, *o)
	return &r.Successes[positionThatAppendWillUse], nil
}

/*
AddSuccessDetails - This method will take in a SuccessDetails type as a pointer
and add it to the Successes property. One would use this if they manually created
their own SuccessDetails type and then wanted to add it to the Successes property.
*/
func (r *Status) AddSuccessDetails(o *StatusDetails) (int, error) {
	positionThatAppendWillUse := len(r.Successes)
	r.Successes = append(r.Successes, *o)
	return positionThatAppendWillUse, nil
}

/*
CreateSuccessDetails - This method will create a SuccessDetails type based on the
data provided and will add it to the Successes property.
*/
func (r *Status) CreateSuccessDetails(id, ver, mesg string) error {
	s, _ := r.NewSuccessDetails()
	s.SetID(id)
	s.SetVersion(ver)
	s.SetMessage(mesg)
	return nil
}

/*
NewFailureDetails - This method will create an empty SuccessDetails type, add it
to the Failures property, and return a pointer to that instance.
*/
func (r *Status) NewFailureDetails() (*StatusDetails, error) {
	o := NewStatusDetails()
	positionThatAppendWillUse := len(r.Failures)
	r.Failures = append(r.Failures, *o)
	return &r.Failures[positionThatAppendWillUse], nil
}

/*
AddFailureDetails - This method will take in a SuccessDetails type as a pointer
and add it to the Failures property. One would use this if they manually created
their own SuccessDetails type and then wanted to add it to the Failures property.
*/
func (r *Status) AddFailureDetails(o *StatusDetails) (int, error) {
	positionThatAppendWillUse := len(r.Failures)
	r.Failures = append(r.Failures, *o)
	return positionThatAppendWillUse, nil
}

/*
CreateFailureDetails - This method will create a FailureDetails type based on the
data provided and will add it to the Failures property.
*/
func (r *Status) CreateFailureDetails(id, ver, mesg string) error {
	s, _ := r.NewFailureDetails()
	s.SetID(id)
	s.SetVersion(ver)
	s.SetMessage(mesg)
	return nil
}

/*
NewPendingDetails - This method will create an empty SuccessDetails type, add it
to the Pendings property, and return a pointer to that instance.
*/
func (r *Status) NewPendingDetails() (*StatusDetails, error) {
	o := NewStatusDetails()
	positionThatAppendWillUse := len(r.Pendings)
	r.Pendings = append(r.Pendings, *o)
	return &r.Pendings[positionThatAppendWillUse], nil
}

/*
AddPendingDetails - This method will take in a SuccessDetails type as a pointer
and add it to the Pendings property. One would use this if they manually created
their own SuccessDetails type and then wanted to add it to the Pendings property.
*/
func (r *Status) AddPendingDetails(o *StatusDetails) (int, error) {
	positionThatAppendWillUse := len(r.Pendings)
	r.Pendings = append(r.Pendings, *o)
	return positionThatAppendWillUse, nil
}

/*
CreatePendingDetails - This method will create a SuccessDetails type based on the
data provided and will add it to the Pending property.
*/
func (r *Status) CreatePendingDetails(id, ver, mesg string) error {
	s, _ := r.NewPendingDetails()
	s.SetID(id)
	s.SetVersion(ver)
	s.SetMessage(mesg)
	return nil
}

// ----------------------------------------------------------------------
// Public Methods - Status Details
// ----------------------------------------------------------------------

/*
SetID - This method will set the ID to the provided value.
*/
func (r *StatusDetails) SetID(s string) error {
	r.ID = s
	return nil
}

/*
SetVersion - This method will set the Version to the provided value.
*/
func (r *StatusDetails) SetVersion(s string) error {
	r.Version = s
	return nil
}

/*
SetMessage - This method will set the Message to the provided value.
*/
func (r *StatusDetails) SetMessage(s string) error {
	r.Message = s
	return nil
}
