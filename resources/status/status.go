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

func (r *Status) SetStatusCompleted() error {
	r.Status = "Completed"
	return nil
}

func (r *Status) SetStatusPending() error {
	r.Status = "Pending"
	return nil
}

func (r *Status) SetRequestTimestamp(s string) error {
	r.RequestTimestamp = s
	return nil
}

func (r *Status) SetRequestTimestampToCurrentTime() error {
	r.RequestTimestamp = timestamp.CurrentTime("micro")
	return nil
}

func (r *Status) SetTotalCount(i int) error {
	r.TotalCount = i
	return nil
}

func (r *Status) IncreaseTotalCount() error {
	r.TotalCount++
	return nil
}

func (r *Status) SetSuccessCount(i int) error {
	r.SuccessCount = i
	return nil
}

func (r *Status) IncreaseSuccessCount() error {
	r.SuccessCount++
	return nil
}

func (r *Status) SetFailureCount(i int) error {
	r.FailureCount = i
	return nil
}

func (r *Status) IncreaseFailureCount() error {
	r.FailureCount++
	return nil
}

func (r *Status) SetPendingCount(i int) error {
	r.PendingCount = i
	return nil
}

func (r *Status) IncreasePendingCount() error {
	r.PendingCount++
	return nil
}

func (r *Status) NewSuccessDetails() (*StatusDetails, error) {
	o := NewStatusDetails()
	positionThatAppendWillUse := len(r.Successes)
	r.Successes = append(r.Successes, *o)
	return &r.Successes[positionThatAppendWillUse], nil
}

func (r *Status) AddSuccessDetails(o *StatusDetails) (int, error) {
	positionThatAppendWillUse := len(r.Successes)
	r.Successes = append(r.Successes, *o)
	return positionThatAppendWillUse, nil
}

func (r *Status) CreateSuccessDetails(id, ver, mesg string) error {
	s, _ := r.NewSuccessDetails()
	s.SetID(id)
	s.SetVersion(ver)
	s.SetMessage(mesg)
	return nil
}

func (r *Status) NewFailureDetails() (*StatusDetails, error) {
	o := NewStatusDetails()
	positionThatAppendWillUse := len(r.Failures)
	r.Failures = append(r.Failures, *o)
	return &r.Failures[positionThatAppendWillUse], nil
}

func (r *Status) AddFailureDetails(o *StatusDetails) (int, error) {
	positionThatAppendWillUse := len(r.Failures)
	r.Failures = append(r.Failures, *o)
	return positionThatAppendWillUse, nil
}

func (r *Status) CreateFailureDetails(id, ver, mesg string) error {
	s, _ := r.NewFailureDetails()
	s.SetID(id)
	s.SetVersion(ver)
	s.SetMessage(mesg)
	return nil
}

func (r *Status) NewPendingDetails() (*StatusDetails, error) {
	o := NewStatusDetails()
	positionThatAppendWillUse := len(r.Pendings)
	r.Pendings = append(r.Pendings, *o)
	return &r.Pendings[positionThatAppendWillUse], nil
}

func (r *Status) AddPendingDetails(o *StatusDetails) (int, error) {
	positionThatAppendWillUse := len(r.Pendings)
	r.Pendings = append(r.Pendings, *o)
	return positionThatAppendWillUse, nil
}

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

func (r *StatusDetails) SetID(s string) error {
	r.ID = s
	return nil
}

func (r *StatusDetails) SetVersion(s string) error {
	r.Version = s
	return nil
}

func (r *StatusDetails) SetMessage(s string) error {
	r.Message = s
	return nil
}
