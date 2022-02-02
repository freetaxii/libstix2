// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package status

import (
	"github.com/freetaxii/libstix2/objects"
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
	objects.IDProperty
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
func (o *Status) SetStatusCompleted() error {
	o.Status = "Completed"
	return nil
}

/*
SetStatusPending - This method will set the status to pending.
*/
func (o *Status) SetStatusPending() error {
	o.Status = "Pending"
	return nil
}

/*
SetRequestTimestamp - This method will set the Request time stamp to the provided
value.
*/
func (o *Status) SetRequestTimestamp(s string) error {
	o.RequestTimestamp = s
	return nil
}

/*
SetRequestTimestampToCurrentTime - This method will set the Request time stamp
to the current time with micro second precision.
*/
func (o *Status) SetRequestTimestampToCurrentTime() error {
	o.RequestTimestamp = objects.GetCurrentTime("micro")
	return nil
}

/*
SetTotalCount - Set Total Count to value provided.
*/
func (o *Status) SetTotalCount(i int) error {
	o.TotalCount = i
	return nil
}

/*
IncreaseTotalCount - Increase Total count number.
*/
func (o *Status) IncreaseTotalCount() error {
	o.TotalCount++
	return nil
}

/*
SetSuccessCount - Set Success Count to value provided.
*/
func (o *Status) SetSuccessCount(i int) error {
	o.SuccessCount = i
	return nil
}

/*
IncreaseSuccessCount - Increase Success count number.
*/
func (o *Status) IncreaseSuccessCount() error {
	o.SuccessCount++
	return nil
}

/*
SetFailureCount - Set Failure Count to value provided.
*/
func (o *Status) SetFailureCount(i int) error {
	o.FailureCount = i
	return nil
}

/*
IncreaseFailureCount - Increase Failure count number.
*/
func (o *Status) IncreaseFailureCount() error {
	o.FailureCount++
	return nil
}

/*
SetPendingCount - Set Pending Count to value provided.
*/
func (o *Status) SetPendingCount(i int) error {
	o.PendingCount = i
	return nil
}

/*
IncreasePendingCount - Increase Pending count number.
*/
func (o *Status) IncreasePendingCount() error {
	o.PendingCount++
	return nil
}

/*
NewSuccessDetails - This method will create an empty SuccessDetails type, add it
to the Successes property, and return a pointer to that instance.
*/
func (o *Status) NewSuccessDetails() (*StatusDetails, error) {
	obj := NewStatusDetails()
	positionThatAppendWillUse := len(o.Successes)
	o.Successes = append(o.Successes, *obj)
	return &o.Successes[positionThatAppendWillUse], nil
}

/*
AddSuccessDetails - This method will take in a SuccessDetails type as a pointer
and add it to the Successes property. One would use this if they manually created
their own SuccessDetails type and then wanted to add it to the Successes property.
*/
func (o *Status) AddSuccessDetails(s *StatusDetails) (int, error) {
	positionThatAppendWillUse := len(o.Successes)
	o.Successes = append(o.Successes, *s)
	return positionThatAppendWillUse, nil
}

/*
CreateSuccessDetails - This method will create a SuccessDetails type based on the
data provided and will add it to the Successes property.
*/
func (o *Status) CreateSuccessDetails(id, ver, mesg string) error {
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
func (o *Status) NewFailureDetails() (*StatusDetails, error) {
	obj := NewStatusDetails()
	positionThatAppendWillUse := len(o.Failures)
	o.Failures = append(o.Failures, *obj)
	return &o.Failures[positionThatAppendWillUse], nil
}

/*
AddFailureDetails - This method will take in a SuccessDetails type as a pointer
and add it to the Failures property. One would use this if they manually created
their own SuccessDetails type and then wanted to add it to the Failures property.
*/
func (o *Status) AddFailureDetails(o *StatusDetails) (int, error) {
	positionThatAppendWillUse := len(o.Failures)
	o.Failures = append(o.Failures, *o)
	return positionThatAppendWillUse, nil
}

/*
CreateFailureDetails - This method will create a FailureDetails type based on the
data provided and will add it to the Failures property.
*/
func (o *Status) CreateFailureDetails(id, ver, mesg string) error {
	s, _ := o.NewFailureDetails()
	s.SetID(id)
	s.SetVersion(ver)
	s.SetMessage(mesg)
	return nil
}

/*
NewPendingDetails - This method will create an empty SuccessDetails type, add it
to the Pendings property, and return a pointer to that instance.
*/
func (o *Status) NewPendingDetails() (*StatusDetails, error) {
	obj := NewStatusDetails()
	positionThatAppendWillUse := len(o.Pendings)
	o.Pendings = append(r.Pendings, *obj)
	return &r.Pendings[positionThatAppendWillUse], nil
}

/*
AddPendingDetails - This method will take in a SuccessDetails type as a pointer
and add it to the Pendings property. One would use this if they manually created
their own SuccessDetails type and then wanted to add it to the Pendings property.
*/
func (o *Status) AddPendingDetails(o *StatusDetails) (int, error) {
	positionThatAppendWillUse := len(o.Pendings)
	o.Pendings = append(o.Pendings, *o)
	return positionThatAppendWillUse, nil
}

/*
CreatePendingDetails - This method will create a SuccessDetails type based on the
data provided and will add it to the Pending property.
*/
func (o *Status) CreatePendingDetails(id, ver, mesg string) error {
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
func (o *StatusDetails) SetID(s string) error {
	o.ID = s
	return nil
}

/*
SetVersion - This method will set the Version to the provided value.
*/
func (o *StatusDetails) SetVersion(s string) error {
	o.Version = s
	return nil
}

/*
SetMessage - This method will set the Message to the provided value.
*/
func (o *StatusDetails) SetMessage(s string) error {
	o.Message = s
	return nil
}
