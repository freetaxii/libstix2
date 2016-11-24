// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package infrastructure

import (
	"github.com/freetaxii/libstix2/objects/common"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type InfrastructureType struct {
	common.CommonPropertiesType
	common.DescriptivePropertiesType
	common.KillChainPhasesType
	First_seen           string `json:"first_seen,omitempty"`
	First_seen_precision string `json:"first_seen_precision,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() InfrastructureType {
	var obj InfrastructureType
	obj.MessageType = "infrastructure"
	obj.Id = obj.NewId("infrastructure")
	obj.Created = obj.GetCurrentTime()
	obj.Modified = obj.Created
	obj.Version = 1
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - InfrastructureType
// ----------------------------------------------------------------------

// SetFirstSeen takes in two parameters and returns and error if there is one
// param: t a timestamp in either time.Time or string format
// param: s a timestamp precision in string format
func (this *InfrastructureType) SetFirstSeen(t interface{}, s string) error {

	ts, err := this.VerifyTimestamp(t)
	if err != nil {
		return err
	}
	this.First_seen = ts

	p, err := this.VerifyPrecision(s)
	if err != nil {
		return err
	}
	this.First_seen_precision = p

	return nil
}
