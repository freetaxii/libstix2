// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package indicator

import (
	"github.com/freetaxii/libstix2/messages/common"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type IndicatorType struct {
	MessageType  string   `json:"type,omitempty"`
	Id           string   `json:"id,omitempty"`
	CreatedAt    string   `json:"created_at,omitempty"`
	Title        string   `json:"title,omitempty"`
	Descriptions []string `json:"descriptions,omitempty"`
}

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

func New() IndicatorType {
	var obj IndicatorType
	obj.MessageType = "indicator"
	obj.Id = common.CreateId("indicator")
	obj.CreatedAt = common.GetCurrentTime()
	return obj
}

// ----------------------------------------------------------------------
// Public Methods - IndicatorType
// ----------------------------------------------------------------------

func (this *IndicatorType) SetTitle(t string) {
	this.Title = t
}

// ----------------------------------------------------------------------
// Private Methods - IndicatorType
// ----------------------------------------------------------------------
